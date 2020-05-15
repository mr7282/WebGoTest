package server

import (
	"net/http"
)

// Router - вызывает обработчики в зависимости от поступившего запроса
func (serv *Server) Router(route *http.ServeMux) {
	route.Handle("/favicon.ico", http.FileServer(http.Dir("./www")))
	route.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	route.HandleFunc("/index", serv.viewBlog)
	route.HandleFunc("/find", serv.viewFind)
	route.HandleFunc("/find/post", serv.responseFind)
	route.HandleFunc("/create", serv.createPostHTML)
	route.HandleFunc("/create/post", serv.createPost)
	// route.HandleFunc("/edit", editPostView)
	// route.HandleFunc("/edit/post", editPost)
}

