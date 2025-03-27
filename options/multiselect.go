package options

import "fmt"

import "github.com/IvanLogvynenko/go-fzf/interfaces"

type MultiselectStruct struct {
	max int
}

func Multiselect() MultiselectStruct {
	return MultiselectStruct{-1}
}

func MultiselectMax(max int) MultiselectStruct {
	return MultiselectStruct{max}
}

func (m MultiselectStruct) Render() []string {
	if m.max == -1 {
		return []string{"--multi"}
	} else {
		return []string{"--multi", fmt.Sprint(m.max)}
	}
}
func (m MultiselectStruct) RenderStruct(opts []interfaces.Struct) []string { return m.Render() }
