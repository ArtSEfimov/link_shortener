package statistic

import (
	"http_server/pkg/di"
	"http_server/pkg/event"
	"log"
)

type ServiceDeps struct {
	EventBus            *event.Bus
	StatisticRepository di.StatisticRepositoryInterface
}

type Service struct {
	EventBus            *event.Bus
	StatisticRepository di.StatisticRepositoryInterface
}

func NewService(deps ServiceDeps) *Service {
	return &Service{
		EventBus:            deps.EventBus,
		StatisticRepository: deps.StatisticRepository,
	}
}

func (service *Service) AddClick() {
	for msg := range service.EventBus.Subscribe() {
		if msg.Type == event.LinkVisited {
			id, ok := msg.Data.(uint)
			if !ok {
				log.Println("Bad EventLinkVisited Data", msg.Data)
				continue
			}
			service.StatisticRepository.AddClick(id)
		}
	}
}
