package quiz

import (
	"errors"

	"github.com/dzikrisyafi/kursusvirtual_quiz-api/src/datasources/mysql/quiz_db"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/logger"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
)

const (
	queryInsertQuiz                 = `INSERT INTO question(question, is_active, activity_id) VALUES(?, ?, ?);`
	queryGetQuiz                    = `SELECT question, is_active, activity_id FROM question WHERE id=?;`
	queryGetAllQuiz                 = `SELECT id, question, is_active, activity_id FROM question;`
	queryGetAllQuestionByActivityID = `SELECT id, question FROM question WHERE activity_id=? AND is_active=1;`
	queryGetAllChoiceByQuestionID   = `SELECT id, choice, is_right FROM choices WHERE question_id=?;`
	queryUpdateQuiz                 = `UPDATE question SET question=?, is_active=?, activity_id=? WHERE id=?;`
	queryDeleteQuiz                 = `DELETE FROM question WHERE id=?;`
	queryDeleteQuizByActivityID     = `DELETE FROM question WHERE activity_id=?;`
)

func (quiz *Quiz) Save(isActive int) rest_errors.RestErr {
	stmt, err := quiz_db.DbConn().Prepare(queryInsertQuiz)
	if err != nil {
		logger.Error("error when trying to prepare save quiz statement", err)
		return rest_errors.NewInternalServerError("error when trying to save quiz", errors.New("database error"))
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(quiz.Question, isActive, quiz.ActivityID)
	if saveErr != nil {
		logger.Error("error when trying to save quiz", saveErr)
		return rest_errors.NewInternalServerError("error when trying to save quiz", errors.New("database error"))
	}

	quizID, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert id after creating a new quiz", err)
		return rest_errors.NewInternalServerError("error when trying to save quiz", errors.New("database error"))
	}
	quiz.ID = int(quizID)

	return nil
}

func (quiz *Quiz) Get() rest_errors.RestErr {
	stmt, err := quiz_db.DbConn().Prepare(queryGetQuiz)
	if err != nil {
		logger.Error("error when trying to prepare get quiz statement by id", err)
		return rest_errors.NewInternalServerError("error when trying to get quiz", errors.New("database error"))
	}
	defer stmt.Close()

	var isactive int
	result := stmt.QueryRow(quiz.ID)
	if getErr := result.Scan(&quiz.Question, &isactive, &quiz.ActivityID); getErr != nil {
		logger.Error("error when trying to get quiz by id", getErr)
		return rest_errors.NewInternalServerError("error when trying to get quiz", errors.New("database error"))
	}

	if isactive == 1 {
		quiz.IsActive = true
	} else {
		quiz.IsActive = false
	}

	return nil
}

func (quiz *Quiz) GetAllQuiz() ([]Quiz, rest_errors.RestErr) {
	stmt, err := quiz_db.DbConn().Prepare(queryGetAllQuiz)
	if err != nil {
		logger.Error("error when trying to prepare get all quiz statement", err)
		return nil, rest_errors.NewInternalServerError("error when trying to get all quiz", errors.New("database error"))
	}
	defer stmt.Close()

	rows, getErr := stmt.Query()
	if getErr != nil {
		logger.Error("error when trying to get all quiz", getErr)
		return nil, rest_errors.NewInternalServerError("error when trying to get all quiz", errors.New("database error"))
	}
	defer rows.Close()

	result := make([]Quiz, 0)
	var isActive int
	for rows.Next() {
		if err := rows.Scan(&quiz.ID, &quiz.Question, &isActive, &quiz.ActivityID); err != nil {
			logger.Error("error when trying to scan quiz rows into quiz struct", err)
			return nil, rest_errors.NewInternalServerError("error when trying to get all quiz", errors.New("database error"))
		}

		if isActive == 1 {
			quiz.IsActive = true
		} else {
			quiz.IsActive = false
		}

		result = append(result, *quiz)
	}

	if len(result) == 0 {
		return nil, rest_errors.NewNotFoundError("no quiz rows in result set")
	}
	return result, nil
}

