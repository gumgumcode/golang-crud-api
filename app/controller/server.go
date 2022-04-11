package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	controller "github.com/omkarscode/restapi/controller/blog"
)

var router *mux.Router

func initHandlers() {

	router.HandleFunc("/api/posts", controller.GetAllPosts).Methods("GET")

	router.HandleFunc("/api/posts/{id}", controller.GetPost).Methods("GET")

	router.HandleFunc("/api/posts", controller.CreatePost).Methods("POST")

	router.HandleFunc("/api/posts", controller.UpdatePost).Methods("PUT")

	router.HandleFunc("/api/posts/{id}", controller.DeletePost).Methods("DELETE")
}

func Start() {

	router = mux.NewRouter()
	
	initHandlers()

	fmt.Printf("Router is listening on port 3200\n")

	log.Fatal(http.ListenAndServe(":3200", router))
}