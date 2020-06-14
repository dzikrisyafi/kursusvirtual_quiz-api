package quiz

type QuizAndChoice struct {
	ID         int      `json:"id"`
	Question   string   `json:"question"`
	ActivityID int      `json:"activity_id"`
	Choices    []Choice `json:"choices"`
}

type QuizsAndChoices []QuizAndChoice

type Choice struct {
	ID      int    `json:"id"`
	Choice  string `json:"choice"`
	IsRight bool   `json:"is_right"`
}
