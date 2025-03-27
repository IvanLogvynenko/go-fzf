package gofzf

import (
	"fmt"
	"os/exec"
	"slices"
	"strings"

	"github.com/IvanLogvynenko/go-fzf/interfaces"
	"github.com/IvanLogvynenko/go-fzf/options"
)

func FzfPrompt(options []string, modes ...options.Mode) ([]string, error) {
	input := strings.Join(options, "\n")
	fzf_cmd := "fzf"
	modesStr := make([]string, 0)
	for _, mode := range modes {
		modesStr = append(modesStr, mode.Render()...)
	}

	fzf := exec.Command(fzf_cmd, modesStr...)
	fzf.Stdin = strings.NewReader(input)
	fmt.Println(fzf.String())

	out, err := fzf.Output()
	if err != nil {
		return nil, err
	}
	return strings.Split(strings.TrimSpace(string(out)), "\n"), nil
}

func FzfPromptStruct(options []interfaces.Struct, modes ...options.Mode) ([]int, error) {
	input := options[0].ToString()
	for _, opt := range options[1:] {
		input += "\n" + opt.ToString()
	}
	fzf_cmd := "fzf"
	modesStr := make([]string, 0)
	for _, mode := range modes {
		modesStr = append(modesStr, mode.RenderStruct(options)...)
	}

	fzf := exec.Command(fzf_cmd, modesStr...)
	fzf.Stdin = strings.NewReader(input)
	fmt.Println(fzf.String())

	out, err := fzf.Output()
	if err != nil {
		return nil, err
	}

	resultStr := strings.Split(strings.TrimSpace(string(out)), "\n")
	result := make([]int, 0)
	for _, res := range resultStr {
		for id, v := range options {
			if v.ToString() == res {
				result = append(result, id)
			}
		}
	}
	slices.Sort(result)
	return result, nil
}
