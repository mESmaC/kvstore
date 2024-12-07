package api

import (
	"encoding/json"
	"fmt"
	"mESmaC/kvstore/db"
	"net/http"
)

func CreateDBHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var requestData struct {
		Name string `json:"name"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil || requestData.Name == "" {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	dbName := requestData.Name

	db.CreateDB(dbName)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Database '%s' created.\n", dbName)
}

func InsertHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var requestData struct {
		Name  string      `json:"name"`
		Key   string      `json:"key"`
		Value interface{} `json:"value"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil ||
		requestData.Name == "" ||
		requestData.Key == "" ||
		requestData.Value == nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	dbName := requestData.Name
	keyStr := requestData.Key
	valFace := requestData.Value

	db.Insert(dbName, keyStr, valFace)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Inserted value '%v' in DB: '%s' at Key: '%s'.\n", valFace, dbName, keyStr)
}
