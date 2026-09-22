package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("До Нового года осталось дней: %d\n", calcDaysToNewYear(now))
}

func calcDaysToNewYear(t time.Time) int {
	start := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
	nextYear := time.Date(t.Year()+1, time.January, 1, 0, 0, 0, 0, time.UTC)
	return int(nextYear.Sub(start) / (24 * time.Hour))
}
