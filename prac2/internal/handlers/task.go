package handlers

import (
	//"fmt"
	"encoding/json"
	"net/http"
	"strconv"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var tasks = []Task{
	{ID: 1, Title: "Write unit tests", Done: false},
	{ID: 2, Title: "Deploy service", Done: true},
}
var nextID = 3

func TasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet {
		handleGet(w, r)
		return
	}
	if r.Method == http.MethodPost {
		handlePost(w, r)
		return
	}
	if r.Method == http.MethodPatch {
		handlePatch(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
func handleGet(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	doneStr := r.URL.Query().Get("done")
	if idStr == "" {
		if doneStr == "" {
			json.NewEncoder(w).Encode(tasks)
			return
		}
		if doneStr != "true" && doneStr != "false" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "invalid done"})
			return
		}
		wDone := (doneStr == "true")
		filtered := []Task{}
		for _, t := range tasks {
			if t.Done == wDone {
				filtered = append(filtered, t)
			}
		}
		json.NewEncoder(w).Encode(filtered)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid id"})
		return
	}
	for _, task := range tasks {
		if task.ID == id {
			json.NewEncoder(w).Encode(task)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "task not found"})
}
func handlePost(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string `json:"title"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil || input.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid title"})
		return
	}
	newTask := Task{
		ID:    nextID,
		Title: input.Title,
		Done:  false,
	}
	nextID++
	tasks = append(tasks, newTask)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}
func handlePatch(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid id"})
		return
	}
	var input struct {
		Done *bool `json:"done"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil || input.Done == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid done"})
		return
	}
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Done = *input.Done
			json.NewEncoder(w).Encode(map[string]bool{"updated": true})
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "task not found"})
}
