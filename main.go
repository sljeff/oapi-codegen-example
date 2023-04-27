package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	oapimw "github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	"github.com/shanbay/gobay/echo/swagger"
	"github.com/sljeff/oapi-codegen-example/gen/oapi"
	"github.com/sljeff/oapi-codegen-example/app"
)

func serve(app *app.App) error {
	// init config / ext

	e := echo.New()

	oapi.RegisterHandlersWithBaseURL(e, app, "/v1")

	// middlewares...
	enableDoc := true
	swg, err := oapi.GetSwagger()
	if err != nil {
		return err
	}
	if enableDoc {
		if swgJSON, err := swg.MarshalJSON(); err == nil {
			// opts := []swagger.Opts{swagger.SetSwaggerAuthorizer(isStaff), swagger.SetSwaggerIsHTTPS(true)}
			e.Pre(swagger.SwaggerDoc("/v1", swgJSON))
		} else {
			return err
		}
	}

	e.Use(oapimw.OapiRequestValidatorWithOptions(swg, &oapimw.Options{}))

	// e.ErrorHandler = func(err error, ctx echo.Context) {

	// start
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	go func(c <-chan os.Signal) {
		<-c
		log.Println("Shutting down...")
		if err := e.Shutdown(context.Background()); err != nil {
			log.Fatal(err)
		}
	}(stopChan)

	err = e.Start(":8080")
	return err
}

func main() {
	if err := serve(&app.App{}); err != nil {
		panic(err)
	}
}
