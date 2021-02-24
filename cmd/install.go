// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var (
	listHeight        = 14
	defaultWidth      = 20
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
	items             = []list.Item{
		item("go1"),
		item("go1.0.1"),
		item("go1.0.2"),
		item("go1.0.3"),
		item("go1.1"),
		item("go1.1.1"),
		item("go1.1.2"),
		item("go1.10"),
		item("go1.10.1"),
		item("go1.10.2"),
		item("go1.10.3"),
		item("go1.10.4"),
		item("go1.10.5"),
		item("go1.10.6"),
		item("go1.10.7"),
		item("go1.10.8"),
		item("go1.10beta1"),
		item("go1.10beta2"),
		item("go1.10rc1"),
		item("go1.10rc2"),
		item("go1.11"),
		item("go1.11.1"),
		item("go1.11.10"),
		item("go1.11.11"),
		item("go1.11.12"),
		item("go1.11.13"),
		item("go1.11.2"),
		item("go1.11.3"),
		item("go1.11.4"),
		item("go1.11.5"),
		item("go1.11.6"),
		item("go1.11.7"),
		item("go1.11.8"),
		item("go1.11.9"),
		item("go1.11beta1"),
		item("go1.11beta2"),
		item("go1.11beta3"),
		item("go1.11rc1"),
		item("go1.11rc2"),
		item("go1.12"),
		item("go1.12.1"),
		item("go1.12.10"),
		item("go1.12.11"),
		item("go1.12.12"),
		item("go1.12.13"),
		item("go1.12.14"),
		item("go1.12.15"),
		item("go1.12.16"),
		item("go1.12.17"),
		item("go1.12.2"),
		item("go1.12.3"),
		item("go1.12.4"),
		item("go1.12.5"),
		item("go1.12.6"),
		item("go1.12.7"),
		item("go1.12.8"),
		item("go1.12.9"),
		item("go1.12beta1"),
		item("go1.12beta2"),
		item("go1.12rc1"),
		item("go1.13"),
		item("go1.13.1"),
		item("go1.13.10"),
		item("go1.13.11"),
		item("go1.13.12"),
		item("go1.13.13"),
		item("go1.13.14"),
		item("go1.13.15"),
		item("go1.13.2"),
		item("go1.13.3"),
		item("go1.13.4"),
		item("go1.13.5"),
		item("go1.13.6"),
		item("go1.13.7"),
		item("go1.13.8"),
		item("go1.13.9"),
		item("go1.13beta1"),
		item("go1.13rc1"),
		item("go1.13rc2"),
		item("go1.14"),
		item("go1.14.1"),
		item("go1.14.10"),
		item("go1.14.11"),
		item("go1.14.12"),
		item("go1.14.13"),
		item("go1.14.14"),
		item("go1.14.15"),
		item("go1.14.2"),
		item("go1.14.3"),
		item("go1.14.4"),
		item("go1.14.5"),
		item("go1.14.6"),
		item("go1.14.7"),
		item("go1.14.8"),
		item("go1.14.9"),
		item("go1.14beta1"),
		item("go1.14rc1"),
		item("go1.15"),
		item("go1.15.1"),
		item("go1.15.10"),
		item("go1.15.11"),
		item("go1.15.12"),
		item("go1.15.13"),
		item("go1.15.14"),
		item("go1.15.15"),
		item("go1.15.2"),
		item("go1.15.3"),
		item("go1.15.4"),
		item("go1.15.5"),
		item("go1.15.6"),
		item("go1.15.7"),
		item("go1.15.8"),
		item("go1.15.9"),
		item("go1.15beta1"),
		item("go1.15rc1"),
		item("go1.15rc2"),
		item("go1.16"),
		item("go1.16.1"),
		item("go1.16.10"),
		item("go1.16.11"),
		item("go1.16.12"),
		item("go1.16.13"),
		item("go1.16.14"),
		item("go1.16.15"),
		item("go1.16.2"),
		item("go1.16.3"),
		item("go1.16.4"),
		item("go1.16.5"),
		item("go1.16.6"),
		item("go1.16.7"),
		item("go1.16.8"),
		item("go1.16.9"),
		item("go1.16beta1"),
		item("go1.16rc1"),
		item("go1.17"),
		item("go1.17.1"),
		item("go1.17.10"),
		item("go1.17.11"),
		item("go1.17.2"),
		item("go1.17.3"),
		item("go1.17.4"),
		item("go1.17.5"),
		item("go1.17.6"),
		item("go1.17.7"),
		item("go1.17.8"),
		item("go1.17.9"),
		item("go1.17beta1"),
		item("go1.17rc1"),
		item("go1.17rc2"),
		item("go1.18"),
		item("go1.18.1"),
		item("go1.18.2"),
		item("go1.18.3"),
		item("go1.18beta1"),
		item("go1.18beta2"),
		item("go1.18rc1"),
		item("go1.19beta1"),
		item("go1.19rc1"),
		item("go1.1rc2"),
		item("go1.1rc3"),
		item("go1.2"),
		item("go1.2.1"),
		item("go1.2.2"),
		item("go1.2rc2"),
		item("go1.2rc3"),
		item("go1.2rc4"),
		item("go1.2rc5"),
		item("go1.3"),
		item("go1.3.1"),
		item("go1.3.2"),
		item("go1.3.3"),
		item("go1.3beta1"),
		item("go1.3beta2"),
		item("go1.3rc1"),
		item("go1.3rc2"),
		item("go1.4"),
		item("go1.4.1"),
		item("go1.4.2"),
		item("go1.4.3"),
		item("go1.4beta1"),
		item("go1.4rc1"),
		item("go1.4rc2"),
		item("go1.5"),
		item("go1.5.1"),
		item("go1.5.2"),
		item("go1.5.3"),
		item("go1.5.4"),
		item("go1.5beta1"),
		item("go1.5beta2"),
		item("go1.5beta3"),
		item("go1.5rc1"),
		item("go1.6"),
		item("go1.6.1"),
		item("go1.6.2"),
		item("go1.6.3"),
		item("go1.6.4"),
		item("go1.6beta1"),
		item("go1.6beta2"),
		item("go1.6rc1"),
		item("go1.6rc2"),
		item("go1.7"),
		item("go1.7.1"),
		item("go1.7.2"),
		item("go1.7.3"),
		item("go1.7.4"),
		item("go1.7.5"),
		item("go1.7.6"),
		item("go1.7beta1"),
		item("go1.7beta2"),
		item("go1.7rc1"),
		item("go1.7rc2"),
		item("go1.7rc3"),
		item("go1.7rc4"),
		item("go1.7rc5"),
		item("go1.7rc6"),
		item("go1.8"),
		item("go1.8.1"),
		item("go1.8.2"),
		item("go1.8.3"),
		item("go1.8.4"),
		item("go1.8.5"),
		item("go1.8.5rc4"),
		item("go1.8.5rc5"),
		item("go1.8.6"),
		item("go1.8.7"),
		item("go1.8beta1"),
		item("go1.8beta2"),
		item("go1.8rc1"),
		item("go1.8rc2"),
		item("go1.8rc3"),
		item("go1.9"),
		item("go1.9.1"),
		item("go1.9.2"),
		item("go1.9.3"),
		item("go1.9.4"),
		item("go1.9.5"),
		item("go1.9.6"),
		item("go1.9.7"),
		item("go1.9beta1"),
		item("go1.9beta2"),
		item("go1.9rc1"),
		item("go1.9rc2"),
	}
)

type item string

func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 0 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
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

type model struct {
	list     list.Model
	items    []item
	choice   string
	quitting bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m model) View() string {
	if m.choice != "" {
		return quitTextStyle.Render(fmt.Sprintf("%s? Sounds good to me.", m.choice))
	}

	if m.quitting {
		return quitTextStyle.Render("Not hungry? That’s cool.")
	}

	return "\n" + m.list.View()
}

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install a go version.",
	Run: func(cmd *cobra.Command, args []string) {
		l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
		l.Title = "Which go version to install?"
		l.SetShowStatusBar(false)
		l.SetFilteringEnabled(false)
		l.Styles.Title = titleStyle
		l.Styles.PaginationStyle = paginationStyle
		l.Styles.HelpStyle = helpStyle

		m := model{list: l}

		if err := tea.NewProgram(m).Start(); err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
