// web_test.go (web-server)

// +build integration

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"
)

var appUrl = []string{
	"http://localhost:8180/stack/",
	"http://localhost:8280/stack/",
	"http://localhost:8180/queue/",
	"http://localhost:8280/queue/",
}

type Body interface{}

type thing struct {
	FName  string `json:"FName"`
	LName  string `json:"LName"`
	Number int    `json:"Number"`
}

func (t thing) toBytes() []byte {
	//log.Println(t)
	return []byte(`{"FName":"` + t.FName + `",LName":"` + t.LName + `","Number":` + strconv.Itoa(t.Number) + `}`)
}

func readFileWithReadString(fn string) []string {

	file, err := os.Open(fn)
	defer file.Close()

	if err != nil {
		log.Println(err)
	}

	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)
	var lines = make([]string, 1)
	var line string
	line, err = reader.ReadString('\n')
	for {
		line, err = reader.ReadString('\n')
		if err != nil {
			break
		}
		lines = append(lines, line[:len(line)-1])
	}

	if err != io.EOF {
		fmt.Printf(" > Failed!: %v\n", err)
	}

	return lines
}

func makeThings() []thing {
	var things []thing
	names := readFileWithReadString("names")
	for i, name := range names {
		num := rand.Int31()
		if name != "" {
			t := thing{
				FName:  name,
				Number: int(num) % 100,
				LName:  names[(i+int(num))%100],
			}
			things = append(things, t)
		}
	}
	return things
}

func TestStack(t *testing.T) {
	var err error
	//push array of items
	//make array
	things := makeThings()
	//log.Println(string(things[1].toBytes()))
	//loop over array and push the item
	//iterate over items and pop items to check equality
	//test pass
	//stackpush
	log.Println("Stack Push " + strconv.Itoa(len(things)) + " things")
	client := &http.Client{}
	for _, tg := range things {
		_, err = client.Post(appUrl[0]+"push", "application/json", strings.NewReader(string(tg.toBytes())))
		_, err = client.Post(appUrl[1]+"push", "application/json", strings.NewReader(string(tg.toBytes())))
		if err != nil {
			t.FailNow()
		}
	}

	log.Println("Pop things to check equallity")
	for i, _ := range things {
		resp1, err := client.Get(appUrl[0] + "pop")
		resp2, err := client.Get(appUrl[1] + "pop")
		//log.Println(resp, err)
		defer resp1.Body.Close()
		defer resp2.Body.Close()
		t1 := new(thing)
		t2 := new(thing)
		log.Println("decode t1", &t1)
		err = json.NewDecoder(resp1.Body).Decode(&t1)
		if err != nil {
			log.Println(err)
		}
		log.Println("decode t2", &t2)
		err = json.NewDecoder(resp2.Body).Decode(&t2)
		if err != nil {
			log.Println(err)
		}
		log.Println(t1, t2, *t1, *t2, things[len(things)-i-1])
		if *t1 != things[len(things)-i-1] || *t2 != things[len(things)-i-1] {
			t.FailNow()
		}
	}

	// err = json.NewDecoder(resp.Body).Decode(&target)
	// //log.Println(target, stack)
	// lastIn := target[0]

	// //stack pop
	// resp, err = client.Get(appUrl[0] + "pop")
	// //log.Println(resp, err)
	// defer resp.Body.Close()
	// targ := new(map[string]interface{})

	// err = json.NewDecoder(resp.Body).Decode(&targ)
	// if err != nil {
	// 	log.Println(err)
	// }
	// // m1 and m2 are the maps we want to compare
	// eq := reflect.DeepEqual(*targ, lastIn)
	// if eq {
	// 	fmt.Println("They're equal.")
	// } else {
	// 	fmt.Println("They're unequal.")
	// }

}

func pushArrayOfItems() {

}

func popItem() {

}

func TestStackPop(t *testing.T) {

}

// func PushToStackAPI(url string, item interface{}) *http.Response {
// 	c := &http.Client{}
// 	stackBase := sling.New().Base(stackUrl).Client(c)
// 	path := "push"

// 	body := &PushRequest{
// 		Name: "Kyle",
// 		Attr: []Attributes{
// 			{
// 				Key:   "id",
// 				Value: "ie69",
// 			},
// 		},
// 	}
// 	req, err := stackBase.New().Post(path).BodyJSON(body).Request()
// 	if err != nil {
// 		log.Println(err)
// 		t.Fail()
// 	}
// 	log.Println(req)
// }

// func PopFromStack() {

// }
