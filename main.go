// Package main запускает утилиту для расчета дней до наступления Нового года.
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func setupRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/calc", handleCalc)
	return mux
}

func handleCalc(w http.ResponseWriter, r *http.Request) {
	dateParam := r.URL.Query().Get("date")
	targetDate := time.Now()
	if dateParam != "" {
		parsed, err := time.Parse("2006-01-02", dateParam)
		if err != nil {
			http.Error(w, "Неверный формат даты. Ожидается YYYY-MM-DD", http.StatusBadRequest)
			return
		}
		targetDate = parsed
	}
	days := calcDaysToNewYear(targetDate)
	resp := TaskResponse{
		Date:          targetDate,
		Title:         "Дни до Нового года",
		DaysToNewYear: days,
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Printf("Ошибка отправки JSON: %v", err)
	}
}

func main() {
	handler := setupRouter()
	log.Println("Сервер работает на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

// calcDaysToNewYear определяет количество календарных дней, оставшихся
// от заданной даты t до 1 января следующего календарного года.
func calcDaysToNewYear(t time.Time) int {
	start := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
	nextYear := time.Date(t.Year()+1, time.January, 1, 0, 0, 0, 0, time.UTC)
	return int(nextYear.Sub(start) / (24 * time.Hour))
}

type TaskResponse struct {
	Date          time.Time `json:"date"`
	Title         string    `json:"title"`
	DaysToNewYear int       `json:"days_to_new_year"`
}
