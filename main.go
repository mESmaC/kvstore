package main

import (
	"fmt"
	"mESmaC/kvstore/api"
)

func main() {
	fmt.Println("===== KVStore API =====")
	api.StartServer()
}
