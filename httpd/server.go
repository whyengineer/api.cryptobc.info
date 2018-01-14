package httpd

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/whyengineer/api.cryptobc.info/caculate"
)

var CalRes *caculate.Cal

func HttpdCT(c *caculate.Cal) {
	CalRes = c
	e := echo.New()
	sio := NewSocketServer()
	e.GET("/socket.io/", echo.WrapHandler(sio))
	e.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/getsecdata", GetSencodeData)
	e.GET("/getstastatus", GetStaStatus)
	e.GET("/getsstadata", GetStaData)
	//e.GET("/getstadatas", GetStaDatas)
	go func() {
		e.Start("0.0.0.0:9700")
	}()
}
