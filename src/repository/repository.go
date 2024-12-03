package repository

import (
	"database/sql"
	"fmt"
	"pro-backend-trainee-assignment/src/models"
)

type Repository interface {
	Generate(models.GenerateValue) (error)
	Retrieve(string) (string,string,error)
	GetCountRequest(int) (int, error)
	UpdateCountRequestAndRetrieveId(int,int) (string,error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Generate(GenValue models.GenerateValue) (error) {
	
    stmt, err := r.db.Prepare("INSERT INTO random_values (guid,values, type,user_agent,requestid,url,countRequest) VALUES ($1, $2,$3,$4,$5,$6,$7)")
    if err != nil {
        return fmt.Errorf("fail to prepare statement: %w", err) 
    }
    defer stmt.Close()

    _,err = stmt.Exec(GenValue.ID,GenValue.Value,GenValue.Type,GenValue.UserAgent,GenValue.RequestId,GenValue.Url,GenValue.CountRequest)
    if err != nil {
        return fmt.Errorf("fail to insert value: %w", err) 
    }

    return nil
}





func (r *repository) Retrieve(id string) (string,string, error) {
	stmt, err := r.db.Prepare("SELECT values,type FROM random_values WHERE guid = $1")
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
    
	return "","", nil

}

func (r *repository) GetCountRequest(requestId int) (int, error) {
	var countRequest int
	err := r.db.QueryRow("SELECT countRequest FROM random_values WHERE requestid = $1", requestId).Scan(&countRequest)

	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, fmt.Errorf("failed to get countRequest: %w", err)
	}
	return countRequest, nil
}

func (r *repository) UpdateCountRequestAndRetrieveId(requestId int, countRequest int) (string,error){
	var id string
    err := r.db.QueryRow("UPDATE random_values SET countRequest = $2 WHERE requestid = $1 RETURNING guid", requestId, countRequest).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("failed to update countRequest: %w", err)
	}
	return id,nil
}


