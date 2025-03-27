package options

import "github.com/IvanLogvynenko/go-fzf/interfaces"

type BindsStruct struct {
	binds string
}

func Binds() BindsStruct {
	return BindsStruct{"tab:down,shift-tab:up"}
}

func CustomBinds(binds string) BindsStruct {
	return BindsStruct{binds}
}

func (b BindsStruct) Render() []string {
	return []string{"--bind", b.binds}
}
func (b BindsStruct) RenderStruct(opts []interfaces.Struct) []string { return b.Render() }
