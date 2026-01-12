package listener

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/adi-QTPi/thestral/internal/config"
	"github.com/adi-QTPi/thestral/internal/model"
	"github.com/adi-QTPi/thestral/internal/proxy"
	"github.com/adi-QTPi/thestral/internal/store"
	"github.com/lib/pq"
)

type Service interface {
	Run() error
	Load() error
}

type service struct {
	cfg   *config.Env
	proxy proxy.Service
	store store.Service
}

func NewService(cfg *config.Env, p proxy.Service, s store.Service) Service {
	return &service{
		cfg:   cfg,
		proxy: p,
		store: s,
	}
}

func (s *service) Run() error {
	reportProblem := func(ev pq.ListenerEventType, err error) {
		if err != nil {
			fmt.Printf("Listener Error: %v", err)
		}
	}

	listener := pq.NewListener(s.cfg.DATABASE_URL, 10*time.Second, time.Minute, reportProblem) // exponential retry intervals

	go func() {
		if err := listener.Listen(model.ListenerName); err != nil {
			fmt.Printf("Failed to listen: %v", err)
			return
		}

		log.Println("Background Listener Started")

		for {
			select {
			case notif, ok := <-listener.Notify:
				if !ok {
					fmt.Println("Listener channel closed. Exiting loop.")
					return
				}

				if notif == nil {
					fmt.Println("Reconnected to DB. You might have missed events.")
					// [TODO] reload all data into local registry.
					continue
				}
				var payload model.EventPayload
				if err := json.Unmarshal([]byte(notif.Extra), &payload); err != nil {
					fmt.Println("Bad Payload detected")
					continue
				}

				switch payload.Action {
				case model.EventCreate:
					{
						s.handleCreateEvent(&payload)
					}
				case model.EventUpdate:
					{
						s.handleUpdateEvent(&payload)
					}
				case model.EventDelete:
					{
						s.handleDeleteEvent(&payload)
					}
				}

			// timer so that connection would not break
			case <-time.After(90 * time.Second):
				go listener.Ping()
			}
		}
	}()

	return nil
}
