package infrastructure

import (
	"sample/interfaces/controllers"
	"github.com/labstack/echo"
)

func Init() {
	// Echo instance
	e := echo.New()

		// Routes
		v1 := e.Group("/api/v1")
		{
			v1.GET("/sample", controllers.Sample())
		}

	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}

