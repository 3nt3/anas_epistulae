package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/nielsdingsbums/anas_epistulae/db"
	"github.com/nielsdingsbums/anas_epistulae/funcs"
	"github.com/nielsdingsbums/anas_epistulae/structs"
	"log"
	"net/http"
	"reflect"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Print("[+] Briefente - Anas Epistulae\n")
	log.Panicf("[-] %v\n", http.ListenAndServe(":8080", nil))
}


// handles stuff
func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("[-] %v\n", err)
	}

	ip := conn.RemoteAddr().String()
	fmt.Printf("[+] New connection established! ip %v\n", ip)

	// establish channel
	msgs := make(chan structs.WsRequest)
	go read(msgs, conn)

	go checkMsgs(conn)

	// send initial info
	conn.WriteJSON(structs.WsRequest{"status", funcs.GetDB()})

	for {
		// initialize emoty response
		cresp := &structs.WsRequest{}

		// wait for user input
		creq := <- msgs

		switch creq.Type {
		case "messageCreate":
			funcs.MessageCreate(creq)
		}

		if cresp.Type != "" {
			conn.WriteJSON(cresp)
		}
	}
}

// reads thingz
func read(msgs chan structs.WsRequest, conn *websocket.Conn) {
	for {
		creq := &structs.WsRequest{}
		if err := conn.ReadJSON(creq); err != nil {
			fmt.Printf("[-] reading error: %v\n", err)
			conn.Close()
			return
		}
		msgs <- *creq
	}
}

// wait for update
func checkMsgs(conn *websocket.Conn) {
	var oldState []structs.Message
	for {
		currentState := db.Messages
		if !reflect.DeepEqual(oldState, currentState) {
			conn.WriteJSON(structs.WsRequest{"status", currentState})
		}
		oldState = currentState
	}
}