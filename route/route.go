package route

import (
	"sample/handler"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	// Routes
	v1 := e.Group("/api/v1")
	{
		v1.GET("/handler", handler.Sample())
	}
	return e
}

