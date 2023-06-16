package main

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"iris.dom/app/config/files/host"
	"iris.dom/app/config/files/log"
	"iris.dom/app/routes/api"
)

func main() {
	app := iris.New()
	api.Setup(app)

	app.Listen(fmt.Sprintf(":%d", host.GetAppPort()), iris.WithLogLevel(log.GetLogLevel().String()))
}

/*
func main(){
	app := iris.New()
	mvc.Configure(app.Party("/ping"))
}
*/
/*
func main() {
	app := iris.Default()
	app.Use(MyMiddleware)

	app.Handle("GET", "/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "pong"})
	})

	app.Listen(":8000")
}

// MyMiddleware middleware function
func MyMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}
*/
