package parsing

import (
	"os"
	"text/template"
)

// Todo is the Model for the Todo Entity.
type Todo struct {
	Title       string
	Description string
	Completed   bool
}

// ParseTextTemplate1 Parses Templates:
func ParseTextTemplate1() {
	td := Todo{
		Title:       "Do Assignment",
		Description: "Finish Golang Assingment",
		Completed:   true,
	}
	newTmpl := template.New("todo")
	// tmpl, err := newTmpl.Parse("Title:\t\t\"{{.Title}}\"\nDescription:\t\"{{.Description}}\" \nCompleted:\t\"{{.Completed}}\"")
	// if err != nil {
	// 	panic(err)
	// }
	tmpl := template.Must(newTmpl.Parse("Title:\t\t\"{{.Title}}\"\nDescription:\t\"{{.Description}}\" \nCompleted:\t\"{{.Completed}}\""))
	// err = tmpl.Execute(os.Stdout, td)
	err := tmpl.Execute(os.Stdout, td)
	if err != nil {
		panic(err)
	}
}

// ParseTextTemplate2 Parses Templates:
func ParseTextTemplate2() {
	td := []Todo{
		Todo{
			Title:       "Do Assignment",
			Description: "Finish Golang Assingment",
			Completed:   true,
		},
		Todo{
			Title:       "Do Assignment",
			Description: "Finish Golang Assingment",
			Completed:   true,
		},
	}
	template := template.Must(template.ParseFiles("./templates/todo.text"))
	// if err != nil {
	// 	panic(err)
	// }
	err := template.Execute(os.Stdout, td)
	if err != nil {
		panic(err)
	}

}
