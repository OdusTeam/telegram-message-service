package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
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
	panic("implement me")
}

func (s *TelegramService) SendFile(token string, chatId int64, fileUrl string) error {
	panic("implement me")
}
