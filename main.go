package main

import (
	"log"
	"net/http"
	"todo-webapp/routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	port := "8080"

	err2 := http.ListenAndServe(":"+port, routes.Init())
	if err2 != nil {
		log.Fatal(err)
	}
}
