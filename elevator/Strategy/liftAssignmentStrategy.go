package Strategy

import "lld/elevator/entity"

type LiftAssignmentStrategy interface {
	AssignLift(floorId int, direction entity.Direction, elevators []entity.Elevator) *entity.Elevator
}
