package ui

import (
	"strings"
	"testing"

	"charm.land/lipgloss/v2"
)

func stripAnsi(s string) string {
	var result strings.Builder
	inEscape := false
	for _, r := range s {
		if r == '\x1b' {
			inEscape = true
			continue
		}
		if inEscape {
			if r == 'm' {
				inEscape = false
			}
			continue
		}
		result.WriteRune(r)
	}
	return result.String()
}

func TestProgressBar(t *testing.T) {
	assertOutput := func(p ProgressBar, want string) {
		t.Helper()
		got := stripAnsi(p.View())
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("zero progress", func(t *testing.T) {
		assertOutput(ProgressBar{Width: 10, Total: 100, Current: 0, Color: lipgloss.Color("#fff")}, "          ")
	})

	t.Run("full progress", func(t *testing.T) {
		assertOutput(ProgressBar{Width: 10, Total: 100, Current: 100, Color: lipgloss.Color("#fff")}, "██████████")
	})

	t.Run("half progress", func(t *testing.T) {
		assertOutput(ProgressBar{Width: 10, Total: 100, Current: 50, Color: lipgloss.Color("#fff")}, "█████     ")
	})

	t.Run("fractional progress", func(t *testing.T) {
		assertOutput(ProgressBar{Width: 10, Total: 100, Current: 37.5, Color: lipgloss.Color("#fff")}, "███▊      ")
	})

	t.Run("clamp over 100%", func(t *testing.T) {
		assertOutput(ProgressBar{Width: 10, Total: 100, Current: 150, Color: lipgloss.Color("#fff")}, "██████████")
	})

	t.Run("zero width", func(t *testing.T) {
		assertOutput(ProgressBar{Width: 0, Total: 100, Current: 50, Color: lipgloss.Color("#fff")}, "")
	})

	t.Run("zero total", func(t *testing.T) {
		assertOutput(ProgressBar{Width: 10, Total: 0, Current: 50, Color: lipgloss.Color("#fff")}, "          ")
	})
}
