package container

import (
	"context"
	"pulsar/internal/core/domain/project"
	"pulsar/internal/ports"
	"time"
)

type IContainerService interface {
	DeployContainerWithStarterCode(ctx context.Context, project *project.Project)
	StartApp(containerId string, successChan chan bool, errChan chan error)
	ChangeAppStatus(ctx context.Context, containerId string) error
}

type containerService struct {
	containerMan      ports.IContainerManager
	fileRepo          ports.IFileRepository
	projectRepo       ports.IProjectRepo
	liveContainers    map[string]*ContainerInfo
	maxContainerAge   time.Duration
	operationsTimeout time.Duration
	start             chan *containerStartArg
	end               chan string
	status            chan string
}

type ContainerInfo struct {
	lastAccessed  time.Time
	server        chan bool
	isServerAlive bool
}

type containerStartArg struct {
	containerId string
	success     chan bool
	error       chan error
}

func NewContainerService(containerMan ports.IContainerManager, fileRepo ports.IFileRepository, projectRepo ports.IProjectRepo) IContainerService {
	service := &containerService{
		containerMan:      containerMan,
		fileRepo:          fileRepo,
		projectRepo:       projectRepo,
		liveContainers:    make(map[string]*ContainerInfo),
		maxContainerAge:   time.Second * 30,
		operationsTimeout: time.Second * 10,
		start:             make(chan *containerStartArg),
		end:               make(chan string),
		status:            make(chan string),
	}

	go service.eventLoop(context.Background())
	return service
}

func (cs *containerService) DeployContainerWithStarterCode(ctx context.Context, newProject *project.Project) {

	sourceDir, err := cs.fileRepo.InstallDefaultProject(newProject)
	if err != nil {
		return
	}

	buildContext, err := cs.fileRepo.CreateBuildContext(sourceDir)
	if err != nil {
		return
	}

	err = cs.containerMan.BuildImage(ctx, buildContext, newProject)
	if err != nil {
		return
	}

	containerId, err := cs.containerMan.CreateContainer(ctx, newProject.Name)
	if err != nil {
		return
	}

	cs.projectRepo.UpdateProject(ctx,
		newProject.ID,
		&project.Project{
			ContainerId: containerId,
		},
	)
}
