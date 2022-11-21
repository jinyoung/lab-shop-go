package customercenter

import (
	"github.com/labstack/echo"
	"net/http"
)

func RouteInit() *echo.Echo {
	e := echo.New()
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "ok")
	})
	myPage := &MyPage{}
	e.GET("/myPages", myPage.Get)
	e.GET("/myPages/:id", myPage.FindById)
	e.POST("/myPages", myPage.Persist)
	e.PUT("/myPages/:id", myPage.Put)
	e.DELETE("/myPages/:id", myPage.Remove)
	return e
}
