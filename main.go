package main

import "fmt"

type task struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}
type ListTask []task

var tasks = ListTask{
	{
		ID:      1,
		Name:    "Task One",
		Content: "Some Content",
	},
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(tasks)
}
