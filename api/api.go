package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	// t := Response{
	// 	Message: "yo",
	// }
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintf(w, "TEST HANDLER IS READY")
}

type SendOneSmsResponse struct {
	Msisdn  string
	Success bool
}

func SmsOneHandler(w http.ResponseWriter, r *http.Request) {
	t := SendOneSmsResponse{
		Msisdn:  "123456789",
		Success: true,
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(t)
}

type SendBulkSmsResponse struct {
	Msisdn  []string
	Success bool
}

func SmsBulkHandler(w http.ResponseWriter, r *http.Request) {
	t := SendBulkSmsResponse{
		Msisdn: []string{"11111", "22222", "33333", "44444"},
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(t)
}

// Преобразуем дату в формат time.Time
func parseDate(input string) (*time.Time, error) {
	layout := "02.01.2006" // Формат DD.MM.YYYY
	t, err := time.Parse(layout, input)
	if err != nil {
		return nil, fmt.Errorf("некорректный формат даты: %v", err)
	}
	return &t, nil
}
