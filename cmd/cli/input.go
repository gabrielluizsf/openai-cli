package cli

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Input interface {
	Value() string
	Blur() tea.Msg
	Update(tea.Msg) (Input, tea.Cmd)
	View() string
}

type ShortAnswerField struct {
	textinput textinput.Model
}

func (saf *ShortAnswerField) Value() string {
	return saf.textinput.Value()
}

func (saf *ShortAnswerField) Update(msg tea.Msg) (Input, tea.Cmd) {
	var cmd tea.Cmd
	saf.textinput, cmd = saf.textinput.Update(msg)
	return saf, cmd
}
func (saf *ShortAnswerField) Blur() tea.Msg {
	return saf.textinput.Blur
}

func (saf *ShortAnswerField) View() string {
	return saf.textinput.View()
}

type LongAnswerField struct {
	textarea textarea.Model
}

func (laf *LongAnswerField) Value() string {
	return laf.textarea.Value()
}

func (laf *LongAnswerField) Blur() tea.Msg {
	return laf.textarea.Blur
}

func (laf *LongAnswerField) View() string {
	return laf.textarea.View()
}

func (laf *LongAnswerField) Update(msg tea.Msg) (Input, tea.Cmd) {
	var cmd tea.Cmd
	laf.textarea, cmd = laf.textarea.Update(msg)
	return laf, cmd
}

const PLACEHOLDER = "Your answer here"

func NewShortAnswerField() *ShortAnswerField {
	ti := textinput.New()
	ti.Placeholder = PLACEHOLDER
	ti.Focus()
	return &ShortAnswerField{ti}
}

func NewLongAnswerField() *LongAnswerField {
	ta := textarea.New()
	ta.Placeholder = PLACEHOLDER
	ta.Focus()
	return &LongAnswerField{ta}
}
