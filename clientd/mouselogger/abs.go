package mouselogger

import (
	//"fmt"
	"os/exec"
	//"time"
	"strconv"
	"strings"
)

type Cords struct {
	X, Y int
}

func GetMouseAbs(ch chan Cords) {
	// Create an *exec.Cmd
	//str:="xdpyinfo  | grep -oP 'dimensions:\s+\K\S+'"
	//cmd1,_ := exec.Command("bash" , "-c", str).Output()
	//fmt.Println(cmd1)

	//for true {
	cmd, _ := exec.Command("xdotool", "getmouselocation").Output()

	s := strings.Split(strings.Replace(string(cmd), ":", " ", -1), " ")
	c := Cords{}
	x, _ := strconv.Atoi(s[1])
	y, _ := strconv.Atoi(s[3])
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
	ch <- c
	//fmt.Println(x,y)
	//fmt.Println(wid)
	//fmt.Println(s[7])
	//time.Sleep(1000 * time.Nanosecond)

	//}

}
