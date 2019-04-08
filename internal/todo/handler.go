package todo

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

type TodoHandler interface {
	CreateTodo(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type todoHandler struct {
	service TodoService
}

func NewTodoHandler(service TodoService) TodoHandler {
	return &todoHandler{
		service,
	}
}
func (h *todoHandler) Get(w http.ResponseWriter, r *http.Request) {
	todos, err := h.service.FindAllTodos()
	if err != nil {
		logrus.WithField("error", err).Error("Unable to find all todos")
		http.Error(w, "Unable to find all todos", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(todos)
	if err != nil {
		logrus.WithField("error", err).Error("Error unmarshalling response")
		http.Error(w, "Unable to get todos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(response); err != nil {
		logrus.WithField("error", err).Error("Error writing response")
	}
}
func (h *todoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var todo Todo
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&todo); err != nil {
		logrus.Error("Unable to decode todo")
		http.Error(w, "Bad format for todo", http.StatusBadRequest)
		return
	}
	if err := h.service.CreateTodo(&todo); err != nil {
		logrus.WithField("error", err).Error("Unable to create todo")
		http.Error(w, "Unable to create todo", http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(todo)
	if err != nil {
		logrus.WithField("error", err).Error("Error unmarshalling response")
		http.Error(w, "Unable to get todo", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	if _, err = w.Write(response); err != nil {
		logrus.WithField("error", err).Error("Error writing response")
	}
}
