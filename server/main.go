package main

import (
	"Gopress/server/routers"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/sessions/sessiondb/redis"
	"github.com/kataras/iris/sessions/sessiondb/redis/service"
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

	// Redis session added...
	db := redis.New(service.Config{
		Network:     service.DefaultRedisNetwork,
		Addr:        service.DefaultRedisAddr,
		Password:    "",
		Database:    "",
		MaxIdle:     0,
		MaxActive:   0,
		IdleTimeout: service.DefaultRedisIdleTimeout,
		Prefix:      ""})
	iris.RegisterOnInterrupt(func() {
		db.Close()
	})

	// Session config added...
	sess := sessions.New(sessions.Config{Cookie: "sessionscookieid", Expires: 45 * time.Minute})
	sess.UseDatabase(db)

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
	case "maintenance":
		app.Get("/", func(ctx iris.Context) {
			ctx.ViewData("title", title)
			ctx.ViewData("description", description)
			ctx.ViewData("keywords", keywords)

			ctx.View("maintenance.html")
		})
	case "active":
		app.PartyFunc("/", routers.AppHandler)
		app.PartyFunc("/api/", routers.ApiHandler)
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
