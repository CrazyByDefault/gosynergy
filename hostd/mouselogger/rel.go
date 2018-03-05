package mouselogger

import (
	"fmt"
	// "io"
	//"io/ioutil"
	"os"
)

// Reading files requires checking most calls for errors.
// This helper will streamline our error checks below.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readMouse(f *os.File) {
	var l, r, mid, xr, yr int
	b1 := make([]byte, 24)
	f.Read(b1)

	//fmt.Println(b1)

	l = int(b1[0] & 0x1)
	r = int(b1[0] & 0x2)
	mid = int(b1[0] & 0x4)
	xr = int(b1[1])
	yr = int(b1[2])

	fmt.Printf("left=%d , right=%d , middle=%d \n", r/2, l, mid/4)
	fmt.Printf("xr=%d , yr=%d \n", xr, yr)

}

// GetMouseRel reads the mouse device and returns the relative motion values, as well as mouseclick events
func GetMouseRel() {

	device := "/dev/input/mouse0"

	// You'll often want more control over how and what
	// parts of a file are read. For these tasks, start
	// by `Open`ing a file to obtain an `os.File` value.
	f, err := os.Open(device)
	check(err)

	for true {
		readMouse(f)
	}

}
