package main

import (
	"testing"
	"time"
)

func TestCalcDaysToNewYear(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected int
	}{
		{
			name:     "Начало календарного года",
			input:    time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
			expected: 365,
		},
		{
			name:     "Конец календарного года",
			input:    time.Date(2023, time.December, 31, 23, 59, 59, 0, time.UTC),
			expected: 1,
		},
		{
			name:     "Високосный год (начало)",
			input:    time.Date(2024, time.January, 1, 12, 0, 0, 0, time.UTC),
			expected: 366,
		},
		{
			name:     "До 29 февраля (28 февраля високосного года)",
			input:    time.Date(2024, time.February, 28, 10, 0, 0, 0, time.UTC),
			expected: 308,
		},
		{
			name:     "В день 29 февраля",
			input:    time.Date(2024, time.February, 29, 0, 0, 0, 0, time.UTC),
			expected: 307,
		},
		{
			name:     "После 29 февраля (1 марта високосного года)",
			input:    time.Date(2024, time.March, 1, 15, 30, 0, 0, time.UTC),
			expected: 306,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calcDaysToNewYear(tt.input)
			if result != tt.expected {
				t.Errorf("%s: got %d, expected %d", tt.name, result, tt.expected)
			}
		})
	}
}
