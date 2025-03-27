package gofzf

import (
	"strings"
)

type BindsStruct struct {
	binds []string
}

func Binds() BindsStruct {
	return BindsStruct{
		[]string{
			"tab:down",
			"shift-tab:up",
			"ctrl-s:select",
			"ctrl-a:select-all",
			"ctrl-d:deselect",
		},
	}
}

func CustomBinds(binds []string) BindsStruct {
	return BindsStruct{binds}
}

func (b BindsStruct) Render() []string {
	return []string{"--bind", strings.Join(b.binds, ",")}
}
func (b BindsStruct) RenderStruct(opts []Struct) []string { return b.Render() }
