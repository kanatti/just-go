package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Add("SERVER", "GoLang")
		res.Write([]byte(`
		<html>
			<head>
			<title>Chat</title>
			</head>
			<body>
			Let's chat!
			</body>
		</html> `))
	})
	http.HandleFunc("/about", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`
		<html>
			<head>
			<title>About</title>
			</head>
			<body>
			Nothing Here
			</body>
		</html> `))
	})
	http.HandleFunc("/redirect", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Add("Location", "/about")
		res.WriteHeader(302)
	})
	http.HandleFunc("/404", func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(404)
		res.Write([]byte(`
		<html>
			<head>
			<title>404</title>
			</head>
			<body>
			Page Not Found
			</body>
		</html> `))
	})
	log.Print("Serving on: ", 3000)
	if err := http.ListenAndServe("localhost:3000", nil); err != nil {
		log.Fatal("ListenAndServe err:", err)
	}
}
