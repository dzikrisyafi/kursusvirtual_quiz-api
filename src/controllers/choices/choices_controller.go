package choices

import (
	"net/http"
	"strconv"

	"github.com/dzikrisyafi/kursusvirtual_quiz-api/src/domain/choices"
	"github.com/dzikrisyafi/kursusvirtual_quiz-api/src/services"
	"github.com/dzikrisyafi/kursusvirtual_utils-go/rest_errors"
	"github.com/gin-gonic/gin"
)

func getChoiceId(choiceIdParam string) (int64, rest_errors.RestErr) {
	choiceID, choiceErr := strconv.ParseInt(choiceIdParam, 10, 64)
	if choiceErr != nil {
		return 0, rest_errors.NewBadRequestError("choice id should be a number")
	}
	return choiceID, nil
}

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

	c.JSON(http.StatusCreated, result)
}

func Get(c *gin.Context) {
	choiceID, idErr := getChoiceId(c.Param("choice_id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	choice, getErr := services.ChoiceService.GetChoice(choiceID)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}

	c.JSON(http.StatusOK, choice)
}

func Update(c *gin.Context) {
	choiceID, idErr := getChoiceId(c.Param("choice_id"))
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

	c.JSON(http.StatusOK, result)
}

func Delete(c *gin.Context) {
	choiceID, idErr := getChoiceId(c.Param("choice_id"))
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
