package services

import (
	"github.com/dzikrisyafi/kursusvirtual_quiz-api/src/domain/answers"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/date_utils"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
)

var (
	AnswerService answerServiceInterface = &answerService{}
)

type answerService struct{}

type answerServiceInterface interface {
	CreateAnswer(answers.Answer) (*answers.Answer, rest_errors.RestErr)
	GetUserAnswer(int, int) (answers.Answers, rest_errors.RestErr)
}

func (s *answerService) CreateAnswer(answer answers.Answer) (*answers.Answer, rest_errors.RestErr) {
	var isRight int
	if answer.IsRight {
		isRight = 1
	} else {
		isRight = 0
	}

	if err := answer.Validate(isRight); err != nil {
		return nil, err
	}

	answer.AnswerTime = date_utils.GetNowDBFormat()
	if err := answer.Save(isRight); err != nil {
		return nil, err
	}
	return &answer, nil
}

func (s *answerService) GetUserAnswer(userID int, sectionID int) (answers.Answers, rest_errors.RestErr) {
	dao := &answers.Answer{UserID: userID}
	return dao.Get(sectionID)
}
