package news

// Request example incoming request.
type Request struct {
	Name string `json:"name" url:"name"`
}

type Author struct {
	Id string `url:"id"`
}
