package view

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/keshavchand/swsimplify/models"
	"github.com/keshavchand/swsimplify/simplifier"
	"github.com/keshavchand/swsimplify/view/calculator"
)

var (
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))

	noStyle    = lipgloss.NewStyle()
	wrongStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("9"))
)

type ViewScreen struct {
	quit      bool
	calculate bool
	input     []textinput.Model
	FromNames map[string]struct{}
	ToNames   map[string]struct{}

	focusIndex    int
	txnFocusIndex int

	TxnInfo []models.RawTransactionInfo
	Service simplifier.Service
}

func IntialMode() *ViewScreen {
	v := ViewScreen{
		FromNames: make(map[string]struct{}),
		ToNames:   make(map[string]struct{}),
	}

	v.reset()
	return &v
}

func IntialModeWithTransactions(txn []models.RawTransactionInfo) *ViewScreen {
	v := ViewScreen{
		FromNames: make(map[string]struct{}),
		ToNames:   make(map[string]struct{}),
		TxnInfo:   txn,
	}

	v.reset()
	return &v
}

func (v *ViewScreen) reset() {
	ph := []string{"From", "To", "Amount", "Description"}
	v.input = make([]textinput.Model, len(ph))
	for i := 0; i < 3; i++ {
		t := textinput.New()
		t.Placeholder = ph[i]
		t.CharLimit = 32
		v.input[i] = t
	}

	// Description
	t := textinput.New()
	t.Placeholder = ph[Description]
	v.input[Description] = t

	v.input[From].ShowSuggestions = true
	var fromNames []string
	for k := range v.FromNames {
		fromNames = append(fromNames, k)
	}
	v.input[From].SetSuggestions(fromNames)

	v.input[To].ShowSuggestions = true
	var toNames []string
	for k := range v.ToNames {
		toNames = append(toNames, k)
	}
	v.input[To].SetSuggestions(toNames)

	v.input[v.focusIndex].Focus()
	v.input[v.focusIndex].PromptStyle = focusedStyle
	v.input[v.focusIndex].TextStyle = focusedStyle
}

func (v *ViewScreen) Init() tea.Cmd {
	return textinput.Blink
}

func (m *ViewScreen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	inputLocChanged := false
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			// We skip the exit to the next update call so the final result can be presented
			return m, tea.Quit
		case "up":
			inputLocChanged = true
			m.focusIndex--
			if (m.focusIndex) < 0 {
				m.focusIndex = len(m.input) - 1
			}
		case "ctrl+j":
			if len(m.TxnInfo) == 0 {
				break
			}
			m.txnFocusIndex += 1
			m.txnFocusIndex %= len(m.TxnInfo)
		case "ctrl+k":
			m.txnFocusIndex -= 1
			if m.txnFocusIndex < 0 {
				m.txnFocusIndex = len(m.TxnInfo) - 1
			}
		case "ctrl+d":
			if len(m.TxnInfo) == 0 {
				break
			}
			m.TxnInfo = append(m.TxnInfo[:m.txnFocusIndex], m.TxnInfo[m.txnFocusIndex+1:]...)
			if m.txnFocusIndex == len(m.TxnInfo) {
				m.txnFocusIndex = 0
			}
		case "down":
			inputLocChanged = true
			m.focusIndex++
			m.focusIndex %= len(m.input)
		case "enter":
			inputLocChanged = true
			if m.focusIndex == len(m.input)-1 {
				m.calculate = true
				m.focusIndex = 0
			} else {
				m.focusIndex++
			}
		}
	}

	if inputLocChanged {
		cmds := make([]tea.Cmd, len(m.input))
		for i := 0; i <= len(m.input)-1; i++ {
			if i == m.focusIndex {
				// Set focused state
				cmds[i] = m.input[i].Focus()
				m.input[i].PromptStyle = focusedStyle
				m.input[i].TextStyle = focusedStyle
				continue
			}
			// Remove focused state
			m.input[i].Blur()
			m.input[i].PromptStyle = noStyle
			m.input[i].TextStyle = noStyle
		}

		return m, tea.Batch(cmds...)
	}

	if m.calculate {
		m.calculate = false

		from := m.input[From].Value()
		to := m.input[To].Value()
		amount := m.input[Amount].Value()
		desc := m.input[Description].Value()

		amt, err := calculator.CalculatePostFix(amount)

		if err != nil {
			m.focusIndex = int(Amount)
			m.input[Amount].SetValue("")
			m.input[Amount].Placeholder = err.Error()
			m.input[Amount].PromptStyle = wrongStyle
			return m, nil
		}

		m.FromNames[from] = struct{}{}
		m.ToNames[to] = struct{}{}

		txnInfo := models.RawTransactionInfo{
			From:        from,
			To:          to,
			Amount:      amt,
			Description: desc,
		}

		m.reset()
		m.TxnInfo = append(m.TxnInfo, txnInfo)
	}

	cmd := make([]tea.Cmd, len(m.input))
	for i := 0; i < len(m.input); i++ {
		m.input[i], cmd[i] = m.input[i].Update(msg)
	}

	return m, tea.Batch(cmd...)
}

func (m *ViewScreen) View() string {
	var buffer strings.Builder
	for i := 0; i < len(m.input); i++ {
		buffer.WriteString(m.input[i].View())
		buffer.WriteString("\n")
	}

	for _, c := range m.TxnInfo {
		var txn strings.Builder
		var style lipgloss.Style

		if c == m.TxnInfo[m.txnFocusIndex] {
			style = focusedStyle
		} else {
			style = noStyle
		}
		txn.WriteString(c.String())
		buffer.WriteString("\n")

		buffer.WriteString(style.Render(txn.String()))
	}
	return buffer.String()
}

var _ tea.Model = &ViewScreen{}
