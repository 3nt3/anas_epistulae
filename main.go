package main

import (
	"fmt"
	"github.com/NielsDingsbums/anas_epistulae/funcs"
	"github.com/NielsDingsbums/anas_epistulae/structs"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
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
	go checkUpdate(conn)

	for {
		// wait for new event
		event := <- events

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
			_ := conn.Close()
			fmt.Printf("[*] connection closed.\n")
			return
		}

		// send to channel
		events <- *event
	}
}


func checkUpdate(conn *websocket.Conn) {
	// Get initial messages
	oldMessages := funcs.GetMessages()
	for {
		// Get current messages
		newMessages := funcs.GetMessages()

		// If the length varies, send update (new messages) to client
		if len(oldMessages) != len(newMessages) {
			err := conn.WriteJSON(newMessages)
			if err != nil {
				fmt.Printf("[-] send message error: %v\n", err)
			}
		}
	}
}