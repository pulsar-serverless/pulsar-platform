package container

import (
	"context"
	"errors"
	"fmt"
	"pulsar/internal/core/services"
	"time"

	"github.com/rs/zerolog/log"
)

func (cs *containerService) changeAppStatus(containerId string) error {
	containerInfo, ok := cs.liveContainers[containerId]

	if !ok {
		log.Info().Str("containerId", containerId).Msg("No app state found: creating a default app state")
		containerInfo = &ContainerInfo{
			lastAccessed:  time.Now(),
			server:        make(chan bool),
			isServerAlive: true,
		}

		cs.liveContainers[containerId] = containerInfo

		go func(containerId string) {
			log.Info().Str("containerId", containerId).Msg(fmt.Sprintf("scheduled app to stop after %v", cs.maxContainerAge))
			time.Sleep(cs.maxContainerAge)
			cs.end <- containerId
		}(containerId)

		return nil
	}

	timeout := time.NewTicker(cs.operationsTimeout)

	select {
	case containerInfo.server <- true:
		log.Info().Str("containerId", containerId).Msg("Announce the app inside the container is alive")
		return nil
	case <-timeout.C:
		log.Info().Str("containerId", containerId).Msg("Timeout listener for app state did not respond")
		return services.NewAppError(services.ErrInternalServer, errors.New("Unable to access app state in time"))
	}
}

func (cs *containerService) ChangeAppStatus(ctx context.Context, subdomain string) error {
	project, err := cs.projectRepo.GetProjectBySubdomain(ctx, subdomain)
	if err != nil {
		return services.NewAppError(services.ErrNotFound, err)
	}

	return cs.changeAppStatus(project.ContainerId)
}
