package mbg

import (
	"math/rand"

	"github.com/tfcolin/dsg"
)

func (b *BattleState) ChangeHP(dhp float32) {
	ohp := b.hp
	odef := b.def
	if dhp > 0 {
		b.hp += dhp
		if b.hp > b.hpmax {
			b.hp = b.hpmax
		}
	} else {
		dhp = -dhp
		if b.def > dhp {
			b.def -= dhp
			if b.def < 1 {
				b.def = 0
			}
		} else {
			dhp -= b.def
			b.def = 0
			b.hp -= dhp
			if b.hp < 1 {
				b.hp = 0
			}
		}
	}

	b.last_ddef = b.def - odef
	b.last_dhp = b.hp - ohp
	if b.hp == 0 {
		b.is_die = true
	} else {
		b.is_die = false
	}
}

func (sst *BattleState) RunTech(tind int, dhp float32, ost *BattleState) (is_run bool, is_die bool) {
	tech := &(techs[tind])

	if sst.latency[tind] != 0 {
		return
	}

	tprob := tech.lprob + (sst.celve/100.0)*(tech.hprob-tech.lprob)
	trn := rand.Float32()

	if trn > tprob {
		return
	}

	tvalue := tech.lvalue + (sst.celve/100.0)*(tech.hvalue-tech.lvalue)
	sst.latency[tind] = tech.latency
	is_run = true

	sdhp, odhp := tech.do(sst, ost, dhp, tvalue)
	sst.ChangeHP(sdhp)
	ost.ChangeHP(odhp)
	is_die = ost.is_die

	return
}

func DoBst(bst []BattleState) {
	for j := 0; j < 2; j++ {
		if bst[j].bst != 0 {
			bst[j].btime--
			if bst[j].btime == 0 {
				bst[j].bst = 0
			}
		}
		if bst[j].fst != 0 {
			bst[j].fst--
		}
		for k := 0; k < TECH_COUNT; k++ {
			if bst[j].latency[k] > 0 {
				bst[j].latency[k]--
			}
		}
	}
}

// return result: 0: 平局. 1: 发动方胜. -1: 被动方胜
func RunBattle(uv UserView, bst []BattleState, scale int) int {

	var max_bt int = BATTLE_ROUND_MIN + BATTLE_ROUND_STEP*scale
	uv.BattleStart(&(bst[0]), &(bst[1]), max_bt)

turn:
	for i := 0; i < max_bt; i++ {
		uv.BattleTurnStart(i)

		var iwinp [2]float32

		for j := 0; j < 2; j++ {
			if bst[j].bst == 0 {
				iwinp[j] = bst[j].winp
			} else {
				iwinp[j] = 0
			}
			if bst[j].fst != 0 {
				bst[j].ChangeHP(-BURN_HURT)
				uv.BattleBurn(j, &(bst[j]))
			}
			if bst[j].bst == 2 {
				bst[j].ChangeHP(-bst[j].power)
				uv.BattleSelfAttack(j, &(bst[j]))
			}
		}

		if bst[0].is_die || bst[1].is_die {
			DoBst(bst)
			break
		}

		tp := iwinp[0] + iwinp[1]
		if tp < 0.01 {
			DoBst(bst)
			continue
		}

		var ws, ls int
		var prob [2]float32 = [2]float32{iwinp[0] / tp, iwinp[1] / tp}
		rn := rand.Float32()
		if rn < prob[0] {
			ws = 0
		} else {
			ws = 1
		}
		ls = 1 - ws

		bst[ls].ChangeHP(-bst[ws].power)
		uv.BattleAttack(ls, &(bst[ls]))
		dhp := bst[ls].last_dhp + bst[ls].last_ddef
		if bst[ls].is_die {
			DoBst(bst)
			break
		}

		for j := 0; j < 2; j++ {
			for _, tind := range bst[j].tech.GetAllLabel() {
				tech := &(techs[tind])
				if tech.scond == 1 && ws != j {
					continue
				}
				if tech.scond == 2 && ls != j {
					continue
				}
				if tech.scond == 3 {
					continue
				}
				is_run, is_die := bst[j].RunTech(tind, dhp, &(bst[1-j]))
				if is_run {
					uv.BattleTech(j, tind, &(bst[j]), &(bst[1-j]))
				}
				if is_die {
					break turn
				}
			}
		}

		DoBst(bst)
	}

	var res int
	switch {
	case bst[0].is_die && bst[1].is_die:
		res = 0
	case bst[0].is_die && !bst[1].is_die:
		res = -1
		for _, tind := range bst[1].tech.GetAllLabel() {
			if techs[tind].scond == 3 {
				is_run, _ := bst[1].RunTech(tind, bst[0].last_dhp, &(bst[0]))
				if is_run {
					uv.BattleTech(1, tind, &(bst[1]), &(bst[0]))
				}
			}
		}
	case bst[1].is_die && !bst[0].is_die:
		res = 1
		for _, tind := range bst[0].tech.GetAllLabel() {
			if techs[tind].scond == 3 {
				is_run, _ := bst[0].RunTech(tind, bst[1].last_dhp, &(bst[1]))
				if is_run {
					uv.BattleTech(0, tind, &(bst[0]), &(bst[1]))
				}
			}
		}
	default:
		res = 0
	}

	uv.BattleEnd(res)
	return res
}

