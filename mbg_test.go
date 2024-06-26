package mbg

import (
	"os"
)

func TestMBG() {

	mbg.Init()
	fin, _ := os.Open("test_game.map")
	d, _ := mbg.LoadMap(fin)
	fin.Close()

      _, nrole, _, _, _, _, _, _, _ := d.GetInfo()

	var uv mbg.TestUserView
	var oi_ []mbg.TestOperationInterface
	var oi []mbg.OperationInterface
	oi_ = make([]mbg.TestOperationInterface, nrole)
	oi = make([]mbg.OperationInterface, nrole)

	for i := 0; i < nrole; i++ {
		oi[i] = &(oi_[i])
	}

	d.ConnectUI(&uv, oi)

	d.Run(10, 0)

}
