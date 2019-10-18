package main

import (
	"fmt"
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
		log.Fatal("listening PORT must be set")
	}
	// Validate it parses to an int
	_, err := strconv.ParseInt(portString, 0, 64)
	if err != nil {
		log.Fatal("PORT must parse to an int")
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/", additionHandler)
	err = http.ListenAndServe(":"+portString, r)
	if err != nil {
		log.Fatal(err)
	}
}

//HTTP handler adding two integear parameters from url routing parameters
func additionHandler(w http.ResponseWriter, r *http.Request) {
	aString := r.FormValue("a")
	bString := r.FormValue("b")
	log.Println(aString, bString)
	a, err := strconv.ParseInt(aString, 0, 64)
	if err != nil {
		http.Error(w, "a must be an int", 500)
		return
	}
	b, err := strconv.ParseInt(bString, 0, 64)
	if err != nil {
		http.Error(w, "b must be an int", 500)
		return
	}
	c := a + b
	log.Println(c)
	w.Write([]byte(fmt.Sprintf("%d", c)))
}

// parseFloatFromRequestBodyJson
// func parseFloatFromRequestBodyJson(key string, )
