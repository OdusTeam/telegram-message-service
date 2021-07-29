package main

import (

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramService struct {
}

func NewTelegramService() *TelegramService {
	return &TelegramService{}
}

func (s *TelegramService) SendMessage(token string, chatId int64, message string) error {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}

	msg := tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID: chatId,
		},
		Text: message,
	}

	_, err = api.Send(msg)
	return err
}

func (s *TelegramService) SendPhoto(token string, chatId int64, fileUrl string) error {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}

	file := tgbotapi.FileURL(fileUrl)

	msg := tgbotapi.NewPhoto(chatId, file)

	_, err = api.Send(msg)
	return err
}

func (s *TelegramService) SendFile(token string, chatId int64, fileUrl string) error {
	panic("implement me")
}
