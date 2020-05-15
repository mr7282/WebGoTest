package server

import (
	// "fmt"
	"encoding/json"
	"homework-3/models"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// var myBlog = models.BlogPage{
// 	Name: "Мой Блог",
// 	Blog: []models.Post{
// 		models.Post{1, "Мой первый пост", "Далее создадим глобальную переменную, опишем в ней простой лист, подготовим переменную, в которую будет считываться шаблон при запуске приложения, и создадим роутер, который будет отдавать нам страницу со списком."},
// 		models.Post{2, "Мой второй пост", "{{year}} — вызов функций шаблона происходит по названию, без указания точки перед ее именем. Функции необходимо добавлять к структуре *Template перед тем, как производить чтение шаблона. "},
// 		models.Post{3, "Мой третий пост", "В шаблон можно встроить общие функции. Например, на сайтах в конце страницы часто указывают копирайт, и в этой строке — текущий год. Создадим функцию, которая будет возвращать в шаблон текущий год."},
// 	},
// }


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

/* func editPostView(wr http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.New("myEdit").ParseFiles("./www/templates/edit.html"))
	if err := tmpl.ExecuteTemplate(wr, "Blog", myBlog); err != nil {
		log.Println(err)
	}
} */

// func editPost(wr http.ResponseWriter, r *http.Request) {
// 	respJSON, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	respUnMarsh := models.RespStruct{}

// 	if err := json.Unmarshal(respJSON, &respUnMarsh); err != nil {
// 		log.Println(err)
// 	}

// 	NewPostEdit := models.Post{
// 		ID: myBlog.Find.ID,
// 		Name: respUnMarsh.NamePost,
// 		Body: respUnMarsh.TextPost,
// 	}

// 	myBlog.Blog[myBlog.Find.ID - 1].Name = NewPostEdit.Name
// 	myBlog.Blog[myBlog.Find.ID - 1].Body = NewPostEdit.Body
// 	myBlog.Find = models.Post{0, "", ""}
// }
