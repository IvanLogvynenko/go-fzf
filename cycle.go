package gofzf

// Yay! empty struct
type CycleStruct struct{}

func Cycle() CycleStruct {
	return CycleStruct{}
}

func (_ CycleStruct) Render() []string {
	return []string{"--cycle"}
}

func (c CycleStruct) RenderStruct(opts []Struct) []string {
	return c.Render()
}
