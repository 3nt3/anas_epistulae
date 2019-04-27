package funcs

import (
	"fmt"
	"github.com/NielsDingsbums/anas_epistulae/db"
	"github.com/NielsDingsbums/anas_epistulae/structs"
)

func HandleEvent(event structs.WSEvent) {
	fmt.Printf("[*] new event: %+v\n", event)
}


func GetMessages() []structs.Message {
	return db.Messages
}