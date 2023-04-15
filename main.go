package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com/"

type todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	UserId    int    `json:"userId"`
}

var form = `
<h1>Todo #{{.ID}}</h1>
<div>{{printf "User %d" .UserId}}</div>
<div>{{printf "%s (completed: %t)" .Title .Completed}}</div>
`

func handler(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get(url + r.URL.Path[1:])
	if err != nil {
		fmt.Fprintln(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	var item todo

	if err = json.NewDecoder(resp.Body).Decode(&item); err != nil {
		fmt.Fprintln(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl := template.New("new")
	tmpl.Parse(form)
	tmpl.Execute(w, item)
}
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
