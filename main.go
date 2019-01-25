package main

import (
	"os"
	_ "golang_login/config"
	R "golang_login/routes"

	"github.com/kataras/iris"
)

func main() {
	app := iris.Default()

	v1 := app.Party("/v1")

	v1.Post("/login", R.APILogin)

	app.Run(iris.Addr(os.Getenv("PORT")))
}
