package systems

import (
	"../ecs"
)

type NetworkingSystem struct {
	ecs.System
}

func (sys *NetworkingSystem) Tick(_ int) {
	for entity := range sys.Manager.GetEntities() {
		if (sys.Manager.HasComponent(entity, byte(1))) {

		}
	}
}