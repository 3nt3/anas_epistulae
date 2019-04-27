package funcs

import (
	"fmt"
	"github.com/NielsDingsbums/anas_epistulae/db"
	"github.com/NielsDingsbums/anas_epistulae/structs"
)


// event handler
func HandleEvent(event structs.WSEvent) {
	fmt.Printf("[*] new event: %+v\n", event)

	switch event.Type {
	case "msg":
		newMessage := &structs.Message{}
		newMessage.Text = event.Data["text"].(string)
		newMessage.Author = event.Data["author"].(string)

		CreateMessage(*newMessage)

	}
}


func GetMessages() []structs.Message {
	return db.Messages
}

func CreateMessage(msg structs.Message)  {
	msg.ID = len(db.Messages)
	db.Messages = append(db.Messages, msg)
}