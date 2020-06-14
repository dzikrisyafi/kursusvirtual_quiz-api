package answers

import "encoding/json"

type PublicAnswer struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
}

type PrivateAnswer struct {
	ID         int  `json:"id"`
	UserID     int  `json:"user_id"`
	QuestionID int  `json:"question_id"`
	ChoiceID   int  `json:"choice_id"`
	IsRight    bool `json:"is_right"`
}

func (answers Answers) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(answers))
	for index, answer := range answers {
		result[index] = answer.Marshall(isPublic)
	}
	return result
}

func (answer *Answer) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicAnswer{
			ID:     answer.ID,
			UserID: answer.UserID,
		}
	}

	answerJson, _ := json.Marshal(answer)
	var privateAnswer PrivateAnswer
	json.Unmarshal(answerJson, &privateAnswer)
	return privateAnswer
}
