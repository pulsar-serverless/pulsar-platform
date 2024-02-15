package container

import "context"

func (ss *containerService) eventLoop(ctx context.Context) {
	for {
		select {
		// start container
		case containerArg := <-ss.start:
			ss.startServerlessApp(ctx, containerArg.containerId, containerArg.success, containerArg.error)
			// stop container
		case containerId := <-ss.end:
			ss.stopServerlessApp(ctx, containerId)
		// change status of an app
		case containerId := <-ss.status:
			ss.changeAppStatus(containerId)
		}
	}
}
