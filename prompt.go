package gofzf

import (
	"os/exec"
	"strings"

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

	out, err := fzf.Output()
	if err != nil {
		return nil, err
	}
	return strings.Split(strings.TrimSpace(string(out)), "\n"), nil
}

// well, struct has to be convertable to string
type Struct interface {
	ToString()
}

func FzfPromptStruct(options []Struct, modes ...options.Mode) ([]Struct, error) {
	return nil, nil
}
