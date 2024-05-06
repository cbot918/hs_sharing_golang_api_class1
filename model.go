package main

type Post struct {
	Id    int
	Title string `json:"title"`
	Body  string `json:"body"`
}
