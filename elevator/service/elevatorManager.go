package service

import (
	"fmt"
	"lld/elevator/Strategy"
	"lld/elevator/entity"
)

type ElevatorManager struct {
	Elevators    []entity.Elevator
	Floors       []entity.Floor
	liftAssigner Strategy.LiftAssignmentStrategy
}

func NewElevatorManager() *ElevatorManager {
	return &ElevatorManager{}
}

func (e *ElevatorManager) AddElevator(elevator entity.Elevator) {
	e.Elevators = append(e.Elevators, elevator)
}

func (e *ElevatorManager) AddFloor(floor entity.Floor) {
	e.Floors = append(e.Floors, floor)
}

func (e *ElevatorManager) AssignElevator(floorid int, direction entity.Direction) (*entity.Elevator, error) {
	elevator := e.liftAssigner.AssignLift(floorid, direction, e.Elevators)
	if elevator == nil {
		return nil, fmt.Errorf("Cannot assign elevator to floor %d", floorid)
	}
	return elevator, nil
}
