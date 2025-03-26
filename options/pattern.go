package options

type PatternStruct struct {
	pattern string
}

func Pattern(pattern string) PatternStruct {
	return PatternStruct{pattern}
}

func (p PatternStruct) Render() []string {
	return []string{"--filter", p.pattern}
}
