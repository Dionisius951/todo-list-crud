package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
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

	fmt.Println("Server Running On PORT 9000")
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func methodHanlder(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTodos(w, r)
	case http.MethodPost:
		addTodos(w, r)
	case http.MethodPut:
		updateTodos(w, r)
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

func updateTodos(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var id = r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
	}
	Id, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "ID is required", http.StatusBadRequest)
	}

	var UdpTodos todo
	err = json.NewDecoder(r.Body).Decode(&UdpTodos)
	if err != nil {
		http.Error(w, "Invalid Body Request", http.StatusBadRequest)
		return
	}

	for index, item := range todos {
		if item.Id == Id {
			if UdpTodos.Value != "" {
				todos[index].Value = UdpTodos.Value
			}
			if todos[index].Status != UdpTodos.Status {
				todos[index].Status = UdpTodos.Status
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode("Berhasil Memperbarui Todo")
			return
		}
	}
}
