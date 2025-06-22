package main_test

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/terloo/almanac"
	"github.com/terloo/xiaochen/service/family"
	"github.com/terloo/xiaochen/service/music"
	"github.com/terloo/xiaochen/service/notify/period"
	"github.com/terloo/xiaochen/thirdparty/wxbot"
	"github.com/terloo/xiaochen/util"
)

func TestReport(t *testing.T) {
	ctx := context.Background()
	wxbot.KeepAlive(ctx)
}

func TestContact(t *testing.T) {
	ctx := context.Background()
	contacts, _ := wxbot.GetContacts(ctx)
	log.Println(contacts)
}

func TestFoo(t *testing.T) {
	parse, _ := time.ParseInLocation(util.DateLayout, "1989-05-11", time.Local)
	log.Println(parse)
}

func TestCal(t *testing.T) {
	month := almanac.NewMonth(2024, 5)
	wxbot.SendMsg(context.Background(), month.FormatCal(), family.TestChatroomWxid)
	log.Println(month.FormatCal())
}

func TestWeather(t *testing.T) {
	weatherNotifier := period.WeatherNotifier{}
	weatherNotifier.Notify(context.Background(), family.TestChatroomWxid)
}

func TestNeteaseMusicLike(t *testing.T) {
	music.PersistentLikeMusic(context.Background())
}

func TestStdioMCPClient(t *testing.T) {
	ctx := context.Background()
	mcpClient, err := client.NewStdioMCPClient("npx", []string{}, "-y", "howtocook-mcp")
	if err != nil {
		t.Fatal(err)
	}
	initialize, err := mcpClient.Initialize(ctx, mcp.InitializeRequest{})
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%s\n", initialize.ServerInfo)
	tools, err := mcpClient.ListTools(ctx, mcp.ListToolsRequest{})
	if err != nil {
		t.Fatal(err)
	}
	for _, tool := range tools.Tools {
		log.Printf("name: %s, description: %s\n", tool.Name, tool.Description)
	}
}
