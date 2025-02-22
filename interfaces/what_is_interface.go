package main

import (
	"fmt"
)

// this is for the tv interface
type RemoteControlar interface {
	TurnOn()
	TurnOff()
}
type SumsungTv struct{}
type SonyTv struct{}

func (s SumsungTv) TurnOn() {
	fmt.Println("Sumsung Tv is Turn On")
}
func (s SumsungTv) TurnOff() {
	fmt.Println("Sumsung Tv is Turn Off")
}

// func main() {
// 	var Tv RemoteControlar
// 	Tv = SumsungTv{}
// 	Tv.TurnOn()
// 	Tv.TurnOff()

// }
