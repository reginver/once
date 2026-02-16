package ui

import (
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithBackground(t *testing.T) {
	bg := color.RGBA{R: 26, G: 27, B: 38, A: 255}
	bgSeq := "\x1b[48;2;26;27;38m"

	t.Run("re-applies background after mid-line reset with visible content following", func(t *testing.T) {
		input := "\x1b[31mred\x1b[m more text"
		result := WithBackground(bg, input)
		expected := "\x1b[31mred\x1b[m" + bgSeq + " more text"
		assert.Equal(t, expected, result)
	})

	t.Run("does not re-apply after trailing reset", func(t *testing.T) {
		input := "\x1b[31mred\x1b[m"
		result := WithBackground(bg, input)
		assert.Equal(t, input, result)
	})

	t.Run("handles explicit zero param reset", func(t *testing.T) {
		input := "\x1b[31mred\x1b[0m more text"
		result := WithBackground(bg, input)
		expected := "\x1b[31mred\x1b[0m" + bgSeq + " more text"
		assert.Equal(t, expected, result)
	})

	t.Run("handles multiple lines independently", func(t *testing.T) {
		input := "\x1b[31mred\x1b[m more\n\x1b[32mgreen\x1b[m"
		result := WithBackground(bg, input)
		expected := "\x1b[31mred\x1b[m" + bgSeq + " more\n\x1b[32mgreen\x1b[m"
		assert.Equal(t, expected, result)
	})

	t.Run("does not touch non-reset SGR sequences", func(t *testing.T) {
		input := "\x1b[31mred\x1b[32mgreen"
		result := WithBackground(bg, input)
		assert.Equal(t, input, result)
	})

	t.Run("passes through plain text unchanged", func(t *testing.T) {
		input := "hello world"
		result := WithBackground(bg, input)
		assert.Equal(t, input, result)
	})
}
