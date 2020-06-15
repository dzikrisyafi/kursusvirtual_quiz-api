package answers

import (
	"github.com/dzikrisyafi/kursusvirtual_oauth-go/oauth"
	"github.com/dzikrisyafi/kursusvirtual_quiz-api/src/domain/answers"
	"github.com/dzikrisyafi/kursusvirtual_quiz-api/src/services"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/controller_utils"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_resp"
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

	resp := rest_resp.NewStatusCreated("success save user answer", result.Marshall(oauth.IsPublic(c.Request)))
	c.JSON(resp.Status(), resp)
}

func Get(c *gin.Context) {
	userID, idErr := controller_utils.GetIDInt(c.Param("user_id"), "user id")
	if idErr != nil {
		restErr := rest_errors.NewBadRequestError("user id should be a number")
		c.JSON(restErr.Status(), restErr)
		return
	}

	activityID, idErr := controller_utils.GetIDInt(c.Param("activity_id"), "activity id")
	if idErr != nil {
		restErr := rest_errors.NewBadRequestError("activity id should be a number")
		c.JSON(restErr.Status(), restErr)
	}

	answer, getErr := services.AnswerService.GetUserAnswer(userID, activityID)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}

	resp := rest_resp.NewStatusOK("success get user answer", answer.Marshall(oauth.IsPublic(c.Request)))
	c.JSON(resp.Status(), resp)
}
