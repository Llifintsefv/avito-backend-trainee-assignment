package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pro-backend-trainee-assignment/src/service"
	"strconv"

	"github.com/gorilla/mux"
)


type Handler struct {
	service service.Service
}

func NewHandler(service service.Service) *Handler{
	return &Handler{service: service}
}

type GenRequest struct {
	Type string `json:"type"`
	Length int  `json:"length,omitempty"`
	Values []string `json:"values,omitempty"`
}



type RetrieveRequest struct {
	id int `json:"id"`
}

type RetrieveResponse struct {
	value interface{} `json:"value"`
}


func (h *Handler)GenerateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w,"method not allowd",http.StatusMethodNotAllowed)
		return
	}
	var req GenRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w,"bad request",http.StatusBadRequest)
		return
	}
	fmt.Println(req)

	if req.Type != "string" && req.Type != "number" && req.Type != "guid" && req.Type != "alphanumeric" && req.Type != "enum" {
		http.Error(w,"invalid type",http.StatusBadRequest)
		return
	}

	var GenResponse int64
	switch req.Type {
	case "string":
		GenResponse,err = h.service.GenerateRandomString(req.Length,req.Type)
	case "number":
		GenResponse,err = h.service.GenerateRandomNumber(req.Length,req.Type)
	case "guid":
		GenResponse,err = h.service.GenerateRandomGUID(req.Type)
	case "alphanumeric":
		GenResponse,err = h.service.GenerateRandomAlphanumeric(req.Length,req.Type)
	case "enum":
		if len(req.Values) == 0 {
			http.Error(w,"values cannot be empthy for enum",http.StatusBadRequest)
		}
		GenResponse,err = h.service.GenerateRandomEnum(req.Values,req.Type)
	}

	if err != nil {
		http.Error(w,"failed to generate " + req.Type,http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type","application/json")
	if err := json.NewEncoder(w).Encode(GenResponse); err != nil {
		http.Error(w,"internal server error",http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}

func (h *Handler)RetrieveHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id,err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w,"bad request",http.StatusBadRequest)
		return
	}

	value,err := h.service.Retrieve(id)
	if err != nil {
		http.Error(w,"internal server error",http.StatusInternalServerError)
		return
	}
	fmt.Println(value)
	var retrieveResponse interface{}
	switch v := value.(type) {
	case string:
		retrieveResponse = v 
	case int:
		retrieveResponse = v
	}


	w.Header().Set("Content-type","application/json")
	if err := json.NewEncoder(w).Encode(retrieveResponse); err != nil {
		http.Error(w,"internal server error",http.StatusInternalServerError)
		return
	}
	
}
