package parsing

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

// IndexPage Parses index.html which is in templates directory.
func IndexPage() {
	template, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		log.Printf("Couldn't Parse The Templates: %v", err)
	}
	template.Execute(os.Stdout, "Barack Obama")
}

// AboutPage Parses about.html which is in templates directory.
func AboutPage() {
	var list []string = []string{
		"Karim",
		"Rahim",
		"Ahmad",
		"Khan",
	}
	fmt.Println("\n\nAbout Page Content:")
	template, err := template.ParseFiles("./templates/about.html")
	if err != nil {
		log.Printf("Couldn't Parse The Templates: %v", err)
	}
	template.Execute(os.Stdout, list)
}
