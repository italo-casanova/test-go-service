package application

import (
	"hello-service/internal/domain"
	"hello-service/internal/domain/entity"
)

type HelloUsecase struct {
	messageRepo domain.MessageRepository
}

func NewHelloUsecase(repo domain.MessageRepository) *HelloUsecase {
	return &HelloUsecase{messageRepo: repo}
}

func (u *HelloUsecase) GetHelloMessage() entity.Message {
	return u.messageRepo.GetMessage()
}
