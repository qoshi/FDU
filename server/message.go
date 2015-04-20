package main

import (
	"linkedin/service/mongodb"

	"gopkg.in/mgo.v2"
)

type Msg struct {
	Name    string
	Country string
	Context string
}

func NewMsg(name string, country string, context string) error {
	msg := Msg{
		Name:    name,
		Country: country,
		Context: context,
	}
	mongodb.Exec("message", func(c *mgo.Collection) error {
		return c.Insert(&msg)
	})
	return nil
}

func getMsg(start, count int) []Msg {
	var result []Msg
	mongodb.Exec("message", func(c *mgo.Collection) error {
		return c.Find(nil).Skip(start).Limit(count).All(&result)
	})
	return nil
}
