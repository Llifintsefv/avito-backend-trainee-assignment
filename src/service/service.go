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
	GenerateRandomGUID(Type string) (int64, error)
	GenerateRandomAlphanumeric(length int, Type string) (int64, error)
	GenerateRandomEnum(values []string, Type string) (int64,error)

	Retrieve(id int)(interface{},error)
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) *service {
	return &service{repo: repo}
}

func (s *service) Retrieve(id int)(interface{},error){
	value,Type,err := s.repo.Retrieve(id)
	if err != nil {
		return "",err
	}
	if Type == "number" {
		valueInt,err := strconv.Atoi(value)
		if err != nil {
			return "",err
		}
		return valueInt,nil
	}

	return value,nil


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


func (s *service) GenerateRandomGUID(Type string) (int64,error) {
	guid := generateGUID()

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

func (s *service) GenerateRandomEnum(values []string,Type string) (int64,error) {
	enum := generateEnum(values)

	id,err := s.repo.Generate(enum,Type)
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
	RuneArr := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())
	var result string
	for i := 0; i < length; i++ {
		result += string(RuneArr[rand.Intn(len(RuneArr))])
	}
	return result
}


func generateGUID() (string) {
	guid := uuid.New().String()

	return guid
}

func generateAlphanumeric(length int) (string) {
	RuneArr := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UnixNano())
	var result string
	for i := 0; i < length; i++ {
		result += string(RuneArr[rand.Intn(len(RuneArr))])
	}
	return result
}

func generateEnum(values []string) (string) {
	randomIndex := rand.Intn(len(values))
    return values[randomIndex]
}