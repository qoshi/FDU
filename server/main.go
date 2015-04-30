package main

import (
	"fmt"
	"io"
	"linkedin/service/mongodb"
	"log"
	"net/http"
	myURL "net/url"

	"encoding/json"
	"strconv"
	"strings"
)

const (
	returnFmt     = "{success:%d,data:%s}"
	statusSuccess = 1
	statusFailed  = 0
)

//new sign in
func newSignin(w http.ResponseWriter, req *http.Request) {
	ip := strings.Split(req.RemoteAddr, ":")[0]
	log.Println(ip)
	name := req.FormValue("name")
	location, _ := strconv.Atoi((req.FormValue("location")))
	callback := req.FormValue("callback")
	err := NewSignin(name, location, ip)
	var result string
	if err != nil {
		log.Println(err, name, location)
		result = fmt.Sprintf(returnFmt, statusFailed, "new Signin failed")
	} else {
		result = fmt.Sprintf(returnFmt, statusSuccess, "new Signin success")
	}
	if callback != "" {
		result = callback + "(" + result + ")"
	}
	log.Println(result)
	io.WriteString(w, result)
}

func getSignin(w http.ResponseWriter, req *http.Request) {
	ip := strings.Split(req.RemoteAddr, ":")[0]
	log.Println(ip)
	u, err := myURL.Parse(req.RequestURI)
	if err == nil {
		log.Println()
	}
	query := u.Query()
	startS, existStart := query["start"]
	countS, existCount := query["count"]
	callback, existCallback := query["callback"]
	var result string
	if !existStart || !existCount {
		result = fmt.Sprintf(returnFmt, statusFailed, "missing params")
	} else {
		start, _ := strconv.Atoi(startS[0])
		count, _ := strconv.Atoi(countS[0])
		ret := GetSignin(start, count)
		tmp, err1 := json.Marshal(ret)
		if err1 != nil {
			log.Println(err, ret)
			result = fmt.Sprintf(returnFmt, statusFailed, "parsing error")
		} else {
			ret := string(tmp)
			result = fmt.Sprintf(returnFmt, statusSuccess, ret)
		}
	}
	if existCallback {
		result = callback[0] + "(" + result + ")"
	}
	log.Println(result)
	io.WriteString(w, result)
}

func newMessage(w http.ResponseWriter, req *http.Request) {
	ip := strings.Split(req.RemoteAddr, ":")[0]
	context := req.FormValue("context")
	name := req.FormValue("name")
	country, _ := strconv.Atoi((req.FormValue("Location")))
	callback := req.FormValue("callback")
	err := NewMsg(name, country, context, ip)
	var result string
	if err != nil {
		log.Println(err, name, context)
		result = fmt.Sprintf(returnFmt, statusFailed, "new msg failed")
	} else {
		result = fmt.Sprintf(returnFmt, statusSuccess, "new msg success")
	}
	if callback != "" {
		result = callback + "(" + result + ")"
	}
	log.Println(result)
	io.WriteString(w, result)
}

func getMessage(w http.ResponseWriter, req *http.Request) {
	ip := strings.Split(req.RemoteAddr, ":")[0]
	log.Println(ip)
	var result string
	u, err := myURL.Parse(req.RequestURI)
	if err == nil {
		log.Println()
	}
	query := u.Query()
	startS, existStart := query["start"]
	countS, existCount := query["count"]
	callback, existCallback := query["callback"]
	if !existStart || !existCount {
		result = fmt.Sprintf(returnFmt, statusFailed, "missing params")
	} else {
		start, _ := strconv.Atoi(startS[0])
		count, _ := strconv.Atoi(countS[0])
		log.Println("~~~~~~~~~~~~~~")
		log.Println(start, count)
		ret := GetMsg(start, count)
		tmp, err1 := json.Marshal(ret)
		if err1 != nil {
			log.Println(err, ret)
			result = fmt.Sprintf(returnFmt, statusFailed, "new Signin failed")
		} else {
			ret := string(tmp)
			result = fmt.Sprintf(returnFmt, statusSuccess, ret)
		}
	}
	if existCallback {
		result = callback[0] + "(" + result + ")"
	}
	log.Println(result)
	io.WriteString(w, result)
}

func main() {
	mongodb.Init("mongodb://127.0.0.1:27017/fudan")
	http.HandleFunc("/nM", newMessage)
	http.HandleFunc("/nS", newSignin)
	http.HandleFunc("/gM", getMessage)
	http.HandleFunc("/gS", getSignin)
	log.Println("listening")
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
