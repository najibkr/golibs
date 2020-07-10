package intro

import "fmt"

// RunIntroduction will print out introductory lesson info to in the terminal:
func RunIntroduction() {
	fmt.Println("Read more at: https://blog.gopheracademy.com/advent-2017/using-go-templates/")
	fmt.Println("Lesson 02: Templates in Go")
	fmt.Println("There are two types of templates in Go: text/template and html/template")
	fmt.Println("Both have the same interface, html though secures the html")
	fmt.Println("\nTo Review the Documentation of template package Please Run \ngo doc template\nin the command line!")
	fmt.Println("\nTo Review the Documentation of template.Template, Please Run \ngo doc template.Template\nin the command line!")
}
