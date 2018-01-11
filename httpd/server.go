package httpd

import (
	"net/http"

	"github.com/labstack/echo"
)

func HttpdCT() {
	e := echo.New()
	//e.GET("/socket.io/",echo.WrapHandler(server))
	e.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	go func() {
		e.Start("127.0.0.1:9700")
	}()
}
