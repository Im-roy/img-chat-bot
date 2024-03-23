package model

type PromptImageModel struct {
	ExtensionName string
	Data          []byte
}

type RequestModel struct {
	Prompt string `json:"data"`
}
