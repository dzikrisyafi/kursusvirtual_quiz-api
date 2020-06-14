package choices

type PublicChoice struct {
	ID         int `json:"id"`
	QuestionID int `json:"question_id"`
}

func (choices Choices) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(choices))
	for index, choice := range choices {
		result[index] = choice.Marshall(isPublic)
	}

	return result
}

func (choice Choice) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicChoice{
			ID:         choice.ID,
			QuestionID: choice.QuestionID,
		}
	}

	return choice
}
