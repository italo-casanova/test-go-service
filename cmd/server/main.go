package main

import (
	"fmt"
	"log"
	"net/http"

	"hello-service/internal/application"
	"hello-service/internal/infrastructure/memory"
	httpInterface "hello-service/internal/interface/http"
)

func main() {
	messageRepo := memory.NewMessageRepository()

	helloUsecase := application.NewHelloUsecase(messageRepo)

	handler := httpInterface.NewHandler(helloUsecase)

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", handler.Hello)

	addr := ":8080"
	fmt.Printf("🚀 Server running at http://localhost%s/hello\n", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
