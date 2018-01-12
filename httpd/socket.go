package httpd


import(
	"github.com/googollee/go-socket.io"
	"log"
)

func NewSocketServer() *socketio.Server {
	//create a connect

	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Println(err)
	}
	return server
}

