package ui

import (
	"image/color"

	"charm.land/lipgloss/v2"
)

type colors struct {
	Primary    color.Color
	Secondary  color.Color
	Background color.Color
	Text       color.Color
	TextDark   color.Color
}

var Colors = colors{
	Primary:    lipgloss.Color("#FF69B4"),
	Secondary:  lipgloss.Color("#9f3"),
	Background: lipgloss.Color("#000000"),
	Text:       lipgloss.Color("#FFFFFF"),
	TextDark:   lipgloss.Color("#000000"),
}

type styles struct {
	Title    lipgloss.Style
	SubTitle lipgloss.Style
}

var Styles = styles{
	Title: lipgloss.NewStyle().
		Background(Colors.Primary).
		Foreground(Colors.TextDark).
		Bold(true),
	SubTitle: lipgloss.NewStyle().
		Foreground(Colors.Secondary).
		Underline(true),
}
