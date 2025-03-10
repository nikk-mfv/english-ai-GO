package handlers

import (
	"english-ai-be/db"
	"english-ai-be/internal/repository"
	"fmt"
	"net/http"
	"time"
)

type TEmployeeHandler struct {
	ID        int
	Name      string
	Email     string
	CreatedAt time.Time
}

var employeeHandler = TEmployeeHandler{}

func (h TEmployeeHandler) getEmployee(w http.ResponseWriter, req *http.Request) {
	// Allow all origins
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	employee, err := repository.GetEmployee(db.GetDB(), 1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, employee)
}
