package server

import (
	"os"
	"strconv"

	"github.com/gorilla/mux"
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

	return server

}

func configServer() (*mux.Router, error) {
	r := mux.NewRouter()

}
