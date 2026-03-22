package main

import (
	"log"
	"net/http"

	"final_project/pkg/api"
	"final_project/pkg/db"
)

func main() {
	// инициализация базы данных
	err := db.Init("scheduler.db")
	if err != nil {
		log.Fatal(err)
	}

	// регистрация API обработчиков
	api.Init()

	// директория с фронтендом
	webDir := "./web"

	// файловый сервер
	http.Handle("/", http.FileServer(http.Dir(webDir)))

	log.Println("Server started on port: 7540")

	// запуск сервера
	err = http.ListenAndServe(":7540", nil)
	if err != nil {
		log.Fatal(err)
	}
}
