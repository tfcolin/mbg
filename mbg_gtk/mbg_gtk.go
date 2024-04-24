package main

import (
	"fmt"
	"os"

	"gitee.com/tfcolin/mbg"
	"gitee.com/tfcolin/mbg/gtk3ui"
)

func main() {

	const max_turn int = 50

	mbg.Init()

	if len(os.Args) < 2 {
		fmt.Printf("Usage: mbg_gtk map_file\n")
		return
	}

	var fin *os.File
	fin, _ = os.Open (os.Args[1])
	var d *mbg.Driver

	var crole int

	fnl := len (os.Args[1])
	if os.Args[1][fnl - 3:] == "sav" {
		d = new (mbg.Driver)
		crole = d.Load (fin)
	} else {
		d = mbg.LoadMap (fin)
		crole = 0
	}

	_, nrole, _, _, _, _, _, _, _ := d.GetInfo()
	fin.Close()

	/*
	// for test
	d.FillCards (10)
	*/

	var uv gtk3ui.GtkUserView
	var oi_ []gtk3ui.GtkOperationInterface
	var oi []mbg.OperationInterface
	oi_ = make([]gtk3ui.GtkOperationInterface, nrole)
	oi = make([]mbg.OperationInterface, nrole)

	for i := 0; i < nrole; i++ {
		oi[i] = &(oi_[i])
	}

	gtk3ui.LoadUI()
	d.ConnectUI(&uv, oi)
	d.Run(max_turn, crole)

}
