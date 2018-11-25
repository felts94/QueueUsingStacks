package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/felts94/QueueUsingStacks/queues"
	"github.com/felts94/QueueUsingStacks/stacks"
	"github.com/gorilla/mux"
)

var mystack stacks.Stack
var myqueue queues.Queue

func PushToStack(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t interface{}
	err := decoder.Decode(&t)
	if err != nil {
		log.Println(err)
		return
	}
	mystack.Push(t)
	//json.NewEncoder(w).Encode(mystack)
}

func PopFromStack(w http.ResponseWriter, r *http.Request) {
	item := mystack.Pop()
	if item == nil {
		log.Println("Nothing in stack")
	}
	json.NewEncoder(w).Encode(item)
}

func PushToQueue(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t interface{}
	err := decoder.Decode(&t)
	if err != nil {
		log.Println(err)
		return
	}
	myqueue.Enqueue(t)
	//json.NewEncoder(w).Encode(myqueue)
}

func PopFromQueue(w http.ResponseWriter, r *http.Request) {
	item := myqueue.Dequeue()
	if item == nil {
		log.Println("Nothing in queue")
		w.Header().Set("Status Code", "404")
	}
	json.NewEncoder(w).Encode(item)
}

func main() {
	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "8180"
	}

	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ok")
	})
	router.HandleFunc("/stack/push", PushToStack).Methods("POST")
	router.HandleFunc("/stack/pop", PopFromStack).Methods("GET")
	router.HandleFunc("/queue/push", PushToQueue).Methods("POST")
	router.HandleFunc("/queue/pop", PopFromQueue).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+PORT, router))
}
