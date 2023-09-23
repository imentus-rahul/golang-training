package main

import (
	"log"
	"rest-api-mux-gorilla-mongo/configs"
	"rest-api-mux-gorilla-mongo/routes"
	"net/http"

	"github.com/gorilla/mux"

	"fmt"
)

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "The Go Server is Live!")
    fmt.Println("Endpoint Hit: /")
}

func main() {
	router := mux.NewRouter()

	//run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(router)

	router.HandleFunc("/", homePage)

	log.Fatal(http.ListenAndServe(":6000", router))
}
