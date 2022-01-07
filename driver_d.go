package mbg

import (
	"fmt"
	"io"

	"github.com/tfcolin/dsg"
)

const (
	TECH_COUNT             = 12   // 战斗技能个数
	CARD_COUNT             = 14   // 卡片个数
	TRAIN_ROUND            = 15   // 完成一次训练所需回合数
	STUDY_SCALE            = 4    // 研究所一回合所完成的点数与研究人员策略值的比例
	P_HP_BASE      float32 = 2000 // 人员最大 HP 的基数
	P_HP_SCALE             = 4    // 人员最大 HP 减基数后与人员所有属性值和的比值
	FENGSHUI_ROUND         = 10   // 每多少回合更新一次城市风水
	MIN_EARN       float32 = 20   // 城市最小收入指数 (当无财务官或财务官的经济属性小于该值时将采用该值)
	MAX_STEP               = 6    // 角色正常行走时的最大步长
	FAST_SCALE             = 2    // 角色急性时最大步长与 MAX_STEP 的比值
	SLOW_SCALE             = 2    // 角色慢行时最大步长与 MAX_STEP 的比值
	ROBBER_STEAL   float32 = 2000 // 未剿灭山贼时被偷的钱数
	TAX_SCALE      float32 = 5    // 城市税收 (角色到达被他人占领城市所付的钱) 与城市每回合正常收入的比例
	OWN_TAX_SCALE  float32 = 2    // 自我城市税收 (角色到达自己所占领城市时收入的钱) 与城市每回合正常收入的比例
	HPPLUS_SCALE   float32 = 2    // 城市每回合恢复 HP 与市长政治属性的比例
	ALLOCATE_TURN          = 10   // 每多少回合进行一次人员调配
)

var zei Officer = Officer{
	name:  "robber",
	role:  -1,
	job:   -1,
	loc:   -1,
	hpmax: 1800,
	hp:    1200,
	prop:  [5]float32{60, 40, 10, 20, 30},
}

var force_quit bool

type Driver struct {
	cities []City
	insts  []Institute
	trains []TrainingRoom
	board  []BoardPoint
	people []Officer
	roles  []Role

	lpset [3](*dsg.Set) // people level set 0:2 低中高

	ngrid   int
	nrole   int
	npeople int
	ncity   int
	ninst   int
	ntrain  int

	uv UserView
	oi []OperationInterface

	turn      int
	win_money float32

	nx, ny int
}

func (d *Driver) GetInfo() (ngrid, nrole, npeople, ncity, ninst, ntrain, turn, nx, ny int) {
	ngrid = d.ngrid
	nrole = d.nrole
	npeople = d.npeople
	ncity = d.ncity
	ninst = d.ninst
	ntrain = d.ntrain
	turn = d.turn
	nx = d.nx
	ny = d.ny
	return
}

func (d *Driver) GetCount() (ngrid, nrole, npeople, ncity, ninst, ntrain int) {
	ngrid = d.ngrid
	nrole = d.nrole
	npeople = d.npeople
	ncity = d.ncity
	ninst = d.ninst
	ntrain = d.ntrain
	return
}

func (d *Driver) GetPeople(i int) *Officer {
	return &(d.people[i])
}

func (d *Driver) GetBoard(i int) *BoardPoint {
	return &(d.board[i])
}

func (d *Driver) GetRole(i int) *Role {
	return &(d.roles[i])
}

func (d *Driver) GetCity(i int) *City {
	return &(d.cities[i])
}

func (d *Driver) GetInst(i int) *Institute {
	return &(d.insts[i])
}

func (d *Driver) GetTrain(i int) *TrainingRoom {
	return &(d.trains[i])
}

func (d *Driver) GetFreePeopleList(level int) []int {
	return d.lpset[level].GetAllLabel()
}

func (d *Driver) Save(fout io.Writer) {
	fmt.Fprintf(fout, "%d %d %d %d %d %d %d %f \n",
		d.ngrid, d.nrole, d.npeople, d.ncity, d.ninst, d.ntrain, d.turn, d.win_money)

	for i := 0; i < d.ncity; i++ {
		d.cities[i].Save(fout)
	}

	for i := 0; i < d.ninst; i++ {
		d.insts[i].Save(fout)
	}

	for i := 0; i < d.ntrain; i++ {
		d.trains[i].Save(fout)
	}

	for i := 0; i < d.ngrid; i++ {
		d.board[i].Save(fout)
	}

	for i := 0; i < d.npeople; i++ {
		d.people[i].Save(fout)
	}

	for i := 0; i < d.nrole; i++ {
		d.roles[i].Save(fout)
	}

	for i := 0; i < 3; i++ {
		d.lpset[i].Save(fout)
	}
}

func (d *Driver) Load(fin io.Reader) {
	fmt.Fscan(fin,
		&d.ngrid, &d.nrole, &d.npeople, &d.ncity, &d.ninst, &d.ntrain, &d.turn, &d.win_money)

	d.cities = make([]City, d.ncity)
	for i := 0; i < d.ncity; i++ {
		d.cities[i].Load(fin)
	}

	d.insts = make([]Institute, d.ninst)
	for i := 0; i < d.ninst; i++ {
		d.insts[i].Load(fin)
	}

	d.trains = make([]TrainingRoom, d.ntrain)
	for i := 0; i < d.ntrain; i++ {
		d.trains[i].Load(fin)
	}

	d.board = make([]BoardPoint, d.ngrid)
	for i := 0; i < d.ngrid; i++ {
		d.board[i].Load(fin)
	}

	d.people = make([]Officer, d.npeople)
	for i := 0; i < d.npeople; i++ {
		d.people[i].Load(fin)
	}

	d.roles = make([]Role, d.nrole)
	for i := 0; i < d.nrole; i++ {
		d.roles[i].Load(fin)
	}

	for i := 0; i < 3; i++ {
		d.lpset[i].Load(fin)
	}
}
