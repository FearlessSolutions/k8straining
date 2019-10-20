package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"net/url"
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

//HTTP handler adding two parameters from url routing parameters
func multiplicationHandler(w http.ResponseWriter, r *http.Request) {
	aString := r.FormValue("a")
	bString := r.FormValue("b")
	log.Println(aString, bString)
	a, err := strconv.Atoi(aString)
	if err != nil {
		http.Error(w, "a must be an int", 422)
		return
	}
	b, err := strconv.Atoi(bString)
	if err != nil {
		http.Error(w, "b must be an int", 422)
		return
	}
	c, err := multiply(a, b, os.Getenv("ADDENDPOINT"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
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

func multiply(a, b int, endpoint string) (int, error) {
	// okay now you have the numbers. what do you do now?
	// Figure out smaller number, to make fewest network calls needed
	aAbs := abs(a)
	bAbs := abs(b)
	min, max := minMax(aAbs, bAbs)
	//Send a request to  /add: looping over min, max times
	out := 0
	for i := 1; i <= min; i++ {
		// call addition endpoint: add max to current out
		resp, err := http.PostForm(endpoint,
			url.Values{"a": {strconv.Itoa(out)}, "b": {strconv.Itoa(max)}})
		if err != nil {
			log.Println("Error calling Addition Endpoint: %w", err)
			return 0, fmt.Errorf("Error calling Addition Endpoint: %w", err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error reading Addition Endpoint response: %w", err)
			return 0, fmt.Errorf("Error reading Addition Endpoint response: %w", err)
		}
		bs := string(body)
		//set out to output of response body
		out, err = strconv.Atoi(bs)
		if err != nil {
			log.Println("Error converting Addition Endpoint response: %w", err)
			return 0, fmt.Errorf("Error converting Addition Endpoint response: %w", err)
		}
	}
	if math.Signbit(float64(a)) != math.Signbit(float64(b)) {
		return -out, nil
	}
	return out, nil
}
