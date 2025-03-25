package state

type Object struct {
	Name string
	X    int
	Y    int
}

func NewObject(name string, x, y int) *Object {
	return &Object{
		X: x, Y: y, Name: name,
	}
}

type World struct {
	Name    string
	Objects []*Object
}

func NewWorld(name string) *World {
	return &World{
		Name: name,
	}
}

func (w *World) SetObject(o *Object) {
	w.Objects = append(w.Objects, o)
}
