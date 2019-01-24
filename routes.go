package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

// set up the handlers, these will do the meat of the app (this is like your middleware)
func IndexHandler(w http.ResponseWriter, req *http.Request) {
	rs := fmt.Sprintf("%#v", req)
	rs=""
	fmt.Println("IndexHandler call" + rs)
	status := "do index"

	wv := WebView {
		Title: status,
		// you can find the WebView definition in gui.go
	}
	// Render page
	RenderRoute(w, wv)
}
func GuideHandler(w http.ResponseWriter, req *http.Request) {
	rs := fmt.Sprintf("%#v", req)
	rs=""
	fmt.Println("GuideHandler call: " + rs)
	status := "do guide"

	wv := WebView {
		Title: status,
		// its what you use when you want to pass vars into the template

	}
	// Render page
	RenderRoute(w, wv)
}
func SearchHandler(w http.ResponseWriter, req *http.Request) {
	rs := fmt.Sprintf("%#v", req)
	rs=""
	fmt.Println("SearchHandler call" + rs)
	status := "do search: "

	fmt.Println("GET got:", req.URL.Query())

	recognitionlanguage := req.FormValue("recognition_language")
	autotranslate := req.FormValue("recognition_lang_tlxd")
	searchquery := req.FormValue("basicbits_ask")

	vars := mux.Vars(req)
	q := vars["q"]
	status = status + " " + q + recognitionlanguage + " " + autotranslate + " " + searchquery 

	// todo: next step is to update the webview, match it to the objects we expect
	wv := WebView {
		Title: status,
		// add your actual stuff here and you can
		// get them by using the .Title style of access
	}
	// Render page
	RenderRoute(w, wv)
}
func RedirectHandler(w http.ResponseWriter, req *http.Request) {
	rs := fmt.Sprintf("%#v", req)
	rs=""
	fmt.Println("RedirectHandler call" + rs)
	vars := mux.Vars(req)
	q := vars["q"]
	
	log.Println("Redirecting to home: " + q)
	http.Redirect(w, req, "/", 302)
	return
}

func CreateRouter(port string) {
	decoder  := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
	r := mux.NewRouter()	
   
    r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/guide/{q}", GuideHandler)
	r.HandleFunc("/search", SearchHandler)
	r.HandleFunc("/api", RedirectHandler)	
	
	// Note: This is the moment I realized how cool Go actually is.
	// It's been fighting me until now to get a basic web page to render.
	// I just wanted it to get out of my may and do what I expected, but that the thing
	// Go isn't like the other ones. It's a completely different mindset.
	// You must be flexible like the branches of a tree lest ye break in the cold canadian winters
	// 
	// Also Sin's Cat/Dog Hypothesis applies here: If you pet a cat like you would a dog, you're gonna
	// get bitten in all probability. If you pet a dog like a cat, well thats an exercise for the reader
	// most large breeds of dog are heavier than I am. It's still a jungle, no matter how much AI
	// we endeavor to build
	//
	// Enough philosophy, let's get back to the coolness of Go
	//
	// so Go is a bit different, your files dont exist unless you say they do.
	// we're going to invoke magic to create directory that doesnt really exist root/css/
	// that we can use in our templates (so all css can live together effectively)
	// same for js. You can use this in a lot more powerful ways. These are just the basics
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/",
		http.FileServer(http.Dir("bbstatic/css/"))))

	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/",
		http.FileServer(http.Dir("bbstatic/js/"))))

	fmt.Println("Starting webserver on port " + port)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":"+port, r))
}