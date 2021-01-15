package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.Default()

	app.Get("/", func(ctx iris.Context) {
		message := "message"
		nick := "anonymous"

		ctx.JSON(iris.Map{
			"status":  "geted",
			"message": message,
			"nick":    nick,
		})
	})
	app.Listen(":8081")
}