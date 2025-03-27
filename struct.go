package gofzf

// well, struct has to be convertable to string. Desc is used only for Info. if you are not using Info you can return ""
type Struct interface {
	ToString() string
	Desc() string
}
