package api

import (
	"net/http"
	"strconv"

	"final_project/pkg/db"
)

type taskResp struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

type tasksResp struct {
	Tasks []taskResp `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	list, err := db.Tasks(50)
	if err != nil {
		writeJSON(w, map[string]any{"error": err.Error()})
		return
	}

	resp := tasksResp{
		Tasks: make([]taskResp, 0, len(list)),
	}

	for _, t := range list {
		resp.Tasks = append(resp.Tasks, taskResp{
			ID:      strconv.FormatInt(t.ID, 10),
			Date:    t.Date,
			Title:   t.Title,
			Comment: t.Comment,
			Repeat:  t.Repeat,
		})
	}

	writeJSON(w, resp)
}
