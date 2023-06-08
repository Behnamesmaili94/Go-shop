package main

import (
	"fmt"
	"mymodule/package/handlers"
	"net/http"
)

const port = "8080"

// this is main handle all function and listen to 8080 port
func main() {
	fmt.Print("server is running:\n")

	http.HandleFunc("/home", handlers.Home)
	http.HandleFunc("/pageOne", handlers.PageOne)
	http.HandleFunc("/aboutUs", handlers.AboutUs)
	_ = http.ListenAndServe(port, nil)

}
