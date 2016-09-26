package world

import (
	"./ecs"
)

// World represents a 2D coordinate plane with entities.
//type World struct {
//	tick     int32
//	entities map[int32]*Entity
//}

//func (w *World) addEntity(entity Entity) bool {
//	_, ok := w.entities[entity.id()]
//	if ok {
//		return false
//	}
//	w.entities[entity.id()] = &entity
//	return true
//}

type World struct {
	tick          int32
	entityManager *ecs.EntityManager
}

func (w *World) Add() {

}

// Snapshot represents a snapshot of a region at a given time.
type Snapshot struct {
}

// AABB is an Axis Aligned Bounding Box in a 2D plane.
type AABB struct {
}
