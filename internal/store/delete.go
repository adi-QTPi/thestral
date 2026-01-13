package store

import (
	"encoding/json"
	"fmt"

	"github.com/adi-QTPi/thestral/internal/admin/dto"
	"github.com/adi-QTPi/thestral/internal/model"
	"gorm.io/gorm"
)

// Deletes a route along with notifying the listening channels, returning any error leads to rollback.
func (s *service) Delete(input dto.DeleteRouteInput) error {
	f := func(tx *gorm.DB) error {
		result := tx.Unscoped().Where("host = ?", input.Host).Delete(&model.Route{}) // gorm hard delete

		if result.Error != nil {
			return fmt.Errorf("%w", result.Error)
		}

		if result.RowsAffected == 0 {
			return nil
		}

		payload := &model.EventPayload{
			Action: model.EventDelete,
			Host:   input.Host,
		}
		data, err := json.Marshal(payload)
		if err != nil {
			return fmt.Errorf("DELETE notif payload marshal error : %w", err)
		}

		if err := tx.Exec(NotifyQuery, model.ListenerName, string(data)).Error; err != nil {
			return fmt.Errorf("failed to emit notification: %w", err)
		}

		return nil
	}

	return s.db.Transaction(f)
}
