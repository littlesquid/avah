package handler

import (
	webhook "avah/webhook/linebot"
	"net/http"
)

func HandleRequests() {
	http.HandleFunc("/reply", webhook.Reply)
	http.HandleFunc("/health", healthCheck)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))
	return
}
