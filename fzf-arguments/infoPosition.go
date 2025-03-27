package fzfarguments

import "github.com/IvanLogvynenko/go-fzf/interfaces"

type Position string

const (
	Up    Position = "up"
	Down  Position = "down"
	Left  Position = "left"
	Right Position = "right"
)

type InfoPositionStruct struct {
	position Position
}

func InfoPosition(pos Position) InfoPositionStruct {
	return InfoPositionStruct{pos}
}

func (ip InfoPositionStruct) Render() []string {
	return []string{"--preview-window=" + string(ip.position)}
}
func (ip InfoPositionStruct) RenderStruct(_ []interfaces.Struct) []string {
	return ip.Render()
}
