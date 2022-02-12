package main

import (
	"avah/config"
	"avah/handler"
	"net/http"
)

func main() {
	config.Init()

	handler.HandleRequests()

	err := http.ListenAndServe(":10000", nil)
	if err != nil {
		panic(err)
	}
}
