package server

import (
	"net/http"

	"log"
)

type Server struct {
	RootDir string
	TemplateDir string
	IndexTemplate string

}

// StartServer - поднимает сервер, обрабатывает ошибку в случае неудачи
func StartServer() {
	route := http.NewServeMux()
	Router(route)

	log.Fatal(http.ListenAndServe(":8080", route))
}

// func NewServer() *Server {
// 	return &Server{
// 		RootDir: "homework-3",
// 		TemplateDir: "www/templates",
// 		IndexTemplate: "./index.html"
// 	}
// }




