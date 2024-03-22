package cli

type Question struct {
	question string
	answer   string
	input    Input
}

func NewQuestion(question string) Question {
	return Question{question: question}
}

func newShortQuestion(question string) Question {
	q := NewQuestion(question)
	field := NewShortAnswerField()
	q.input = field
	return q
}
