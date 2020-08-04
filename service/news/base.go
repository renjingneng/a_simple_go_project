package news

type Base struct {
	Name string
}

func NewBase(name string) *Base {
	a := &Base{Name: name}
	return a
}
func (a *Base) FetchName() string {
	return a.Name
}
