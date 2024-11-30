package repository

import (
	"database/sql"
	"fmt"
)

type Repository interface {
	Generate(string,string) (int64,error)
	Retrieve(int) (string,string,error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Generate(randomNumber string, Type string) (int64, error) {
	
    stmt, err := r.db.Prepare("INSERT INTO random_values (values, type) VALUES ($1, $2) RETURNING id")
    if err != nil {
        return 0, fmt.Errorf("fail to prepare statement: %w", err) 
    }
    defer stmt.Close()

	var id int64
    err = stmt.QueryRow(randomNumber, Type).Scan(&id)
    if err != nil {
        return 0, fmt.Errorf("fail to insert value: %w", err) 
    }


    return id, nil
}
func (r *repository) Retrieve(id int) (string,string, error) {
	stmt, err := r.db.Prepare("SELECT values,type FROM random_values WHERE id = $1")
	if err != nil {
		return "","", fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return "","", fmt.Errorf("failed to query value: %w", err)
	}
	defer rows.Close() 

	if rows.Next() {
		var value string
		var Type string
		err := rows.Scan(&value,&Type)
		if err != nil {
			return "","", fmt.Errorf("failed to scan value: %w", err)
		}
		return value,Type, nil
	} 

	if err := rows.Err(); err != nil {
        return "","", fmt.Errorf("error during rows iteration: %w", err)
    }
    
	return "","", fmt.Errorf("value with id %d not found", id) 
}