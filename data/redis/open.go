package redis

type Open struct {
	Base
}

func NewOpen() *Open {
	return &Open{
		Base: NewBase("RedisOpen", "Single"),
	}
}
