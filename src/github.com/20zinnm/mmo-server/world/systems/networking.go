package systems

import (
	"../ecs"
	"../components"
)

type NetworkingSystem struct {
	ecs.System
}

func (sys *NetworkingSystem) Tick(_ int) {
	for entity := range sys.Manager.GetEntities() {
		if (sys.Manager.HasComponent(entity, components.CLIENT_COMPONENT)) {

		}
	}
}