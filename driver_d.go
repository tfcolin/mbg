package mbg

import (
	"fmt"
	"io"

	"gitee.com/tfcolin/dsg"
)

const (
	TECH_COUNT             = 12   // 战斗技能个数
	CARD_COUNT             = 14   // 卡片个数
	TRAIN_ROUND            = 6    // 完成一次训练所需回合数
      TRAIN_PLUS             = 3    // 完成一次修炼所获得的最大能力提升值
	STUDY_SCALE            = 1    // 研究所一回合所完成的点数与研究人员策略值的比例
	P_HP_BASE      float32 = 1000 // 人员最大 HP 的基数
	P_HP_SCALE     float32 = 20   // 人员最大 HP 减基数后与人员平均属性值的比值
	P_PRICE_SCALE  float32 = 30   // 雇佣价格与参考属性值 (单一: 0-100) 的比例.
	P_POISON_PROP_SCALE float32 = 0.75 // 处于中毒状态的人物能力占健康能力的比例
	DEF_SMALL	   float32 = 500  // 小型城市最大防御
	DEF_MEDIUM     float32 = 800  // 中型城市最大防御
	DEF_BIG  	   float32 = 1200 // 大型城市最大防御
	DEF_PLUS_SCALE float32 = 0.05 // 城市每回合恢复防御占总防御的比例
	FENGSHUI_ROUND         = 5    // 每多少回合更新一次城市风水
	MIN_EARN       float32 = 20   // 城市最小收入指数 (当无财务官或财务官的经济属性小于该值时将采用该值)
	MAX_STEP               = 6    // 角色正常行走时的最大步长
	FAST_SCALE     float32 = 2    // 角色急性时最大步长与 MAX_STEP 的比值
	SLOW_SCALE     float32 = 2    // 角色慢行时 MAX_STEP 与最大步长的比值
	ROBBER_STEAL   float32 = 2000 // 未剿灭山贼时被偷的钱数
	TAX_SCALE      float32 = 6    // 城市税收 (角色到达被他人占领城市所付的钱) 与城市每回合正常收入的比例
	OWN_TAX_SCALE  float32 = 3    // 自我城市税收 (角色到达自己所占领城市时收入的钱) 与城市每回合正常收入的比例
	HPPLUS_SCALE   float32 = 1.5  // 城市每回合恢复 HP 与市长政治属性的比例
	ALLOCATE_TURN          = 8    // 每多少回合进行一次人员调配

	BATTLE_ROUND_MIN          = 40 // 战斗回合限制基数
	BATTLE_ROUND_STEP         = 30 // 战斗回合限制增量 (随战斗规模增加)
	BURN_HURT         float32 = 20 // 战时燃烧伤害

	CARD_STEAL_MONEY    float32 = 1500 // 通过卡片偷钱的数量
	CARD_HP_RECOVER_SCALE float32 = 0.1 // 通过卡片恢复 HP 的比例
)

var zei Officer = Officer{
	name:  "robber",
	role:  -1,
	job:   -1,
	loc:   -1,
	hpmax: 1800,
	hp:    1500,
	prop:  [5]float32{60, 50, 10, 20, 30},
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

// crole: 当前操作角色
func (d *Driver) Save (fout io.Writer, crole int) {
	fmt.Fprintf(fout, "%d %d %d %d %d %d %d %f %d %d %d\n",
		d.ngrid, d.nrole, d.npeople, d.ncity, d.ninst, d.ntrain, d.turn, d.win_money, d.nx, d.ny, crole)

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

// 返回当前操作角色
func (d *Driver) Load(fin io.Reader) (crole int) {
	fmt.Fscan(fin,
		&d.ngrid, &d.nrole, &d.npeople, &d.ncity, &d.ninst, &d.ntrain, &d.turn, &d.win_money, &d.nx, &d.ny, &crole)

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
		d.lpset[i] = dsg.LoadSet(fin)
	}

	return 
}
