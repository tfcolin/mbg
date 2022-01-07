package mbg

import (
	"fmt"
	"io"

	"github.com/tfcolin/dsg"
)

type BoardPoint struct {
	class   int // 0: 空地 1:城 2:会馆 3:谋略研究所 4:修炼所 5: 机会
	base    int
	roles   []bool
	barrier bool
	robber  bool
	x, y    int
}

type City struct {
	name           string
	scope          int // 0:小 1:中 2:大
	hpmax          float32
	hp             float32
	role           int      // -1: 未被占领
	mos            *dsg.Set // 占领人员 (Empty: 未被占领)
	mayor          int      // -1: 未指定
	treasurer      int      // -1: 未指定 (收益基础 20)
	fengshui       float32  //
	x0, y0, x1, y1 int
}

type Institute struct {
	name           string
	tech           []int     // 支持的研发策略编号
	mos            []int     // [nrole] 驻扎人员. -1: 未驻扎
	on_study       []int     // [nrole] 正在研发的项目 -1: 未研发
	point          []float32 // [nrole] 研发剩余点数
	x0, y0, x1, y1 int
}

type TrainingRoom struct {
	name           string
	item           Property // 0-4: 武术, 战术, 谋略, 政治, 经济
	mos            []int    // 驻扎人员(每个 role): -1: 未驻扎
	round          []int    // 训练剩余回合数
	x0, y0, x1, y1 int
}

func (d *Driver) GetBoardInfo(ind int) (class int, base int, roles []bool, barrier bool, robber bool) {
	if ind == -1 {
		return
	}
	b := &(d.board[ind])
	class = b.class
	base = b.base
	roles = make([]bool, d.nrole)
	copy(roles, b.roles)
	barrier = b.barrier
	robber = b.robber
	return
}

func (d *Driver) GetCityInfo(cind int) (
	name string, scope int, hpmax, hp float32,
	role int, mos []int, mayor int, treasurer int, fengshui float32,
	hpplus float32, earn float32) {
	if cind == -1 {
		return
	}
	c := &(d.cities[cind])
	name = c.name
	scope = c.scope
	hpmax = c.hpmax
	hp = c.hp
	role = c.role
	mos = c.mos.GetAllLabel()
	mayor = c.mayor
	treasurer = c.treasurer
	fengshui = c.fengshui
	hpplus = c.GetHPPlus(d)
	earn = c.GetEarn(d)

	return
}

func (d *Driver) GetInstInfo(iind int, rind int) (name string, tech []int, mos int, on_study int, point float32, round int) {
	if iind == -1 {
		return
	}
	inst := &(d.insts[iind])
	name = inst.name
	tech = make([]int, len(inst.tech))
	copy(tech, inst.tech)
	if rind == -1 {
		return
	}
	mos = inst.mos[rind]
	on_study = inst.on_study[rind]
	point = inst.point[rind]
	if mos != -1 {
		v := d.people[mos].GetProp(Mou)
		round = (int)(point/(STUDY_SCALE*v)) + 1
	}
	return
}

func (d *Driver) GetTrainInfo(tind int, rind int) (name string, item Property, mos int, round int) {
	if tind == -1 {
		return
	}
	train := &(d.trains[tind])
	name = train.name
	item = train.item
	if rind == -1 {
		return
	}
	mos = train.mos[rind]
	round = train.round[rind]
	return
}

func (city *City) Save(fout io.Writer) {
	fmt.Fprintf(fout, "%s %d %f %f %d %d %d %f \n",
		city.name, city.scope, city.hpmax, city.hp, city.role, city.mayor, city.treasurer, city.fengshui)
	city.mos.Save(fout)
}

func (city *City) Load(fin io.Reader) {
	fmt.Fscan(fin, &city.name, &city.scope, &city.hpmax, &city.hp, &city.role, &city.mayor, &city.treasurer, &city.fengshui)
	city.mos.Load(fin)
}

func (inst *Institute) Save(fout io.Writer) {
	fmt.Fprintf(fout, "%s %d \n", inst.name, len(inst.tech))
	for _, i := range inst.tech {
		fmt.Fprintf(fout, "%d ", i)
	}
	fmt.Fprintf(fout, "\n")

	fmt.Fprintf(fout, "%d ", len(inst.mos))
	for _, i := range inst.mos {
		fmt.Fprintf(fout, "%d ", i)
	}
	fmt.Fprintf(fout, "\n")

	fmt.Fprintf(fout, "%d ", len(inst.on_study))
	for _, i := range inst.on_study {
		fmt.Fprintf(fout, "%d ", i)
	}
	fmt.Fprintf(fout, "\n")

	fmt.Fprintf(fout, "%d ", len(inst.point))
	for _, i := range inst.point {
		fmt.Fprintf(fout, "%f ", i)
	}
	fmt.Fprintf(fout, "\n")

}

func (inst *Institute) Load(fin io.Reader) {
	var n int
	fmt.Fscan(fin, &inst.name)

	fmt.Fscan(fin, &n)
	inst.tech = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(fin, &inst.tech[i])
	}

	fmt.Fscan(fin, &n)
	inst.mos = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(fin, &inst.mos[i])
	}

	fmt.Fscan(fin, &n)
	inst.on_study = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(fin, &inst.on_study[i])
	}

	fmt.Fscan(fin, &n)
	inst.point = make([]float32, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(fin, &inst.point[i])
	}

}

func (t *TrainingRoom) Save(fout io.Writer) {
	fmt.Fprintf(fout, "%s %d \n", t.name, int(t.item))

	fmt.Fprintf(fout, "%d ", len(t.mos))
	for _, i := range t.mos {
		fmt.Fprintf(fout, "%d ", i)
	}
	fmt.Fprintf(fout, "\n")

	fmt.Fprintf(fout, "%d ", len(t.round))
	for _, i := range t.round {
		fmt.Fprintf(fout, "%d ", i)
	}
	fmt.Fprintf(fout, "\n")

}

func (t *TrainingRoom) Load(fin io.Reader) {
	fmt.Fscan(fin, &t.name, &t.item)
	var n int

	fmt.Fscan(fin, &n)
	t.mos = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(fin, &t.mos[i])
	}

	fmt.Fscan(fin, &n)
	t.round = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(fin, &t.round[i])
	}

}

func (bp *BoardPoint) Save(fout io.Writer) {
	fmt.Fprintf(fout, "%d %d %t %t \n", bp.class, bp.base, bp.barrier, bp.robber)

	fmt.Fprintf(fout, "%d ", len(bp.roles))
	for _, i := range bp.roles {
		fmt.Fprintf(fout, "%d ", i)
	}
	fmt.Fprintf(fout, "\n")

}

func (bp *BoardPoint) Load(fin io.Reader) {
	fmt.Fscan(fin, &bp.class, &bp.base, &bp.barrier, &bp.robber)
	var n int

	fmt.Fscan(fin, &n)
	bp.roles = make([]bool, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(fin, &bp.roles[i])
	}

}
