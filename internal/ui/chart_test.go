package ui

import (
	"fmt"
	"testing"
)

func TestChartView(t *testing.T) {
	// 80 data points = 40 chars wide (2 points per character)
	data := []float64{
		10, 15, 25, 20, 15, 25, 40, 35, 35, 40,
		50, 45, 45, 50, 30, 40, 55, 50, 60, 55,
		45, 55, 70, 65, 65, 70, 80, 75, 75, 80,
		90, 85, 85, 90, 70, 80, 95, 90, 100, 95,
		85, 90, 75, 80, 90, 85, 80, 85, 70, 75,
		85, 80, 75, 80, 65, 70, 80, 75, 70, 75,
		60, 65, 75, 70, 65, 70, 55, 60, 70, 65,
		60, 65, 50, 55, 65, 60, 55, 60, 45, 50,
	}

	chart := NewChart(40, 8, data)
	output := chart.View()

	fmt.Println("\nChart output:")
	fmt.Println(output)

	if output == "" {
		t.Error("expected non-empty chart output")
	}
}
