package mbg

import (
	"fmt"
	"io"
	"math/rand"
	"time"

	"gitee.com/tfcolin/dsg"
)

func Init() {
	rand.Seed(time.Now().UnixNano())
	InitTech()
	InitCard()
	force_quit = false
}

func LoadMap(mfile io.Reader) (d *Driver) {
	d = new(Driver)

	var err error

	var nx, ny int
	var x0, y0, x1, y1 int

	fmt.Fscan(mfile, &d.ngrid, &nx, &ny, &d.win_money)
	d.nx = nx
	d.ny = ny
	d.turn = 1

	d.board = make([]BoardPoint, d.ngrid)
	for i := 0; i < d.ngrid; i++ {
		var x, y int
		_, err = fmt.Fscan(mfile, &(d.board[i].class))
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		_, err = fmt.Fscan(mfile, &(d.board[i].base))
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		if d.board[i].class == 0 || d.board[i].class == 2 || d.board[i].class == 5 {
			d.board[i].base = 0
		}
		_, err = fmt.Fscan(mfile, &x, &y)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		d.board[i].x = x
		d.board[i].y = y
		d.board[i].barrier = false
		d.board[i].robber = false
	}

	fmt.Fscan(mfile, &d.npeople)
	for i := 0; i < 3; i++ {
		d.lpset[i] = dsg.InitSet(d.npeople)
	}
	d.people = make([]Officer, d.npeople)
	for i := 0; i < d.npeople; i++ {
		p := &(d.people[i])
		fmt.Fscan(mfile, &(p.name))
		var tp float32 = 0
		for j := 0; j < 5; j++ {
			fmt.Fscan(mfile, &(p.prop[j]))
			tp += p.prop[j]
		}
		p.role = -1
		p.job = -1
		p.loc = -1
		p.hpmax = P_HP_BASE + tp * 0.2 * P_HP_SCALE
		p.hp = p.hpmax

		p.lst = 0
		p.pst = 0

		p.level = p.GetLevel()
		d.lpset[p.level].SetLabel(i, true)
	}

	_, err = fmt.Fscan(mfile, &d.ncity)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	d.cities = make([]City, d.ncity)
	for i := 0; i < d.ncity; i++ {
		city := &(d.cities[i])
		fmt.Fscan(mfile, &(city.name), &(city.scope))
		city.role = -1
		var nmos, imos int
		fmt.Fscan(mfile, &(nmos))
		city.mos = dsg.InitSet(d.npeople)
		for j := 0; j < nmos; j++ {
			fmt.Fscan(mfile, &(imos))
			city.mos.SetLabel(imos, true)
		}
		switch city.scope {
		case 0:
			city.hpmax = DEF_SMALL
		case 1:
			city.hpmax = DEF_MEDIUM
		case 2:
			city.hpmax = DEF_BIG
		}
		city.hp = city.hpmax
		fmt.Fscan(mfile, &(city.mayor), &(city.treasurer))
		city.fengshui = 0
		fmt.Fscan(mfile, &x0, &y0, &x1, &y1)
		city.x0 = x0
		city.y0 = y0
		city.x1 = x1
		city.y1 = y1
	}

	fmt.Fscan(mfile, &d.nrole)
	d.roles = make([]Role, d.nrole)
	for i := 0; i < d.nrole; i++ {
		r := &(d.roles[i])
		fmt.Fscan(mfile, &(r.name), &(r.loc))

		var ntech, itech int
		fmt.Fscan(mfile, &(ntech))
		r.tech = dsg.InitSet(TECH_COUNT)
		for j := 0; j < ntech; j++ {
			fmt.Fscan(mfile, &(itech))
			r.tech.SetLabel(itech, true)
		}

		var nmos, imos int
		fmt.Fscan(mfile, &(nmos))
		r.mos = dsg.InitSet(d.npeople)
		for j := 0; j < nmos; j++ {
			fmt.Fscan(mfile, &(imos))
			r.mos.SetLabel(imos, true)
		}

		var nmcs, imcs int
		fmt.Fscan(mfile, &(nmcs))
		r.mcs = dsg.InitSet(d.ncity)
		for j := 0; j < nmcs; j++ {
			fmt.Fscan(mfile, &(imcs))
			r.mcs.SetLabel(imcs, true)
		}

		r.sos = dsg.InitSet(d.npeople)
		r.tos = dsg.InitSet(d.npeople)

		fmt.Fscan(mfile, &(r.money))
		fmt.Fscan(mfile, &(r.dir))

		r.mst = 0
		r.mtime = 0
		r.cst = false
		r.ast = -1
		r.atime = 0
	}

	fmt.Fscan(mfile, &d.ninst)
	d.insts = make([]Institute, d.ninst)
	for i := 0; i < d.ninst; i++ {
		inst := &(d.insts[i])
		var ntech int
		fmt.Fscan(mfile, &(inst.name), &ntech)
		inst.tech = make([]int, ntech)
		for j := 0; j < ntech; j++ {
			fmt.Fscan(mfile, &(inst.tech[j]))
		}
		inst.mos = make([]int, d.nrole)
		inst.on_study = make([]int, d.nrole)
		inst.point = make([]float32, d.nrole)
		for j := 0; j < d.nrole; j++ {
			inst.mos[j] = -1
			inst.on_study[j] = -1
			inst.point[j] = 0
		}
		fmt.Fscan(mfile, &x0, &y0, &x1, &y1)
		inst.x0 = x0
		inst.y0 = y0
		inst.x1 = x1
		inst.y1 = y1
	}

	fmt.Fscan(mfile, &d.ntrain)
	d.trains = make([]TrainingRoom, d.ntrain)
	for i := 0; i < d.ntrain; i++ {
		train := &(d.trains[i])
		fmt.Fscan(mfile, &(train.name), &(train.item))
		train.mos = make([]int, d.nrole)
		train.round = make([]int, d.nrole)
		for j := 0; j < d.nrole; j++ {
			train.mos[j] = -1
			train.round[j] = 0
		}
		fmt.Fscan(mfile, &x0, &y0, &x1, &y1)
		train.x0 = x0
		train.y0 = y0
		train.x1 = x1
		train.y1 = y1
	}

	for i := 0; i < d.ngrid; i++ {
		d.board[i].roles = make([]bool, d.nrole)
	}
	for i := 0; i < d.nrole; i++ {
		l := d.roles[i].loc
		if l >= 0 {
			d.board[l].roles[i] = true
		}

		mos := d.roles[i].mos.GetAllLabel()
		for _, ind := range mos {
			p := &(d.people[ind])
			p.role = i
		}

		mcs := d.roles[i].mcs.GetAllLabel()
		for _, j := range mcs {
			c := &(d.cities[j])
			c.role = i
			pinc := c.mos.GetAllLabel()
			for _, k := range pinc {
				d.people[k].role = i
			}
		}
	}

	for i := 0; i < d.ncity; i++ {
		c := d.cities[i]
		pinc := c.mos.GetAllLabel()
		for _, j := range pinc {
			d.people[j].loc = i
			d.people[j].job = 2
		}
		if c.mayor != -1 {
			d.people[c.mayor].job = 0
		}
		if c.treasurer != -1 {
			d.people[c.treasurer].job = 1
		}
	}

	for i := 0; i < d.npeople; i++ {
		p := &(d.people[i])
		if p.role == -1 && p.loc != -1 {
			d.POutCity(i, false)
		}
		if p.role != -1 {
			d.lpset[p.level].SetLabel(i, false)
		}
	}

	return d
}

