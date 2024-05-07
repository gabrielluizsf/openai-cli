package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/Simplou/goxios"
	"github.com/Simplou/openai"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type openaiCLI struct {
	index     int
	questions []Question
	width     int
	height    int
	styles    *Styles
	done      bool
}

func New(questions []Question) *openaiCLI {
	styles := DefaultStyles()
	return &openaiCLI{
		questions: questions,
		styles:    styles,
	}
}

func (m openaiCLI) Init() tea.Cmd {
	return nil
}

func (m openaiCLI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	current := &m.questions[m.index]
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			if m.index == len(m.questions)-1 {
				m.done = true
			}
			current.answer = current.input.Value()
			m.Next()
			return m, current.input.Blur
		}
	}
	current.input, cmd = current.input.Update(msg)
	return m, cmd
}

func (m openaiCLI) View() string {
	current := m.questions[m.index]
	if m.done {
		var data string
		var gptResponse string
		for _, question := range m.questions {
			title := fmt.Sprintf("Question:%s, Answer: %s", question.question, question.answer)
			gptResponse += GPT("correct the users' answers by informing if they got it right or wrong", title)
			data += fmt.Sprintf("\n%s\nGPT Answer: %s\n", title, gptResponse)
		}
		return lipgloss.Place(
			m.width,
			m.height,
			lipgloss.Center,
			lipgloss.Center,
			lipgloss.JoinVertical(
				lipgloss.Center,
				"Your Results",
				m.styles.InputField.Render(data),
			),
		)
	}
	if m.width == 0 {
		return "loading..."
	}
	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Center,
			m.questions[m.index].question,
			m.styles.InputField.Render(current.input.View()),
		),
	)
}

func (m *openaiCLI) Next() int {
	if m.index < len(m.questions)-1 {
		m.index++
		return m.index
	}
	m.index = 0
	return m.index
}

func GPT(system, userMessage string) string {
	ctx := context.Background()
	openaiClient := openai.New(ctx, os.Getenv("OPENAI_API_KEY"))
	httpClient := goxios.New(ctx)
	res, err := openai.ChatCompletion(openaiClient, httpClient, &openai.CompletionRequest{
		Model: "gpt-3.5-turbo",
		Messages: []openai.Message{
			{Role: "system", Content: system},
			{Role: "user", Content: userMessage},
		},
	})
	if err != nil {
		panic(err)
	}
	gptQuestion := res.Choices[0].Message.Content
	return gptQuestion
}
