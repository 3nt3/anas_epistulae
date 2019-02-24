package structs

import "time"

type WsRequest struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type Message struct {
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	Author    Author    `json:"author"`
}

type Author struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
