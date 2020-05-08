package server

import (
	"net/http"
)

// Router - вызывает обработчики в зависимости от поступившего запроса
func Router(route *http.ServeMux) {
	route.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	route.HandleFunc("/", viewBlog)
	route.HandleFunc("/find", viewFind)
	route.HandleFunc("/find/post", responseFind)
	route.HandleFunc("/create", createPostHTML)
	route.HandleFunc("/create/post", createPost)
	route.HandleFunc("/edit", editPostView)
	// route.HandleFunc("/edit/post", editPost)
}

