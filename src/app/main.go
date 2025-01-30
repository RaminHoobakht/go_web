package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	my_router := mux.NewRouter()

	my_router.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "the name of the book is: %s on page: %s", title, page)
	})

	fmt.Println("server bootstrapping is running ...")
	http.ListenAndServe(":8080", my_router)

	fmt.Println("\n #(17:02:06): The End ...")
}
