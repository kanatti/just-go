package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
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

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  socketBufferSize,
	WriteBufferSize: socketBufferSize,
}

type roomHandler struct {
	room *Room
}

// Upgrades HTTP connection to socket connection
// Joins the common room
// Leaves the room when connection is closed
// Reads and Writes concurrently into the websocket
func (r *roomHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(res, req, nil)

	if err != nil {
		log.Fatal("ServeHttp:", err)
		return
	}

	user := newUser(socket, r.room)

	r.room.join <- user
	defer func() {
		r.room.leave <- user
	}()

	go user.write()
	user.read()
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

	commonRoomHandler := &roomHandler{room: newRoom()}

	router := mux.NewRouter()
	router.Handle("/", chatHandler)
	router.Handle("/404", pageNotFoundHandler)
	router.Handle("/about", aboutHandler)
	router.HandleFunc("/redirect", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Add("Location", "/about")
		res.WriteHeader(302)
	})
	router.Handle("/room", commonRoomHandler)
	http.Handle("/", router)

	go commonRoomHandler.room.run()

	log.Print("Serving on: ", 3000)
	if err := http.ListenAndServe("localhost:3000", nil); err != nil {
		log.Fatal("ListenAndServe err:", err)
	}
}
