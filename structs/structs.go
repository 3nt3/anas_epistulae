package structs

type WSEvent struct {
	Type string      `json:"type"`
	Data map[string]interface{} `json:"data"`
}

type Message struct {
	ID     int    `json:"id"`
	Text   string `json:"text"`
	Author string `json:"author"`
}
