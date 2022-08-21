package models

type Answer struct {
	ID         uint   `json:"id"`
	QuestionID uint   `json:"question_id"`
	Content    string `json:"content"`
	IsProper   bool   `json:"is_proper"`
}

type PublicAnswer struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
}
