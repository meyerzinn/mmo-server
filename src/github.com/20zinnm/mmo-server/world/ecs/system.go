package ecs

type System struct {
	Manager *EntityManager
}

type ISystem interface {
	Tick(delta int)
}