package ui

import (
	"strings"

	"charm.land/lipgloss/v2"
)

// Chart renders a histogram-style chart using braille characters.
// Each data point is one character wide, and the height scales dynamically
// so the maximum value fills the available height.
type Chart struct {
	Width  int
	Height int
	Data   []float64
	Color  lipgloss.Style
}

// braille bit patterns for left and right columns
// Each column has 4 dots, allowing 2 data points per character.
// Left column dots (bottom to top): 7, 3, 2, 1
// Right column dots (bottom to top): 8, 6, 5, 4
var (
	leftDots  = [4]rune{0x40, 0x04, 0x02, 0x01} // dots 7, 3, 2, 1
	rightDots = [4]rune{0x80, 0x20, 0x10, 0x08} // dots 8, 6, 5, 4
)

func NewChart(width, height int, data []float64) Chart {
	return Chart{
		Width:  width,
		Height: height,
		Data:   data,
		Color:  lipgloss.NewStyle().Foreground(Colors.Secondary),
	}
}

func (c Chart) View() string {
	if len(c.Data) == 0 || c.Width == 0 || c.Height == 0 {
		return ""
	}

	maxVal := c.maxValue()
	if maxVal == 0 {
		maxVal = 1
	}

	// Each character row represents 4 vertical dots
	dotsHeight := c.Height * 4

	// Calculate the height in dots for each data point
	heights := make([]int, len(c.Data))
	for i, v := range c.Data {
		heights[i] = int((v / maxVal) * float64(dotsHeight))
		if v > 0 && heights[i] == 0 {
			heights[i] = 1 // ensure non-zero values show at least 1 dot
		}
	}

	// Build the chart row by row, from top to bottom
	// Each character holds 2 data points (left and right columns)
	var rows []string
	for row := range c.Height {
		var sb strings.Builder
		rowBottomDot := (c.Height - 1 - row) * 4
		rowTopDot := rowBottomDot + 4

		for col := range c.Width {
			dataIdxLeft := col * 2
			dataIdxRight := col*2 + 1

			var char rune = 0x2800 // braille base character

			// Left column (first data point)
			if dataIdxLeft < len(heights) {
				char |= brailleColumn(heights[dataIdxLeft], rowBottomDot, rowTopDot, leftDots)
			}

			// Right column (second data point)
			if dataIdxRight < len(heights) {
				char |= brailleColumn(heights[dataIdxRight], rowBottomDot, rowTopDot, rightDots)
			}

			sb.WriteRune(char)
		}
		rows = append(rows, c.Color.Render(sb.String()))
	}

	return strings.Join(rows, "\n")
}

// brailleColumn returns the braille bits for a single column based on height
func brailleColumn(h, rowBottom, rowTop int, dots [4]rune) rune {
	if h <= rowBottom {
		return 0
	}

	var bits rune
	dotsToFill := min(h-rowBottom, 4)
	for i := range dotsToFill {
		bits |= dots[i]
	}
	return bits
}

func (c Chart) maxValue() float64 {
	if len(c.Data) == 0 {
		return 0
	}
	max := c.Data[0]
	for _, v := range c.Data[1:] {
		if v > max {
			max = v
		}
	}
	return max
}
