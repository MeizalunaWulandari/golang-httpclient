package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

type Post struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId int    `json:"userId"`
}

var BASE_URL = "https://jsonplaceholder.typicode.com"

func Index(w http.ResponseWriter, r *http.Request) {

	var posts []Post

	response, err := http.Get(BASE_URL + "/posts")
	if err != nil {
		log.Print(err)
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&posts); err != nil {
		log.Print(err)
	}

	data := map[string]interface{}{
		"posts": posts,
	}

	temp, _ := template.ParseFiles("views/index.html")
	temp.Execute(w, data)
}
func Create(w http.ResponseWriter, r *http.Request) {

	var post Post
	var data map[string]interface{}

	id := r.URL.Query().Get("id")
	if id != "" {
		response, err := http.Get(BASE_URL + "/posts/" + id)
		if err != nil {
			log.Print(err)
		}
		defer response.Body.Close()

		decoder := json.NewDecoder(response.Body)
		if err := decoder.Decode(&post); err != nil {
			log.Print(err)
		}

		data = map[string]interface{}{
			"post": post,
		}

	}

	temp, _ := template.ParseFiles("views/create.html")
	temp.Execute(w, data)

}
func Store(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	id := r.Form.Get("post_id")

	idInt, _ := strconv.ParseInt(id, 10, 64)

	newPost := Post{
		Id:     int(idInt),
		Title:  r.Form.Get("post_title"),
		Body:   r.Form.Get("post_body"),
		UserId: 1,
	}

	jsonValue, _ := json.Marshal(newPost)
	buff := bytes.NewBuffer(jsonValue)
	var request *http.Request
	var err error

	if id != "" {
		// Update Request
		fmt.Println("Proses Update")
		request, err = http.NewRequest(http.MethodPut, BASE_URL+"/posts/"+id, buff)

	} else {
		// Create Request
		fmt.Println("Proses Create")
		request, err = http.NewRequest(http.MethodPost, BASE_URL+"/posts", buff)
		if err != nil {
			log.Print(err)

		}
	}

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	httpClient := &http.Client{}
	response, err := httpClient.Do(request)
	if err != nil {
		log.Print(err)
	}
	defer response.Body.Close()

	var postResponse Post

	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&postResponse); err != nil {
		log.Print(err)
	}

	//	fmt.Println(response.StatusCode)
	//	fmt.Println(response.Status)
	//	fmt.Println(postResponse)
	if response.StatusCode == 201 || response.StatusCode == 200 {
		http.Redirect(w, r, "/posts", http.StatusSeeOther)
	}

}
func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	request, err := http.NewRequest(http.MethodDelete, BASE_URL+"/posts/"+id, nil)
	if err != nil {
		log.Print(err)
	}

	httpClient := &http.Client{}
	response, err := httpClient.Do(request)
	if err != nil {
		log.Print("error disini", err)
	}
	defer response.Body.Close()

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)

	if response.StatusCode == 200 {
		http.Redirect(w, r, "/posts", http.StatusSeeOther)
	}

}