func (quiz *Quiz) Update(isActive int) rest_errors.RestErr {
	stmt, err := quiz_db.DbConn().Prepare(queryUpdateQuiz)
	if err != nil {
		logger.Error("error when trying to prepare update quiz statement", err)
		return rest_errors.NewInternalServerError("error when trying to update quiz", errors.New("database error"))
	}
	defer stmt.Close()

	_, err = stmt.Exec(quiz.Question, isActive, quiz.ActivityID, quiz.ID)
	if err != nil {
		logger.Error("error when trying to update quiz", err)
		return rest_errors.NewInternalServerError("error when trying to update user", errors.New("database error"))
	}

	return nil
}

func (quiz *Quiz) Delete() rest_errors.RestErr {
	stmt, err := quiz_db.DbConn().Prepare(queryDeleteQuiz)
	if err != nil {
		logger.Error("error when trying to prepare delete quiz statement", err)
		return rest_errors.NewInternalServerError("error when trying to delete quiz", errors.New("database error"))
	}
	defer stmt.Close()

	if _, err = stmt.Exec(quiz.ID); err != nil {
		logger.Error("error when trying to delete quiz by id", err)
		return rest_errors.NewInternalServerError("error when trying to delete quiz", errors.New("database error"))
	}

	return nil
}

func (quiz *QuizAndChoice) GetAllQuestionByActivityID() ([]QuizAndChoice, rest_errors.RestErr) {
	stmt, err := quiz_db.DbConn().Prepare(queryGetAllQuestionByActivityID)
	if err != nil {
		logger.Error("error when trying to prepare get all question by activity id statement", err)
		return nil, rest_errors.NewInternalServerError("error when trying to get all question", errors.New("database error"))
	}
	defer stmt.Close()

	rows, getErr := stmt.Query(quiz.ActivityID)
	if getErr != nil {
		logger.Error("error when trying to get all question by activity id", getErr)
		return nil, rest_errors.NewInternalServerError("error when trying to get all question", errors.New("database error"))
	}
	defer rows.Close()

	result := make([]QuizAndChoice, 0)
	for rows.Next() {
		if err := rows.Scan(&quiz.ID, &quiz.Question); err != nil {
			logger.Error("error when trying to scan question rows into question struct", err)
			return nil, rest_errors.NewInternalServerError("error when trying to get all question", errors.New("database error"))
		}

		result = append(result, *quiz)
	}

	if len(result) == 0 {
		return nil, rest_errors.NewNotFoundError("no question rows in result set")
	}
	return result, nil
}

func (choice *Choice) GetAllChoiceByQuestionID(quiz *QuizAndChoice) rest_errors.RestErr {
	stmt, err := quiz_db.DbConn().Prepare(queryGetAllChoiceByQuestionID)
	if err != nil {
		logger.Error("error when trying to prepare get choice by id statement", err)
		rest_errors.NewInternalServerError("error when trying to get choice", errors.New("database error"))
	}
	defer stmt.Close()

	rows, err := stmt.Query(quiz.ID)
	if err != nil {
		logger.Error("error when trying to get choice by question id", err)
		rest_errors.NewInternalServerError("error when trying to get choice", errors.New("database error"))
	}
	defer rows.Close()

	var isRight int
	for rows.Next() {
		if err := rows.Scan(&choice.ID, &choice.Choice, &isRight); err != nil {
			logger.Error("error when trying to scan row into choice struct", err)
			rest_errors.NewInternalServerError("error when trying to get choice", errors.New("database error"))
		}

		if isRight == 1 {
			choice.IsRight = true
		} else {
			choice.IsRight = false
		}

		quiz.Choices = append(quiz.Choices, *choice)
	}

	return nil
}

func (quiz *Quiz) DeleteQuestionByActivityID() rest_errors.RestErr {
	stmt, err := quiz_db.DbConn().Prepare(queryDeleteQuizByActivityID)
	if err != nil {
		logger.Error("error when trying to prepare delete question by activity id statement", err)
		return rest_errors.NewInternalServerError("error when trying to delete question", errors.New("database error"))
	}
	defer stmt.Close()

	if _, err := stmt.Exec(quiz.ActivityID); err != nil {
		logger.Error("error when trying to delete question by activity id", err)
		return rest_errors.NewInternalServerError("error when trying to delete question", errors.New("database error"))
	}

	return nil
}
