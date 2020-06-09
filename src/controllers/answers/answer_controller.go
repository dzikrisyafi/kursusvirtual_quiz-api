package answers

import (
	"net/http"
	"strconv"

	"github.com/dzikrisyafi/kursusvirtual_oauth-go/oauth"
	"github.com/dzikrisyafi/kursusvirtual_quiz-api/src/domain/answers"
	"github.com/dzikrisyafi/kursusvirtual_quiz-api/src/services"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var answer answers.Answer
	if err := c.ShouldBindJSON(&answer); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	result, saveErr := services.AnswerService.CreateAnswer(answer)
	if saveErr != nil {
		c.JSON(saveErr.Status(), saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func Get(c *gin.Context) {
	userID, idErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if idErr != nil {
		restErr := rest_errors.NewBadRequestError("user id should be a number")
		c.JSON(restErr.Status(), restErr)
		return
	}

	sectionID, idErr := strconv.Atoi(c.Param("section_id"))
	if idErr != nil {
		restErr := rest_errors.NewBadRequestError("section id should be a number")
		c.JSON(restErr.Status(), restErr)
	}

	answer, getErr := services.AnswerService.GetUserAnswer(userID, sectionID)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}

	c.JSON(http.StatusOK, answer.Marshall(oauth.IsPublic(c.Request)))
}
