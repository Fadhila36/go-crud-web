package main

import (
	"go-web-native/config"
	"go-web-native/controllers/categorycontroller"
	"go-web-native/controllers/homecontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	// homepage
	http.HandleFunc("/", homecontroller.Welcome)

	// categories
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	log.Println("Server runing on port 8080")
	http.ListenAndServe(":8080", nil)
}
