package cli

import (
	tea "github.com/charmbracelet/bubbletea"
)

func Execute() {
	questions := []Question{}
	maxQuestions := 3
	for range maxQuestions {
		gptQuestion := GPT(
			"you are a question generator about programming in golang, when the user says start you should generate a question",
			"start",
		)
		questions = append(questions, newShortQuestion(gptQuestion))
	}
	m := New(questions)
	logFile, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	openaiCLIProgram := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := openaiCLIProgram.Run(); err != nil {
		panic(err)
	}
}
