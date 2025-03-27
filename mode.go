package gofzf

type Mode interface {
	Render() []string
	RenderStruct(opts []Struct) []string
}
