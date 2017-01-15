package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

var templates *template.Template
var homeTemplate *template.Template

type Comment struct {
	ID      string
	Content string
	Author  string
	Created string
}

type Task struct {
	ID       string
	Title    string
	Content  string
	Created  string
	Comments []Comment
}

type Context struct {
	Tasks []Task
}

var comments []Comment
var tasks []Task

func main() {
	PopulateTemplate()
	http.Handle("/static/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/completed/", ShowCompletedTasksFunc)
	http.HandleFunc("/login/", LoginFunc)
	http.HandleFunc("/", HomeFunc)
	fmt.Println("Server running on 8081")
	http.ListenAndServe("0.0.0.0:8081", nil)
}

// PopulateTemplate reads the ./templates folder and parse all the html files inside it
// and  it stores it in the templates variable which will be looked up by other variables
func PopulateTemplate() {
	templates, err := template.ParseGlob("./templates/*.html")

	if err != nil {
		fmt.Println("main.go: PopulateTemplate: ", err)
		os.Exit(1)
	}

	homeTemplate = templates.Lookup("task.html")
}

// ShowCompletedTasksFunc populates the /completed URL for the GET request for the currently loggedin user.
func ShowCompletedTasksFunc(w http.ResponseWriter, r *http.Request) {
	context := Context{}
	homeTemplate.Execute(w, context)
}

// HomeFunc handles the / URL and asks the name of the user in German
func HomeFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		comment1 := Comment{
			ID:      "1",
			Content: "First Comment",
			Author:  "Mehre",
			Created: "15 Jan 2017",
		}

		comment2 := Comment{
			ID:      "2",
			Content: "Second Comment",
			Author:  "Mehre",
			Created: "15 Jan 2017",
		}

		comments = append(comments, comment1)
		comments = append(comments, comment2)

		task1 := Task{
			ID:       "1",
			Title:    "Title of First Task",
			Content:  "Golang",
			Created:  "Mehre",
			Comments: comments,
		}

		task2 := Task{
			ID:       "2",
			Title:    "Title of Second Task",
			Content:  "Watching twitch.tv",
			Created:  "Mehre 2",
			Comments: comments,
		}

		task3 := Task{
			ID:       "3",
			Title:    "Title of Third Task",
			Content:  "Sleep all day",
			Created:  "Mehre 3",
			Comments: comments,
		}

		tasks = append(tasks, task1)
		tasks = append(tasks, task2)
		tasks = append(tasks, task3)

		context := Context{Tasks: tasks}

		homeTemplate.Execute(w, context)
	}
}

// LoginFunc handles the /login URL and shows the profile page of the logged in user on a GET request
// handles the Login process on the post request
func LoginFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You are on the profile page")
}
