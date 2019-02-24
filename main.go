package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/nielsdingsbums/anas_epistulae/funcs"
	"github.com/nielsdingsbums/anas_epistulae/structs"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Print("[+] Briefente - Anas Epistulae\n")
	fmt.Printf("[-] %v\n", http.ListenAndServe(":8080", nil))
}


// handles stuff
func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("[-] %v\n", err)
	}

	ip := conn.RemoteAddr()
	fmt.Printf("[+] New connection established! ip %v\n", ip)

	// establish channel
	msgs := make(chan structs.WsRequest)
	go read(msgs, conn)

	// send initial info
	conn.WriteJSON(structs.WsRequest{"status", funcs.GetDB()})

	for {
		// initialize emoty response
		cresp := &structs.WsRequest{}

		// wait for user input
		creq := <- msgs

		switch creq.Type {

		}

		conn.WriteJSON(cresp)
	}
}

// reads thingz
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
