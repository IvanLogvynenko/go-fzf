package options

import "github.com/IvanLogvynenko/go-fzf/interfaces"

// Yay! empty struct
type ReverseStruct struct{}

func Reverse() ReverseStruct {
	return ReverseStruct{}
}

func (_ ReverseStruct) Render() []string {
	return []string{"--reverse"}
}

func (r ReverseStruct) RenderStruct(opts []interfaces.Struct) []string {
	return r.Render()
}
