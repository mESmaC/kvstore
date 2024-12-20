package api

import (
	"encoding/json"
	"fmt"
	"github.com/mESmaC/kvstore/db"
	"net/http"
)

var requestData struct {
	Name  string      `json:"name"`
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

func CreateDBHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed. HINT: POST", http.StatusMethodNotAllowed)
		return
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
		http.Error(w, "Method not allowed. HINT: POST", http.StatusMethodNotAllowed)
		return
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

func QueryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed. HINT: GET", http.StatusMethodNotAllowed)
		return
	}

	dbName := r.URL.Query().Get("name")
	keyStr := r.URL.Query().Get("key")

	if dbName == "" || keyStr == "" {
		http.Error(w, "Missing 'name' or 'key' query parameter", http.StatusBadRequest)
		return
	}

	value, exists := db.Query(dbName, keyStr)
	if !exists {
		http.Error(w, fmt.Sprintf("Key '%s' not found in database '%s'", keyStr, dbName), http.StatusNotFound)
		return
	}

	response := map[string]interface{}{
		"database": dbName,
		"key":      keyStr,
		"value":    value,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func DeleteKeyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed. HINT: POST", http.StatusMethodNotAllowed)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil ||
		requestData.Name == "" ||
		requestData.Key == "" {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	dbName := requestData.Name
	keyStr := requestData.Key

	db.DeleteKeyFromDB(dbName, keyStr)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Deleted key: '%s' in DB: '%s'.\n", keyStr, dbName)
}

func DeleteDBHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed. HINT: POST", http.StatusMethodNotAllowed)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil ||
		requestData.Name == "" ||
		requestData.Key == "" {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	dbName := requestData.Name
	keyStr := requestData.Key

	db.DeleteKeyFromDB(dbName, keyStr)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Deleted Key and value for: '%s' in DB: '%s'.\n", keyStr, dbName)
}

func DeleteValueAtKeyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed. HINT: POST", http.StatusMethodNotAllowed)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil || requestData.Name == "" || requestData.Key == "" {
		http.Error(w, "Invalid JSON payload. Required fields: 'name', 'key'.", http.StatusBadRequest)
		return
	}

	dbName := requestData.Name
	keyStr := requestData.Key

	db.DeleteValueAtKey(dbName, keyStr)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Cleared value at Key '%s' in Database '%s'.\n", keyStr, dbName)
}

func CreateValueAtEmptyKeyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed. HINT: POST", http.StatusMethodNotAllowed)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil || requestData.Name == "" || requestData.Key == "" || requestData.Value == nil {
		http.Error(w, "Invalid JSON payload. Required fields: 'name', 'key', 'value'.", http.StatusBadRequest)
		return
	}

	dbName := requestData.Name
	keyStr := requestData.Key
	value := requestData.Value

	db.CreateValueAtEmptyKey(dbName, keyStr, value)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Created value '%v' at Key '%s' in Database '%s'.\n", value, keyStr, dbName)
}
