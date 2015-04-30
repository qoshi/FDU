package main

import (
	"fmt"
	"linkedin/service/mongodb"

	"time"

	"gopkg.in/mgo.v2"
)

type Msg struct {
	Name     string
	Location int
	Context  string
	Time     time.Time
	Ip       string
}

func NewMsg(name string, country int, context string, ip string) error {
	msg := Msg{
		Name:     name,
		Location: country,
		Context:  context,
		Time:     time.Now(),
	}
	fmt.Println(msg)
	mongodb.Exec("message", func(c *mgo.Collection) error {
		return c.Insert(&msg)
	})
	return nil
}

func GetMsg(start, count int) []Msg {
	var result []Msg
	fmt.Println(start, count)
	mongodb.Exec("message", func(c *mgo.Collection) error {
		return c.Find(nil).Skip(start).Limit(count).All(&result)
	})
	return result
}
