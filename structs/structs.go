package structs

type WSEvent struct {
	Type string `json:"type"`
	Data interface{} `json:"data"`
}
