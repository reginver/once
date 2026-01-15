package ui

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type EmptyState struct {
	width, height int
}

func NewEmptyState() EmptyState {
	return EmptyState{}
}

func (m EmptyState) Init() tea.Cmd {
	return nil
}

func (m EmptyState) Update(msg tea.Msg) (Component, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
	}
	return m, nil
}

func (m EmptyState) View() string {
	title := Styles.Title.Width(m.width).Align(lipgloss.Center).Render("Amar")

	message := lipgloss.NewStyle().
		Width(m.width).
		Align(lipgloss.Center).
		Render("No applications deployed yet")

	return title + "\n\n" + message
}
