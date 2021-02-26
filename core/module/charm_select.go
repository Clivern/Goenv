// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Example
//
// ```
// import (
//	  "github.com/clivern/goenv/core/module"
// )
//
//
// list := module.NewCharmSelect("Which go version to install?", module.golangReleases)
//
// if err := list.Start(); err != nil {
// 	  fmt.Println("Error running program:", err)
// 	  os.Exit(1)
// }
//
// if module.SelectedValue == "" {
// 	  os.Exit(1)
// }
// ```

// NewCharmSelect Creates a new instance of tea.Program
func NewCharmSelect(title string, versions []string) *tea.Program {
	items := []list.Item{}

	for i := 0; i < len(versions); i++ {
		items = append(items, item(versions[i]))
	}

	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.Title = title
	l.SetShowStatusBar(true)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	m := listModel{list: l}

	return tea.NewProgram(m)
}

// SelectedValue the list selected value
var SelectedValue string

// listHeight variable
var listHeight = 14

// defaultWidth variable
var defaultWidth = 20

// titleStyle variable
var titleStyle = lipgloss.NewStyle().MarginLeft(2)

// itemStyle variable
var itemStyle = lipgloss.NewStyle().PaddingLeft(4)

// selectedItemStyle variable
var selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))

// paginationStyle variable
var paginationStyle = list.DefaultStyles().PaginationStyle.PaddingLeft(4)

// helpStyle variable
var helpStyle = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)

// quitTextStyle variable
var quitTextStyle = lipgloss.NewStyle().Margin(1, 0, 2, 4)

// item type
type item string

// FilterValue method
func (i item) FilterValue() string {
	return ""
}

// itemDelegate type
type itemDelegate struct{}

// Height method
func (d itemDelegate) Height() int {
	return 1
}

// Spacing method
func (d itemDelegate) Spacing() int {
	return 0
}

// Update method
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	return nil
}

// Render method
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s string) string {
			return selectedItemStyle.Render("> " + s)
		}
	}

	fmt.Fprintf(w, fn(str))
}

// listModel type
type listModel struct {
	list     list.Model
	items    []item
	choice   string
	quitting bool
}

// Init init the list model
func (m listModel) Init() tea.Cmd {
	return nil
}

// Update update the list
func (m listModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = string(i)
			}
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

// View show the list
func (m listModel) View() string {
	if m.choice != "" {
		SelectedValue = m.choice
		return ""
	}

	if m.quitting {
		return quitTextStyle.Render("Not hungry? That’s cool.")
	}

	return "\n" + m.list.View()
}
