package models

// BlogPage - Страница блога
type BlogPage struct {
	Name string
	Blog []Post
	Find Post
}