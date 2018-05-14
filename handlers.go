package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Message struct {
	Content string `json: "content"`
}

func sendHandler(w http.ResponseWriter, r *http.Request, ch chan Message) {

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var message Message

	if err = json.Unmarshal(b, &message); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ch <- message

	fmt.Fprintf(w, "%v", "Message sent")
}

func receiveHandler(w http.ResponseWriter, r *http.Request, ch chan Message) {

	var (
		msgs       []string
		noMessages bool = false
	)

	for {

		if noMessages {
			break
		}

		fmt.Println("Loop start")

		select {

		case resp := <-ch:

			msg := Message(resp)
			msgs = append(msgs, msg.Content)

		default:
			noMessages = true
		}
	}

	fmt.Fprintf(w, "%v", strings.Join(msgs, "\n"))

	return
}
