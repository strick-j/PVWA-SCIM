package views

/*
Holds the non insert/update/delete related view handlers
*/

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

//PopulateTemplates is used to parse all templates present in
//the templates folder
func PopulateTemplates() {
	var allFiles []string
	templatesDir := "./templates/"
	files, err := ioutil.ReadDir(templatesDir)
	if err != nil {
		log.Println(err)
		os.Exit(1) // No point in running app if templates aren't read
	}
	for _, file := range files {
		filename := file.Name()
		if strings.HasSuffix(filename, ".html") {
			allFiles = append(allFiles, templatesDir+filename)
		}
	}

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	templates, err = template.ParseFiles(allFiles...)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	usersTemplate = templates.Lookup("users.html")
	groupsTemplate = templates.Lookup("groups.html")
	homeTemplate = templates.Lookup("home.html")

}
