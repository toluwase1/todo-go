package routes

import (
	"todo-webapp/controllers"

	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	route := mux.NewRouter()

	route.HandleFunc("/", controllers.List)
	route.HandleFunc("/create", controllers.CreateTodo).Methods("POST")
	route.HandleFunc("/delete/{id}", controllers.Delete)
	route.HandleFunc("/update/{id}", controllers.CompleteTodo)

	return route
}
