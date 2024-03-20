package container

import (
	"context"
	"io"
	"pulsar/internal/core/domain/analytics"
	"pulsar/internal/core/domain/project"
	resource "pulsar/internal/core/services/analytics"
	service "pulsar/internal/core/services/log"
	"pulsar/internal/ports"
	"time"
)

type IContainerService interface {
	DeployContainer(ctx context.Context, project *project.Project, buildContext io.Reader) (string, error)
	StartApp(project *project.Project, successChan chan bool, errChan chan error)
	ChangeAppStatus(ctx context.Context, containerId string) error
	AccessResource() *analytics.RuntimeResourceObj
}

type containerService struct {
	containerMan      ports.IContainerManager
	fileRepo          ports.IFileRepository
	projectRepo       ports.IProjectRepo
	logService        service.ILogService
	resourceService   resource.IResourceService
	liveContainers    map[string]*ContainerInfo
	maxContainerAge   time.Duration
	operationsTimeout time.Duration
	start             chan *containerStartArg
	end               chan *project.Project
	status            chan *project.Project
	resource          *analytics.RuntimeResourceObj
	monitor           *analytics.RuntimeResMonitor
}

type ContainerInfo struct {
	lastAccessed  time.Time
	server        chan bool
	isServerAlive bool
}

type containerStartArg struct {
	project *project.Project
	success chan bool
	error   chan error
}

func NewContainerService(containerMan ports.IContainerManager, logService service.ILogService, fileRepo ports.IFileRepository, projectRepo ports.IProjectRepo, resourceService resource.IResourceService) IContainerService {
	service := &containerService{
		containerMan:      containerMan,
		fileRepo:          fileRepo,
		projectRepo:       projectRepo,
		logService:        logService,
		resourceService:   resourceService,
		liveContainers:    make(map[string]*ContainerInfo),
		maxContainerAge:   time.Second * 30,
		operationsTimeout: time.Second * 10,
		start:             make(chan *containerStartArg),
		end:               make(chan *project.Project),
		status:            make(chan *project.Project),
	}

	go service.eventLoop(context.Background())
	return service
}
