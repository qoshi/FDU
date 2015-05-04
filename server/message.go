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

type Tips struct {
	Totalcount  int64
	count       int64
	countryID   int64
	areaID      int64
	countryName string
}

func NewMsg(name string, country int, context string, ip string) error {
	// dup := Msg{}
	// find := mongodb.Exec("message", func(c *mgo.Collection) error {
	// 	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!")
	// 	fmt.Println(c.Find(bson.M{"ip": ip}))
	// 	return c.Find(bson.M{"ip": ip}).One(&dup)
	// })
	// if find {
	// 	return errors.New("一个IP只能提交一次哦^^")
	// }
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
