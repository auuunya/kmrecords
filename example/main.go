package main

import (
	"fmt"
	"kmrecords"
)

func main() {
	mhook := kmrecords.SetHooks(kmrecords.WH_MOUSE_LL, kmrecords.HookMProcCallback)
	fmt.Printf("mhook: %v\n", mhook)
	khook := kmrecords.SetHooks(kmrecords.WH_KEYBOARD_LL, kmrecords.HookKProcCallback)
	fmt.Printf("khook: %v\n", khook)
	var msg *kmrecords.MSG
	for kmrecords.GetMessageW(msg) != 0 {
		fmt.Printf("msg: %v\n", msg)
	}
}
