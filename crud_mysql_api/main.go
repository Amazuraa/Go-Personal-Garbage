package main

import (
	"fmt"
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	// router.HandleFunc("/getproducts", returnAllProducts).Methods("GET")
	http.Handle("/", router)
	fmt.Println("Connected to port 1234")
	log.Fatal(http.ListenAndServe(":1234", router))

}