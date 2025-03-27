# go-fzf

This is a wrapper for an existing `fzf` cli written in GOlang to make cli applications access power of fuzzy finding.

There are only two functions. 
"""go
	FzfPrompt(options []string, modes ...fzfarguments.Modes) ([]string, error)
	FzfPromptStruct(options []Struct, modes ...fzfarguments.Modes) ([]Struct, error)
"""

first one can be considered a case of the second. Second only requieres you to pass a struct with defined method `ToString` and `Desc`(used only if Info is passed, so can return "" if info is not used). Functions return a list of matched options. In default case (no options passed) that list will have only one entry. If arguments like `fzfarguments.Multiselect` or `fzfarguments.Pattern` are passed, fzf was able to match several entries. The top one is guaranteed to be the first(from the `options []string`) and the best (fzf sorts result) match of all. 

Options:
- Pattern: Provide a fzf with a pattern to search for. Interactive UI will not b loaded. Userful to omit typos during argument parsing
- Height: Self-explanatory. Nice case to use is when you don't want fzf to take full terminal. Best to be paired with `Reverse`
- Reverse: Self-explanatory. (If not just try it and you be shameful)
- Header: Sets a header (I didn't want to write Self-explanatory again. Wops I did it again). Use to display a message
- Prompt: Used to change a prompt in the input line (default: >)
- Multiselect: Guess what? And you can limit maximal selection
- Cycle: Cycling through given options (default: when you reach the end you won't be sent to the first again)
- Info: You can provide some description for each option and it will be displayed when the option is selected. Lenght and indexes must match. If no description is needed pass an empty string. You can also store description in your structs if you are using FzfPromptStruct. If so, override `Desc` in your struct and make it return your description
- InfoPosition: [up/down/right(default)/left] Choose where to display your info window
- Binds: (default: tab:down,shift-tab:up,ctrl-s:select,ctrl-a:select-all,ctrl-d:deselect) They are mine. You can provide your own that will override that
- Style: [default/full/minimal]

Default modes: Binds

Feel free to Pull Request
