package wxbot

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/url"
	"strings"

	"github.com/terloo/xiaochen/client"
)

var host = "http://tx:8080/api/"

func GetWxid(ctx context.Context) (string, error) {
	b, err := client.HttpGet(ctx, host+"checklogin", nil, nil)
	if err != nil {
		return "", err
	}
	loginState := &LoginState{}
	err = json.Unmarshal(b, loginState)
	if err != nil {
		return "", err
	}
	if loginState.Code != 200 {
		log.Println("not login: " + loginState.Msg)
		return "", err
	}
	return loginState.Data.Wxid, nil
}

func GetContacts(ctx context.Context) (*Contacts, error) {
	b, err := client.HttpGet(ctx, host+"dbcontacts", nil, nil)
	if err != nil {
		return nil, err
	}
	contacts := &Contacts{}
	err = json.Unmarshal(b, contacts)
	if err != nil {
		return nil, err
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
		return nil, err
	}
	return chatroom, nil
}

func SendMsg(ctx context.Context, wxid string, msg string) error {
	msg = strings.TrimSpace(msg)
	wxmsg := &WxMsg{
		Wxid:    wxid,
		Content: msg,
		Atlist:  make([]string, 0),
	}
	b, err := client.HttpPost(ctx, host+"sendtxtmsg", nil, nil, wxmsg)
	if err != nil {
		return err
	}
	result := &BaseBody{}
	err = json.Unmarshal(b, result)
	if err != nil {
		return err
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
		return err
	}
	log.Println(result)
	return nil
}
