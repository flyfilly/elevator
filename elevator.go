package elevator

import (
	"sort"
)

// Direction Constants
const (
	Up   = 1
	Idle = 0
	Down = -1
)

// Elevator houses the elevator object and methods
type Elevator struct {
	CurrentFloor  int
	Direction     int
	upwardQueue   []int
	downwardQueue []int
}

// PressButton adds to the appropriate Queue
func (elevator *Elevator) PressButton(start, end, direction int) {
	switch direction {
	case Up:
		if elevator.Direction == Idle {
			elevator.Direction = Up
		}
		elevator.upwardQueue = append(elevator.upwardQueue, start, end)
		sort.Slice(elevator.upwardQueue, func(i, j int) bool {
			return elevator.upwardQueue[i] < elevator.upwardQueue[j]
		})
	case Down:
		if elevator.Direction == Idle {
			elevator.Direction = Down
		}
		elevator.downwardQueue = append(elevator.downwardQueue, start, end)
		sort.Slice(elevator.downwardQueue, func(i, j int) bool {
			return elevator.downwardQueue[i] > elevator.downwardQueue[j]
		})
	}

	elevator.Move()
}

// NextFloor returns the next floor in the respective queue
func (elevator *Elevator) NextFloor() int {
	return 0
}

//Move simulates the elevator movement
func (elevator *Elevator) Move() {

}
