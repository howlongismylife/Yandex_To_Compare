package api

import (
	"net/http"
	"time"

	"final_project/pkg/db"
)

func doneTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": "id is required"})
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		if err.Error() == "task not found" {
			writeJSON(w, http.StatusNotFound, map[string]any{"error": "task not found"})
			return
		}

		writeJSON(w, http.StatusInternalServerError, map[string]any{"error": err.Error()})
		return
	}

	// одноразовая задача
	if task.Repeat == "" {
		if err := db.DeleteTask(id); err != nil {
			if err.Error() == "task not found" {
				writeJSON(w, http.StatusNotFound, map[string]any{"error": "task not found"})
				return
			}

			writeJSON(w, http.StatusInternalServerError, map[string]any{"error": err.Error()})
			return
		}

		writeJSON(w, http.StatusOK, map[string]any{})
		return
	}

	// периодическая задача
	next, err := NextDate(time.Now(), task.Date, task.Repeat)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		return
	}

	if err := db.UpdateDate(next, id); err != nil {
		if err.Error() == "task not found" {
			writeJSON(w, http.StatusNotFound, map[string]any{"error": "task not found"})
			return
		}

		writeJSON(w, http.StatusInternalServerError, map[string]any{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{})
}

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": "id is required"})
		return
	}

	if err := db.DeleteTask(id); err != nil {
		if err.Error() == "task not found" {
			writeJSON(w, http.StatusNotFound, map[string]any{"error": "task not found"})
			return
		}

		writeJSON(w, http.StatusInternalServerError, map[string]any{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{})
}
