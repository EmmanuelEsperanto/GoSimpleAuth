package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the about page.")
}

func main() {
	const connStr = "postgres://user:password@localhost:5432/dbname"
	db, err := sql.Open("postgres", connStr)
	// Регистрируем обработчики для разных маршрутов
	fmt.Println(db.Stats())
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)

	// Запускаем сервер
	fmt.Println("Starting server at port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
