package log

type Log struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	Message string `json:"message"`
}
