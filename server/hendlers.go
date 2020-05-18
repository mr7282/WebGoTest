package server

import (
	"encoding/json"
	"homework-3/models"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var myFindPage = &models.BlogPage{}

func (serv *Server) viewBlog(wr http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.New("myBlog").ParseFiles("./www/templates/index.html"))

	myBlog, err := models.ShowAll(serv.db)
	myBlogPage := models.BlogPage{
		Name: "Мой блог",
		Blog: myBlog,
	}
	if err != nil {
		serv.lg.WithError(err).Println("can't show all posts")
	}


	if err := tmpl.ExecuteTemplate(wr, "Blog", myBlogPage); err != nil {
		log.Println(err)
	}
}

func (serv *Server) viewFind(wr http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.New("myFind").ParseFiles("./www/templates/find.html"))
	myFindPage.Name = "Поиск"
	if err := tmpl.ExecuteTemplate(wr, "Find", myFindPage); err != nil {
		serv.lg.WithError(err).Fatal("This page can't be displayed.  ./find.html")
	}
}

func (serv *Server) responseFind(wr http.ResponseWriter, r *http.Request) {
	reqJSON, err := ioutil.ReadAll(r.Body)
	if err != nil {
		serv.lg.Warning(err)
	}
	findRequest := ""
	if err := json.Unmarshal(reqJSON, &findRequest); err != nil {
		serv.lg.Warning(err)
	}
	fr, _ := strconv.Atoi(findRequest)
	myFindPage.Find, err = models.FindPost(serv.db, fr)
	if err != nil {
		serv.lg.WithError(err).Info("Проблемы с SELECT запросом")
	}
}


func (serv *Server) createPostHTML(wr http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.New("myCreate").ParseFiles("./www/templates/create.html"))
	myCreatePage := models.BlogPage{Name: "Новый пост"}
	if err := tmpl.ExecuteTemplate(wr, "Blog", myCreatePage); err != nil {
		serv.lg.WithError(err).Warning("This page can't be displayed.  ./create.html")
	}
}

func (serv *Server) createPost(wr http.ResponseWriter, r *http.Request) {
	reqJSON, err := ioutil.ReadAll(r.Body)
	if err != nil {
		serv.lg.WithError(err).Warning("can't read r.Body")
	}

	myResponse := models.RespStruct{}
	if err := json.Unmarshal(reqJSON, &myResponse); err != nil {
		serv.lg.Warning(err)
	}
	newPost := models.Post{
		Name: myResponse.NamePost,
		Post: myResponse.TextPost,
	}

	err = newPost.CreatePost(serv.db)
	if err != nil {
		serv.lg.WithError(err).Warning("can't create new post")
	}

}

func (serv *Server) editPostView(wr http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.New("myEdit").ParseFiles("./www/templates/edit.html"))
	myFindPage.Name = "Редактирование"
	serv.lg.Info(myFindPage.Find)
	if err := tmpl.ExecuteTemplate(wr, "Blog", myFindPage); err != nil {
		log.Println(err)
	}
}

func (serv *Server) editPost(wr http.ResponseWriter, r *http.Request) {
	respJSON, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	respUnMarsh := models.RespStruct{}

	if err := json.Unmarshal(respJSON, &respUnMarsh); err != nil {
		log.Println(err)
	}

	NewPostEdit := models.Post{
		ID: myFindPage.Find.ID,
		Name: respUnMarsh.NamePost,
		Post: respUnMarsh.TextPost,
	}

	if err = NewPostEdit.EditPost(serv.db); err != nil {
		serv.lg.WithError(err).Warning("can't update your post!")
	}
}
