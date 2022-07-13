package server

import (
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jcmartins81/iesesEstoque/server/routes"
)

type Server struct {
	port   int
	router *mux.Router
}

func InitServer() Server {
	var server Server
	var err error

	server.port, err = strconv.Atoi(os.Getenv("PORT_API"))
	if err != nil {
		server.port = 9000
	}
	router := mux.NewRouter()
	server.router = routes.ConfigRouter(router)

	return server
}
