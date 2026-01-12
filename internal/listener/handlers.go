package listener

import (
	"fmt"

	"github.com/adi-QTPi/thestral/internal/model"
	"gorm.io/gorm"
)

func (s *service) handleCreateEvent(payload *model.EventPayload) {
	filter := &model.Route{
		Model: gorm.Model{
			ID: payload.ID,
		},
	}
	result, err := s.store.FindOneRoute(filter)
	if err != nil {
		fmt.Printf("error executing : %v", err)
	}

	if err := s.proxy.Create(result); err != nil {
		fmt.Printf("error adding route : %v", err)
	}

}

func (s *service) handleUpdateEvent(payload *model.EventPayload) {
	// [TODO] registry op after updation in db
}

func (s *service) handleDeleteEvent(payload *model.EventPayload) {
	// [TODO] registry op after deletion in db
}
