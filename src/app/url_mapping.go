package app

import (
	"github.com/dzikrisyafi/kursusvirtual_quiz-api/src/controllers/answers"
	"github.com/dzikrisyafi/kursusvirtual_quiz-api/src/controllers/choices"
	"github.com/dzikrisyafi/kursusvirtual_quiz-api/src/controllers/quiz"
)

func mapUrls() {
	router.POST("/quiz", quiz.Create)
	router.GET("/quiz/:quiz_id", quiz.Get)
	router.GET("/quiz", quiz.GetAll)
	router.PUT("/quiz/:quiz_id", quiz.Update)
	router.PATCH("/quiz/:quiz_id", quiz.Update)
	router.DELETE("/quiz/:quiz_id", quiz.Delete)
	router.GET("internal/quiz/:section_id", quiz.GetAllBySectionID)

	router.POST("/choice", choices.Create)
	router.GET("/choice/:choice_id", choices.Get)
	router.PUT("/choice/:choice_id", choices.Update)
	router.PATCH("/choice/:choice_id", choices.Update)
	router.DELETE("/choice/:choice_id", choices.Delete)

	router.POST("/answer", answers.Create)
	router.GET("/answer/:user_id/:section_id", answers.Get)
}
