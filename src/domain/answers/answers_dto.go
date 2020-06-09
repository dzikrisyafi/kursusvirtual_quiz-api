package answers

import "github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"

type Answer struct {
	ID         int64  `json:"id"`
	UserID     int64  `json:"user_id"`
	QuestionID int64  `json:"question_id"`
	ChoiceID   int64  `json:"choice_id"`
	IsRight    bool   `json:"is_right"`
	AnswerTime string `json:"answer_time"`
}

type Answers []Answer

func (answer Answer) Validate(isRight int) rest_errors.RestErr {
	if answer.UserID <= 0 {
		return rest_errors.NewBadRequestError("invalid user id")
	}

	if answer.QuestionID <= 0 {
		return rest_errors.NewBadRequestError("invalid question id")
	}

	if answer.ChoiceID <= 0 {
		return rest_errors.NewBadRequestError("invalid choice id")
	}

	if isRight < 0 || isRight > 1 {
		return rest_errors.NewBadRequestError("invalid status value")
	}
	return nil
}
