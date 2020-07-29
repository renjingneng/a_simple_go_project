package news

type base struct {
	Name string
}

func NewBase(name string) *base {
	a := &base{Name: name}
	return a
}
func (a *base) FetchBase(length int) interface{} {
	return length
}
