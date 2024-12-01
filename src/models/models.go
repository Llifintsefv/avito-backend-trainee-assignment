package models

type GenRequest struct {
	Type         string   `json:"type"`
	Length       int      `json:"length,omitempty"`
	Values       []string `json:"values,omitempty"`
	UserAgent    string   `json:"user-agent,omitempty"`
	RequestId    int      `json:"requestId,omitempty"`
	Url          string   `json:"url,omitempty"`
}


type RetrieveRequest struct {
	ID           string `json:"id"`
	UserAgent    string   `json:"user-agent,omitempty"`
	RequestId    int      `json:"requestId,omitempty"`
	Url          string   `json:"url,omitempty"`
}

type Response struct {
	Id string        `json:"id"`
	Value interface{} `json:"value"`
}

type GenerateValue struct {
	ID           string `json:"id"`
	Type         string `json:"type"`
	Length       int    `json:"length,omitempty"`
	Value        string `json:"value,omitempty"`
	UserAgent    string `json:"user-agent,omitempty"`
	RequestId    int    `json:"requestId,omitempty"`
	Url          string `json:"url,omitempty"`
	CountRequest int    `json:"countRequest,omitempty"`
}