package main

import (
	"fmt"
	"os"

	"github.com/tfcolin/mbg"
	"github.com/tfcolin/mbg/gtkui"
)

func main() {

	mbg.Init()

	if len(os.Args) < 2 {
		fmt.Printf("Usage: mbg_gtk map_file\n")
		return
	}

	var fin *os.File
	fin, _ = os.Open(os.Args[1])
	var d *mbg.Driver
	d = mbg.LoadMap(fin)
	_, nrole, _, _, _, _, _, _, _ := d.GetInfo()
	fin.Close()

	/*
	   // for test
	   d.FillCards (10)
	*/

	var uv gtkui.GtkUserView
	var oi_ []gtkui.GtkOperationInterface
	var oi []mbg.OperationInterface
	oi_ = make([]gtkui.GtkOperationInterface, nrole)
	oi = make([]mbg.OperationInterface, nrole)

	for i := 0; i < nrole; i++ {
		oi[i] = &(oi_[i])
	}

	gtkui.LoadUI()

	d.ConnectUI(&uv, oi)

	d.Run(100)

	gtkui.FreeUI()
}
