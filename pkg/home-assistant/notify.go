package haas

import (
	"fmt"
	"log"
)

type HomeAssistantNotifyRequest struct {
	Message string `json:"message"`
	Title   string `json:"title"`
}

func (s *HomeAssistantAPI) Notify(device string, request HomeAssistantNotifyRequest) {
	log.Printf("sending notification to %s | %s: %s", device, request.Title, request.Message)

	_, err := s.Post(fmt.Sprintf("api/services/notify/%s", device), request)

	if err == nil {
		log.Printf("sent to %s | %s: %s", device, request.Title, request.Message)
	}
}
