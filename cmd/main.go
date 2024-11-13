package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/KidiXDev/GoSimpleApi/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("Starting the server")
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	content, err := os.ReadFile("ascii-art.txt")
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	var url string = "localhost:8000"
	fmt.Println(string(content))
	fmt.Println("listening on " + url + "\n")

	err = http.ListenAndServe(url, r)
	if err != nil {
		log.Error(err)
	}

}
