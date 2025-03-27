package options

import "github.com/IvanLogvynenko/go-fzf/interfaces"

type InfoStruct struct {
	descriptions []string
}

func Info() InfoStruct {
	return InfoStruct{nil}
}
func InfoDescriptions(descriptions []string) InfoStruct {
	return InfoStruct{descriptions}
}

func (i InfoStruct) Render() []string {
	return []string{""}
}
func (i InfoStruct) RenderStruct(opts []interfaces.Struct) []string {
	return nil
}
