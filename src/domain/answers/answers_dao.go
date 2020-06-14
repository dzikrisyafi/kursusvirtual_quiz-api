package answers

import (
	"errors"

	"github.com/dzikrisyafi/kursusvirtual_quiz-api/src/datasources/mysql/quiz_db"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/logger"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
)

const (
	queryInsertAnswer  = `INSERT INTO answer(user_id, question_id, choice_id, is_right, answer_time) VALUES(?, ?, ?, ?, ?);`
	queryGetUserAnswer = `SELECT answer.id, user_id, question_id, choice_id, answer.is_right FROM answer INNER JOIN question ON question_id=question.id WHERE user_id=? AND section_id=? AND is_active=1;`
)

func (answer *Answer) Save(isRight int) rest_errors.RestErr {
	stmt, err := quiz_db.DbConn().Prepare(queryInsertAnswer)
	if err != nil {
		logger.Error("error when trying to prepare save answer statement", err)
		return rest_errors.NewInternalServerError("error when trying to save answer", errors.New("database error"))
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(answer.UserID, answer.QuestionID, answer.ChoiceID, isRight, answer.AnswerTime)
	if saveErr != nil {
		logger.Error("error when trying to save answer", saveErr)
		return rest_errors.NewInternalServerError("error when trying to save answer", errors.New("database error"))
	}

	answerID, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert id after creating a new answer", err)
		return rest_errors.NewInternalServerError("error when trying to save answer", errors.New("database error"))
	}
	answer.ID = int(answerID)

	return nil
}

func (answer *Answer) Get(sectionID int) ([]Answer, rest_errors.RestErr) {
	stmt, err := quiz_db.DbConn().Prepare(queryGetUserAnswer)
	if err != nil {
		logger.Error("error when trying to prepare get user answer statement", err)
		return nil, rest_errors.NewInternalServerError("error when trying to get user answer", errors.New("database error"))
	}
	defer stmt.Close()

	rows, getErr := stmt.Query(answer.UserID, sectionID)
	if getErr != nil {
		logger.Error("error when trying to get user answer", err)
		return nil, rest_errors.NewInternalServerError("error when trying to get user answer", errors.New("database error"))
	}
	defer rows.Close()

	result := make([]Answer, 0)
	var isRight int
	for rows.Next() {
		if err := rows.Scan(&answer.ID, &answer.UserID, &answer.QuestionID, &answer.ChoiceID, &isRight); err != nil {
			logger.Error("error when trying to scan answer rows into answer struct", getErr)
			return nil, rest_errors.NewInternalServerError("error when trying to get user answer", errors.New("database error"))
		}

		if isRight == 1 {
			answer.IsRight = true
		} else {
			answer.IsRight = false
		}

		result = append(result, *answer)
	}

	if len(result) == 0 {
		return nil, rest_errors.NewNotFoundError("no answer rows in result set")
	}

	return result, nil
}
