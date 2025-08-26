package ddd

type IDer interface {
	ID() string
}

type Entity struct {
	id   string
	name string
}
