package mouselogger

import (
	//"fmt"
	"fmt"
	"os/exec"
	"time"
	//"time"
	"strconv"
	"strings"
)

type Cords struct {
	X, Y int
}

func BoundaryCheck(chAbs chan Cords, lim int) {

	b := chAbs
	if b.X >= lim-10 {
		fmt.Print("boundary")
	}

}

func GetMouseAbs(ch chan Cords) {
	// Create an *exec.Cmd
	//str:="xdpyinfo  | grep -oP 'dimensions:\s+\K\S+'"
	//cmd1,_ := exec.Command("bash" , "-c", str).Output()
	//fmt.Println(cmd1)

	//for true {
	cmd, _ := exec.Command("xdotool", "getmouselocation").Output()

	s := strings.Split(strings.Replace(fmt.Sprintf("%s", cmd), ":", " ", -1), " ")
	c := Cords{}
	x, _ := strconv.Atoi(s[1])
	y, _ := strconv.Atoi(s[3])
	fmt.Println(x, y)
	//wid,_:= strconv.Atoi(s[7])
	c.X = x
	c.Y = y

	//strings.Replace(s, ":", " ", -1)
	//s:=strings.Split(s, " ")
	// // Stdout buffer
	// cmdOutput := &bytes.Buffer{}
	// // Attach buffer to command

	// cmd.Stdout = cmdOutput

	// Execute command
	//ch <- c
	fmt.Println(x, y)
	//fmt.Println(wid)
	//fmt.Println(s[7])
	time.Sleep(1000 * time.Nanosecond)

	//}

}
