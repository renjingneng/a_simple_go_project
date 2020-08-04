package redis

type Jiafu struct {
	*Base
}

func NewJiafu() *Jiafu {
	return &Jiafu{
		Base: NewBase("RedisJiafu", "Single"),
	}
}

func (j *Jiafu) FetchRowNew() {

}
