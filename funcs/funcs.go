package funcs

import (
	"github.com/nielsdingsbums/anas_epistulae/db"
	"github.com/nielsdingsbums/anas_epistulae/structs"
	"time"
)

func GetDB() []structs.Message {
	return db.Messages
}

func MessageCreate(request structs.WsRequest) {
	msg := toMessage(request.Data.(map[string]interface{}))

	db.Messages = append(db.Messages, msg)
}

func toMessage(raw map[string]interface{}) structs.Message {
	var msg structs.Message
	msg.ID = len(db.Messages)
	msg.Text = raw["text"].(string)
	msg.CreatedAt = time.Now()
	msg.Author = raw["author"].(string)

	return msg
}