package main

import (
	"fmt"
	"github.com/NielsDingsbums/anas_epistulae/structs"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}


func main() {

}


func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		event := &structs.WSEvent{}
		err := conn.ReadJSON(event)
		if err != nil {
			fmt.Printf("[-] reading error: %v\n", err)
			continue
		}

		fmt.Printf("[*] new ws event: %+v", event)
	}

}
