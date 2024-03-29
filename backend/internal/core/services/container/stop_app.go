package container

import (
	"context"
	"fmt"
	domain "pulsar/internal/core/domain/log"
	"pulsar/internal/core/domain/project"
	"time"
)

func (cs *containerService) stopServerlessApp(ctx context.Context, project *project.Project) (bool, error) {
	containerInfo, ok := cs.liveContainers[project.ContainerId]

	if ok && time.Now().After(containerInfo.lastAccessed.Add(cs.maxContainerAge)) {
		// stop the serverless container
		err := cs.containerMan.StopContainer(ctx, project.ContainerId)
		if err != nil {
			cs.logService.CreateLogEvent(context.Background(), domain.NewAppLog(
				project.ID,
				domain.Error,
				fmt.Sprintf("Unable to stop container: %v", err.Error()),
			))
			return false, err
		}

		// stop collecting stats
		cs.containerMan.StopContainerStats(ctx, cs.monitor)

		cs.logService.CreateLogEvent(context.Background(), domain.NewAppLog(
			project.ID,
			domain.INFO,
			fmt.Sprintf("Stopped collecting container stats, id: %v", project.ContainerId)))

		// save container stats
		err = cs.resourceService.CreateResourceUtil(ctx, cs.resource, project)
		if err != nil {
			cs.logService.CreateLogEvent(context.Background(), domain.NewAppLog(
				project.ID,
				domain.Error,
				fmt.Sprintf("Unable to save stats: %v", err)))
		}

		cs.logService.CreateLogEvent(context.Background(), domain.NewAppLog(
			project.ID,
			domain.Error,
			"Container stopped.",
		))

		delete(cs.liveContainers, project.ContainerId)
		return true, nil
	}

	cs.logService.CreateLogEvent(context.Background(), domain.NewAppLog(
		project.ID,
		domain.Error,
		"Deadline extended; container not stopped.",
	))
	return false, nil
}
