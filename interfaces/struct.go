package interfaces

// well, struct has to be convertable to string. Desc can return ""
type Struct interface {
	ToString() string
	Desc() string
}