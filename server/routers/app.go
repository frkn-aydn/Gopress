package routers

import (
	"fmt"

	"github.com/kataras/iris"
)

var title string = "Muhammed Furkan Aydın"
var description string = "Just a developer from Earth"
var keywords string = "Muhammed Furkan Aydın, front-end developer, back-end developer, javascript, javascript developer, golang developer"
// AppHandler function serving all HTML files with SEO friendly URLs
func AppHandler(app iris.Party) {
	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("title", title)
		ctx.ViewData("description", description)
		ctx.ViewData("keywords", keywords)
		ctx.View("index.html")
	})

	app.Get("/resume", func(ctx iris.Context) {
		ctx.ViewData("title", title + " | Resume")
		ctx.ViewData("description", description)
		ctx.ViewData("keywords", keywords)
		ctx.View("resume.html")
	})

	app.Get("/blog", func(ctx iris.Context) {
		ctx.ViewData("title", title + " | Blog")
		ctx.ViewData("description", description)
		ctx.ViewData("keywords", keywords)
		ctx.View("blog.html")
	})

	app.Get("/blog/{url:string}", func(ctx iris.Context) {
		postURL := ctx.Params().GetDecoded("url")
		fmt.Println(postURL)
		ctx.View("article.html")
	})

	app.Get("/contact", func(ctx iris.Context) {
		ctx.ViewData("title", title + " | Contact")
		ctx.ViewData("description", description)
		ctx.ViewData("keywords", keywords)
		ctx.View("contact.html")
	})

	app.Get("/login", func(ctx iris.Context) {
		ctx.ViewData("title", title + " | Login")
		ctx.ViewData("description", description)
		ctx.ViewData("keywords", keywords)
		ctx.View("login.html")
	})

	app.Get("/register", func(ctx iris.Context) {
		ctx.ViewData("title", title + " | Register")
		ctx.ViewData("description", description)
		ctx.ViewData("keywords", keywords)
		ctx.View("register.html")
	})
}
