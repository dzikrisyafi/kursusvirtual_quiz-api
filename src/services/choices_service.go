package services

import (
	"github.com/dzikrisyafi/kursusvirtual_quiz-api/src/domain/choices"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
)

var (
	ChoiceService choiceServiceInterface = &choiceService{}
)

type choiceService struct{}

type choiceServiceInterface interface {
	CreateChoice(choices.Choice) (*choices.Choice, rest_errors.RestErr)
	GetChoice(int64) (*choices.Choice, rest_errors.RestErr)
	UpdateChoice(bool, choices.Choice) (*choices.Choice, rest_errors.RestErr)
	DeleteChoice(int64) rest_errors.RestErr
}

func (s *choiceService) CreateChoice(choice choices.Choice) (*choices.Choice, rest_errors.RestErr) {
	var isRight int
	if choice.IsRight {
		isRight = 1
	} else {
		isRight = 0
	}

	if err := choice.Validate(isRight); err != nil {
		return nil, err
	}

	if err := choice.Save(isRight); err != nil {
		return nil, err
	}

	return &choice, nil
}

func (s *choiceService) GetChoice(choiceID int64) (*choices.Choice, rest_errors.RestErr) {
	result := &choices.Choice{ID: choiceID}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *choiceService) UpdateChoice(isPartial bool, choice choices.Choice) (*choices.Choice, rest_errors.RestErr) {
	current, err := s.GetChoice(choice.ID)
	if err != nil {
		return nil, err
	}

	var isRight int
	if choice.IsRight {
		isRight = 1
	} else {
		isRight = 0
	}

	if isPartial {
		if choice.Choice != "" {
			current.Choice = choice.Choice
		}

		if choice.IsRight == true || choice.IsRight == false {
			current.IsRight = choice.IsRight
		}
	} else {
		if err := choice.Validate(isRight); err != nil {
			return nil, err
		}

		current.Choice = choice.Choice
		current.IsRight = choice.IsRight
		current.QuestionID = choice.QuestionID
	}

	if err := current.Update(isRight); err != nil {
		return nil, err
	}

	return current, nil
}

func (s *choiceService) DeleteChoice(choiceID int64) rest_errors.RestErr {
	dao := &choices.Choice{ID: choiceID}
	return dao.Delete()
}
