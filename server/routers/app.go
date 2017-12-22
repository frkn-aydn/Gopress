package routers

import (
	"github.com/kataras/iris"
)

func AppHandler(app iris.Party) {
	app.Get("/", func(ctx iris.Context) {
		ctx.View("index.html")
	})

	app.Get("/demo", func(ctx iris.Context) {
		ctx.View("demo.html")
	})

	app.Get("/register", func(ctx iris.Context) {
		ctx.View("register.html")
	})

	app.Get("/dashboard", func(ctx iris.Context) {
		ctx.View("dashboard.html")
	})
}
