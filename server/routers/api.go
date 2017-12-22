package routers

import (
	"github.com/kataras/iris"
)

func ApiHandler(api iris.Party) {
	api.Get("/", func(ctx iris.Context) {
		ctx.View("index.html")
	})
}
