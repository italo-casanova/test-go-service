package domain

import "hello-service/internal/domain/entity"

type MessageRepository interface {
	GetMessage() entity.Message
}
