package handler

import (
	"math/rand"
	"net/http"
	"time"
)

func generateHandler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(1000000)
	
}