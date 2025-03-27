package fzfarguments

import "github.com/IvanLogvynenko/go-fzf/interfaces"

type Styles string

const (
	Minimal    Styles = "minimal"
	Full  Styles = "full"
	Default  Styles = "default")

type StyleStruct struct {
	style Styles
}

func Style(style Styles) StyleStruct {
	return StyleStruct{style}
}

func (s StyleStruct) Render() []string {
	return []string{"--style=" + string(s.style)}
}
func (s StyleStruct) RenderStruct(_ []interfaces.Struct) []string {
	return s.Render()
}
