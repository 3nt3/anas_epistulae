package main

import (
	"fmt"
	"github.com/NielsDingsbums/anas_epistulae/funcs"
	"github.com/NielsDingsbums/anas_epistulae/structs"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	fmt.Printf("-=-=-= Anas Epistulae =-=-=-\n")
	fmt.Printf("[*] Starting server\n")

	http.HandleFunc("/", handler)
	fmt.Println(http.ListenAndServe(":8008", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("[+] Connection established to %v\n", conn.RemoteAddr())

	events := make(chan structs.WSEvent)
	go read(conn, events)

	initial := true
	go checkUpdate(conn, &initial)

	for {
		// wait for new event
		event := <-events

		// do stuff with it
		funcs.HandleEvent(event)
	}
}

func read(conn *websocket.Conn, events chan structs.WSEvent) {
	for {
		// wait for new event
		event := &structs.WSEvent{}
		err := conn.ReadJSON(event)

		// error handling
		if err != nil {
			fmt.Printf("[-] Reading error: %v\n", err)
			_ = conn.Close()
			fmt.Printf("[*] connection closed.\n")
			return
		}

		// send to channel
		events <- *event
	}
}

func checkUpdate(conn *websocket.Conn, initial *bool) {
	// Get initial messages
	oldMessages := funcs.GetMessages()
	for {
		time.Sleep(500)
		// Get current messages
		newMessages := funcs.GetMessages()

		// If the length varies, send update (new messages) to client
		if len(oldMessages) != len(newMessages) || *initial {

			// create event
			event := &structs.WSEvent{
				Type: "update",
				Data: map[string]interface{}{"msgs": newMessages},
			}

			// send event
			err := conn.WriteJSON(event)
			if err != nil {
				fmt.Printf("[-] send message error: %v\n", err)
			}
		}

		oldMessages = newMessages
		*initial = false
	}
}
