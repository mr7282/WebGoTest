package server

import (
	"fmt"
	"encoding/json"
	"homework-3/models"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)



var myBlog = models.BlogPage{
	Name: "Мой Блог",
	Blog: []models.Post{
		models.Post{1, "Мой первый пост", "Далее создадим глобальную переменную, опишем в ней простой лист, подготовим переменную, в которую будет считываться шаблон при запуске приложения, и создадим роутер, который будет отдавать нам страницу со списком."},
		models.Post{2, "Мой второй пост", "{{year}} — вызов функций шаблона происходит по названию, без указания точки перед ее именем. Функции необходимо добавлять к структуре *Template перед тем, как производить чтение шаблона. "},
		models.Post{3, "Мой третий пост", "В шаблон можно встроить общие функции. Например, на сайтах в конце страницы часто указывают копирайт, и в этой строке — текущий год. Создадим функцию, которая будет возвращать в шаблон текущий год."},
	},
}



func viewBlog(wr http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.New("myBlog").ParseFiles("./www/templates/index.html"))
	if err := tmpl.ExecuteTemplate(wr, "Blog", myBlog); err != nil {
		log.Println(err)
	}
}

func viewFind(wr http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.New("myFind").ParseFiles("./www/templates/find.html"))
	if err := tmpl.ExecuteTemplate(wr, "Blog", myBlog); err != nil {
		log.Println(err)
	}
}

func responseFind(wr http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.New("myFind").ParseFiles("./www/templates/find.html"))
	reqJSON, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	findRequest := ""
	if err := json.Unmarshal(reqJSON, &findRequest); err != nil {
		log.Println(err)
	}
	fr, _ := strconv.Atoi(findRequest)
	myBlog.Find = myBlog.Blog[fr-1]
	fmt.Println(myBlog.Find)
	if err := tmpl.ExecuteTemplate(wr, "Blog", myBlog); err != nil {
		log.Println(err)
	}
}

func createPostHTML(wr http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.New("myFind").ParseFiles("./www/templates/create.html"))
	if err := tmpl.ExecuteTemplate(wr, "Blog", myBlog); err != nil {
		log.Println(err)
	}
}

func createPost(wr http.ResponseWriter, r *http.Request) {
	reqJSON, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	myRespStruct := models.RespStruct{}

	if err := json.Unmarshal(reqJSON, &myRespStruct); err != nil {
		log.Println(err)
	}
	NewPost := models.Post{
		ID: len(myBlog.Blog)+1,
		Name: myRespStruct.NamePost,
		Body: myRespStruct.TextPost,

	}
	myBlog.Blog = append(myBlog.Blog, NewPost)

	fmt.Println(myRespStruct)
	fmt.Println(myBlog)
}

func editPostView(wr http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.New("myEdit").ParseFiles("./www/templates/edit.html"))
	if err := tmpl.ExecuteTemplate(wr, "Blog", myBlog); err != nil {
		log.Println(err)
	}
}


func editPost(wr http.ResponseWriter, r *http.Request) {
	respJSON, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	respUnMarsh := models.RespStruct{}

	if err := json.Unmarshal(respJSON, &respUnMarsh); err != nil {
		log.Println(err)
	}

	NewPostEdit := models.Post{
		ID: myBlog.Find.ID,
		Name: respUnMarsh.NamePost,
		Body: respUnMarsh.TextPost,
	}

	myBlog.Blog[myBlog.Find.ID - 1].Name = NewPostEdit.Name
	myBlog.Blog[myBlog.Find.ID - 1].Body = NewPostEdit.Body
	myBlog.Find = models.Post{0, "", ""}
}