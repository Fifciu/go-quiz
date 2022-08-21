package models

type Question struct {
	ID      uint   `json:"id"`
	TestID  uint   `json:"test_id"`
	Content string `json:"content"`
}

type PublicQuestion struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
}
