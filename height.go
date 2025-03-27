package gofzf

import "fmt"

type HeightStruct struct {
	height int
}

// Returns num is between low and high, or the nearest border
func border(num, low, high int) int {
	temp := max(low, high)
	low = min(low, high)
	high = temp

	if num < low {
		return low
	} else if num > high {
		return high
	} else {
		return num
	}
}

// height: persent of the screen
func Height(height int) HeightStruct {
	return HeightStruct{border(height, 0, 100)}
}

func (h HeightStruct) Render() []string {
	return []string{"--height", fmt.Sprint(h.height)}
}

func (h HeightStruct) RenderStruct(_ []Struct) []string {
	return h.Render()
}
