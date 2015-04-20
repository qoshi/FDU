package main

import (
	"fmt"
	"io"
	"linkedin/service/mongodb"
	"log"
	"net/http"
	myURL "net/url"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func newSignin(w http.ResponseWriter, req *http.Request) {

}

func getSignin(w http.ResponseWriter, req *http.Request) {
	u, err := myURL.Parse(req.RequestURI)
	if err == nil {
		log.Println()
	}
	query := u.Query()
	start, existStart := query["start"]
	count, existCount := query["count"]
	callback, existCallback := query["callback"]
	if !existStart || !existCount || !existCallback {
		return
	}
	log.Println(start, count, callback)
}

func newMessage(w http.ResponseWriter, req *http.Request) {

}

func getMessage(w http.ResponseWriter, req *http.Request) {
	u, err := myURL.Parse(req.RequestURI)
	if err == nil {
		log.Println()
	}
	query := u.Query()
	start, existStart := query["start"]
	count, existCount := query["count"]
	callback, existCallback := query["callback"]
	if !existStart || !existCount || !existCallback {
		return
	}
	log.Println(start, count, callback)
}

func main() {
	mongodb.Init("mongodb://127.0.0.1:27017/fudan")
	http.HandleFunc("/hello", HelloServer)
	http.HandleFunc("/gM", getMessage)
	http.HandleFunc("/gS", getSignin)
	fmt.Println("listening")
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
