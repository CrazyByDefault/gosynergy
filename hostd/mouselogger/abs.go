package mouselogger

import (
	"fmt"
	"os/exec"
	"time"
)

func getMouseAbs() {
	// Create an *exec.Cmd
	//str:="xdpyinfo  | grep -oP 'dimensions:\s+\K\S+'"
	//cmd1,_ := exec.Command("bash" , "-c", str).Output()
	//fmt.Println(cmd1)

	for true {
		cmd, _ := exec.Command("xdotool", "getmouselocation").Output()

		// // Stdout buffer
		// cmdOutput := &bytes.Buffer{}
		// // Attach buffer to command

		// cmd.Stdout = cmdOutput

		// Execute command
		fmt.Println(string(cmd))
		time.Sleep(1000 * time.Nanosecond)

	}

}
