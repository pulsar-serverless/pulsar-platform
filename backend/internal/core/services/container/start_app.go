package container

import (
	"context"
	"errors"
	"fmt"
	"pulsar/internal/core/services"
	"time"

	"github.com/rs/zerolog/log"
)

func (ss *containerService) startServerlessApp(ctx context.Context, containerId string, success chan bool, errorChan chan error) {
	containerInfo, ok := ss.liveContainers[containerId]

	if !ok {
		// start the serverless container
		log.Info().Str("containerId", containerId).Msg("Starting app container: container not started.")
		err := ss.containerMan.StartContainer(ctx, containerId)
		if err != nil {
			log.Error().Str("containerId", containerId).Msg("Unable to start app container.")
			errorChan <- services.NewAppError(services.ErrInternalServer, err)
		}

		containerInfo = &ContainerInfo{
			lastAccessed:  time.Now(),
			server:        make(chan bool),
			isServerAlive: false,
		}

		ss.liveContainers[containerId] = containerInfo
	} else {
		log.Info().Str("containerId", containerId).Msg("Starting app container; container already started.")
		containerInfo.lastAccessed = time.Now()
	}

	// schedule app closing
	go func(containerId string) {
		log.Info().Str("containerId", containerId).Msg(fmt.Sprintf("Scheduled app to stop after %v.", ss.maxContainerAge))
		time.Sleep(ss.maxContainerAge)
		ss.end <- containerId
	}(containerId)

	// await for the app to announce it's live
	go func() {
		if !(containerInfo.isServerAlive) {
			timeout := time.NewTicker(ss.operationsTimeout)
			select {
			case isServerAlive := <-containerInfo.server:
				containerInfo.isServerAlive = isServerAlive
				log.Info().Str("containerId", containerId).Msg("Serverless app started.")
				success <- isServerAlive
			case <-timeout.C:
				log.Error().Str("containerId", containerId).Msg("unable to start app in time")
				errorChan <- services.NewAppError(services.ErrInternalServer, errors.New("Unable to start app in time"))
			}
		} else {
			success <- true
		}
	}()
}

func (ss *containerService) StartApp(containerId string, successChan chan bool, errChan chan error) {
	ss.start <- &containerStartArg{containerId, successChan, errChan}
}
