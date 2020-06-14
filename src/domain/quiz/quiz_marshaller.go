package quiz

type PublicQuiz struct {
	ID         int `json:"id"`
	ActivityID int `json:"activity_id"`
}

type PublicQuizAndActivity struct {
	ID         int `json:"id"`
	ActivityID int `json:"activity_id"`
}

func (quizs Quizs) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(quizs))
	for index, quiz := range quizs {
		result[index] = quiz.Marshall(isPublic)
	}

	return result
}

func (quiz Quiz) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicQuiz{
			ID:         quiz.ID,
			ActivityID: quiz.ActivityID,
		}
	}

	return quiz
}

func (quizs QuizsAndChoices) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(quizs))
	for index, quiz := range quizs {
		result[index] = quiz.Marshall(isPublic)
	}

	return result
}

func (quiz QuizAndChoice) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicQuizAndActivity{
			ID:         quiz.ID,
			ActivityID: quiz.ActivityID,
		}
	}

	return quiz
}
