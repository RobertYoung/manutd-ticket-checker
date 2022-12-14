package haas

import (
	"fmt"
	"log"
)

type HomeAssistantStateUpdateRequest struct {
	State     string      `json:"state"`
	Attribute interface{} `json:"attributes"`
}

func (s *HomeAssistantAPI) StateUpdate(entity_id string, request HomeAssistantStateUpdateRequest) {
	log.Printf("updating state for %s | %s", entity_id, request.State)

	s.Post(fmt.Sprintf("api/states/%s", entity_id), request)
}
