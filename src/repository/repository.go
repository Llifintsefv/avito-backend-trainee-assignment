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

func (r *repository) Generate(randomNumber int) (error) {
	stmt,err := r.db.Prepare("INSERT INTO random_values (value) VALUES (?)")
	if err != nil {
		return fmt.Errorf("failte to prepare statement: %w",err)
	}
	defer stmt.Close()
	_,err = stmt.Exec(randomNumber)
	if err != nil {
		return fmt.Errorf("failte to insert value: %w",err)
	}
	return nil
} 

func (r *repository) Retrieve(id int) (int, error) {
	stmt, err := r.db.Prepare("SELECT values FROM random_values WHERE id = ?")
	if err != nil {
		return 0, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return 0, fmt.Errorf("failed to query value: %w", err)
	}
	defer rows.Close() 

	if rows.Next() {
		var randomValue int
		err := rows.Scan(&randomValue)
		if err != nil {
			return 0, fmt.Errorf("failed to scan value: %w", err)
		}
		return randomValue, nil
	}

	if err := rows.Err(); err != nil {
        return 0, fmt.Errorf("error during rows iteration: %w", err)
    }
    
	return 0, fmt.Errorf("value with id %d not found", id) 
}