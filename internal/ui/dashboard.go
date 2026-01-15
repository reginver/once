package ui

import (
	"fmt"
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"

	"github.com/basecamp/amar/internal/docker"
)

type Dashboard struct {
	app           *docker.Application
	width, height int
}

type dashboardTickMsg struct{}

func NewDashboard(app *docker.Application) Dashboard {
	return Dashboard{
		app: app,
	}
}

func (m Dashboard) Init() tea.Cmd {
	return tea.Tick(time.Second, func(time.Time) tea.Msg { return dashboardTickMsg{} })
}

func (m Dashboard) Update(msg tea.Msg) (Component, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
	case dashboardTickMsg:
		return m, tea.Tick(time.Second, func(time.Time) tea.Msg { return dashboardTickMsg{} })
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

	if m.app.Running && !m.app.RunningSince.IsZero() {
		stateDisplay += fmt.Sprintf(" (up %s)", formatDuration(time.Since(m.app.RunningSince)))
	}

	content := lipgloss.NewStyle().PaddingLeft(2).Render(stateDisplay)

	return title + "\n\n" + content
}

// Helpers

func formatDuration(d time.Duration) string {
	if d < time.Minute {
		return fmt.Sprintf("%ds", int(d.Seconds()))
	}
	if d < time.Hour {
		return fmt.Sprintf("%dm", int(d.Minutes()))
	}
	if d < 24*time.Hour {
		hours := int(d.Hours())
		mins := int(d.Minutes()) % 60
		if mins == 0 {
			return fmt.Sprintf("%dh", hours)
		}
		return fmt.Sprintf("%dh %dm", hours, mins)
	}
	days := int(d.Hours()) / 24
	hours := int(d.Hours()) % 24
	if hours == 0 {
		return fmt.Sprintf("%dd", days)
	}
	return fmt.Sprintf("%dd %dh", days, hours)
}
