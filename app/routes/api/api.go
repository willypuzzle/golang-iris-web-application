package api

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"iris.dom/app/config/files/environment"
	"iris.dom/app/controllers/http"
	"iris.dom/app/services"
)

type Api struct {
	app *iris.Application
}

func (api *Api) pingSetup() *Api {
	mvc.Configure(api.app.Party("/ping"), func(app *mvc.Application) {
		app.Register(
			environment.GetEnvironment(),
			services.NewPingService,
		)

		app.Handle(new(http.PingController))
	})

	return api
}

func buildApi(app *iris.Application) *Api {
	return &Api{
		app: app,
	}
}

// Setup this function setups the api router group
func Setup(app *iris.Application) *Api {
	api := buildApi(app)
	api.pingSetup()

	return api
}
