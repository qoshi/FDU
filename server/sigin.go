package main

import (
	"linkedin/service/mongodb"
	"time"

	mgo "gopkg.in/mgo.v2"
)

type Signin struct {
	Name     string
	Location int
	Time     time.Time
	Token    string
	Ip       string
}

func NewSignin(name string, location int, ip string) error {
	signin := Signin{
		Name:     name,
		Location: location,
		Time:     time.Now(),
	}
	mongodb.Exec("signin", func(c *mgo.Collection) error {
		return c.Insert(&signin)
	})
	return nil
}

func GetSignin(start, count int) []Msg {
	var result []Msg
	mongodb.Exec("signin", func(c *mgo.Collection) error {
		return c.Find(nil).Skip(start).Limit(count).All(&result)
	})
	return result
}
