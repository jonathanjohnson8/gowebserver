package main

import (
	"fmt"
	"log"

	//net/http is a package that contains both the http server and the http client
	"net/http"
	"strconv"
	"sync"
)

var counter int
var mutex = &sync.Mutex{}

// here, we create a function that will write "hello" on the main page of the web server - http://localhost:8081/
func echoString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

// here we create a function that will tell us how many times the button has been clicked
func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()                          // here we lock the index to begin (to prevent Go's multi-threaded nature from changing the counter prematurely)
	counter++                             // here, we increment the counter to display the number of site visits, but I'm not sure how this
	fmt.Fprintf(w, strconv.Itoa(counter)) // here we print the counter to show how many times the page has been visited
	mutex.Unlock()                        // here we unlock the index to allow Go's multi-threaded nature to set a new value
}

func main() {
	http.HandleFunc("/", echoString) // here we

	http.HandleFunc("/increment", incrementCounter)

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h2>Hello WebServer Vistors!</h2>"))
		fmt.Fprintf(w, "Hi")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
