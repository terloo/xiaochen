package wxbot

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/pkg/errors"
	"github.com/terloo/xiaochen/service/family"

	"github.com/terloo/xiaochen/client"
	"github.com/terloo/xiaochen/config"
)

var host = fmt.Sprintf("http://%s/api/", config.NewLoader("main.wxBotHost").Get())

func KeepAlive(ctx context.Context) {
	_ = SendMsg(ctx, "ping", family.TestChatroomWxid)
}

func CheckAlive(ctx context.Context) (bool, error) {
	contacts, err := GetContacts(ctx)
	if err != nil {
		return false, err
	}
	if len(contacts.Data.Contacts) == 0 {
		return false, nil
	}
	return true, nil
}

func GetWxid(ctx context.Context) (string, error) {
	b, err := client.HttpGet(ctx, host+"checklogin", nil, nil)
	if err != nil {
		return "", err
	}

	loginState := &LoginState{}
	err = json.Unmarshal(b, loginState)
	if err != nil {
		return "", errors.WithStack(err)
	}

	if loginState.Code != 200 {
		return "", errors.Errorf("not login: %s", loginState.Msg)
	}
	return loginState.Data.Wxid, nil
}

var selfWxid string

func GetWxidWithCache(ctx context.Context) (string, error) {
	if selfWxid == "" {
		wxid, err := GetWxid(ctx)
		if err != nil {
			return "", err
		}
		selfWxid = wxid
	}
	return selfWxid, nil
}

func GetContacts(ctx context.Context) (*Contacts, error) {
	b, err := client.HttpGet(ctx, host+"dbcontacts", nil, nil)
	if err != nil {
		return nil, err
	}

	contacts := &Contacts{}
	err = json.Unmarshal(b, contacts)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return contacts, nil
}

func GetChatroom(ctx context.Context, wxid string) (*ChatRoom, error) {
	if !strings.Contains(wxid, "@chatroom") {
		return nil, errors.New("not chatroom wxid")
	}
	param := url.Values{
		"wxid": []string{wxid},
	}
	b, err := client.HttpGet(ctx, host+"dbchatroom", nil, param)
	if err != nil {
		return nil, err
	}

	chatroom := &ChatRoom{}
	err = json.Unmarshal(b, chatroom)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return chatroom, nil
}

func SendMsg(ctx context.Context, msg string, wxids ...string) error {
	for _, wxid := range wxids {
		msg = strings.TrimSpace(msg)
		wxmsg := &WxMsg{
			Wxid:    wxid,
			Content: msg,
			Atlist:  make([]string, 0),
		}
		b, err := client.HttpPost(ctx, host+"sendtxtmsg", nil, nil, wxmsg)
		if err != nil {
			log.Println("send wxmsg err:", err)
			continue
		}
		result := &BaseBody{}
		err = json.Unmarshal(b, result)
		if err != nil {
			log.Printf("send wxmsg err: %v, resp: %s\n", err, string(b))
			continue
		}
		if result.Code != 200 {
			log.Printf("send wxmsg err, message: %s\n", result.Msg)
		}
	}
	return nil
}

func RegistryCallback(ctx context.Context, callbackURL string) error {
	callback := &Callback{
		Url:     callbackURL,
		Timeout: 10000,
		Type:    "public-msg",
	}
	b, err := client.HttpPost(ctx, host+"syncurl", nil, nil, callback)
	if err != nil {
		return err
	}
	result := &BaseBody{}
	err = json.Unmarshal(b, result)
	if err != nil {
		return errors.WithStack(err)
	}
	log.Println(result)
	return nil
}
