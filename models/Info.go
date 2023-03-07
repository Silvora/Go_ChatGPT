package models

import "github.com/sashabaranov/go-openai"

type InfoData struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

//var openAiModel string

func (data *InfoData) ModelType() string {
	//CodexCodeDavinci002 = "code-davinci-002"
	//CodexCodeCushman001 = "code-cushman-001"
	//CodexCodeDavinci001 = "code-davinci-001"
	//GPT3Dot5Turbo0301       = "gpt-3.5-turbo-0301"
	//GPT3Dot5Turbo           = "gpt-3.5-turbo"
	//GPT3TextDavinci003      = "text-davinci-003"
	//GPT3TextDavinci002      = "text-davinci-002"
	//GPT3TextCurie001        = "text-curie-001"
	//GPT3TextBabbage001      = "text-babbage-001"
	//GPT3TextAda001          = "text-ada-001"
	//GPT3TextDavinci001      = "text-davinci-001"
	//GPT3DavinciInstructBeta = "davinci-instruct-beta"
	//GPT3Davinci             = "davinci"
	//GPT3CurieInstructBeta   = "curie-instruct-beta"
	//GPT3Curie               = "curie"
	//GPT3Ada                 = "ada"
	//GPT3Babbage             = "babbage"
	var openAiModel string
	model := data.Model
	switch {
	case model == "code-davinci-002":
		openAiModel = openai.CodexCodeDavinci002
	case model == "code-cushman-001":
		openAiModel = openai.CodexCodeCushman001
	case model == "code-davinci-001":
		openAiModel = openai.CodexCodeDavinci001
	case model == "gpt-3.5-turbo-0301":
		openAiModel = openai.GPT3Dot5Turbo0301
	case model == "gpt-3.5-turbo":
		openAiModel = openai.GPT3Dot5Turbo
	case model == "text-davinci-003":
		openAiModel = openai.GPT3TextDavinci003
	case model == "text-davinci-002":
		openAiModel = openai.GPT3TextDavinci002
	case model == "text-davinci-001":
		openAiModel = openai.GPT3TextDavinci001
	case model == "text-curie-001":
		openAiModel = openai.GPT3TextCurie001
	case model == "text-babbage-001":
		openAiModel = openai.GPT3TextBabbage001
	case model == "text-ada-001":
		openAiModel = openai.GPT3TextAda001
	case model == "davinci-instruct-beta":
		openAiModel = openai.GPT3DavinciInstructBeta
	case model == "davinci":
		openAiModel = openai.GPT3Davinci
	case model == "curie-instruct-beta":
		openAiModel = openai.GPT3CurieInstructBeta
	case model == "curie":
		openAiModel = openai.GPT3Curie
	case model == "ada":
		openAiModel = openai.GPT3Ada
	case model == "babbage":
		openAiModel = openai.GPT3Babbage
	default:
		openAiModel = openai.GPT3Dot5Turbo

	}

	return openAiModel
}
