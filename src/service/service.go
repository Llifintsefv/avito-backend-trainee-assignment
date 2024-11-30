package service

import (
	"math/rand"
	"pro-backend-trainee-assignment/src/repository"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type Service interface {
	GenerateRandomNumber(length int, Type string) (int64, error)
	GenerateRandomString(length int, Type string) (int64, error)
	GenerateRandomGUID(length int, Type string) (int64, error)
	GenerateRandomAlphanumeric(length int, Type string) (int64, error)
	GenerateRandomEnum(length int, Type string, values []string) (int64, error)
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) *service {
	return &service{repo: repo}
}

func (s *service) GenerateRandomNumber(length int, Type string) (int64, error) {
	number := generateNumber(length)

	id,err := s.repo.Generate(number,Type) 
	if err != nil {
		return  0,err
	}
	return id,nil
	
}

func (s *service) GenerateRandomString(length int, Type string) (int64, error) {
	valueString := generateString(length)

	id,err := s.repo.Generate(valueString,Type) 
	if err != nil {
		return  0,err
	}
	return id,nil

}


func (s *service) GenerateRandomGUID(length int, Type string) (int64,error) {
	guid := generateGUID(length)

	id,err := s.repo.Generate(guid,Type)
	if err != nil {
		return  0,err
	}
	return id,nil

}

func (s *service) GenerateRandomAlphanumeric(length int, Type string) (int64,error) {
	alphamuric := generateAlphanumeric(length)

	id,err := s.repo.Generate(alphamuric,Type)
	if err != nil {
		return  0,err
	}
	return id,nil
}

func (s *service) GenerateRandomEnum(length int, Type string, values []string) (int64,error) {
	alphamuric := generateAlphanumeric(length)

	id,err := s.repo.Generate(alphamuric,Type)
	if err != nil {
		return  0,err
	}
	return id,nil
}


func generateNumber(length int) (string) {
	rand.Seed(time.Now().UnixNano())
	if length != 0{
		strInt := strconv.Itoa(rand.Intn(length))
		return strInt
	}

	strInt := strconv.Itoa(rand.Intn(100000000))
	return strInt
}

func generateString(length int) (string) {
	RuneArr := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UnixNano())
	var result string
	for i := 0; i < length; i++ {
		result += string(RuneArr[rand.Intn(len(RuneArr))])
	}
	return result
}


func generateGUID(length int) (int, error) {
	guid := uuid.New().String()

	return guid, nil
}

func generateAlphanumeric(length int) (string, error) {
	RuneArr := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())
	var result string
	for i := 0; i < length; i++ {
		result += string(RuneArr[rand.Intn(len(RuneArr))])
	}
	return result
}