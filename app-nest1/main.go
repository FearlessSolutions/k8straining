package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

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

	addEndpoint := os.Getenv("NEST2ENDPOINT")
	if addEndpoint == "" {
		log.Fatal("NEST2ENDPOINT must be set")
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/nest1", nestHandler)
	r.Get("/health", healthHandler)
	err = http.ListenAndServe(":"+portString, r)
	if err != nil {
		log.Fatal(err)
	}
}

//respond with 200 to health checks
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}

func nestHandler(w http.ResponseWriter, r *http.Request) {
	endpoint := os.Getenv("NEST2ENDPOINT")
	if endpoint == "" {
		log.Println("ERROR: NEST2ENDPOINT must be set")
		http.Error(w, "env var not set", 500)
	}
	resp, err := http.Get(endpoint)
	if err != nil {
		log.Printf("Error calling Nest2 Endopoint: %v", err.Error())
		http.Error(w, err.Error(), 500)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response: %v", err.Error())
		http.Error(w, err.Error(), 500)
	}
	log.Println("Response: ", string(body))
	w.Write(body)
	w.WriteHeader(200)
}
