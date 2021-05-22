package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/gorilla/mux"
)

func getProjectPath() string {
	currentDir, _ := os.Getwd()
	return filepath.Join(currentDir, "books", "go-blueprints", "chat")
}

type handlerMiddleWare func(res http.ResponseWriter, req *http.Request)

// Handler that responds with templates
type templateHandler struct {
	preWriter handlerMiddleWare
	filename  string
	once      sync.Once
	templ     *template.Template
}

// ServeHttp for Handler interface
func (t *templateHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// Do a lazy parsing of the template once
	t.once.Do(func() {
		log.Print("Compiling template: ", t.filename)
		templatePath := filepath.Join(getProjectPath(), "templates", t.filename)
		log.Print("Template path is ", templatePath)
		t.templ = template.Must(template.ParseFiles(templatePath))
	})

	// Prewriter Middleware
	if t.preWriter != nil {
		log.Print("Executing prewriter")
		t.preWriter(res, req)
	} else {
		log.Print("No prewriter, skipping")
	}

	// Write template response
	if t.templ != nil {
		log.Print("Writing template response")
		t.templ.Execute(res, nil)
	} else {
		log.Print("No compiled template found")
	}
}

func main() {
	chatHandler := &templateHandler{
		filename: "chat.html",
		preWriter: func(res http.ResponseWriter, req *http.Request) {
			res.Header().Add("SERVER", "GoLang")
		},
	}

	aboutHandler := &templateHandler{filename: "about.html"}

	pageNotFoundHandler := &templateHandler{
		filename: "404.html",
		preWriter: func(res http.ResponseWriter, req *http.Request) {
			res.WriteHeader(404)
		},
	}

	router := mux.NewRouter()
	router.Handle("/", chatHandler)
	router.Handle("/404", pageNotFoundHandler)
	router.Handle("/about", aboutHandler)
	router.HandleFunc("/redirect", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Add("Location", "/about")
		res.WriteHeader(302)
	})
	http.Handle("/", router)

	log.Print("Serving on: ", 3000)
	if err := http.ListenAndServe("localhost:3000", nil); err != nil {
		log.Fatal("ListenAndServe err:", err)
	}
}
