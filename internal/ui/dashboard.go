package ui

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"

	"github.com/basecamp/amar/internal/docker"
)

type Dashboard struct {
	app           *docker.Application
	width, height int
}

func NewDashboard(app *docker.Application) Dashboard {
	return Dashboard{
		app: app,
	}
}

func (m Dashboard) Init() tea.Cmd {
	return nil
}

func (m Dashboard) Update(msg tea.Msg) (Component, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
	}
	return m, nil
}

func (m Dashboard) View() string {
	title := Styles.Title.Width(m.width).Align(lipgloss.Center).Render(m.app.Settings.Name)

	status := "stopped"
	statusColor := lipgloss.Color("#ff5555")
	if m.app.Running {
		status = "running"
		statusColor = lipgloss.Color("#50fa7b")
	}

	stateStyle := lipgloss.NewStyle().Foreground(statusColor)
	stateDisplay := fmt.Sprintf("State: %s", stateStyle.Render(status))

	content := lipgloss.NewStyle().PaddingLeft(2).Render(stateDisplay)

	return title + "\n\n" + content
}
