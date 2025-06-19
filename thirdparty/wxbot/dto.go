package wxbot

type BaseBody struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 登录状态
type LoginState struct {
	BaseBody `json:",inline"`
	Data     LoginStateData `json:"data,omitempty"`
}

type LoginStateData struct {
	Status int    `json:"status,omitempty"`
	Wxid   string `json:"wxid,omitempty"`
}

// 通讯录
type Contacts struct {
	BaseBody `json:",inline"`
	Data     ContactsData `json:"data,omitempty"`
}

type ContactsData struct {
	Contacts []ContactsContent `json:"contacts,omitempty"`
	Total    int               `json:"total,omitempty"`
}

type ContactsContent struct {
	Alias    string `json:"Alias,omitempty"`
	NickName string `json:"NickName,omitempty"`
	UserName string `json:"UserName,omitempty"`
}

// 微信群聊信息
type ChatRoom struct {
	BaseBody `json:",inline"`
	Data     ChatRoomData `json:"data,omitempty"`
}

type ChatRoomData struct {
	Announcement            string            `json:"Announcement"`
	AnnouncementEditor      string            `json:"AnnouncementEditor"`
	AnnouncementPublishTime string            `json:"AnnouncementPublishTime"`
	ChatRoomFlag            string            `json:"ChatRoomFlag"`
	ChatRoomStatus          string            `json:"ChatRoomStatus"`
	DisplayNameList         string            `json:"DisplayNameList"`
	InfoVersion             string            `json:"InfoVersion"`
	IsShowName              string            `json:"IsShowName"`
	Members                 map[string]Member `json:"Members"`
	SelfDisplayName         string            `json:"SelfDisplayName"`
	UserNameList            string            `json:"UserNameList"`
}

type Member struct {
	Alias               string `json:"Alias"`
	BigHeadImgURL       string `json:"BigHeadImgUrl"`
	ChatRoomNotify      string `json:"ChatRoomNotify"`
	ChatRoomType        string `json:"ChatRoomType"`
	DelFlag             string `json:"DelFlag"`
	DomainList          string `json:"DomainList"`
	EncryptUserName     string `json:"EncryptUserName"`
	ExtraBuf            string `json:"ExtraBuf"`
	HeadImgMd5          string `json:"HeadImgMd5"`
	LabelIDList         string `json:"LabelIDList"`
	NickName            string `json:"NickName"`
	PYInitial           string `json:"PYInitial"`
	QuanPin             string `json:"QuanPin"`
	Remark              string `json:"Remark"`
	RemarkPYInitial     string `json:"RemarkPYInitial"`
	RemarkQuanPin       string `json:"RemarkQuanPin"`
	SmallHeadImgURL     string `json:"SmallHeadImgUrl"`
	Type                string `json:"Type"`
	UserName            string `json:"UserName"`
	VerifyFlag          string `json:"VerifyFlag"`
	ProfilePicture      string `json:"profilePicture"`
	ProfilePictureSmall string `json:"profilePictureSmall"`
}

// 微信消息请求体
type WxMsg struct {
	Wxid    string   `json:"wxid,omitempty"`
	Content string   `json:"content,omitempty"`
	Atlist  []string `json:"atlist,omitempty"`
}

// 微信消息ws
type WxGeneralMsg struct {
	Data  []WxGeneralMsgData `json:"data"`
	Total int                `json:"total"`
	Wxid  string             `json:"wxid"`
}

type WxGeneralMsgData struct {
	BytesExtra     string `json:"BytesExtra"`
	BytesTrans     string `json:"BytesTrans"`
	Content        string `json:"Content"`
	CreateTime     string `json:"CreateTime"`
	DisplayContent string `json:"DisplayContent"`
	FlagEx         string `json:"FlagEx"`
	IsSender       string `json:"IsSender"`
	MsgSequence    string `json:"MsgSequence"`
	MsgServerSeq   string `json:"MsgServerSeq"`
	MsgSvrID       string `json:"MsgSvrID"`
	Reserved0      string `json:"Reserved0"`
	Reserved1      string `json:"Reserved1"`
	Sender         string `json:"Sender"`
	Sequence       string `json:"Sequence"`
	Status         string `json:"Status"`
	StatusEx       string `json:"StatusEx"`
	StrContent     string `json:"StrContent"`
	StrTalker      string `json:"StrTalker"`
	SubType        string `json:"SubType"`
	TalkerID       string `json:"TalkerId"`
	Type           string `json:"Type"`
	LocalID        string `json:"localId"`
}

// 注册回调
type Callback struct {
	Url     string `json:"url"`
	Timeout int    `json:"timeout"`
	Type    string `json:"type"`
}
