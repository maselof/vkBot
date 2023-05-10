package botsLP

type UpdateResponse struct {
	Ts      string   `json:"ts"`
	Updates []Update `json:"updates"`
}

type HistoryResponse struct {
	Response *History `json:"response"`
}

type History struct {
	Items []Item `json:"items"`
}

type Item struct {
	Text string `json:"text"`
}

type Update struct {
	Type   string  `json:"type"`
	Object *Object `json:"object"`
}

type Object struct {
	Message *IncomingMessage `json:"message"`
}

type IncomingMessage struct {
	Text   string `json:"text"`
	FromId int    `json:"from_id"`
}
