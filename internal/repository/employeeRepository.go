package repository

import (
	"database/sql"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Email     string
	CreatedAt time.Time
}

func GetEmployee(db *sql.DB, id int) (Employee, error) {
	return Employee{}, nil
}
