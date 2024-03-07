package container

import (
	"bufio"
	"context"
	"encoding/json"
	"time"

	"strings"

	"pulsar/internal/core/domain/log"
	"pulsar/internal/core/domain/project"

	zeroLog "github.com/rs/zerolog/log"
)

func (ss *containerService) saveContainerLogs(proj *project.Project) {
	reader, err := ss.containerMan.GetContainerLogs(context.Background(), proj.ContainerId)
	if err != nil {
		zeroLog.Error().
			Str("AppID", proj.ID).
			Err(err).
			Msg("Unable to read project container log")
		return
	}
	defer reader.Close()

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		newLog := parseLog(scanner.Text(), proj.ID)
		go ss.logService.CreateLogEvent(context.Background(), newLog)
	}
}

func parseLog(rawLog string, projectId string) *log.AppLog {
	defaultLog := log.NewAppLog(
		projectId,
		log.WARNING,
		rawLog,
	)
	parts := strings.Split(rawLog, " ")

	if len(parts) != 2 {
		zeroLog.Warn().
			Str("AppID", projectId).
			Msgf("Unrecognized log format from container: %v", rawLog)
		return defaultLog
	}

	parsedTime, err := time.Parse(parts[0], parts[0])
	if err != nil {
		zeroLog.Error().
			Str("AppID", projectId).
			Err(err).
			Msgf("Unrecognized timestamp in a container log: %v", parts[0])
		return defaultLog
	}

	var newLog = &log.AppLog{CreatedAt: parsedTime}
	err = json.Unmarshal([]byte(parts[1]), newLog)

	if err != nil {
		zeroLog.Error().
			Str("AppID", projectId).
			Err(err).
			Msgf("Unrecognized timestamp in a container log: %v", parts[0])
		return defaultLog
	}

	return newLog
}
