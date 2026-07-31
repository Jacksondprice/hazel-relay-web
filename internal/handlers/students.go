package handlers

import (
	"encoding/json"
	"net/http"
)

type Student struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func ListStudents(w http.ResponseWriter, _ *http.Request) {
	items := []Student{{ID: 1, Name: "sample"}}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(items)
}
