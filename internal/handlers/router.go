package handlers

import (
	"net/http"
)

func Router() {
	http.HandleFunc("/api/v1/employee", employeeHandler.getEmployee)
	http.ListenAndServe(":8090", nil)
}
