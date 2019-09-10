package elevator

import (
	"testing"
)

func BasicTest(t *testing.T) {
	elevator := Elevator{}
	elevator.AddToQueue(12, 15, Up)
	// expectedFloor := 2
	// actualFloor := NextFloor(
	// 	[]int{1, 9, 3, 4, 5, 6},
	// 	[]int{9, 22, 3, 4, 6, 5},
	// )

	// if expectedStrong != actualStrong || expectedWeak != actualWeak {
	// 	t.Errorf("\nTest case failed!\n\tExpected Strong Matches: %s\n\tActual Strong Matches: %s\n\n\tExpected Weak Matches: %s\n\tActual Weak Matches: %s\n",
	// 		strconv.Itoa(expectedStrong),
	// 		strconv.Itoa(actualStrong),
	// 		strconv.Itoa(expectedWeak),
	// 		strconv.Itoa(actualWeak),
	// 	)
	// }
}
