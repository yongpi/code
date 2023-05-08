package dynamic

import (
	"fmt"
	"testing"
)

func TestRobot(t *testing.T) {
	fmt.Println(Robot(5, 2, 3, 3))
	fmt.Println(RobotDynamic(5, 2, 3, 3))
	fmt.Println(RobotDynamic2(5, 2, 3, 3))
}
