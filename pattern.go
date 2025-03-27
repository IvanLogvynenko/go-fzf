package gofzf

type PatternStruct struct {
	pattern string
}

func Pattern(pattern string) PatternStruct {
	return PatternStruct{pattern}
}

func (p PatternStruct) Render() []string {
	return []string{"--filter", p.pattern}
}

func (p PatternStruct) RenderStruct(opts []Struct) []string { return p.Render() }
