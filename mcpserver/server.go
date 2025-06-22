package mcpserver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/pkg/errors"
	"github.com/terloo/xiaochen/thirdparty/gd"
)

type MCPServer struct {
	server *server.MCPServer
}

func NewMCPServer() *MCPServer {
	s := server.NewMCPServer(
		"xiaochen_mcp",
		"0.0.1",
		server.WithLogging(),
		server.WithRecovery(),
	)
	m := &MCPServer{
		server: s,
	}

	xiaochenSearchMusic := mcp.NewTool(
		"xiaochen_search_music",
		mcp.WithDescription("通过歌曲的名字、演唱者、专辑来搜索歌曲"),
		mcp.WithString("name", mcp.Description("歌曲名字"), mcp.Required()),
		mcp.WithArray("artist", mcp.Description("歌曲演唱者"), mcp.Items(map[string]any{
			"type": "string",
		})),
		mcp.WithString("album", mcp.Description("歌曲专辑")),
	)
	s.AddTool(xiaochenSearchMusic, m.searchMusic)

	xiaochenDownloadMusic := mcp.NewTool(
		"xiaochen_download_music",
		mcp.WithDescription("通过歌曲的名字、演唱者、专辑等找到并下载歌曲，需要先搜索歌曲获得id、source"),
		mcp.WithString("id", mcp.Description("歌曲id"), mcp.Required()),
		mcp.WithString("source", mcp.Description("歌曲来源"), mcp.Required()),
		mcp.WithString("name", mcp.Description("歌曲名字"), mcp.Required()),
		mcp.WithArray("artist", mcp.Description("歌曲演唱者"), mcp.Items(map[string]any{
			"type": "string",
		})),
		mcp.WithString("album", mcp.Description("歌曲专辑")),
	)
	s.AddTool(xiaochenDownloadMusic, m.downloadMusic)

	return m
}

func (m *MCPServer) Start() error {
	file, err := os.OpenFile(filepath.Join(os.TempDir(), "mcp_server.log"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	err = server.ServeStdio(
		m.server,
		server.WithErrorLogger(log.New(file, "", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)),
	)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (m *MCPServer) searchMusic(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	name, err := request.RequireString("name")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	artist := request.GetStringSlice("artist", []string{})
	album := request.GetString("album", "")
	musics, err := gd.SearchMusic(ctx, "kuwo", fmt.Sprintf("%s %s %s", name, strings.Join(artist, " "), album), 30, 1)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	result, err := json.Marshal(musics)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: string(result),
			},
		},
		IsError: len(musics) == 0,
	}, nil
}

func (m *MCPServer) downloadMusic(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := request.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	source, err := request.RequireString("source")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	name, err := request.RequireString("name")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	artist := request.GetStringSlice("artist", []string{})
	album := request.GetString("album", "")

	music, err := gd.GetMusic(ctx, id, source, name, artist, album)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	fileName, md5, err := gd.PersistentMusic(ctx, *music)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: fmt.Sprintf("已保存歌曲到: %s, md5: %s", fileName, md5),
			},
		},
		IsError: false,
	}, nil
}
