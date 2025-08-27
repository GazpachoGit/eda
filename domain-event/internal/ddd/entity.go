package ddd

type IDer interface {
	ID() string
}

type Entity struct {
	id   string
	name string
}

func NewEntity(id, name string) Entity {
	return Entity{
		id:   id,
		name: name,
	}
}

func (e Entity) ID() string             { return e.id }
func (e Entity) EntityName() string     { return e.name }
func (e Entity) Equals(other IDer) bool { return e.id == other.ID() }

func (e *Entity) setID(id string)     { e.id = id }
func (e *Entity) setName(name string) { e.name = name }
