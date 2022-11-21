package delivery

import (
	"github.com/labstack/echo"
	"net/http"
)

func RouteInit() *echo.Echo {
	e := echo.New()
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "ok")
	})
	delivery := &Delivery{}
	e.GET("/deliveries", delivery.Get)
	e.GET("/deliveries/:id", delivery.FindById)
	e.POST("/deliveries", delivery.Persist)
	e.PUT("/deliveries/:id", delivery.Put)
	e.DELETE("/deliveries/:id", delivery.Remove)
	return e
}
