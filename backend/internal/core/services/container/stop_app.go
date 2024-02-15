package container

import (
	"context"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
)

func (cs *containerService) stopServerlessApp(ctx context.Context, containerId string) (bool, error) {
	containerInfo, ok := cs.liveContainers[containerId]

	if ok && time.Now().After(containerInfo.lastAccessed.Add(cs.maxContainerAge)) {
		// stop the serverless container
		err := cs.containerMan.StopContainer(ctx, containerId)
		if err != nil {
			log.Error().Str("containerId", containerId).Msg(fmt.Sprintf("Unable to stop container: %v", err.Error()))
			return false, err
		}

		log.Info().Str("containerId", containerId).Msg("container stopped.")
		delete(cs.liveContainers, containerId)
		return true, nil
	}

	log.Info().Str("containerId", containerId).Msg("Deadline extended; container not stopped.")
	return false, nil
}
