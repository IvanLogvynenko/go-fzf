package gofzf

type InfoStruct struct {
	info []string
}

func Info(info []string) InfoStruct {
	return InfoStruct{info}
}

func (i InfoStruct) Render() []string {
	previewCmd := "opts=("

	previewCmd += "\"" + i.info[0] + "\""
	for _, value := range i.info[1:] {
		previewCmd += " \"" + value + "\""
	}
	previewCmd += ") && echo ${opts[{n}+1]}"

	return []string{"--preview", previewCmd}
}

func (i InfoStruct) RenderStruct(opts []Struct) []string {
	if i.info != nil || len(i.info) != 0 {
		return i.Render()
	}
	previewCmd := "opts=("

	previewCmd += "\"" + opts[0].Desc() + "\""
	for _, value := range opts[1:] {
		previewCmd += " \"" + value.Desc() + "\""
	}
	previewCmd += ") && echo ${opts[{n}+1]}"

	return []string{"--preview", previewCmd}
}
