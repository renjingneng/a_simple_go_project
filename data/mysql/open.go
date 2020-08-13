package mysql

type Open struct {
	*Base
}

func NewOpen() *Open {
	return &Open{
		Base: NewBase("MysqlOpen"),
	}
}

func (open *Open) FetchRowNew() {

}
