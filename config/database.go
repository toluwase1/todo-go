package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Database() *sql.DB {
	host := "localhost"
	user := "root"
	password := "toluwase"
	credentials := fmt.Sprintf("%s:%s@(%s:3306)/?charset=utf8&parseTime=True", user, password, host)
	log.Println(credentials)
	database, err := sql.Open("mysql", credentials)
	if err != nil {
		panic(err)
	} else {
		log.Println("mysql connection was successful")
	}

	//note: in order not to get an error while running your app a second time
	// execute this query on line 26 instead: `CREATE DATABASE IF NOT EXISTS todoapp;` 
	_, err2 := database.Exec(`CREATE DATABASE todoapp`)
	if err2 != nil {
		log.Println("an error occured when trying to create your database", err2)
		panic(err)
	}

	_, err3 := database.Exec(`USE todoapp`)

	if err3 != nil {
		log.Println(err)
	}

	database.Exec(`
	CREATE TABLE todos (
		id INT AUTO_INCREMENT,
		item TEXT NOT NULL,
		completed BOOLEAN DEFAULT FALSE,
		PRIMARY KEY (id)
	);
	`)

	return database
}
