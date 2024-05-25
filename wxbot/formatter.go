package wxbot

import (
	"encoding/xml"
	"strings"
)

type FormattedMessage struct {
	Self         bool   // 是否自己发送
	Chat         string // 对话id
	Sender       string // 消息发送者id
	Content      string
	ReferMessage string
	ReferSender  string
	Chatroom     bool
	At           bool
	Command      bool
	CommandName  string
	CommandArgs  []string
}

type WxMessageContent struct {
	Appmsg Appmsg `xml:"appmsg"`
	// 消息发送人
	Fromusername string `xml:"fromusername"`
	Scene        int    `xml:"scene"`
	Commenturl   string `xml:"commenturl"`
}

type Appmsg struct {
	Title    string   `xml:"title"`    // 消息内容
	Refermsg Refermsg `xml:"refermsg"` // 被引用的消息
}

type Refermsg struct {
	Type        int    `xml:"type"`
	Svrid       int64  `xml:"svrid"`
	Fromusr     string `xml:"fromusr"`
	Chatusr     string `xml:"chatusr"` // 被引用消息的发送人
	Displayname string `xml:"displayname"`
	Msgsource   string `xml:"msgsource"`
	Content     string `xml:"content"` // 被引用的消息内容
	Strid       string `xml:"strid"`
	Createtime  int    `xml:"createtime"`
}

func FormatMessage(msg WxGeneralMsgData) (FormattedMessage, error) {
	result := FormattedMessage{
		Chat: msg.StrTalker,
	}

	result.Self = msg.IsSender == "1"

	if len(msg.Sender) == 0 {
		// 私聊
		result.Sender = msg.StrTalker
	} else {
		// 群聊
		result.Sender = msg.Sender
		result.Chatroom = true
	}

	if len(msg.StrContent) == 0 {
		// 引用，解析引用消息
		wxMessageContent := &WxMessageContent{}
		msg.Content = strings.Replace(msg.Content, "\\n", "", -1)
		msg.Content = strings.Replace(msg.Content, "\\t", "", -1)
		msg.Content = strings.Replace(msg.Content, "\\\"", "\"", -1)
		msg.Content = strings.Replace(msg.Content, "<?xml version=\"1.0\"?>", "", -1)
		err := xml.Unmarshal([]byte(msg.Content), &wxMessageContent)
		if err != nil {
			return FormattedMessage{}, err
		}
		result.Content = wxMessageContent.Appmsg.Title
		result.ReferMessage = wxMessageContent.Appmsg.Refermsg.Content
		result.ReferSender = wxMessageContent.Appmsg.Refermsg.Chatusr
	} else {
		// 非引用
		result.Content = msg.StrContent
	}

	if result.Chatroom {
		if strings.HasPrefix(result.Content, "@xiaochen ") {
			result.At = true
		}
		result.Content = strings.TrimPrefix(result.Content, "@xiaochen ")
	}
	if strings.HasPrefix(result.Content, "/") {
		result.Command = true
		split := strings.Split(result.Content, " ")
		if len(split) == 0 {
			result.Command = false
		} else {
			result.CommandName = strings.TrimPrefix(split[0], "/")
			if len(split) >= 2 {
				result.CommandArgs = split[1:]
			}
		}
	}
	result.Content = strings.TrimPrefix(result.Content, "/")

	return result, nil
}
