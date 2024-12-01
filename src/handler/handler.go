package handler

import (
	"encoding/json"
	"net/http"
	"pro-backend-trainee-assignment/src/models"
	"pro-backend-trainee-assignment/src/service"

	"github.com/gorilla/mux"
)


type Handler struct {
	service service.Service
}

func NewHandler(service service.Service) *Handler{
	return &Handler{service: service}
}

var (
	GenRequest models.GenRequest
	RetrieveRequest models.RetrieveRequest
	RetrieveResponse models.Response
)


func (h *Handler)GenerateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w,"method not allowd",http.StatusMethodNotAllowed)
		return
	}
	var req models.GenRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w,"bad request",http.StatusBadRequest)
		return
	}

	if req.Type != "string" && req.Type != "number" && req.Type != "guid" && req.Type != "alphanumeric" && req.Type != "enum" {
		http.Error(w,"invalid type",http.StatusBadRequest)
		return
	}

	req.UserAgent = r.Header.Get("User-Agent")
	req.Url = r.URL.String()

	GenResponse,err := h.service.GenerateNumber(req)


	if err != nil {
		http.Error(w,"failed to generate " + req.Type,http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(GenResponse); err != nil {
		http.Error(w,"internal server error",http.StatusInternalServerError)
		return
	}

}

func (h *Handler)RetrieveHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id:= params["id"]

	 if id == "" {
         http.Error(w, "missing id parameter", http.StatusBadRequest)
         return
     }

	Response,err := h.service.Retrieve(id)
	if err != nil {
		http.Error(w,"internal server error",http.StatusInternalServerError)
		return
	}
	
	if err := json.NewEncoder(w).Encode(Response); err != nil {
		http.Error(w,"internal server error",http.StatusInternalServerError)
		return
	}
	
}
