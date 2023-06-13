package main

type RequestHomepage struct {
	Limit int `json:"limit"`
}

type ResponseError struct {
	Error string `json:"error"`
}

type ResponseUser struct {
	Id   int
	Name string
	Age  int
}
