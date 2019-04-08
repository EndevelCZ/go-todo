package todo

type Todo struct {
	ID   int64  `json:"id" db:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}
