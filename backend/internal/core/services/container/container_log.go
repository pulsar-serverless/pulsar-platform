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
	index := strings.Index(rawLog, " ")

	if index == -1 {
		zeroLog.Warn().
			Str("AppID", projectId).
			Msgf("Unrecognized log format from container: %v", rawLog)
		return nil
	}

	timePart := rawLog[:index]
	messagePart := rawLog[index:]

	parsedTime, err := time.Parse(time.RFC3339Nano, timePart)
	if err != nil {
		zeroLog.Error().
			Str("AppID", projectId).
			Err(err).
			Msgf("Unrecognized timestamp in a container log: %v", timePart)
		return nil
	}

	var newLog = &log.AppLog{CreatedAt: parsedTime}
	err = json.Unmarshal([]byte(messagePart), newLog)

	if err != nil {
		zeroLog.Error().
			Str("AppID", projectId).
			Err(err).
			Msgf("Unrecognized timestamp in a container log: %v", messagePart)
		return nil
	}

	return newLog
}
