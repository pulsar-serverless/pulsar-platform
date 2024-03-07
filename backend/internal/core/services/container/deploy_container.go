package container

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"pulsar/internal/core/domain/log"
	"pulsar/internal/core/domain/project"
)

func (cs *containerService) DeployContainer(ctx context.Context, newProject *project.Project, buildContext io.Reader) (string, error) {
	buildOutput, err := cs.containerMan.BuildImage(ctx, buildContext, newProject)

	if err != nil {
		cs.logService.CreateLogEvent(context.Background(), log.NewAppLog(
			newProject.ID,
			log.Error,
			fmt.Sprintf("Unable to build app Image: %v", err)))
		return "", err
	}

	defer buildOutput.Close()
	cs.saveContainerBuildOutput(newProject, buildOutput)

	containerId, err := cs.containerMan.CreateContainer(ctx, newProject.Name)
	if err != nil {
		cs.logService.CreateLogEvent(context.Background(), log.NewAppLog(
			newProject.ID,
			log.Error,
			fmt.Sprintf("Unable to build app Container: %v", err)))
	}

	return containerId, err
}

func (ss *containerService) saveContainerBuildOutput(proj *project.Project, reader io.Reader) error {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		newLog, err := parseBuildOutput(scanner.Bytes(), proj.ID)
		go ss.logService.CreateLogEvent(context.Background(), newLog)
		if err != nil {
			return err
		}
	}

	return nil
}

func parseBuildOutput(buildOutput []byte, projectId string) (*log.AppLog, error) {
	var output map[string]interface{}
	err := json.Unmarshal(buildOutput, &output)

	if err != nil {
		return log.NewAppLog(projectId, log.INFO, string(buildOutput)), nil
	}

	message, ok := output["stream"]
	if ok {
		return log.NewAppLog(projectId, log.INFO, message.(string)), nil
	}

	errMessage, ok := output["error"]
	if ok {
		return log.NewAppLog(projectId, log.Error, errMessage.(string)), errors.New(errMessage.(string))
	}

	return log.NewAppLog(projectId, log.INFO, ""), nil
}
