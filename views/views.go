package views

/*Holds the fetch related view handlers*/

import (
	"html/template"
	"net/http"
)

var templates *template.Template
var usersTemplate *template.Template
var homeTemplate *template.Template
var groupsTemplate *template.Template

var message string //message will store the message to be shown as notification
var err error

//ShowHomeFunc is used to populate the "/home/" URL
func ShowMainFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		homeTemplate.Execute(w, nil)
	}
}

//ShowUserFunc is used to populate the "/users/" URL
func ShowUserFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		usersTemplate.Execute(w, nil)
	}
}

//ShowGroupsFunc is used to populate the "/groups/" URL
func ShowGroupsFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		groupsTemplate.Execute(w, nil)
	}
}
