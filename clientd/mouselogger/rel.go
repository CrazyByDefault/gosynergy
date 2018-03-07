package mousemover

import (
	//"fmt"
	// "io"
	//"io/ioutil"

	"os/exec"

	"../../hostd/mouselogger"
)

// type Activity struct {
// 	Rx, Ry, Ri, Le, Mid int
// }

// Reading files requires checking most calls for errors.
// This helper will streamline our error checks below.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readMouse(ch chan mouselogger.Activity) {
	//var l, r, mid, xr, yr int
	var x, y string
	// ch.Ri = int(b1[0] & 0x1)
	// ch.Le = int(b1[0]&0x2) / 2
	// ch.Mid = int(b1[0]&0x4) / 4
	// a.Rx = int(b1[1])
	// a.Ry = int(b1[2])

	if ch.Rx < 128 {
		x = "1"
	} else {
		x = "-1"
	}
	if ch.Ry < 128 {
		y = "-1"
	} else {
		y = "1"
	}
	if ch.Ri == 1 {
		exec.Command("xdotool", "mousemove_relative", "--", x, y, "click", "3").Run()

	} else if ch.Le == 1 {
		exec.Command("xdotool", "mousemove_relative", "--", x, y, "click", "1").Run()

	} else if ch.Mid == 1 {
		exec.Command("xdotool", "mousemove_relative", "--", x, y, "click", "2").Run()

	} else {
		exec.Command("xdotool", "mousemove_relative", "--", x, y).Run()
	}
	// fmt.Printf("left=%d , right=%d , middle=%d \n", r/2, l, mid/4)
	// fmt.Printf("xr=%d , yr=%d \n", xr, yr)

}
