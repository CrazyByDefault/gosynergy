package main

import (
	"fmt"
	"sync"
    "os/exec"
	"./keylogger"
	"./mouselogger"
)

	func getRes() {
	toExec := "xdpyinfo  | grep -oP 'dimensions:\\s+\\K\\S+'"
	cmd, _ := exec.Command("bash", "-c", toExec).Output()
	fmt.Printf("%s", cmd)
}

func keebListener(chKey chan uint16) {
	devs, err := keylogger.NewDevices()
	if err != nil {
		fmt.Println(err)
		return
	}	


	// for _, val := range devs {
	// 	fmt.Println("Id->", val.Id, "Device->", val.Name)
	// }

	//keyboard device file, on your system it will be diffrent!
	rd := keylogger.NewKeyLogger(devs[4])

	for true {
		go keylogger.KeyActivity(chKey, rd)
		fmt.Print(int(<-chKey))
	}
}

func mouseListener(ch_abs chan mouselogger.Cords, ch_act chan mouselogger.Activity) {

	for true {
		go mouselogger.GetMouseAbs(ch_abs)
		go mouselogger.GetMouseRel(ch_act)
		c_a, c_r := <-ch_abs, <-ch_act

		fmt.Println(c_a.X, c_a.Y)
		fmt.Println(c_r.Rx, c_r.Ry)
		fmt.Println("l=", c_r.Le, "r=", c_r.Ri, "mid=", c_r.Mid)
	}

}

// func mouse(){

// 		mouselogger.GetMouseAbs()
// }
// type Coords struct {

//     X,Y  int
// }

func main() {
	var wg sync.WaitGroup

	ch_abs := make(chan mouselogger.Cords)
	ch_act := make(chan mouselogger.Activity)
	ch_key := make(chan uint16)

	//go mouse()
	// var c mouselogger.Cords
	//mouselogger.GetMouseAbs(ch)
	wg.Add(2)
	go mouseListener(ch_abs, ch_act)
	go keebListener(ch_key)
	// mouselogger.GetMouseAbs(ch_abs)
	// mouselogger.GetMouseRel(ch_act)

	// c_a, c_r := <-ch_abs, <-ch_act
	// c_a, c_r := <-ch_abs, <-ch_act

	// // should change this

	wg.Wait()
	close(ch_act)
	close(ch_abs)
	close(ch_key)
}

// func sendToActiveDevice(cords Cords) {
// 	cords.
// }
