package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type todo struct {
	Id     int    `json:"id"`
	Value  string `json:"value"`
	Status bool   `json:"status"`
}

var todos = []todo{
	{
		Id:     1,
		Value:  "Belajar Fisika",
		Status: false,
	},
}

func main() {
	http.HandleFunc("/todos", methodHanlder)
	// http.HandleFunc("/todos", updateTodos)
	// http.HandleFunc("/todos", deleteTodos)

	fmt.Println("Server Running On PORT 9000")
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func methodHanlder(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTodos(w,r)
	case http.MethodPost:
		addTodos(w,r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(todos)

	if err != nil {
		http.Error(w, "Failed to encode todos", http.StatusInternalServerError)
		return
	}
}

func addTodos(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var newTodos todo

	err := json.NewDecoder(r.Body).Decode(&newTodos)
	if err != nil {
		http.Error(w, "Invalid Body Request", http.StatusBadRequest)
		return
	}

	if len(todos) > 0 {
		newTodos.Id = len(todos) + 1
	} else {
		newTodos.Id = 1
	}

	newTodos.Status = false

	todos = append(todos, newTodos)
	json.NewEncoder(w).Encode("Berhasil Menambahkan Todo")
}
