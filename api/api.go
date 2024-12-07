package api

import (
	"fmt"
	"net/http"
)

func StartServer() {

	http.HandleFunc("/create_db", CreateDBHandler)
	http.HandleFunc("/insert_pair", InsertHandler)

	fmt.Println("Server is started on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server", err)
	}
}
