package memory

import "hello-service/internal/domain/entity"

type MessageRepository struct{}

func NewMessageRepository() *MessageRepository {
	return &MessageRepository{}
}

func (r *MessageRepository) GetMessage() entity.Message {
	return entity.Message{Text: "Hello, World!"}
}
