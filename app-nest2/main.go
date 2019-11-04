package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT must be set")
	}
	// Validate it parses to an int
	_, err := strconv.ParseInt(portString, 0, 64)
	if err != nil {
		log.Fatal("PORT must parse to an int")
	}

	whatToSay := os.Getenv("WHATTOSAY")
	if whatToSay == "" {
		log.Fatal("WHATTOSAY must be set")
	}

	time.Sleep(3 * time.Second)

	r := chi.NewRouter()
	r.Use(middleware.Heartbeat("/health"))
	r.Use(middleware.Logger)
	r.Get("/nest2", nestHandler)
	err = http.ListenAndServe(":"+portString, r)
	if err != nil {
		log.Fatal(err)
	}
}

func nestHandler(w http.ResponseWriter, r *http.Request) {
	whatToSay := os.Getenv("WHATTOSAY")
	if whatToSay == "" {
		log.Fatal("WHATTOSAY must be set")
	}
	log.Println(whatToSay)
	w.Write([]byte(whatToSay))
	w.WriteHeader(200)
}
