package quiz

type QuizAndChoice struct {
	ID        int64    `json:"id"`
	Question  string   `json:"question"`
	SectionID int64    `json:"section_id"`
	Choices   []Choice `json:"choices"`
}

type QuizsAndChoices []QuizAndChoice

type Choice struct {
	ID      int64  `json:"id"`
	Choice  string `json:"choice"`
	IsRight bool   `json:"is_right"`
}
