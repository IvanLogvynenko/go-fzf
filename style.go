package gofzf

type Styles string

const (
	Minimal Styles = "minimal"
	Full    Styles = "full"
	Default Styles = "default"
)

type StyleStruct struct {
	style Styles
}

func Style(style Styles) StyleStruct {
	return StyleStruct{style}
}

func (s StyleStruct) Render() []string {
	return []string{"--style=" + string(s.style)}
}
func (s StyleStruct) RenderStruct(_ []Struct) []string {
	return s.Render()
}
