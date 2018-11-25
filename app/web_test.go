// web_test.go (web-server)

// +build integration

package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"
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
	return []byte(`{"FName":"` + t.FName + `","LName":"` + t.LName + `","Number":` + strconv.Itoa(t.Number) + `}`)
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
	rand.Seed(time.Now().UTC().UnixNano())
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
	//make array
	//loop over array and push the item
	//iterate over items and pop items to check equality
	//test pass

	things := makeThings()
	things = things[:900]
	log.Println("Stack Push " + strconv.Itoa(len(things)) + " things")
	client := &http.Client{}
	for _, tg := range things {
		res1, err := client.Post(appUrl[0]+"push", "application/json", bytes.NewReader(tg.toBytes()))
		defer res1.Body.Close()
		if err != nil {
			log.Println(err)
			t.FailNow()
		}
		res2, err := client.Post(appUrl[1]+"push", "application/json", bytes.NewReader(tg.toBytes()))
		defer res2.Body.Close()
		if err != nil {
			log.Println(err)
			t.FailNow()
		}
	}

	log.Println("Pop things to check equallity")
	for i, _ := range things {
		resp1, err := client.Get(appUrl[0] + "pop")
		resp2, err := client.Get(appUrl[1] + "pop")
		defer resp1.Body.Close()
		defer resp2.Body.Close()
		var t1 *thing
		var t2 *thing
		err = json.NewDecoder(resp1.Body).Decode(&t1)
		if err != nil {
			log.Println(err)
		}
		err = json.NewDecoder(resp2.Body).Decode(&t2)
		if err != nil {
			log.Println(err)
		}
		log.Println(*t1, *t2, "==", things[len(things)-i-1])
		if *t1 != things[len(things)-i-1] || *t2 != things[len(things)-i-1] {
			t.FailNow()
		}
	}

}

func TestQueue(t *testing.T) {
	//make array
	//loop over array and push the item
	//iterate over items and pop items to check equality
	//test pass

	things := makeThings()
	things = things[:900]
	log.Println("Queue Push " + strconv.Itoa(len(things)) + " things")
	client := &http.Client{}
	for ij, tg := range things {
		res1, err := client.Post(appUrl[2]+"push", "application/json", bytes.NewReader(tg.toBytes()))
		defer res1.Body.Close()
		if err != nil {
			log.Println(err)
			t.FailNow()
		}
		res2, err := client.Post(appUrl[3]+"push", "application/json", bytes.NewReader(tg.toBytes()))
		defer res2.Body.Close()
		if err != nil {
			log.Println(err)
			t.FailNow()
		}
	}

	log.Println("Pop things to check equallity")
	for i, _ := range things {
		resp1, err := client.Get(appUrl[2] + "pop")
		resp2, err := client.Get(appUrl[3] + "pop")
		defer resp1.Body.Close()
		defer resp2.Body.Close()
		var t1 *thing
		var t2 *thing
		err = json.NewDecoder(resp1.Body).Decode(&t1)
		if err != nil {
			log.Println(err)
		}
		err = json.NewDecoder(resp2.Body).Decode(&t2)
		if err != nil {
			log.Println(err)
		}
		log.Println(*t1, *t2, "==", things[i])
		if *t1 != things[i] || *t2 != things[i] {
			t.FailNow()
		}
	}

}
