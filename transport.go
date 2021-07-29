package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Service interface {
	SendMessage(token string, chatId int64, message string) error
	SendPhoto(token string, chatId int64, fileUrl string) error
	SendFile(token string, chatId int64, fileUrl string) error
}

type Transport struct {
	service Service
}

func NewTransport(service Service) *Transport {
	return &Transport{service: service}
}

func (t *Transport) Routes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.SetHeader("Content-Type", "application/json"))
	r.Post(`/send-message`, t.sendMessage)
	r.Post(`/send-photo`, t.sendPhoto)

	return r
}

func (t *Transport) sendMessage(w http.ResponseWriter, r *http.Request) {
	var req sendMessageRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := t.service.SendMessage(req.Token, req.ChatId, req.Message); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := serviceResponse{
		Ok:      true,
		Message: "Message sent",
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (t *Transport) sendPhoto(w http.ResponseWriter, r *http.Request) {
	var req sendFileRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := t.service.SendPhoto(req.Token, req.ChatId, req.FileUrl); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := serviceResponse{
		Ok:      true,
		Message: "Message with a photo sent",
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type serviceResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

type sendMessageRequest struct {
	ChatId  int64  `json:"chat_id"`
	Token   string `json:"token"`
	Message string `json:"message"`
}

type sendFileRequest struct {
	ChatId  int64  `json:"chat_id"`
	Token   string `json:"token"`
	FileUrl string `json:"file_url"`
}
