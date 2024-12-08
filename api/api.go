package api

import (
	"fmt"
	"net/http"
)

func StartServer() {

	http.HandleFunc("/create_db", CreateDBHandler)
	http.HandleFunc("/insert_pair", InsertHandler)
	http.HandleFunc("/query_db", QueryHandler)
	http.HandleFunc("/delete_key", DeleteKeyHandler)
	http.HandleFunc("/delete_db", DeleteDBHandler)
	http.HandleFunc("/delete_value", DeleteValueAtKeyHandler)
	http.HandleFunc("/insert_value", CreateValueAtEmptyKeyHandler)

	fmt.Println("Server is started on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server", err)
	}
}
