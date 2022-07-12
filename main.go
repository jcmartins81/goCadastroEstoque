package main

import (
	"fmt"
	"log"

	"github.com/jcmartins81/iesesEstoque/server"

	"github.com/joho/godotenv"
)

func init() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	server.InitServer()
	fmt.Println("Main rodando")
}
