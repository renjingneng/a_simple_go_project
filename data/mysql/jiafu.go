package mysql

type Jiafu struct {
	*base
}

func NewJiafu() *Jiafu {
	return &Jiafu{
		base: NewBase("LocalJiafu"),
	}
}

func (j *Jiafu) FetchRowNew() {

}
