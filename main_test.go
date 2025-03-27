package gofzf_test

import (
	"fmt"
	"testing"

	gofzf "github.com/IvanLogvynenko/go-fzf"
	"github.com/IvanLogvynenko/go-fzf/interfaces"
	"github.com/IvanLogvynenko/go-fzf/options"
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
	fruits := []interfaces.Struct{
		Fruit{"Apple", "A sweet red fruit"},
		Fruit{"Banana", "A yellow tropical fruit"},
		Fruit{"Cherry", "A small red fruit"},
	}

	selected, err := gofzf.FzfPromptStruct(fruits, options.Binds(), options.Cycle(), options.Info(nil))
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

	selected, err := gofzf.FzfPrompt(optionsStr, options.Binds(), options.Cycle(), options.Info(descriptions))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Selected:", selected)
}
