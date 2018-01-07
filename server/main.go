package main

import (
	"Gopress/server/routers"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"runtime"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
)

func main() {
	// Maximum proccess number...
	runtime.GOMAXPROCS(4)

	// Config file inserting...
	type ConfigType struct {
		Port   string `json:"port"`
		Status string `json:"status"`
	}

	configFile, err := ioutil.ReadFile("server/config.json")
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}

	var config ConfigType
	json.Unmarshal(configFile, &config)
	fmt.Println("Config file inserted...")

	// Iris started...
	app := iris.New()
	app.Use(logger.New())

	// Static files served with iris...
	app.StaticServe("server/public", "/")

	// Templates registered...
	app.RegisterView(iris.HTML("server/views", ".html"))

	// Getting website variables from env variable...
	var (
		title       string = os.Getenv("GOPRESS_TITLE")
		description string = os.Getenv("GOPRESS_DESCRIPTION")
		keywords    string = os.Getenv("GOPRESS_KEYWORDS")
	)

	// Looking for website status...
	switch config.Status {
	case "coming-soon":
		app.Get("/", func(ctx iris.Context) {
			ctx.ViewData("title", title)
			ctx.ViewData("description", description)
			ctx.ViewData("keywords", keywords)

			ctx.View("coming-soon.html")
		})
		app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
			ctx.View("not-found.html")
		})
	case "maintenance":
		app.Get("/", func(ctx iris.Context) {
			ctx.ViewData("title", title)
			ctx.ViewData("description", description)
			ctx.ViewData("keywords", keywords)

			ctx.View("maintenance.html")
		})
		app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
			ctx.View("not-found.html")
		})
	case "active":
		app.PartyFunc("/", routers.AppHandler)
		app.PartyFunc("/api/", routers.ApiHandler)
		app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
			matched, err := regexp.MatchString("api", ctx.Path())
			if err != nil {
				// [TODO] Do somthin gwhen regex gives an error
				fmt.Println(err)
			}
			if matched {
				type Response struct {
					Success bool   `json:"success"`
					Error   string `json:"error"`
				}
				ctx.JSON(Response{
					Success: false,
					Error:   "Endpoint not found.",
				})
			} else {
				ctx.View("not-found.html")
			}
		})
	default:
		app.Get("/", func(ctx iris.Context) {
			ctx.ViewData("title", title)
			ctx.ViewData("description", description)
			ctx.ViewData("keywords", keywords)

			ctx.View("coming-soon.html")
		})
	}

	// Starting server...
	app.Run(
		iris.Addr("localhost:"+config.Port),
		iris.WithoutVersionChecker,
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations, // enables faster json serialization and more
	)
}
