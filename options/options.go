package options

type Mode interface {
	// well... naming being easy)
	Render() []string
}
