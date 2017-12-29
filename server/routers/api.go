package routers

import (
	"Gopress/server/jsonWebToken"
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
)

func ApiHandler(api iris.Party) {
	api.Get("/make", func(ctx iris.Context) {
		token, err := jsonWebToken.Make()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(token)
		ctx.View("index.html")
	})
	api.Get("/parse", func(ctx iris.Context) {
		token, err := jsonWebToken.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoib3RpYWkxMCIsImFnZSI6MzB9.Rs-LpJmqeg8dvj7ft4K1FS7y73kd2BcN4NmsEap31yU")
		if err != nil {
			fmt.Println(err)
		}
		claim, ok := token.(jwt.MapClaims)
		if !ok {
			fmt.Println(err)
		}
		fmt.Println(claim)
		ctx.View("index.html")
	})
}
