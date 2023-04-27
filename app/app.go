package app

import (
	"github.com/labstack/echo/v4"
	"github.com/sljeff/oapi-codegen-example/gen/oapi"
)

var _ oapi.ServerInterface = (*App)(nil)

type App struct {
	// ...
}

// health check
// (GET /health)
func (app *App) HealthCheck(ctx echo.Context, params oapi.HealthCheckParams) error {
	return nil
}

// (POST /notes)
func (app *App) NotesPost(ctx echo.Context) error {
	body := oapi.Notes{}
	if err := ctx.Bind(&body); err != nil {
		return err
	}

	return ctx.JSON(200, oapi.Notes{
		Objects: &[]oapi.Note{
			{},
		},
	})
}
