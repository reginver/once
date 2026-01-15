package ui

import (
	"image/color"
	"math/rand/v2"
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type ProgressBusy struct {
	Width int
	Color color.Color

	pattern []rune
}

type progressBusyTickMsg struct{}

func NewProgressBusy(width int, clr color.Color) ProgressBusy {
	return ProgressBusy{
		Width:   width,
		Color:   clr,
		pattern: generateBraillePattern(width),
	}
}

func (p ProgressBusy) Init() tea.Cmd {
	return p.tick()
}

func (p ProgressBusy) Update(msg tea.Msg) (ProgressBusy, tea.Cmd) {
	switch msg.(type) {
	case progressBusyTickMsg:
		p.pattern = generateBraillePattern(p.Width)
		return p, p.tick()
	}
	return p, nil
}

func (p ProgressBusy) View() string {
	if p.Width <= 0 {
		return ""
	}

	style := lipgloss.NewStyle().Foreground(p.Color)
	return style.Render(string(p.pattern))
}

// Private

func (p ProgressBusy) tick() tea.Cmd {
	return tea.Tick(50*time.Millisecond, func(t time.Time) tea.Msg {
		return progressBusyTickMsg{}
	})
}

func generateBraillePattern(width int) []rune {
	pattern := make([]rune, width)
	for i := range pattern {
		// Braille patterns: U+2800 to U+28FF (256 patterns)
		pattern[i] = rune(0x2800 + rand.IntN(256))
	}
	return pattern
}
