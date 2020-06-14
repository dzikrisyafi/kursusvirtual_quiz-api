package quiz

import (
	"html"
	"strings"

	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
)

type Quiz struct {
	ID         int    `json:"id"`
	Question   string `json:"question"`
	IsActive   bool   `json:"is_active"`
	ActivityID int    `json:"activity_id"`
}

type Quizs []Quiz

func (quiz Quiz) Validate(isActive int) rest_errors.RestErr {
	quiz.Question = html.EscapeString(strings.TrimSpace(quiz.Question))
	if quiz.Question == "" {
		return rest_errors.NewBadRequestError("invalid question name")
	}

	if isActive < 0 || isActive > 1 {
		return rest_errors.NewBadRequestError("invalid status")
	}

	if quiz.ActivityID <= 0 {
		return rest_errors.NewBadRequestError("invalid activity id")
	}

	return nil
}
