package gofzf

type HeaderStruct struct {
	header string
}

func Header(header string) HeaderStruct {
	return HeaderStruct{header}
}

func (h HeaderStruct) Render() []string {
	return []string{"--header=" + h.header}
}
func (h HeaderStruct) RenderStruct(opts []Struct) []string { return h.Render() }
