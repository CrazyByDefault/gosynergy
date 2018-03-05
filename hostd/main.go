package main

import (
	"./mouselogger"
	"fmt"
)


// func mouse(){

// 		mouselogger.GetMouseAbs()	
// }
type Coords struct {
   
    X,Y  int
}

func main() {

	for true {

		ch_abs := make(chan mouselogger.Cords)
		ch_act:= make(chan mouselogger.Activity)
		//go mouse()
		// var c mouselogger.Cords
		//mouselogger.GetMouseAbs(ch)

		go mouselogger.GetMouseAbs(ch_abs)
		go mouselogger.GetMouseRel(ch_act)
		c_a,c_r := <-ch_abs,<-ch_act

		fmt.Println(c_a.X, c_a.Y)
		fmt.Println(c_r.Rx, c_r.Ry)
		fmt.Println("l=",c_r.Le,"r=",c_r.Ri,"mid=",c_r.Mid)
		
	}	



		
}

// func sendToActiveDevice(cords Cords) {
// 	cords.
// }