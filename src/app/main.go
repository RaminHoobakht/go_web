package main

import (
	"fmt"
	"net/http"

)

func main() {
	
	http.HandleFunc("/hello", func(resp http.ResponseWriter, req *http.Request){
		fmt.Fprintf(resp,"Hello Hinata, How is learning Go programming?")
	})

	fmt.Println("Start Hinata Server ...")
	http.ListenAndServe(":8080", nil)
	
	fmt.Println("\n The End ...")
}
