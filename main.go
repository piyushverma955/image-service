package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/razorpay/image-service/controller"
	"net/http"
)

var Port = ":8080"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/convert", controller.ConvertImage).Methods("POST")
	err := http.ListenAndServe(Port, r)
	if err != nil {
		fmt.Println("error while running server ", err)
	}
}
