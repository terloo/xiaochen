package main

import (
	"log"

	"github.com/terloo/xiaochen/mcpserver"
)

func main() {
	server := mcpserver.NewMCPServer()
	err := server.Start()
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
}
