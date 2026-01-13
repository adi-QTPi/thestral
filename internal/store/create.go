package store

import (
	"encoding/json"
	"fmt"

	"github.com/adi-QTPi/thestral/internal/admin/dto"
	"github.com/adi-QTPi/thestral/internal/model"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

// Creates a new route along with notifying the listening channels, returning any error leads to rollback.
func (s *service) Create(input dto.CreateRouteInput) error {

	// [TODO] some kind of host and target validation (ping maybe?)

	newRoute := &model.Route{
		Host:    input.Host,
		Targets: pq.StringArray(input.Targets),
	}
	f := func(tx *gorm.DB) error {
		if err := tx.Create(newRoute).Error; err != nil {
			// [TODO] return custom error on cases like key already exists, etc.
			return fmt.Errorf("%w", err)
		}

		payload := &model.EventPayload{
			Action: model.EventCreate,
			ID:     newRoute.ID,
			Host:   newRoute.Host,
		}
		data, err := json.Marshal(payload)
		if err != nil {
			return fmt.Errorf("CREATE notif payload marshal : %w", err)
		}

		if err := tx.Exec(NotifyQuery, model.ListenerName, string(data)).Error; err != nil {
			return fmt.Errorf("failed to emit notification: %w", err)
		}

		return nil
	}
	return s.db.Transaction(f)
}
