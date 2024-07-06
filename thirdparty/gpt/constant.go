package gpt

import (
	"github.com/sashabaranov/go-openai"

	"github.com/terloo/xiaochen/config"
)

var openKey = config.NewLoader("thirdparty.gpt.openKey")
var openHost = "https://api.gpt.ge/"
var ModelName = openai.GPT4Turbo

// var openKey = config.NewLoader("thirdparty.gpt.openFreeKey")
// var openHost = "https://free.gpt.ge/"
// var ModelName = openai.GPT3Dot5Turbo
