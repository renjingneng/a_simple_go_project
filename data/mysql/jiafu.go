package mysql

type Jiafu struct {
	*Base
}

func NewJiafu() *Jiafu {
	return &Jiafu{
		Base: NewBase("MysqlJiafu"),
	}
}

func (j *Jiafu) FetchRowNew() {

}
