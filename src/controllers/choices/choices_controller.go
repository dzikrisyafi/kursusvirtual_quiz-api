package choices

import (
	"net/http"

	"github.com/dzikrisyafi/kursusvirtual_oauth-go/oauth"
	"github.com/dzikrisyafi/kursusvirtual_quiz-api/src/domain/choices"
	"github.com/dzikrisyafi/kursusvirtual_quiz-api/src/services"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/controller_utils"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_resp"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var choice choices.Choice
	if err := c.ShouldBindJSON(&choice); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	result, saveErr := services.ChoiceService.CreateChoice(choice)
	if saveErr != nil {
		c.JSON(saveErr.Status(), saveErr)
		return
	}

	resp := rest_resp.NewStatusCreated("success create new choice", result.Marshall(oauth.IsPublic(c.Request)))
	c.JSON(resp.Status(), resp)
}

func Get(c *gin.Context) {
	choiceID, idErr := controller_utils.GetIDInt(c.Param("choice_id"), "choice id")
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	choice, getErr := services.ChoiceService.GetChoice(choiceID)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}

	resp := rest_resp.NewStatusOK("success get choice data", choice.Marshall(oauth.IsPublic(c.Request)))
	c.JSON(resp.Status(), resp)
}

func Update(c *gin.Context) {
	choiceID, idErr := controller_utils.GetIDInt(c.Param("choice_id"), "choice id")
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	var choice choices.Choice
	if err := c.ShouldBindJSON(&choice); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	choice.ID = choiceID
	isPartial := c.Request.Method == http.MethodPatch
	result, saveErr := services.ChoiceService.UpdateChoice(isPartial, choice)
	if saveErr != nil {
		c.JSON(saveErr.Status(), saveErr)
		return
	}

	resp := rest_resp.NewStatusOK("success update choice", result.Marshall(oauth.IsPublic(c.Request)))
	c.JSON(resp.Status(), resp)
}

func Delete(c *gin.Context) {
	choiceID, idErr := controller_utils.GetIDInt(c.Param("choice_id"), "choice id")
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	if err := services.ChoiceService.DeleteChoice(choiceID); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"message": "success deleted choice", "status": http.StatusOK})
}
