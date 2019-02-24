package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/nielsdingsbums/anas_epistulae/structs"
	"net/http"
)



var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {

}

func handler(w http.ResponseWriter, r *http.Request)  {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("[-] %v\n", err)
	}

	msgs := make(chan structs.WsRequest)

	go read(msgs, conn)

	for {
		// initialize emoty response
		cresp := &structs.WsRequest{}

		// wait for user input
		creq := <- msgs

		switch creq.Type {
		case "hello":

		}
	}
}

func read(msgs chan structs.WsRequest, conn *websocket.Conn) {
	for {
		creq := &structs.WsRequest{}
		if err := conn.ReadJSON(creq); err != nil {
			fmt.Printf("[-] reading error: %v\n", err)
			continue
		}
		msgs <- *creq
	}
}

func broadcast() {

}