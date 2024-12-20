package main

import (
	"fmt"
	"github.com/mESmaC/kvstore/api"
)

func StartServer() {
	fmt.Println("PONG")
	go func() {
		fmt.Println("+---------------------------+")
		fmt.Println("| Starting Service: KVStore |")
		fmt.Println("+---------------------------+")
		api.StartServer()
	}()
}

func main() {
	fmt.Println("===== KVStore API =====")
	api.StartServer()
}
