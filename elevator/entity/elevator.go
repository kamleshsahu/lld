package entity

import "container/list"

type ElevatorState string

const (
	GOINGUP      ElevatorState = "UP"
	GOINGDOWN    ElevatorState = "DOWN"
	IDLE         ElevatorState = "IDLE"
	MAINTAINANCE ElevatorState = "MAINTAINANCE"
)

type Direction string

const (
	UP   Direction = "UP"
	DOWN Direction = "DOWN"
)

type Elevator struct {
	Floors        list.List
	DownRequests  DownHeap
	UpRequests    UpHeap
	ElevatorId    int
	ElevatorState ElevatorState
	CurrentFloor  int
}

type DownHeap []int
type UpHeap []int

type InternalController struct {
}

type ExternalController struct {
}

type Floor struct {
	Id          int
	Controllers map[int]ExternalController
}
