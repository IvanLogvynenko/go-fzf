package fzfarguments

import "github.com/IvanLogvynenko/go-fzf/interfaces"

type PromptStruct struct {
	prompt string
}

func Prompt(prompt string) PromptStruct {
	return PromptStruct{prompt}
}

func (p PromptStruct) Render() []string {
	return []string{"--prompt=" + p.prompt}
}

func (p PromptStruct) RenderStruct(opts []interfaces.Struct) []string {
	return p.Render()
}
