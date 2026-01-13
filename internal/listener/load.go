package listener

import (
	"fmt"

	"github.com/adi-QTPi/thestral/internal/model"
)

func (s *service) Load() error {

	data, err := s.store.FindManyRoutes(&model.Route{})
	if err != nil {
		fmt.Println("Listener Loading Error: ", err)
	}

	s.proxy.BulkLoad(data)

	return nil
}
