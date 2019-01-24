package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"io/ioutil"
	"html/template"
)


// WebView is a framework to send objects & data to a Web view
type WebView struct {
	Redirect    bool
	Title string
	CategoryMap map[string]int
	Counter    []int
	Flashes []interface{}
}


// string helper functions because you always end up needing them
// and why waste time on stack overflow when you dont have to

func SplitLines(s string) []string {
	sli := strings.Split(s, "/n")
	return sli
}

func sliceString(s string, i int) string {

	l := len(s)

	if l > i {
		return s[:i] + "..."
	}
	return s[:l]
}

func subtract(a, b int) int {
	return a - b
}

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) int {
	return a / b
}

func percent(a, b int) float32 {
	return (float32(a) / float32(b)) * 100.0
}

func isIn(s []int, t int) bool {
	for _, n := range s {
		if n == t {
			return true
		}
	}
	return false
}

func isInString(s []string, t string) bool {
	for _, n := range s {
		if n == t {
			return true
		}
	}
	return false
}

// This is where you will fire up and execute your template files. This is like your view, or graphical
// user interface

var templates *template.Template
func InitalizeGuiTemplates() {
    var allFiles []string
    files, err := ioutil.ReadDir("./bbtemplates")
    if err != nil {
        fmt.Println(err)
    }
    for _, file := range files {
        filename := file.Name()
        if strings.HasSuffix(filename, ".tmpl") {
            allFiles = append(allFiles, "./bbtemplates/"+filename)
        }
    }
	templates, err = template.ParseFiles(allFiles...) //parses all .tmpl files in the 'templates' folder
}

func RenderRoute(w http.ResponseWriter, data interface{}) {
	t := templates.Lookup("base.html.tmpl")
	if err := t.ExecuteTemplate(w, "base", data); err != nil {
        log.Printf("Failed to execute template: %v", err)
    }	
}
