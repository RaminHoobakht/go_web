package main

import (
	"fmt"
	"net/http"
)

func main() {

	/* create an endpoint to say hello */
	http.HandleFunc("/hello", say_hello)

	/* create an endpoint to say good bye */
	http.HandleFunc("/bye", say_goodbye)

	/* message to inform for starting server  */
	fmt.Println("http://localhost:8080 server is starting ...")

	/* listen to the incoming request from http://localhost:8080 */
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("\n #(14:10:27): The End ...")
}

/* -------------------------------------------------------------------------------- */

func say_hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Hinata, How is learning Go programming ...\n")
}

func say_goodbye(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Good Bye, Take care ...\n")
}
