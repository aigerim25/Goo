package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"practice5/internal/repository"
)

type UserHandler struct {
	repo *repository.UserRepository
}

func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}
func (h *UserHandler) GetPaginatedUsersHandler(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("pageSize")
	orderBy := r.URL.Query().Get("orderBy")

	page := 1
	pageSize := 10

	if pageStr != "" {
		if ps, err := strconv.Atoi(pageStr); err == nil {
			pageSize = ps
		}
	}
	if pageSizeStr != "" {
		if ps, err := strconv.Atoi(pageSizeStr); err == nil {
			pageSize = ps
		}
	}
	filters := map[string]string{}
	if id := r.URL.Query().Get("id"); id != "" {
		filters["id"] = id
	}
	if name := r.URL.Query().Get("name"); name != "" {
		filters["name"] = name
	}
	if email := r.URL.Query().Get("email"); email != "" {
		filters["email"] = email
	}
	if gender := r.URL.Query().Get("gender"); gender != "" {
		filters["gender"] = gender
	}
	if birthDate := r.URL.Query().Get("birth_date"); birthDate != "" {
		filters["birth_date"] = birthDate
	}
	result, err := h.repo.GetPaginatedUsers(page, pageSize, filters, orderBy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
func (h *UserHandler) GetCommonFriendsHandler(w http.ResponseWriter, r *http.Request) {
	user1Str := r.URL.Query().Get("user1")
	user2Str := r.URL.Query().Get("user2")

	user1, err := strconv.Atoi(user1Str)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user2, err := strconv.Atoi(user2Str)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := h.repo.GetCommonFriends(user1, user2)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
