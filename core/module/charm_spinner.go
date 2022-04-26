// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Example
//
// ```
// import (
//	  "github.com/clevenio/goenv/core/module"
// )
//
//
// spinner := module.NewCharmSpinner("Getting Environment Ready!")
//
// go func() {
// 	 // Start a long running job
// 	 spinner.Quit()
// }()
//
// if err := spinner.Start(); err != nil {
// 	 fmt.Println(err)
// 	 os.Exit(1)
// }
// ```

// NewCharmSpinner Creates a new instance of tea.Program
func NewCharmSpinner(loadingMsg string) *tea.Program {
	return tea.NewProgram(initialSpinnerModel(loadingMsg))
}

// errMsg type
type errMsg error

// spinnerModel spinner type
type spinnerModel struct {
	spinner    spinner.Model
	spin       bool
	quitting   bool
	err        error
	loadingMsg string
}

// initialSpinnerModel creats a spinner model
func initialSpinnerModel(loadingMsg string) spinnerModel {
	s := spinner.New()

	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return spinnerModel{
		spinner:    s,
		spin:       true,
		loadingMsg: loadingMsg,
	}
}

// Init init the spinner model
func (m spinnerModel) Init() tea.Cmd {
	return m.spinner.Tick
}

// Update update the spinner
func (m spinnerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		default:
			return m, nil
		}

	case errMsg:
		m.err = msg
		return m, nil

	case spinner.TickMsg:
		var cmd tea.Cmd
		// If m.spin is false, don't update on the spinner, effectively stopping it.
		if m.spin {
			m.spinner, cmd = m.spinner.Update(msg)
		}
		return m, cmd

	default:
		return m, nil
	}
}

// View show the spinner
func (m spinnerModel) View() string {
	if m.err != nil {
		return m.err.Error()
	}

	str := fmt.Sprintf("\n   %s %s ...\n", m.spinner.View(), m.loadingMsg)

	if m.quitting {
		return str + "\n"
	}

	return str
}
