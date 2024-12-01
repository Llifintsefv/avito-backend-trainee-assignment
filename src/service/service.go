package service

import (
	"errors"
	"math/rand"
	"pro-backend-trainee-assignment/src/models"
	"pro-backend-trainee-assignment/src/repository"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Service interface {
	GenerateNumber(models.GenRequest) (models.Response,error)
	Retrieve(models.RetrieveRequest)(models.Response,error)
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) *service {
	return &service{repo: repo}
}

func (s *service) Retrieve(RetrieveRequest models.RetrieveRequest) (models.Response, error) {
	Response,err := s.GetValueById(RetrieveRequest.ID)
	if err != nil {
		return models.Response{}, err
	}

	var GenValue models.GenerateValue

	GenValue.ID = Response.Id
	GenValue.UserAgent = RetrieveRequest.UserAgent
	GenValue.Url = RetrieveRequest.Url
	GenValue.RequestId = RetrieveRequest.RequestId
	err = s.repo.Generate(GenValue)
	
	return Response, nil 
}

func (s *service) GetValueById(id string) (models.Response, error) {
	value, Type, err := s.repo.Retrieve(id)
	if err != nil {
		return models.Response{}, err
	}

	response := models.Response{Id: id}

	if Type == "number" {
		valueInt, err := strconv.Atoi(value)
		if err != nil {
			return models.Response{}, err 
		}
		response.Value = valueInt 
	} else {
		response.Value = value 
	}

	return response, nil 
}


func (s *service) GenerateNumber(req models.GenRequest) (models.Response,error) {
	var GenValue models.GenerateValue
	var err error
	
	GenValue.RequestId = req.RequestId

	GenValue.CountRequest,err = s.repo.GetCountRequest(GenValue.RequestId)
	if err != nil {
		return models.Response{}, errors.New("Error to find countRequest")
	}

	if GenValue.CountRequest != 0 {
		id := s.repo.UpdateCountRequestAndRetrieveId(GenValue.RequestId,GenValue.CountRequest+1)
		var Response models.Response
		Response,err = s.GetValueById(id)
		if err != nil {
			return Response, errors.New("Error to find value")
		}
		return Response, nil
		// вернуть надо уже сгенерированное
	}

	GenValue.ID = generateGUID()
	GenValue.Type = req.Type
	GenValue.UserAgent = req.UserAgent
	GenValue.Url = req.Url


	switch req.Type {
	case "string":
		GenValue.Value = generateString(req.Length)
	case "number":
		GenValue.Value = generateNumber(req.Length)
	case "guid":
		GenValue.Value = generateGUID()
	case "alphanumeric":
		GenValue.Value= generateAlphanumeric(req.Length)
	case "enum":
		if len(req.Values) == 0 {
			return models.Response{}, errors.New("values cannot be empthy for enum")
		}
		GenValue.Value = generateEnum(req.Values)
	}

	err = s.repo.Generate(GenValue) 
	if err != nil {
		return models.Response{}, errors.New("Error to generate value")
	}
	
	var Response models.Response
	Response = models.Response{Id: GenValue.ID, Value: GenValue.Value}
	return Response, nil


}


func generateNumber(length int) (string) {
	rand.Seed(time.Now().UnixNano())
	var result string
	for i := 0; i < length; i++ {
		result += strconv.Itoa(rand.Intn(10))
	}
	return result
	
}

func generateString(length int) (string) {
	rand.Seed(time.Now().UnixNano())
	RuneArr := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
    var sb strings.Builder
    sb.Grow(length)
    for i := 0; i < length; i++ {
        sb.WriteRune(RuneArr[rand.Intn(len(RuneArr))])
    }
    return sb.String()
}


func generateGUID() (string) {
	guid := uuid.New().String()

	return guid
}

func generateAlphanumeric(length int) (string) {
	RuneArr := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UnixNano())
	var sb strings.Builder
    sb.Grow(length)
    for i := 0; i < length; i++ {
        sb.WriteRune(RuneArr[rand.Intn(len(RuneArr))])
    }
    return sb.String()
}

func generateEnum(values []string) (string) {
	randomIndex := rand.Intn(len(values))
    return values[randomIndex]
}