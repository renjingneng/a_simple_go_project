package main

import (
	"github.com/kataras/iris/v12"
	"github.com/renjingneng/a_simple_go_project/lib/utility"
)

func main() {
	app := iris.New()

	booksAPI := app.Party("/books")
	{

		// GET: http://localhost:8080/books
		booksAPI.Get("/", list)
	}

	app.Listen(":8080")
}

// Book example.
type Book struct {
	Title string `json:"title"`
}

func list(ctx iris.Context) {
	books := []Book{
		{"Mastering Concurrency in Go"},
		{"Go Design Patterns"},
		{"Black Hat Go"},
	}

	utility.ApiFormatData(10005, books, ctx)
	//ctx.JSON(books)
	// TIP: negotiate the response between server's prioritizes
	// and client's requirements, instead of ctx.JSON:
	// ctx.Negotiation().JSON().MsgPack().Protobuf()
	// ctx.Negotiate(books)
}
