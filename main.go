package main

import (
	"log"
	"net/http"
	"os"

	"final_project/pkg/api"
	"final_project/pkg/db"
)

func main() {
	// инициализация базы данных
	err := db.Init("scheduler.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.DB.Close()

	// регистрация API обработчиков
	api.Init()

	// директория с фронтендом
	webDir := "./web"

	// файловый сервер
	http.Handle("/", http.FileServer(http.Dir(webDir)))

	// порт (исправление замечания ревьюера)
	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = "7540"
	}

	log.Println("Server started on port:", port)

	// запуск сервера
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
