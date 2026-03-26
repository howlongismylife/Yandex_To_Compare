package api

import "net/http"

func Init(pass string) {
	http.HandleFunc("/api/nextdate", nextDayHandler)

	http.HandleFunc("/api/signin", signinHandler(pass))

	http.HandleFunc("/api/task", auth(pass, taskHandler))
	http.HandleFunc("/api/task/done", auth(pass, doneTaskHandler))
	http.HandleFunc("/api/tasks", auth(pass, tasksHandler))
}
