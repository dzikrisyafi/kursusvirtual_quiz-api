package services

import (
	"github.com/dzikrisyafi/kursusvirtual_quiz-api/src/domain/quiz"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
)

var (
	QuizService quizServiceInterface = &quizService{}
)

type quizService struct{}

type quizServiceInterface interface {
	CreateQuiz(quiz.Quiz) (*quiz.Quiz, rest_errors.RestErr)
	GetQuiz(int) (*quiz.Quiz, rest_errors.RestErr)
	GetAllQuiz() (quiz.Quizs, rest_errors.RestErr)
	GetAllQuestionByActivityID(int) (quiz.QuizsAndChoices, rest_errors.RestErr)
	GetAllChoiceByQuestionID(quiz *quiz.QuizAndChoice) rest_errors.RestErr
	UpdateQuiz(bool, quiz.Quiz) (*quiz.Quiz, rest_errors.RestErr)
	DeleteQuiz(int) rest_errors.RestErr
	DeleteQuestionByCourseID(int) rest_errors.RestErr
}

func (s *quizService) CreateQuiz(quiz quiz.Quiz) (*quiz.Quiz, rest_errors.RestErr) {
	var isActive int
	if quiz.IsActive {
		isActive = 1
	} else {
		isActive = 0
	}

	if err := quiz.Validate(isActive); err != nil {
		return nil, err
	}

	if err := quiz.Save(isActive); err != nil {
		return nil, err
	}
	return &quiz, nil
}

func (s *quizService) GetQuiz(quizID int) (*quiz.Quiz, rest_errors.RestErr) {
	result := &quiz.Quiz{ID: quizID}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *quizService) GetAllQuiz() (quiz.Quizs, rest_errors.RestErr) {
	dao := &quiz.Quiz{}
	return dao.GetAllQuiz()
}

func (s *quizService) GetAllQuestionByActivityID(activityID int) (quiz.QuizsAndChoices, rest_errors.RestErr) {
	dao := &quiz.QuizAndChoice{ActivityID: activityID}
	allQuestion, err := dao.GetAllQuestionByActivityID()
	if err != nil {
		return nil, err
	}

	results := make([]quiz.QuizAndChoice, 0)
	for _, question := range allQuestion {
		if err := s.GetAllChoiceByQuestionID(&question); err != nil {
			return nil, err
		}
		results = append(results, question)
	}

	return results, nil
}

func (s *quizService) GetAllChoiceByQuestionID(question *quiz.QuizAndChoice) rest_errors.RestErr {
	dao := &quiz.Choice{}
	return dao.GetAllChoiceByQuestionID(question)
}

func (s *quizService) UpdateQuiz(isPartial bool, quiz quiz.Quiz) (*quiz.Quiz, rest_errors.RestErr) {
	current, err := s.GetQuiz(quiz.ID)
	if err != nil {
		return nil, err
	}

	var isActive int
	if quiz.IsActive {
		isActive = 1
	} else {
		isActive = 0
	}

	if isPartial {
		if quiz.Question != "" {
			current.Question = quiz.Question
		}

		if isActive >= 0 {
			current.IsActive = quiz.IsActive
		}
	} else {
		if err := quiz.Validate(isActive); err != nil {
			return nil, err
		}

		current.Question = quiz.Question
		current.IsActive = quiz.IsActive
		current.ActivityID = quiz.ActivityID
	}

	if err := current.Update(isActive); err != nil {
		return nil, err
	}

	return current, nil
}

func (s *quizService) DeleteQuiz(quizID int) rest_errors.RestErr {
	dao := &quiz.Quiz{ID: quizID}
	return dao.Delete()
}

func (s *quizService) DeleteQuestionByCourseID(courseID int) rest_errors.RestErr {
	dao := &quiz.Quiz{CourseID: courseID}
	return dao.DeleteQuestionByCourseID()
}
