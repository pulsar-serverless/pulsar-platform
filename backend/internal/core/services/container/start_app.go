package container

import (
	"context"
	"fmt"
	"pulsar/internal/core/domain/analytics"
	domain "pulsar/internal/core/domain/log"
	"pulsar/internal/core/domain/project"
	"pulsar/internal/core/services"
	"time"

	"github.com/docker/go-connections/nat"
)

func (cs *containerService) startServerlessApp(ctx context.Context, project *project.Project, success chan *nat.PortBinding, errorChan chan error) {
	containerInfo, ok := cs.liveContainers[project.ContainerId]

	if !ok {
		// start the serverless container
		cs.logService.CreateLogEvent(context.Background(), domain.NewAppLog(
			project.ID,
			domain.INFO,
			"Container not started; Starting app container.",
		))

		portBinding, err := cs.containerMan.StartContainer(ctx, project.ContainerId)

		if err != nil {
			cs.logService.CreateLogEvent(context.Background(), domain.NewAppLog(
				project.ID,
				domain.Error,
				fmt.Sprintf("Unable to start app container: %v", err),
			))
			errorChan <- services.NewAppError(services.ErrInternalServer, err)
		}

		cs.logService.CreateLogEvent(context.Background(), domain.NewAppLog(
			project.ID,
			domain.INFO,
			fmt.Sprintf("Starting collecting container stats, id: %v", project.ContainerId)))

		cs.resource = analytics.NewRuntimeResObj()
		cs.monitor = analytics.NewRuntimeResMonitor()

		go cs.containerMan.GetContainerStats(ctx, project.ContainerId, cs.resource, cs.monitor)
		go cs.saveContainerLogs(project)

		containerInfo = &ContainerInfo{
			lastAccessed:  time.Now(),
			server:        make(chan bool),
			isServerAlive: false,
			portBinding:   portBinding,
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
				success <- containerInfo.portBinding
			case <-timeout.C:
				cs.logService.CreateLogEvent(context.Background(), domain.NewAppLog(
					project.ID,
					domain.Error,
					"Timeout: Unable to start app in time.",
				))
			}
		} else {
			success <- containerInfo.portBinding
		}
	}()
}

func (ss *containerService) StartApp(proj *project.Project, successChan chan *nat.PortBinding, errChan chan error) {
	ss.start <- &containerStartArg{proj, successChan, errChan}
}
