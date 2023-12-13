package main

import (
	"fmt"
	"kmrecords"
)

func main() {
	mh := kmrecords.NewHookData(kmrecords.WH_MOUSE_LL, kmrecords.HookMProcCallback)
	mhook := kmrecords.SetHooks(mh)
	fmt.Printf("mhook: %v\n", mhook)
	defer kmrecords.UnWindowHook(mhook)
	kh := kmrecords.NewHookData(kmrecords.WH_KEYBOARD_LL, kmrecords.HookKProcCallback)
	khook := kmrecords.SetHooks(kh)
	fmt.Printf("khook: %v\n", khook)
	defer kmrecords.UnWindowHook(khook)
	go func() {
		for {
			select {
			case event := <-kmrecords.MouseEventChan:
				// 处理鼠标事件
				fmt.Printf("Mouse keyboard event: %v\n", event)
			}
		}
	}()

	for {
		msg := kmrecords.MSG{}
		kmrecords.GetMessageA(&msg)
		fmt.Printf("msg: %#v\n", msg)
	}
}
