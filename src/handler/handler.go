package handler

import (
	"encoding/json"
	"net/http"
	"pro-backend-trainee-assignment/src/models"
	rabbitmq "pro-backend-trainee-assignment/src/rabbitMQ"
	"pro-backend-trainee-assignment/src/service"
	"pro-backend-trainee-assignment/src/utils"

	"github.com/gorilla/mux"
)


type Handler struct {
	service service.Service
	publisher *rabbitmq.Publisher
}

func NewHandler(service service.Service, publisher *rabbitmq.Publisher) *Handler{
	return &Handler{
		service: service,
		publisher: publisher,
	}
}


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

	req.ID = utils.GenerateGUID()
	req.UserAgent = r.Header.Get("User-Agent")
	req.Url = r.URL.String()


	reqJson,err := json.Marshal(req)
	if err != nil {
		http.Error(w,"internal server error",http.StatusInternalServerError)
		return
	}
	err = h.publisher.PublishGenerateValue(reqJson)
	if err != nil {
		http.Error(w,"internal server error",http.StatusInternalServerError)
		return
	}


	if err := json.NewEncoder(w).Encode(req.ID); err != nil {
		http.Error(w,"internal server error",http.StatusInternalServerError)
		return
	}

}

func (h *Handler)RetrieveHandler(w http.ResponseWriter, r *http.Request) {
	var RetrieveRequest models.RetrieveRequest

	params := mux.Vars(r)
	RetrieveRequest.ID = params["id"]

	 if RetrieveRequest.ID == "" {
         http.Error(w, "missing id parameter", http.StatusBadRequest)
         return
     }
	
	RetrieveRequest.UserAgent = r.Header.Get("User-Agent")
	RetrieveRequest.Url = r.URL.String()

	Response,err := h.service.Retrieve(RetrieveRequest)
	if err != nil {
		http.Error(w,"internal server error",http.StatusInternalServerError)
		return
	}
	
	if err := json.NewEncoder(w).Encode(Response); err != nil {
		http.Error(w,"internal server error",http.StatusInternalServerError)
		return
	}
	
}
