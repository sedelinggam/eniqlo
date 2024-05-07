package v1

import (
	staffControllers "eniqlo/internal/delivery/http/v1/controller/staff"
	staffService "eniqlo/internal/service/staff"

	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func Init(app *echo.Echo, db *sqlx.DB, val *validator.Validate) {
	var (
		staffSvc = staffService.New(db)
	)
	v1 := app.Group("/v1")
	staffControllers.Init(v1, val, staffSvc)
}
