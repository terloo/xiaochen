package mcp

import (
	"context"
	"encoding/json"
	"log"

	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/pkg/errors"
	"github.com/sashabaranov/go-openai"
)

type ClientManager struct {
	// clientName -> client
	clientMap map[string]*client.Client
	// toolName -> tool
	toolsMap map[string][]*mcp.Tool
	// toolName -> client 保存tool和client的从属关系
	toolsClientMap map[string]*client.Client
}

func NewClientManger() *ClientManager {
	return &ClientManager{
		clientMap:      make(map[string]*client.Client),
		toolsMap:       make(map[string][]*mcp.Tool),
		toolsClientMap: make(map[string]*client.Client),
	}
}

func (m *ClientManager) InitializeAll(ctx context.Context) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	c1, err := client.NewStdioMCPClient("npx", []string{}, "-y", "howtocook-mcp")
	if err != nil {
		return errors.Wrapf(err, "create mcp client error")
	}
	c1Resp, err := c1.Initialize(ctx, mcp.InitializeRequest{})
	if err != nil {
		return errors.Wrapf(err, "initialize mcp client error")
	}
	log.Printf("initialize mcp server success, name: %s, version: %s \n", c1Resp.ServerInfo.Name, c1Resp.ServerInfo.Version)
	m.clientMap[c1Resp.ServerInfo.Name] = c1

	c2, err := client.NewStdioMCPClient("xiaochen_mcp", []string{}, "")
	if err != nil {
		return errors.Wrapf(err, "create mcp client error")
	}
	c2Resp, err := c2.Initialize(ctx, mcp.InitializeRequest{})
	if err != nil {
		return errors.Wrapf(err, "initialize mcp client error")
	}
	log.Printf("initialize mcp server success, name: %s, version: %s \n", c2Resp.ServerInfo.Name, c2Resp.ServerInfo.Version)
	m.clientMap[c2Resp.ServerInfo.Name] = c2

	return nil
}

func (m *ClientManager) GetAllTools(ctx context.Context) ([]openai.Tool, error) {
	var result []openai.Tool
	for _, c := range m.clientMap {
		tools, err := c.ListTools(ctx, mcp.ListToolsRequest{})
		if err != nil {
			return nil, err
		}
		for _, tool := range tools.Tools {
			m.toolsClientMap[tool.Name] = c
		}
		result = append(result, mcpToolsToOpenaiTools(tools.Tools)...)
	}
	return result, nil
}

func (m *ClientManager) CallTool(ctx context.Context, toolName string, arguments string) (*mcp.CallToolResult, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	c, ok := m.toolsClientMap[toolName]
	if !ok {
		return nil, errors.Errorf("tool %s not found", toolName)
	}

	// 处理参数
	var argumentsJson map[string]any
	err := json.Unmarshal([]byte(arguments), &argumentsJson)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	log.Printf("mcp call start. name: %s, args %s\n", toolName, arguments)
	toolResult, err := c.CallTool(ctx, mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Name:      toolName,
			Arguments: argumentsJson,
		},
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if toolResult.IsError {
		log.Printf("mcp call error. result %s\n", toolResult.Content[0])
	} else {
		log.Printf("mcp call finish. result %s\n", toolResult.Content[0])
	}
	return toolResult, nil
}

func mcpToolsToOpenaiTools(mcpTools []mcp.Tool) []openai.Tool {
	var openaiTools []openai.Tool
	if len(mcpTools) > 0 {
		for _, tool := range mcpTools {
			openaiTools = append(openaiTools, openai.Tool{
				Type: openai.ToolTypeFunction,
				Function: &openai.FunctionDefinition{
					Name:        tool.GetName(),
					Description: tool.Description,
					Parameters:  tool.InputSchema,
				},
			})
		}
	}
	return openaiTools
}
