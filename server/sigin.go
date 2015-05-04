package main

import (
	"bufio"
	"fmt"
	"linkedin/service/mongodb"
	"log"
	"time"

	"os"

	"strconv"
	"strings"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Signin struct {
	Name     string
	Location int
	Time     time.Time
	Token    string
	Ip       string
}

type Area struct {
	Name   string
	Count  int
	Total  int
	Parent int
}

type Location struct {
	Name   string
	Total  int
	Count  int
	Parent int
}

type Result struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

var areas []Area
var locations []Location
var total int
var china []Area
var nameMap map[string]string

func NewSignin(name string, location int, ip string) error {
	// dup := Signin{}
	// find := mongodb.Exec("message", func(c *mgo.Collection) error {
	// 	t := c.Find(bson.M{"ip": ip})
	// 	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	// 	fmt.Print(t)
	// 	fmt.Println(dup)
	// 	return c.Find(bson.M{"ip": ip}).One(&dup)
	// })
	// if find {
	// 	return errors.New("一个IP只能提交一次哦^^")
	// }
	signin := Signin{
		Name:     name,
		Location: location,
		Time:     time.Now(),
	}
	mongodb.Exec("signin", func(c *mgo.Collection) error {
		return c.Insert(&signin)
	})
	area := locations[location].Parent
	areas[area].Count += 1
	total += 1
	return nil
}

func GetSignin(start, count int) []Msg {
	var result []Msg
	mongodb.Exec("signin", func(c *mgo.Collection) error {
		return c.Find(nil).Skip(start).Limit(count).All(&result)
	})
	return result
}

func getLocationName(location int) string {
	return locations[location].Name
}

func getAreaName(area int) string {
	return areas[area].Name
}

func getT() interface{} {
	top := make([]int, 3)
	top[0] = 0
	top[1] = 1
	top[2] = 2
	for i := range areas {
		if areas[top[0]].Count < areas[i].Count {
			top[0] = i
		}
	}
	for i := range areas {
		if areas[top[1]].Count < areas[i].Count && top[0] != i && areas[i].Count <= areas[top[0]].Count {
			top[1] = i
		}
	}
	for i := range areas {
		if areas[top[1]].Count < areas[i].Count && top[1] != i && top[0] != i && areas[i].Count <= areas[top[0]].Count {
			top[2] = i
		}
	}
	c := 0
	type re struct {
		Axis  []string
		Count []int
	}
	r := re{}
	for j := range top {
		r.Axis = append(r.Axis, areas[top[j]].Name)
		r.Count = append(r.Count, areas[top[j]].Count)
		c += areas[top[j]].Count
	}
	r.Axis = append(r.Axis, "其他")
	r.Count = append(r.Count, total-c)
	return r
}

func Init() {
	//read range
	var scanner *bufio.Scanner
	// aindex := 0
	fa, err := os.Open("../funny/rangeWorld.txt")
	if err != nil {
		log.Panic(err)
	}
	scanner = bufio.NewScanner(fa)
	for scanner.Scan() {
		tmp := strings.Split(scanner.Text(), " ")
		num, err := strconv.Atoi(tmp[1])
		area := Area{
			Name:  tmp[0],
			Count: 0,
			Total: num,
		}
		areas = append(areas, area)
		if err != nil {
			log.Panic(err)
		}
	}
	//read china
	fc, err := os.Open("../funny/china.txt")
	if err != nil {
		log.Panic(err)
	}
	scanner = bufio.NewScanner(fc)
	for scanner.Scan() {
		tmp := strings.Split(scanner.Text(), " ")
		num, err := strconv.Atoi(tmp[1])
		if err != nil {
			log.Panic(err)
		}
		location := Location{
			Total:  num,
			Count:  0,
			Parent: 0,
			Name:   tmp[0],
		}
		locations = append(locations, location)
	}
	//read world
	fw, err := os.Open("../funny/world.txt")
	if err != nil {
		log.Println(err)
		log.Panic(err)
	}
	scanner = bufio.NewScanner(fw)
	for scanner.Scan() {
		tmp := strings.Split(scanner.Text(), " ")
		num, err := strconv.Atoi(tmp[1])
		if err != nil {
			log.Panic(err)
		}
		location := Location{
			Total:  num,
			Count:  0,
			Parent: 0,
			Name:   tmp[0],
		}
		locations = append(locations, location)
	}
	//read NameMap
	nameMap = make(map[string]string)
	nm, err := os.Open("../funny/mapping.txt")
	scanner = bufio.NewScanner(nm)
	for scanner.Scan() {
		tmp := strings.Split(scanner.Text(), " ")
		nameMap[tmp[0]] = tmp[1]
	}
	for i := range locations {
		mongodb.Exec("signin", func(c *mgo.Collection) error {
			t, err := c.Find(bson.M{"location": i}).Count()
			if err == nil {
				locations[i].Count = t
				areas[locations[i].Parent].Count += t
				total += t
			}
			return err
		})
	}
}

func getW() []Result {
	var r []Result
	l := len(locations)
	for i := 34; i < l; i++ {
		t := Result{
			Name:  nameMap[locations[i].Name],
			Value: float64(areas[locations[i].Parent].Count) / float64(areas[locations[i].Parent].Total),
		}
		if t.Value > 1 {
			t.Value = 1
		}
		r = append(r, t)
	}
	fmt.Println(areas)
	fmt.Println(r)
	return r
}

func getC() []Result {
	var r []Result
	for i := 0; i < 34; i++ {
		t := Result{
			Name:  locations[i].Name,
			Value: float64(locations[i].Count) / float64(locations[i].Total),
		}
		if t.Value > 1 {
			t.Value = 1
		}
		r = append(r, t)
	}
	fmt.Println(r)
	return r
}