// 作战: r = -1: robber; loc = -1: 随身; scale: 规模 0-2 小中大
// return: 0: 和平结束 -1: 1 方失败; 1: 1 方获胜
func (d *Driver) Battle(r1 int, loc1 int, r2 int, loc2 int, scale int) int {

	var rind [2]int = [2]int{r1, r2}
	var pind [2][2]int
	var gen [2][2](*Officer)
	var cind [2]int = [2]int{loc1, loc2}
	var city [2](*City)

	var bst [2]BattleState

	var res int

	for i := 0; i < 2; i++ {
		if rind[i] == -1 {
			gen[i][0] = new(Officer)
			*gen[i][0] = zei
			gen[i][1] = nil
			pind[i][0] = -1
			pind[i][1] = -1
		} else {
			var sel_pset []int
			if cind[i] == -1 {
				for _, j := range d.roles[rind[i]].mos.GetAllLabel() {
					if d.people[j].lst != 0 {
						continue
					}
					sel_pset = append(sel_pset, j)
				}
			} else {
				for _, j := range d.cities[cind[i]].mos.GetAllLabel() {
					if d.people[j].lst != 0 {
						continue
					}
					sel_pset = append(sel_pset, j)
				}
			}
			main, vice := d.oi[rind[i]].SelGeneral(sel_pset)
			if main == -1 {
				vice = -1
			}
			if main == vice {
				vice = -1
			}
			if main == -1 {
				pind[i][0] = -1
			} else {
				pind[i][0] = sel_pset[main]
			}
			if vice == -1 {
				pind[i][1] = -1
			} else {
				pind[i][1] = sel_pset[vice]
			}
			for j := 0; j < 2; j++ {
				if pind[i][j] != -1 {
					gen[i][j] = &(d.people[pind[i][j]])
				}
			}
		}
		if gen[i][0] == nil {
			d.uv.BattleNoGerenal(rind[i], i)
			return (i*2 - 1)
		}
		if cind[i] != -1 {
			city[i] = &(d.cities[cind[i]])
		}
	}

	for i := 0; i < 2; i++ {
		bst[i].role = rind[i]
		bst[i].winp = gen[i][0].GetProp(Zhan)
		bst[i].power = gen[i][0].GetProp(Wu)
		bst[i].celve = gen[i][0].GetProp(Mou)

		if gen[i][1] != nil {
			if gen[i][1].GetProp(Wu) > bst[i].power {
				bst[i].power = gen[i][1].GetProp(Wu)
			}
			if gen[i][1].GetProp(Mou) > bst[i].celve {
				bst[i].celve = gen[i][1].GetProp(Mou)
			}
		}

		bst[i].hpmax = gen[i][0].hp
		bst[i].hp = bst[i].hpmax
		if city[i] != nil {
			bst[i].def = city[i].hp
		} else {
			bst[i].def = 0
		}

		bst[i].tech = dsg.InitSet(TECH_COUNT)
		if rind[i] != -1 {
			bst[i].tech.CopyFrom(d.roles[rind[i]].tech)
		}
	}

	res = RunBattle(d.uv, bst[:], scale)

	for i := 0; i < 2; i++ {
		gen[i][0].hp = bst[i].hp
		gen[i][0].is_quit = gen[i][0].is_quit || bst[i].is_quit
		if city[i] != nil {
			city[i].hp = bst[i].def
		}
	}

	return res
}