func (d *Driver) ConnectUI(uv UserView, oi []OperationInterface) {
	if len(oi) < d.nrole {
		panic("not enough operation interface")
	}

	d.uv = uv
	uv.InitMap(d.ngrid, d.nrole, d.npeople, d.ncity, d.ninst, d.ntrain, TECH_COUNT, CARD_COUNT, d.nx, d.ny)
	for j, city := range d.cities {
		uv.ShowCity(j, city.name, city.scope, city.x0, city.y0, city.x1, city.y1)
	}
	for j, inst := range d.insts {
		uv.ShowInstitute(j, inst.name, inst.x0, inst.y0, inst.x1, inst.y1)
	}
	for j, train := range d.trains {
		uv.ShowTrainingRoom(j, train.name, train.item, train.x0, train.y0, train.x1, train.y1)
	}
	for j, m := range d.board {
		uv.ShowMap(j, m.class, m.base, m.x, m.y)
	}

	d.oi = oi[0:d.nrole]
}

// return: -1: 正常; >0: 游戏结束, 返回胜利角色号. -2: 全部破产
func (d *Driver) RoundEnd() int {
	for i := 0; i < d.ninst; i++ {
		d.RoundInst(i)
	}
	for i := 0; i < d.ntrain; i++ {
		d.RoundTrain(i)
	}
	for i := 0; i < d.ncity; i++ {
		d.RoundCity(i)
	}
	for i := 0; i < d.npeople; i++ {
		d.RoundPeople(i)
	}

	var live int = -1
	var nlive int = 0
	for i := 0; i < d.nrole; i++ {
		ires := d.RoundRole(i)
		if ires == 1 {
			return i
		}
		if d.roles[i].loc != -1 {
			live = i
			nlive++
		}
	}
	if nlive == 0 {
		d.uv.NoteAllLoss()
		return -2
	}
	if nlive == 1 {
		d.uv.NoteRoleWin(live)
		return live
	}

	d.uv.EndTurn(d.turn)

	if d.turn % ALLOCATE_TURN == 0 {
		for rind := 0; rind < d.nrole; rind++ {
			r := &(d.roles[rind])
			if r.loc == -1 {continue}
		sel:
			for {
				av_locs := make ([]bool, d.ngrid)
				for i, b := range d.board {
					switch b.class {
					case 1:
						if d.cities[b.base].role == rind {
							av_locs[i] = true
						}
					case 3:
						if d.insts[b.base].mos[rind] != -1 {
							av_locs[i] = true
						}
					case 4:
						if d.trains[b.base].mos[rind] != -1 {
							av_locs[i] = true
						}
					}
				}
				d.oi[rind].StartAllocate (av_locs)
				st, obj := d.oi[rind].SelAllocObj()
				if obj == -1 {
					st = -1
				}
				switch st {
				case -1:
					break sel
				case 0:
					c := &(d.cities[obj])
					if c.role != rind {
						continue
					}
					spl := r.mos.GetAllLabel()
					opl := c.mos.GetAllLabel()
					il, ol := d.oi[rind].ExchangeCityPeople(obj, spl, opl)
					for _, i := range il {
						d.PInCity(spl[i], obj, true)
					}
					for _, i := range ol {
						d.POutCity(opl[i], true)
					}
					pllist := c.mos.GetAllLabel()
					if c.role != -1 {
						mayor := d.oi[rind].SelMayor(obj, pllist, c.mayor, c.treasurer)
						if mayor != -1 {
							d.PAsMayer(pllist[mayor], true)
						}
						trea := d.oi[rind].SelTreasurer(obj, pllist, c.mayor, c.treasurer)
						if trea != -1 {
							d.PAsTreasurer(pllist[trea], true)
						}
					}
				case 1:
					inst := &(d.insts[obj])
					if inst.mos[rind] == -1 {
						continue
					}
					if d.oi[rind].IsCancelStudy(inst.mos[rind], inst.on_study[rind], inst.point[rind]) {
						d.POutInst(inst.mos[rind], true)
					} else {
						continue
					}
				case 2:
					train := &(d.trains[obj])
					if train.mos[rind] == -1 {
						continue
					}
					if d.oi[rind].IsCancelTrain(train.mos[rind], train.item, train.round[rind]) {
						d.POutTrain(train.mos[rind], true)
					} else {
						continue
					}
				}
			}
		}
	}

	d.turn++

	return -1
}

