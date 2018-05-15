package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Message struct {
	Content string `json: "content"`
}

var exitMessage = "/exit"

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

	if message.Content == exitMessage {
		close(ch)
		fmt.Println("Connection closed. \n\nTerminating program...")
		os.Exit(0)
	}

	ch <- message

	fmt.Fprintf(w, "%v", "Message sent")
}

func listener(ch chan Message) {

	for {

		resp, ok := <-ch

		if !ok {
			break
		}

		msg := Message(resp)
		fmt.Println(msg.Content)
	}
}
