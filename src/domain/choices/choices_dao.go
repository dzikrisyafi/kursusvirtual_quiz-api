package choices

import (
	"errors"

	"github.com/dzikrisyafi/kursusvirtual_quiz-api/src/datasources/mysql/quiz_db"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/logger"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
)

const (
	queryInsertChoice = `INSERT INTO choices(choice, is_right, question_id) VALUES(?, ?, ?);`
	queryGetChoice    = `SELECT choice, is_right, question_id FROM choices WHERE id=?;`
	queryUpdateChoice = `UPDATE choices SET choice=?, is_right=?, question_id=? WHERE id=?;`
	queryDeleteChoice = `DELETE FROM choices WHERE id=?;`
)

func (choice *Choice) Save(isRight int) rest_errors.RestErr {
	stmt, err := quiz_db.DbConn().Prepare(queryInsertChoice)
	if err != nil {
		logger.Error("error when trying to prepare save choice statement", err)
		return rest_errors.NewInternalServerError("error when trying to save choice", errors.New("database error"))
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(choice.Choice, isRight, choice.QuestionID)
	if saveErr != nil {
		logger.Error("error when trying to save choice", saveErr)
		return rest_errors.NewInternalServerError("error when trying to save choice", errors.New("database error"))
	}

	choiceID, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert id after creating a new choice", err)
		return rest_errors.NewInternalServerError("error when trying to save choice", errors.New("database error"))
	}
	choice.ID = int(choiceID)

	return nil
}

func (choice *Choice) Get() rest_errors.RestErr {
	stmt, err := quiz_db.DbConn().Prepare(queryGetChoice)
	if err != nil {
		logger.Error("error when trying to prepare get choice by id statement", err)
		return rest_errors.NewInternalServerError("error when trying to get choice", errors.New("database error"))
	}
	defer stmt.Close()

	var isRight int
	result := stmt.QueryRow(choice.ID)
	if getErr := result.Scan(&choice.Choice, &isRight, &choice.QuestionID); getErr != nil {
		logger.Error("error when trying to get choice by id", getErr)
		return rest_errors.NewInternalServerError("error when trying to get choice", errors.New("database error"))
	}

	if isRight == 1 {
		choice.IsRight = true
	} else {
		choice.IsRight = false
	}

	return nil
}

func (choice *Choice) Update(isRight int) rest_errors.RestErr {
	stmt, err := quiz_db.DbConn().Prepare(queryUpdateChoice)
	if err != nil {
		logger.Error("error when trying to prepare update choice statement", err)
		return rest_errors.NewInternalServerError("error when trying to update choice", errors.New("database error"))
	}
	defer stmt.Close()

	_, err = stmt.Exec(choice.Choice, isRight, choice.QuestionID, choice.ID)
	if err != nil {
		logger.Error("error when trying to update choice", err)
		return rest_errors.NewInternalServerError("error when trying to update choice", errors.New("database error"))
	}

	return nil
}

func (choice *Choice) Delete() rest_errors.RestErr {
	stmt, err := quiz_db.DbConn().Prepare(queryDeleteChoice)
	if err != nil {
		logger.Error("error when trying to prepare delete choice by id statement", err)
		return rest_errors.NewInternalServerError("error when trying to delete choice", errors.New("database error"))
	}
	defer stmt.Close()

	if _, err = stmt.Exec(choice.ID); err != nil {
		logger.Error("error when trying to delete choice by id", err)
		return rest_errors.NewInternalServerError("error when trying to delete choice", errors.New("database error"))
	}

	return nil
}
