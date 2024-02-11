package container

import (
	"context"
	"pulsar/internal/core/domain/project"
	"pulsar/internal/ports"
)

type IContainerService interface {
	DeployContainerWithStarterCode(project *project.Project)
}

type containerService struct {
	containerMan ports.IContainerManager
	fileRepo     ports.IFileRepository
}

func NewContainerService(containerMan ports.IContainerManager, fileRepo ports.IFileRepository) IContainerService {
	return &containerService{containerMan, fileRepo}
}

func (cs *containerService) DeployContainerWithStarterCode(project *project.Project) {
	context := context.Background()

	sourceDir, err := cs.fileRepo.InstallDefaultProject(project)
	if err != nil {
		return
	}

	buildContext, err := cs.fileRepo.CreateBuildContext(sourceDir)
	if err != nil {
		return
	}

	err = cs.containerMan.BuildImage(context, buildContext, project)
	if err != nil {
		return
	}

	cs.containerMan.CreateContainer(context, project.Name)
}
