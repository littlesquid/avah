package main

import (
	"net/http"

	"avah/config"
	"avah/handler"
)

func main() {
	config.Init()

	handler.HandleRequests()

	err := http.ListenAndServe(":10000", nil)
	if err != nil {
		panic(err)
	}
}
