package main

import (
	"fmt"
	"log"
	"math"
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

	addEndpoint := os.Getenv("ADDENDPOINT")
	if addEndpoint == "" {
		log.Fatal("ADDENDPOINT must be set")
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/multiply", multiplicationHandler)
	err = http.ListenAndServe(":"+portString, r)
	if err != nil {
		log.Fatal(err)
	}
}

//HTTP handler adding two parameters from url routing parameters
func multiplicationHandler(w http.ResponseWriter, r *http.Request) {
	aString := r.FormValue("a")
	bString := r.FormValue("b")
	log.Println(aString, bString)
	a, err := strconv.Atoi(aString)
	if err != nil {
		http.Error(w, "a must be an int", 500)
		return
	}
	b, err := strconv.Atoi(bString)
	if err != nil {
		http.Error(w, "b must be an int", 500)
		return
	}
	c := multiply(a, b, os.Getenv("ADDENDPOINT"))
	log.Println(c)
	w.Write([]byte(fmt.Sprintf("%d", c)))
}

// Abs returns the absolute value of x.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minMax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func multiply(a, b int, endpoint string) int {
	// okay now you have the numbers. what do you do now?
	// check if they share a sign, then do multiplication. if they don't, flip the result
	if math.Signbit(float64(a)) == math.Signbit(float64(b)) {
		aAbs := abs(a)
		bAbs := abs(b)
		min, max := minMax(aAbs, bAbs)
		//Send a request to  /add: looping over min, max times
		out := 0
		for i := 0; i <= min; i++ {
			// TODO: this needs to be a call to another http endpoint
			out = out + max
		}
		return out
	}
	aAbs := abs(a)
	bAbs := abs(b)
	min, max := minMax(aAbs, bAbs)
	//Send a request to  /add: looping over min, max times
	out := 0
	for i := 0; i <= min; i++ {
		// TODO: this needs to be a call to another http endpoint
		out = out + max
	}
	return -out
}
