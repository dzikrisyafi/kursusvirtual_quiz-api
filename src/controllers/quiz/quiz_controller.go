package quiz

import (
	"net/http"

	"github.com/dzikrisyafi/kursusvirtual_oauth-go/oauth"
	"github.com/dzikrisyafi/kursusvirtual_quiz-api/src/domain/quiz"
	"github.com/dzikrisyafi/kursusvirtual_quiz-api/src/services"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/controller_utils"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_resp"
	"github.com/gin-gonic/gin"
)

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

	resp := rest_resp.NewStatusCreated("success created quiz", result.Marshall(oauth.IsPublic(c.Request)))
	c.JSON(resp.Status(), resp)
}

func Get(c *gin.Context) {
	quizID, idErr := controller_utils.GetIDInt(c.Param("quiz_id"), "quiz id")
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	quiz, getErr := services.QuizService.GetQuiz(quizID)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}

	resp := rest_resp.NewStatusOK("success get quiz", quiz)
	c.JSON(resp.Status(), resp)
}

func GetAll(c *gin.Context) {
	quiz, err := services.QuizService.GetAllQuiz()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	resp := rest_resp.NewStatusOK("success get quiz", quiz.Marshall(oauth.IsPublic(c.Request)))
	c.JSON(resp.Status(), resp)
}

func GetAllByActivityID(c *gin.Context) {
	activityID, err := controller_utils.GetIDInt(c.Param("activity_id"), "activity id")
	if err != nil {
		restErr := rest_errors.NewBadRequestError("activity id should be a number")
		c.JSON(restErr.Status(), restErr)
		return
	}

	questions, getErr := services.QuizService.GetAllQuestionByActivityID(activityID)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}

	resp := rest_resp.NewStatusOK("success get quiz", questions.Marshall(oauth.IsPublic(c.Request)))
	c.JSON(resp.Status(), resp)
}

func Update(c *gin.Context) {
	quizID, idErr := controller_utils.GetIDInt(c.Param("quiz_id"), "quiz id")
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

	resp := rest_resp.NewStatusOK("success updated quiz", result.Marshall(oauth.IsPublic(c.Request)))
	c.JSON(resp.Status(), resp)
}

func Delete(c *gin.Context) {
	quizID, idErr := controller_utils.GetIDInt(c.Param("quiz_id"), "quiz id")
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

func DeleteAll(c *gin.Context) {
	courseID, idErr := controller_utils.GetIDInt(c.Param("course_id"), "course id")
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	if err := services.QuizService.DeleteQuestionByCourseID(courseID, c.Query("access_token")); err != nil {
		c.JSON(err.Status(), err)
		return
	}
}
