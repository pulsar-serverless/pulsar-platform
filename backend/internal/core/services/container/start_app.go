package container

import (
	"context"
	"fmt"
	domain "pulsar/internal/core/domain/log"
	"pulsar/internal/core/domain/project"
	"pulsar/internal/core/services"
	"time"
)

func (cs *containerService) startServerlessApp(ctx context.Context, project *project.Project, success chan bool, errorChan chan error) {
	containerInfo, ok := cs.liveContainers[project.ContainerId]

	if !ok {
		// start the serverless container
		cs.logService.CreateLogEvent(context.Background(), domain.NewAppLog(
			project.ID,
			domain.INFO,
			"Container not started; Starting app container.",
		))

		err := cs.containerMan.StartContainer(ctx, project.ContainerId)
		if err != nil {
			cs.logService.CreateLogEvent(context.Background(), domain.NewAppLog(
				project.ID,
				domain.Error,
				fmt.Sprintf("Unable to start app container: %v", err),
			))
			errorChan <- services.NewAppError(services.ErrInternalServer, err)
		}

		go cs.saveContainerLogs(project)

		containerInfo = &ContainerInfo{
			lastAccessed:  time.Now(),
			server:        make(chan bool),
			isServerAlive: false,
		}

		cs.liveContainers[project.ContainerId] = containerInfo
	} else {
		cs.logService.CreateLogEvent(context.Background(), domain.NewAppLog(
			project.ID,
			domain.Error,
			"Starting app container; container already started.",
		))
		containerInfo.lastAccessed = time.Now()
	}

	// schedule app closing
	go func() {
		cs.logService.CreateLogEvent(context.Background(), domain.NewAppLog(
			project.ID,
			domain.WARNING,
			fmt.Sprintf("Scheduled app to stop after %v.", cs.maxContainerAge),
		))
		time.Sleep(cs.maxContainerAge)
		cs.end <- project
	}()

	// await for the app to announce it's live
	go func() {
		if !(containerInfo.isServerAlive) {
			timeout := time.NewTicker(cs.operationsTimeout)
			select {
			case isServerAlive := <-containerInfo.server:
				containerInfo.isServerAlive = isServerAlive
				cs.logService.CreateLogEvent(context.Background(), domain.NewAppLog(
					project.ID,
					domain.WARNING,
					"Serverless app started.",
				))
				success <- isServerAlive
			case <-timeout.C:
				cs.logService.CreateLogEvent(context.Background(), domain.NewAppLog(
					project.ID,
					domain.Error,
					"Timeout: Unable to start app in time.",
				))
			}
		} else {
			success <- true
		}
	}()
}

func (ss *containerService) StartApp(proj *project.Project, successChan chan bool, errChan chan error) {
	ss.start <- &containerStartArg{proj, successChan, errChan}
}
