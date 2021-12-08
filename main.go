package main

/**
 * This is the main file for the Task application
 * License: MIT
 **/
import (
	"flag"
	"log"
	"net/http"
	"strings"

	"github.com/strick-j/PVWA-SCIM/config"
	"github.com/strick-j/PVWA-SCIM/views"
)

func main() {
	values, err := config.ReadConfig("config.json")
	var port *string

	if err != nil {
		port = flag.String("port", "", "IP address")
		flag.Parse()

		//User is expected to give :8080 like input, if they give 8080
		//we'll append the required ':'
		if !strings.HasPrefix(*port, ":") {
			*port = ":" + *port
			log.Println("port is " + *port)
		}

		values.ServerPort = *port
	}

	views.PopulateTemplates()

	//these handlers are for users
	http.HandleFunc("/users/", views.RequiresLogin(views.ShowUserFunc))

	//these handlers are for groups
	http.HandleFunc("/groups/", views.RequiresLogin(views.ShowGroupsFunc))

	//these handlers are for main
	http.HandleFunc("/home/", views.RequiresLogin(views.ShowMainFunc))

	log.Println("running server on ", values.ServerPort)
	log.Fatal(http.ListenAndServe(values.ServerPort, nil))
}
