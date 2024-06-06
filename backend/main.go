package main

import (
	"go-vue-docker/api"
	"go-vue-docker/chatbot"
	"go-vue-docker/config"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()

	go api.ServerGo()
	// Start the chat bot in a separate Goroutine
	go chatbot.StartBot(cfg)

	// Serve static files
	fs := http.FileServer(http.Dir("../frontend/dist"))
	http.Handle("/", fs)

	log.Println("Server started on :8083")

	err := http.ListenAndServe(":8083", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
