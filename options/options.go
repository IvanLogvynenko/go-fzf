package options

import "github.com/IvanLogvynenko/go-fzf/interfaces"

type Mode interface {
	Render() []string
	RenderStruct(opts []interfaces.Struct) []string
}