// 返回: -3: 强制退出. -2: ui 或 oi 未准备好. -1: 平局. >=0: 胜利阵营号
// max_turn: 最大回合数
// crole: 首个操作角色
func (d *Driver) Run (max_turn int, crole int) int {

	if d.uv == nil {
		return -2
	}
	for i := 0; i < d.nrole; i++ {
		if d.oi[i] == nil {
			return -2
		}
	}

	d.uv.StartGame(d)
	for i := 0; i < d.nrole; i++ {
		d.oi[i].StartGame(d, i)
	}

	var res int
	is_first_round := true
	for {
		for i := 0; i < d.nrole; i++ {
			if is_first_round && i < crole {
				continue
			}
			if d.RoleAction(i) {
				d.roles[i].cst = false
				i--
			}
			if force_quit {
				d.uv.NoteForceQuit()
				return -3
			}
		}
		res = d.RoundEnd()
		is_first_round = false
		if res == -2 {
			return -1
		}
		if res >= 0 {
			return res
		}
		if d.turn > max_turn {
			break
		}
	}

	var imax_money int = -1
	var mm float32 = 0
	for i := 0; i < d.nrole; i++ {
		if imax_money == -1 || d.roles[i].money > mm {
			imax_money = i
			mm = d.roles[i].money
		}
	}
	d.uv.NoteRoleWin(imax_money)

	res = imax_money

	return res
}

// only for test
func (d *Driver) FillCards(ncards int) {
	for i := 0; i < d.nrole; i++ {
		role := &(d.roles[i])
		for j := 0; j < CARD_COUNT; j++ {
			role.cards[j] = ncards
		}
	}
}
