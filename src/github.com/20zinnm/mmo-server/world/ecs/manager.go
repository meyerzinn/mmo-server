package ecs

import "sync"

type EntityManager struct {
	entities map[Entity]map[uint8]*Component
	id       *int64
	idLock   *sync.Mutex
}

func (em *EntityManager) Add(entity *Entity) {
	em.entities[entity] = entity
}

func (em *EntityManager) GetComponent(entity Entity, componentId byte) Component {
	return em.entities[entity][componentId]
}

func (em *EntityManager) CreateEntity() *Entity {
	em.idLock.Lock()
	defer em.idLock.Unlock()
	var entity Entity = Entity(em.id)
	em.entities[entity] = map[uint8]*Component{}
	em.id++
	return entity
}

func (em *EntityManager) GetEntities() []Entity {
	return em.entities
}

func (em *EntityManager) HasComponent(entity Entity, id byte) bool {
	_, ok := em.entities[entity][id]
	return ok
}