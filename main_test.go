package gofzf_test

import (
	"fmt"
	"testing"

	gofzf "github.com/IvanLogvynenko/go-fzf"
)

type Fruit struct {
	Name        string
	Description string
}

func (f Fruit) ToString() string {
	return f.Name
}
func (f Fruit) Desc() string {
	return f.Description
}

func TestFzfPromptStruct(t *testing.T) {
	fruits := []gofzf.Struct{
		Fruit{"Apple", "A sweet red fruit"},
		Fruit{"Banana", "A yellow tropical fruit"},
		Fruit{"Cherry", "A small red fruit"},
	}

	selected, err := gofzf.FzfPromptStruct(fruits, gofzf.Binds(), gofzf.Cycle(), gofzf.Info(nil), gofzf.InfoPosition(gofzf.Up))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Selected:", selected)
}

func TestFzfPrompt(t *testing.T) {
	optionsStr := []string{
		"Apple", "Banana", "Cherry",
	}
	descriptions := []string{
		"A sweet red fruit",
		"A yellow tropical fruit",
		"A small red fruit",
	}

	selected, err := gofzf.FzfPrompt(optionsStr, gofzf.Binds(), gofzf.Cycle(), gofzf.Info(descriptions))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Selected:", selected)
}
