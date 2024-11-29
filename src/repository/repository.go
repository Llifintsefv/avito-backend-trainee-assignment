package repository

import (
	"database/sql"
	"fmt"
)

type Repository interface {
	Generate(int) (error)
	Retrieve(int) (int,error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Generate() (error) {
	stmt,err := r.db.Prepare("INSERT INTO random_values (value) VALUES (?)")
	if err != nil {
		return fmt.Errorf("failte to prepare statement: %w",err)
	}
	defer stmt.Close()

} 