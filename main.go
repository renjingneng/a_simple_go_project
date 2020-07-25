package main

import (
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	testInt := 123123

	tmpl := iris.HTML("./templates", ".html")
	tmpl.Layout("layouts/layout.html")
	tmpl.AddFunc("greet", func(s string) string {
		return "Greetings " + s + "!"
	})

	app.RegisterView(tmpl)

	app.Get("/", func(ctx iris.Context) {
		if err := ctx.View("page1.html"); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)

			ctx.Writef(err.Error())
		}
	})

	// remove the layout for a specific route
	app.Get("/nolayout", func(ctx iris.Context) {
		ctx.ViewLayout(iris.NoLayout)
		if err := ctx.View("page1.html"); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Writef(err.Error())
		}
	})

	// set a layout for a party, .Layout should be BEFORE any Get or other Handle party's method
	my := app.Party("/my").Layout("layouts/mylayout.html")
	{ // both of these will use the layouts/mylayout.html as their layout.
		my.Get("/", func(ctx iris.Context) {
			ctx.View("page1.html")
		})
		my.Get("/other", func(ctx iris.Context) {
			firstname := ctx.URLParamDefault("var1", "var1")
			//lastname := ctx.URLParamDefault("var2", "var2")
			is_ren := firstname == "ren"
			is_zeng := firstname == "zeng"
			maptest := []map[string]string{{
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
			}}
			ctx.ViewData("int", testInt)
			ctx.ViewData("firstname", firstname)
			ctx.ViewData("is_ren", is_ren)
			ctx.ViewData("is_zeng", is_zeng)
			ctx.ViewData("maptest", maptest)

			ctx.View("page1.html")
		})
	}

	// http://localhost:8080
	// http://localhost:8080/nolayout
	// http://localhost:8080/my
	// http://localhost:8080/my/other
	app.Listen(":8080")
}
