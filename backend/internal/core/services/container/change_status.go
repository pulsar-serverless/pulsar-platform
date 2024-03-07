package container

import (
	"context"
	"errors"
	"fmt"
	domain "pulsar/internal/core/domain/log"
	"pulsar/internal/core/domain/project"
	"pulsar/internal/core/services"
	"time"
)

func (cs *containerService) changeAppStatus(project *project.Project) error {
	containerInfo, ok := cs.liveContainers[project.ContainerId]

	if !ok {
		cs.logService.CreateLogEvent(context.Background(), domain.NewAppLog(
			project.ID,
			domain.WARNING,
			"App state not found; creating a default app state.",
		))

		containerInfo = &ContainerInfo{
			lastAccessed:  time.Now(),
			server:        make(chan bool),
			isServerAlive: true,
		}

		cs.liveContainers[project.ContainerId] = containerInfo

		go func(containerId string) {
			cs.logService.CreateLogEvent(context.Background(), domain.NewAppLog(
				project.ID,
				domain.WARNING,
				fmt.Sprintf("Scheduled app to stop after %v.", cs.maxContainerAge),
			))
			time.Sleep(cs.maxContainerAge)
			cs.end <- project
		}(project.ContainerId)

		return nil
	}

	timeout := time.NewTicker(cs.operationsTimeout)

	select {
	case containerInfo.server <- true:
		return nil
	case <-timeout.C:
		return services.NewAppError(services.ErrInternalServer, errors.New("unable to access app state in time"))
	}
}

func (cs *containerService) ChangeAppStatus(ctx context.Context, appId string) error {
	project, err := cs.projectRepo.GetProject(ctx, appId)
	if err != nil {
		cs.logService.CreateLogEvent(context.Background(), domain.NewAppLog(
			project.ID,
			domain.Error,
			"Timeout listener for app state did not respond. Unable to access app state in time.",
		))
		return services.NewAppError(services.ErrNotFound, err)
	}

	return cs.changeAppStatus(project)
}
