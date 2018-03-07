package mousemover

import (
	"fmt"
	// "io"
	//"io/ioutil"

	"os/exec"
)

type Activity struct {
	Rx, Ry, Ri, Le, Mid int
}

// Reading files requires checking most calls for errors.
// This helper will streamline our error checks below.
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

// ReadMouse reads and executes the mouse events in the 'current' Activity struct
func ReadMouse(current Activity, leftState int, rightState int, midState int) (int, int, int) {
	// fmt.Print("Called ReadMouse")
	//var l, r, mid, xr, yr int

	var x, y string

	// fmt.Print("Read from chan")
	// ch.Ri = int(b1[0] & 0x1)
	// ch.Le = int(b1[0]&0x2) / 2
	// ch.Mid = int(b1[0]&0x4) / 4
	// a.Rx = int(b1[1])
	// a.Ry = int(b1[2])

	if current.Rx > 0 && current.Rx < 50 {
		x = "1"
	} else if current.Rx > 0 && current.Rx > 200 {
		x = "-1"
		current.Rx = 256 - current.Rx
	} else {
		x = "0"
	}
	if current.Ry > 0 && current.Ry < 50 {
		y = "-1"
	} else if current.Ry > 0 && current.Ry > 200 {
		y = "1"
		current.Ry = 256 - current.Ry
	} else {
		y = "0"
	}

	if current.Ri != leftState {
		var event string
		leftState = current.Ri
		if leftState == 0 {
			event = "mouseup"
		} else {
			event = "mousedown"
		}

		exec.Command("xdotool", event, "1").Run()
		fmt.Println(event + "1")

	} else if current.Le != rightState {
		var event string
		rightState = current.Le
		if current.Le == 0 {
			event = "mouseup"
		} else {
			event = "mousedown"
		}

		exec.Command("xdotool", event, "3").Run()
		fmt.Println(event + "2")

	} else if current.Mid != midState {
		var event string
		midState = current.Mid
		if current.Mid == 0 {
			event = "mouseup"
		} else {
			event = "mousedown"
		}

		exec.Command("xdotool", event, "2").Run()
		fmt.Println(event + "3")
	}
	m := max(current.Rx, current.Ry)

	for i := 0; i < m*10; i++ {

		exec.Command("xdotool", "mousemove_relative", "--", x, y).Run()
	}
	return leftState, rightState, midState
}
