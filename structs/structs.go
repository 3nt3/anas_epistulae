package structs

type WsRequest struct {
	Type string `json:"type"`
	Data map[string]interface{} `json:"data"`
}
