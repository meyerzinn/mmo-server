package ecs

type Entity uint64

type Identifiable interface {
	ID() uint64
}