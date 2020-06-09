package choices

import (
	"strings"

	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
	"golang.org/x/net/html"
)

type Choice struct {
	ID         int64  `json:"id"`
	Choice     string `json:"choice"`
	IsRight    bool   `json:"is_right"`
	QuestionID int64  `json:"question_id"`
}

type Choices []Choice

func (choice Choice) Validate(isRight int) rest_errors.RestErr {
	choice.Choice = html.EscapeString(strings.TrimSpace(choice.Choice))
	if choice.Choice == "" {
		return rest_errors.NewBadRequestError("invalid choice")
	}

	if isRight < 0 || isRight > 1 {
		return rest_errors.NewBadRequestError("invalid status")
	}

	if choice.QuestionID <= 0 {
		return rest_errors.NewBadRequestError("invalid question id")
	}

	return nil
}
