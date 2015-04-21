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
}

func NewMsg(name string, country int, context string) error {
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
	mongodb.Exec("message", func(c *mgo.Collection) error {
		return c.Find(nil).Skip(start).Limit(count).All(&result)
	})
	return nil
}
