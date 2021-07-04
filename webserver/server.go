package webserver

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jtieri/Nostalgia/webserver/router"
)

type WebServer struct {
	Router *gin.Engine
}

func New() *WebServer {
	r := router.SetupRouter()

	return &WebServer{
		Router: r,
	}
}

func (server *WebServer) Start(address string, port int) (err error) {
	err = server.Router.Run(fmt.Sprintf("%s:%d", address, port))
	if err != nil {
		return err
	}
	return nil
}
