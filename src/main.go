package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"pro-backend-trainee-assignment/src/handler"
	rabbitmq "pro-backend-trainee-assignment/src/rabbitMQ"
	"pro-backend-trainee-assignment/src/repository"
	"pro-backend-trainee-assignment/src/service"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InitDB(connStr string) (*sql.DB,error) {
	

	db,err := sql.Open("postgres",connStr)
	if err != nil {
		return nil,fmt.Errorf("failed to open db connection: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil,fmt.Errorf("failed to ping db: %w", err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS random_values (id SERIAL PRIMARY KEY,guid TEXT, values TEXT, type TEXT, user_agent TEXT,requestid INTEGER,url TEXT,countRequest INTEGER)")
	if err != nil {
		return nil,fmt.Errorf("failed to create table: %w", err)
	}
	return db,nil
}



func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSL_MODE")

	if dbHost == "" || dbPort == "" || dbUser == "" || dbPassword == "" || dbName == "" || dbSSLMode == ""{
		 log.Fatal("DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME, DB_SSL_MODE must be set")
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPassword, dbHost, dbPort, dbName, dbSSLMode)
	
	db,err := InitDB(connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ch,err := rabbitmq.InitRabbitMQ()
	
	publisher,err := rabbitmq.NewPublisher(ch)
	if err != nil {
		log.Fatal(err)
	}

	
	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service,publisher)

	consumer,err := rabbitmq.NewConsumer(ch,service)
	if err != nil {
		log.Fatal(err)
	}
	go consumer.ConsumeGeneratedValue()


	r := mux.NewRouter()

	r.HandleFunc("/api/generate",handler.GenerateHandler).Methods("POST")
	r.HandleFunc("/api/retrieve/{id}",handler.RetrieveHandler).Methods("GET")

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}