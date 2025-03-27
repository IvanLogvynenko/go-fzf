package gofzf

// Yay! empty struct
type ReverseStruct struct{}

func Reverse() ReverseStruct {
	return ReverseStruct{}
}

func (_ ReverseStruct) Render() []string {
	return []string{"--reverse"}
}

func (r ReverseStruct) RenderStruct(opts []Struct) []string {
	return r.Render()
}