func InitTech() {

	techs[0] = Tech{
		name:    "弓箭",
		study:   5000,
		scond:   0,
		lprob:   0,
		hprob:   0.8,
		lvalue:  20,
		hvalue:  30,
		latency: 2,
	}

	techs[0].do = func(sst *BattleState, ost *BattleState, dhp float32, value float32) (sdhp, odhp float32) {
		odhp = -value
		return
	}

	techs[1] = Tech{
		name:    "火箭",
		study:   6000,
		scond:   0,
		lprob:   0,
		hprob:   0.6,
		lvalue:  15,
		hvalue:  15,
		latency: 3,
	}

	techs[1].do = func(sst *BattleState, ost *BattleState, dhp float32, value float32) (sdhp, odhp float32) {
		odhp = -value
		ost.fst = 4
		return
	}

	techs[2] = Tech{
		name:    "医疗",
		study:   7500,
		scond:   0,
		lprob:   0,
		hprob:   0.7,
		lvalue:  20,
		hvalue:  30,
		latency: 3,
	}

	techs[2].do = func(sst *BattleState, ost *BattleState, dhp float32, value float32) (sdhp, odhp float32) {
		sdhp = value
		return
	}

	techs[3] = Tech{
		name:    "盾牌",
		study:   8000,
		scond:   2,
		lprob:   0,
		hprob:   0.3,
		lvalue:  1,
		hvalue:  1,
		latency: 5,
	}

	techs[3].do = func(sst *BattleState, ost *BattleState, dhp float32, value float32) (sdhp, odhp float32) {
		if dhp < 0 {
			sdhp = -dhp * value
		} else {
			sdhp = 0
		}
		return
	}

	techs[4] = Tech{
		name:    "转移",
		study:   9000,
		scond:   2,
		lprob:   0,
		hprob:   0.3,
		lvalue:  0.5,
		hvalue:  0.5,
		latency: 5,
	}

	techs[4].do = func(sst *BattleState, ost *BattleState, dhp float32, value float32) (sdhp, odhp float32) {
		if dhp < 0 {
			sdhp = -dhp * value
			odhp = dhp * value
		}
		return
	}

	techs[5] = Tech{
		name:    "强击",
		study:   8500,
		scond:   1,
		lprob:   0,
		hprob:   0.4,
		lvalue:  0.5,
		hvalue:  1.0,
		latency: 5,
	}

	techs[5].do = func(sst *BattleState, ost *BattleState, dhp float32, value float32) (sdhp, odhp float32) {
		if dhp < 0 {
			odhp = dhp * value
		}
		return
	}

	techs[6] = Tech{
		name:    "离间",
		study:   7000,
		scond:   1,
		lprob:   0,
		hprob:   0.4,
		lvalue:  0.2,
		hvalue:  0.3,
		latency: 5,
	}

	techs[6].do = func(sst *BattleState, ost *BattleState, dhp float32, value float32) (sdhp, odhp float32) {
		if dhp < 0 {
			sdhp = -dhp * value
		}
		return
	}

	techs[7] = Tech{
		name:    "雷击",
		study:   9000,
		scond:   0,
		lprob:   0.1,
		hprob:   0.1,
		lvalue:  100,
		hvalue:  180,
		latency: 6,
	}

	techs[7].do = func(sst *BattleState, ost *BattleState, dhp float32, value float32) (sdhp, odhp float32) {
		odhp = -value
		return
	}

	techs[8] = Tech{
		name:    "混乱",
		study:   9500,
		scond:   1,
		lprob:   0,
		hprob:   0.5,
		lvalue:  0,
		hvalue:  0,
		latency: 5,
	}

	techs[8].do = func(sst *BattleState, ost *BattleState, dhp float32, value float32) (sdhp, odhp float32) {
		ost.bst = 1
		ost.btime = 3
		return
	}

	techs[9] = Tech{
		name:    "内讧",
		study:   13000,
		scond:   0,
		lprob:   0,
		hprob:   0.25,
		lvalue:  0,
		hvalue:  0,
		latency: 6,
	}

	techs[9].do = func(sst *BattleState, ost *BattleState, dhp float32, value float32) (sdhp, odhp float32) {
		ost.bst = 2
		ost.btime = 3
		return
	}

	techs[10] = Tech{
		name:    "攻城",
		study:   10000,
		scond:   1,
		lprob:   1,
		hprob:   1,
		lvalue:  0.5,
		hvalue:  3,
		latency: 0,
	}

	techs[10].do = func(sst *BattleState, ost *BattleState, dhp float32, value float32) (sdhp, odhp float32) {
		if ost.def > 0 && dhp < 0 {
			odhp = value * dhp
			if ost.def+odhp < 0 {
				odhp = -ost.def
			}
		}
		return
	}

	techs[11] = Tech{
		name:    "流放",
		study:   11000,
		scond:   3,
		lprob:   0,
		hprob:   0.5,
		lvalue:  0,
		hvalue:  0,
		latency: 0,
	}

	techs[11].do = func(sst *BattleState, ost *BattleState, dhp float32, value float32) (sdhp, odhp float32) {
		ost.is_quit = true
		return
	}

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
