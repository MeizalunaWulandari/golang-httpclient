package main

import (
	post "github.com/MeizalunaWulandari/golang-httpclient/controllers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", post.Index)
	http.HandleFunc("/posts", post.Index)
	http.HandleFunc("/post/create", post.Create)
	http.HandleFunc("/post/store", post.Store)
	http.HandleFunc("/post/delete", post.Delete)

	log.Print("Server started on: http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
