package service

import (
	"Go_ChatGPT/models"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	"net/http"
)

func Chat(ctx *gin.Context) {
	var data models.InfoData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//fmt.Println(data)

	c := openai.NewClient(data.Token)
	con := context.Background()
	req := openai.CompletionRequest{
		Model:     data.ModelType(),
		MaxTokens: 2048,
		Prompt:    data.Prompt,
	}
	resp, err := c.CreateCompletion(con, req)
	if err != nil {
		return
	}
	fmt.Println(resp.Choices[0].Text)

	ctx.JSON(http.StatusOK, gin.H{
		"code": "200",
		"data": resp.Choices[0].Text,
	})
}

func Help(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":               200,
		"code-cushman-001":   "几乎与 Davinci Codex 一样强大，但速度稍快。这种速度优势可能使其成为实时应用程序的首选。",
		"code-davinci-002":   "功能最强大的 Codex 型号。特别擅长将自然语言翻译成代码。除了补全代码，还支持在代码中插入补全。",
		"code-davinci-001":   "针对代码完成任务进行了优化",
		"gpt-3.5-turbo":      "功能最强大的 GPT-3.5 模型，并针对聊天进行了优化，成本仅为text-davinci-003. 将使用我们最新的模型迭代进行更新。",
		"gpt-3.5-turbo-0301": "2023 年 3 月 1 日的快照gpt-3.5-turbo。与 不同的是gpt-3.5-turbo，此模型不会收到更新，并且只会在 2023 年 6 月 1 日结束的三个月内提供支持。",
		"text-davinci-003":   "可以以比居里、巴贝奇或 ada 模型更好的质量、更长的输出和一致的指令遵循来完成任何语言任务。还支持在文本中插入补全。",
		"text-davinci-002":   "类似的能力，text-davinci-003但训练有监督的微调而不是强化学习",
		"text-curie-001":     "非常有能力，比达芬奇更快，成本更低。",
		"text-babbage-001":   "能够执行简单的任务，速度非常快，成本更低。",
		"text-ada-001":       "能够执行非常简单的任务，通常是 GPT-3 系列中最快的型号，而且成本最低。",
		"text-davinci-001":   "功能最强大的 GPT-3 模型。通常质量更高。",
		"davinci":            "功能最强大的 GPT-3 模型。可以完成其他模型可以完成的任何任务，而且通常质量更高。",
		"curie":              "非常有能力，但比达芬奇更快，成本更低。",
		"ada":                "能够执行非常简单的任务，通常是 GPT-3 系列中最快的型号，而且成本最低。",
		"babbage":            "能够执行简单的任务，速度非常快，成本更低。",
	})
}
