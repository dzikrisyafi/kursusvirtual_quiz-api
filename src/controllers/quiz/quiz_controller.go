package quiz

import (
	"net/http"
	"strconv"

	"github.com/dzikrisyafi/kursusvirtual_quiz-api/src/domain/quiz"
	"github.com/dzikrisyafi/kursusvirtual_quiz-api/src/services"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
	"github.com/gin-gonic/gin"
)

func getQuizId(quizIdParam string) (int64, rest_errors.RestErr) {
	quizID, quizErr := strconv.ParseInt(quizIdParam, 10, 64)
	if quizErr != nil {
		return 0, rest_errors.NewBadRequestError("quiz id should be a number")
	}
	return quizID, nil
}

func Create(c *gin.Context) {
	var quiz quiz.Quiz
	if err := c.ShouldBindJSON(&quiz); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	result, saveErr := services.QuizService.CreateQuiz(quiz)
	if saveErr != nil {
		c.JSON(saveErr.Status(), saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func Get(c *gin.Context) {
	quizID, idErr := getQuizId(c.Param("quiz_id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	quiz, getErr := services.QuizService.GetQuiz(quizID)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}

	c.JSON(http.StatusOK, quiz)
}

func GetAll(c *gin.Context) {
	quiz, err := services.QuizService.GetAllQuiz()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, quiz)
}

func Update(c *gin.Context) {
	quizID, idErr := getQuizId(c.Param("quiz_id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	var quiz quiz.Quiz
	if err := c.ShouldBindJSON(&quiz); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	quiz.ID = quizID
	isPartial := c.Request.Method == http.MethodPatch
	result, saveErr := services.QuizService.UpdateQuiz(isPartial, quiz)
	if saveErr != nil {
		c.JSON(saveErr.Status(), saveErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

func Delete(c *gin.Context) {
	quizID, idErr := getQuizId(c.Param("quiz_id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	if err := services.QuizService.DeleteQuiz(quizID); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"message": "success deleted quiz", "status": http.StatusOK})
}

func GetAllBySectionID(c *gin.Context) {
	sectionID, err := strconv.ParseInt(c.Param("section_id"), 10, 64)
	if err != nil {
		restErr := rest_errors.NewBadRequestError("section id should be a number")
		c.JSON(restErr.Status(), restErr)
		return
	}

	questions, getErr := services.QuizService.GetAllQuestionBySectionID(sectionID)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}

	c.JSON(http.StatusOK, questions)
}
