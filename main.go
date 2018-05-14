package main

import (
	"log"
	"net/http"
)

func main() {

	ch := make(chan Message, 100)

	http.HandleFunc("/send", middleWare(sendHandler, ch, "POST"))
	http.HandleFunc("/receive", middleWare(receiveHandler, ch, "GET"))

	log.Println("Listening to localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func middleWare(
	fn func(http.ResponseWriter, *http.Request, chan Message),
	ch chan Message, method string,
) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if method == r.Method {

			// Actual Handler Function
			fn(w, r, ch)

		} else {

			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
