package mouselogger

import (
	//"fmt"
	// "io"
	//"io/ioutil"
	"os"
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

func readMouse(f *os.File, ch chan Activity) {
	//var l, r, mid, xr, yr int
	b1 := make([]byte, 24)
	f.Read(b1)

	//fmt.Println(b1)
	a := Activity{}
	a.Ri = int(b1[0] & 0x1)
	a.Le = int(b1[0]&0x2) / 2
	a.Mid = int(b1[0]&0x4) / 4
	a.Rx = int(b1[1])
	a.Ry = int(b1[2])

	ch <- a

	// fmt.Printf("left=%d , right=%d , middle=%d \n", r/2, l, mid/4)
	// fmt.Printf("xr=%d , yr=%d \n", xr, yr)

}

func GetMouseRel(ch chan Activity) {

	device := "/dev/input/mouse0"

	// You'll often want more control over how and what
	// parts of a file are read. For these tasks, start
	// by `Open`ing a file to obtain an `os.File` value.
	f, err := os.Open(device)
	check(err)

	// You'll often want more control over how and what
	// parts of a file are read. For these tasks, start
	// by `Open`ing a file to obtain an `os.File` value.

	//	for true {
	readMouse(f, ch)
	//	}

}
