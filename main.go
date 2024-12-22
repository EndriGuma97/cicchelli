package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}


func main() {
	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", index)
	http.HandleFunc("/blog", blog)
	http.HandleFunc("/blog/", blogPost)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":80", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}
func blog(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "blog.html", BlogPosts)
}
func blogPost(w http.ResponseWriter, req *http.Request) {
	// Extract the slug from the URL
	slug := req.URL.Path[len("/blog/"):]
	for _, post := range BlogPosts {
		if post.Slug == slug {
			tpl.ExecuteTemplate(w, "post.html", post)
			return
		}
	}
	http.NotFound(w, req)
}