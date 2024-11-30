package repository

import (
	"database/sql"
	"fmt"
)

type Repository interface {
	Generate(string,string) (int64,error)
	Retrieve(int) (int,string,error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Generate(randomNumber string, Type string) (int64, error) {
    stmt, err := r.db.Prepare("INSERT INTO random_values (value, type) VALUES (?, ?)")
    if err != nil {
        return 0, fmt.Errorf("fail to prepare statement: %w", err) 
    }
    defer stmt.Close()

    result, err := stmt.Exec(randomNumber, Type)
    if err != nil {
        return 0, fmt.Errorf("fail to insert value: %w", err) 
    }

    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("fail to get last insert id: %w", err)
    }

    return id, nil
}
func (r *repository) Retrieve(id int) (int,string, error) {
	stmt, err := r.db.Prepare("SELECT values,type FROM random_values WHERE id = ?")
	if err != nil {
		return 0,"", fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return 0,"", fmt.Errorf("failed to query value: %w", err)
	}
	defer rows.Close() 

	if rows.Next() {
		var randomValue int
		var Type string
		err := rows.Scan(&randomValue,&Type)
		if err != nil {
			return 0,"", fmt.Errorf("failed to scan value: %w", err)
		}
		return randomValue,Type, nil
	} 

	if err := rows.Err(); err != nil {
        return 0,"", fmt.Errorf("error during rows iteration: %w", err)
    }
    
	return 0,"", fmt.Errorf("value with id %d not found", id) 
}