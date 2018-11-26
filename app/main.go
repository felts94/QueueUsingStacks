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

type Status struct {
	Length    int           `json:"length,omitempty"`
	NextThree []interface{} `json:"next_three,omitempty"`
	LastThree []interface{} `json:"last_three,omitempty"`
}

func StackStatus(w http.ResponseWriter, r *http.Request) {
	if len(mystack) > 3 {
		json.NewEncoder(w).Encode(Status{
			Length:    len(mystack),
			NextThree: mystack[0:3],
			LastThree: mystack[len(mystack)-4:],
		})
	} else {
		json.NewEncoder(w).Encode(Status{
			Length:    len(mystack),
			NextThree: mystack,
		})
	}
}

func QueueStatus(w http.ResponseWriter, r *http.Request) {
	if len(myqueue.S) > 3 && !myqueue.Inverted {
		json.NewEncoder(w).Encode(Status{
			Length:    len(myqueue.S),
			NextThree: myqueue.S[0:3],
			LastThree: myqueue.S[len(myqueue.S)-4:],
		})
	} else if len(myqueue.S) > 3 {
		json.NewEncoder(w).Encode(Status{
			Length:    len(myqueue.S),
			NextThree: myqueue.S[len(myqueue.S)-4:],
			LastThree: myqueue.S[0:3],
		})
	} else if !myqueue.Inverted {
		//invert queue
		var item []int
		item = append(item, 1)

		myqueue.Enqueue(item)
		_ = myqueue.Dequeue()

		json.NewEncoder(w).Encode(Status{
			Length:    len(myqueue.S),
			NextThree: myqueue.S,
		})
	} else {
		json.NewEncoder(w).Encode(Status{
			Length:    len(myqueue.S),
			NextThree: myqueue.S,
		})
	}
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
	//stack endpoints
	router.HandleFunc("/stack/push", PushToStack).Methods("POST")
	router.HandleFunc("/stack/pop", PopFromStack).Methods("GET")
	router.HandleFunc("/stack/status", StackStatus).Methods("GET")

	//queue endpoints
	router.HandleFunc("/queue/push", PushToQueue).Methods("POST")
	router.HandleFunc("/queue/pop", PopFromQueue).Methods("GET")
	router.HandleFunc("/queue/status", QueueStatus).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+PORT, router))
}
