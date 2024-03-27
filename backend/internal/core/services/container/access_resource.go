package container

import "pulsar/internal/core/domain/analytics"

func (cs *containerService) AccessResource() *analytics.RuntimeResourceObj {
	return cs.resource
}
