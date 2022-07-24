package controllers

import (
	"html/template"
	"log"
	"net/http"
	"todo-webapp/config"
	"todo-webapp/models"

	"github.com/gorilla/mux"
)

var (
	id        int
	item      string
	completed int
	view      = template.Must(template.ParseFiles("./views/index.html"))
	database  = config.Database()
)

func CreateTodo(resp http.ResponseWriter, req *http.Request) {
	item := req.FormValue("item")
	_, err := database.Exec(`INSERT INTO todos (item) VALUES (?)`, item)
	if err != nil {
		log.Println("an error occurred while inserting the item intp the table")
	}

	http.Redirect(resp, req, "/", 301)

}

func Delete(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	_, err := database.Exec(`DELETE FROM todos WHERE id = ?`, id)
	if err != nil {
		log.Println("an error occurred while inserting the item intp the table")
	}
	http.Redirect(resp, req, "/", 301)
}

func CompleteTodo(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	_, err := database.Exec(`UPDATE todos SET completed = 1 WHERE id = ?`, id)
	if err != nil {
		log.Println("an error occurred while inserting the item intp the table")
	}
	http.Redirect(resp, req, "/", 301)
}

func List(resp http.ResponseWriter, req *http.Request) {
	statement, err := database.Query("SELECT * FROM todos")

	if err != nil {
		log.Println("could not fetch todos: ", err)
	}

	todos := []models.Todo{}

	for statement.Next() {
		err := statement.Scan(&id, &item, &completed)
		if err != nil {
			log.Println("no more items on the list: ", err)
		}
		todo := models.Todo{
			Id:        id,
			Completed: completed,
			Item:      item,
		}
		todo.Id = id
		todo.Completed = completed
		todo.Item = item

		todos = append(todos, todo)

	}
	data := models.View{
		Todos: todos,
	}

	view.Execute(resp, data)
}
