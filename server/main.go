package main

import (
	"fmt"
	"io"
	"linkedin/service/mongodb"
	"log"
	"net/http"
	myURL "net/url"

	"github.com/linkedin-inc/redis"

	"github.com/googollee/go-socket.io"

	"encoding/json"
	"strconv"
	"strings"
)

const (
	returnFmt     = "{success:%d,data:%s}"
	statusSuccess = 1
	statusFailed  = 0
	msgChannel    = "msg"
	signinChannel = "sigin"
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
	publish(signinChannel, result)
	if callback != "" {
		result = callback + "(" + result + ")"
	}
	log.Println(result)
	w.Header().Set("Content-Type", "application/json")
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
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, result)
}

func newMessage(w http.ResponseWriter, req *http.Request) {
	ip := strings.Split(req.RemoteAddr, ":")[0]
	context := req.FormValue("context")
	name := req.FormValue("name")
	location, _ := strconv.Atoi((req.FormValue("Location")))
	callback := req.FormValue("callback")
	err := NewMsg(name, location, context, ip)
	var result string
	if err != nil {
		log.Println(err, name, context)
		result = fmt.Sprintf(returnFmt, statusFailed, "new msg failed")
	} else {
		result = fmt.Sprintf(returnFmt, statusSuccess, "new msg success")
	}
	publish(msgChannel, result)
	if callback != "" {
		result = callback + "(" + result + ")"
	}

	log.Println(result)
	w.Header().Set("Content-Type", "application/json")
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
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, result)
}

func getTop(w http.ResponseWriter, req *http.Request) {
	var result string
	u, err := myURL.Parse(req.RequestURI)
	if err == nil {
		log.Println()
	}
	query := u.Query()
	callback, existCallback := query["callback"]
	top := getT()
	tmp, err := json.Marshal(top)
	if err != nil {
		log.Println(err)
		result = fmt.Sprintf(returnFmt, statusFailed, "get top failed")
	} else {
		ret := string(tmp)
		result = fmt.Sprintf(returnFmt, statusSuccess, ret)
	}
	if existCallback {
		result = callback[0] + "(" + result + ")"
	}
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, result)
}

func getChina(w http.ResponseWriter, req *http.Request) {
	var result string
	u, err := myURL.Parse(req.RequestURI)
	if err == nil {
		log.Println()
	}
	query := u.Query()
	fmt.Print(query)
	callback, existCallback := query["callback"]
	tmp, err := json.Marshal(getC())
	if err != nil {
		log.Println(err)
		result = fmt.Sprintf(returnFmt, statusFailed, "get china failed")
	} else {
		ret := string(tmp)
		result = fmt.Sprintf(returnFmt, statusSuccess, ret)
	}
	if existCallback {
		result = callback[0] + "(" + result + ")"
	}
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, result)
}

func getWorld(w http.ResponseWriter, req *http.Request) {
	var result string
	u, err := myURL.Parse(req.RequestURI)
	fmt.Println("~~~~~~~~~~~~")
	fmt.Println(req.RequestURI)
	if err == nil {
		log.Println()
	}
	query := u.Query()
	fmt.Println("!!!!!!!!!!!!!!!!!")
	fmt.Print(query)
	callback, existCallback := query["callback"]
	tmp, err := json.Marshal(getW())
	if err != nil {
		log.Println(err)
		result = fmt.Sprintf(returnFmt, statusFailed, "get world failed")
	} else {
		ret := string(tmp)
		result = fmt.Sprintf(returnFmt, statusSuccess, ret)
	}
	if existCallback {
		result = callback[0] + "(" + result + ")"
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Println(result)
	io.WriteString(w, result)
}

func socketConnect(so socketio.Socket) {
	log.Println("socket connected")
	sc := redis.NewTCPClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		DB:       0,
		PoolSize: 100,
	}).PubSub()
	so.Emit("open", "success")
	so.On(signinChannel, func() {
		so.Join(signinChannel)
		log.Println("socket", signinChannel)
		defer sc.Close()
		sc.Subscribe(signinChannel)
		for {
			n, _ := sc.Receive()
			switch n.(type) {
			case *redis.Message:
				t := n.(*redis.Message)
				so.BroadcastTo(signinChannel, "data", t.Payload)
			}
		}
	})
	so.On(msgChannel, func() {
		so.Join(msgChannel)
		log.Println("socket", msgChannel)
		defer sc.Close()
		sc.Subscribe(msgChannel)
		for {
			n, _ := sc.Receive()
			switch n.(type) {
			case *redis.Message:
				t := n.(*redis.Message)
				so.BroadcastTo(msgChannel, "data", t.Payload)
			}
		}
	})
	so.On("disconnection", func() {
		log.Println("socket disconnected")
		sc.Close()
	})
}

func publish(channel, message string) {
	redisClient := redis.NewTCPClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		DB:       0,
		PoolSize: 100,
	})
	defer redisClient.Close()
	_ = redisClient.Publish(channel, message)
}

func socketError(so socketio.Socket, err error) {
	log.Println("socketError", err)
}

func socketServer() *socketio.Server {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.On("connection", socketConnect)
	server.On("error", socketServer)
	return server
}

func main() {
	mongodb.Init("localhost:27017")
	Init()
	http.Handle("/socket.io/", socketServer())
	http.HandleFunc("/nM", newMessage)
	http.HandleFunc("/nS", newSignin)
	http.HandleFunc("/gM", getMessage)
	http.HandleFunc("/gS", getSignin)
	http.HandleFunc("/gC", getChina)
	http.HandleFunc("/gW", getWorld)
	http.HandleFunc("/gT", getTop)
	log.Println("listening")
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
