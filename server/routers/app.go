package routers

import (
	"github.com/kataras/iris"
)

func AppHandler(app iris.Party) {
	app.Get("/", func(ctx iris.Context) {
		ctx.View("index.html")
	})

	app.Get("/resume", func(ctx iris.Context) {
		ctx.View("resume.html")
	})

	app.Get("/blog", func(ctx iris.Context) {
		ctx.View("blog.html")
	})

	app.Get("/contact", func(ctx iris.Context) {
		ctx.View("contact.html")
	})

	app.Get("/hire", func(ctx iris.Context) {
		ctx.View("hire.html")
	})
}
