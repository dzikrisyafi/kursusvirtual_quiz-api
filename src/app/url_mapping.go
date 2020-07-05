package app

import (
	"github.com/dzikrisyafi/kursusvirtual_middleware/middleware"
	"github.com/dzikrisyafi/kursusvirtual_quiz-api/src/controllers/answers"
	"github.com/dzikrisyafi/kursusvirtual_quiz-api/src/controllers/choices"
	"github.com/dzikrisyafi/kursusvirtual_quiz-api/src/controllers/quiz"
)

func mapUrls() {
	// quiz end point
	quizGroup := router.Group("/quiz")
	quizGroup.Use(middleware.Auth())
	{
		quizGroup.POST("", quiz.Create)
		quizGroup.GET("/:quiz_id", quiz.Get)
		quizGroup.GET("", quiz.GetAll)
		quizGroup.PUT("/:quiz_id", quiz.Update)
		quizGroup.PATCH("/:quiz_id", quiz.Update)
		quizGroup.DELETE("/:quiz_id", quiz.Delete)
	}

	internalGroup := router.Group("/internal")
	internalGroup.Use(middleware.Auth())
	{
		internalGroup.GET("/quiz/:activity_id", quiz.GetAllByActivityID)
		internalGroup.DELETE("/quiz/:course_id", quiz.DeleteAll)
	}

	// choice end point
	choicesGroup := router.Group("/choices")
	choicesGroup.Use(middleware.Auth())
	{
		choicesGroup.POST("", choices.Create)
		choicesGroup.GET("/:choice_id", choices.Get)
		choicesGroup.PUT("/:choice_id", choices.Update)
		choicesGroup.PATCH("/:choice_id", choices.Update)
		choicesGroup.DELETE("/:choice_id", choices.Delete)
	}

	// answer end point
	answersGroup := router.Group("/answers")
	answersGroup.Use(middleware.Auth())
	{
		answersGroup.POST("", answers.Create)
		answersGroup.GET("/:user_id/:activity_id", answers.Get)
	}
}
