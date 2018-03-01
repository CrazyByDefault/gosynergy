// Tejas, add your code here

package main

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
	var l,r,mid int 
	b1 := make([]byte, 24)
	f.Read(b1)

	fmt.Println(b1)

	l=int(b1[0] & 0x1)
	r=int(b1[0] & 0x2)
	mid=int(b1[0] & 0x4)

	fmt.Printf("left=%d , right=%d , middle=%d \n",l,r/2,mid/4)
}

func main() {

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