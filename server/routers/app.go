package routers

import (
	"fmt"

	"github.com/kataras/iris"
)

// AppHandler function serving all HTML files with SEO friendly URLs
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

	app.Get("/blog/{url:string}", func(ctx iris.Context) {
		postURL := ctx.Params().GetDecoded("url")
		fmt.Println(postURL)
		ctx.View("article.html")
	})

	app.Get("/contact", func(ctx iris.Context) {
		ctx.View("contact.html")
	})

	app.Get("/login", func(ctx iris.Context) {
		ctx.View("login.html")
	})

	app.Get("/register", func(ctx iris.Context) {
		ctx.View("register.html")
	})
}
