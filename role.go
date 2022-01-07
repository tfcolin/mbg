package mbg

import (
	"math/rand"

	"github.com/tfcolin/dsg"
)

// atime resatime : 0: 正常; 1: 减为零
func (r *Role) ChangeMoney(dm float32) (real_dm float32, res int) {
	if r.money+dm < 0 {
		real_dm = -r.money
		res = 1
	} else {
		real_dm = dm
		res = 0
	}
	r.money += real_dm
	return
}

// 阵营回合结算,
// 返回 0: 正常; 1: 胜利; 2: 破产
func (d *Driver) RoundRole(rind int) int {
	r := &(d.roles[rind])
	if r.mst != 0 {
		st := r.mst
		r.mtime--
		if r.mtime == 0 {
			r.mst = 0
			d.uv.NoteRoleStEnd(rind, st)
		}
	}
	if r.ast != -1 {
		st := r.ast
		r.atime--
		if r.atime == 0 {
			r.ast = -1
			if st == -2 {
				d.uv.NoteRoleCeEnd(rind)
			} else if rind < st {
				d.uv.NoteRoleAlignEnd(rind, st)
			}
		}
	}
	if r.cst {
		r.cst = false
	}

	if r.money >= d.win_money {
		d.uv.NoteRoleWin(rind)
		return 1
	}

	if r.money <= 0 {
		// 死亡结算
		r.money = 0
		r.loc = -1
		d.uv.NoteRoleLoss(rind)
		for _, pind := range r.mos.GetAllLabel() {
			d.PQuit(pind, true)
		}
		for _, cind := range r.mcs.GetAllLabel() {
			c := &(d.cities[cind])
			for _, pind := range c.mos.GetAllLabel() {
				d.PQuit(pind, true)
			}
		}
		for _, pind := range r.sos.GetAllLabel() {
			d.PQuit(pind, true)
		}
		for _, pind := range r.tos.GetAllLabel() {
			d.PQuit(pind, true)
		}
		r.mst = 0
		r.mtime = 0
		r.cst = false
		if r.ast >= 0 {
			find := r.ast
			friend := &(d.roles[find])
			r.ast = -1
			r.atime = 0
			friend.ast = -1
			friend.atime = 0
			if rind < find {
				d.uv.NoteRoleAlignEnd(rind, find)
			} else {
				d.uv.NoteRoleAlignEnd(find, rind)
			}
		}
		return 2
	}
	return 0
}

func (r *Role) GetMaxStep() int {
	switch r.mst {
	case 1:
		return MAX_STEP / SLOW_SCALE
	case 2:
		return MAX_STEP * FAST_SCALE
	case 3:
		return 0
	case 0:
		return MAX_STEP
	default:
		panic("wrong mst value")
	}
}

func (d *Driver) RoleAction(rind int) (is_cont bool) {
	r := &(d.roles[rind])
	if r.loc == -1 {
		return
	}
	if r.mst == 3 {
		d.uv.RoleForbidAction(rind)
		if r.cst {
			d.uv.RoleContinueAction(rind)
		}
		return r.cst
	}
	d.uv.RoleStartAction(rind)
sel:
	for {
		action := d.oi[rind].SelRoleAction(r.cards[:])
		if action == -2 {
			force_quit = true
			return false
		}
		if action == -1 {
			max_step := r.GetMaxStep()
			var step int
			if max_step == 0 {
				step = 0
			} else {
				step = rand.Intn(max_step) + 1
			}

			d.board[r.loc].roles[rind] = false
			old_loc := r.loc
			for i := 0; i < step; i++ {
				d.uv.NoteRoleMoveOneStep(rind, r.loc, r.dir)
				r.loc += r.dir
				if r.loc >= d.ngrid {
					r.loc -= d.ngrid
				}
				if r.loc < 0 {
					r.loc += d.ngrid
				}
				if d.board[r.loc].barrier || d.board[r.loc].robber {
					break
				}
			}
			d.uv.NoteRoleMove(rind, old_loc, r.loc, step)

			b := &(d.board[r.loc])
			b.roles[rind] = true

			if b.barrier {
				b.barrier = false
				d.uv.UnSetBarrier(r.loc)
			}
			if b.robber {
				b.robber = false
				d.uv.UnSetRobber(r.loc)
				res := d.Battle(rind, -1, -1, -1, 1)
				if res != 1 {
					r.ChangeMoney(-ROBBER_STEAL)
					d.uv.StealMoney(-1, rind, -ROBBER_STEAL)
				}
			}

			for i := 0; i < d.nrole; i++ {
				if i == rind {
					continue
				}
				if b.roles[i] {
					if r.ast == i { // 结盟
						d.oi[rind].SkipBattleByAlign(i)
						continue
					}
					battle_scale := d.oi[rind].BattleWithRole(i)
					if battle_scale == -1 { // 公愤强制战斗
						if r.ast == -2 {
							d.oi[rind].ForceBattle(i)
							battle_scale = 0
						}
						if d.roles[i].ast == -2 {
							d.oi[i].ForceBattle(rind)
							battle_scale = 0
						}
					}
					if battle_scale != -1 {
						d.Battle(rind, -1, i, -1, battle_scale)
					}
				}
			}

			switch b.class {
			case 1:
				d.CityAction(rind, b.base)
			case 2:
				d.SaloonAction(rind)
			case 3:
				d.InstAction(rind, b.base)
			case 4:
				d.TrainAction(rind, b.base)
			case 5:
				d.ChanceAction(rind)
			}
			break
		} else {
			if r.cards[action] == 0 {
				continue
			}
			card := cards[action]
			switch card.otype {
			case 0:
				if d.oi[rind].ConfirmCard(action) != 0 {
					continue
				}
				if card.do(d, rind, -1) != 0 {
					continue
				}
			case 1:
				for {
					oc_id := d.oi[rind].SelCardObjCity(action)
					if oc_id == -1 {
						continue sel
					}
					oc_rid := d.cities[oc_id].role
					if card.odir == 1 && oc_rid != rind {
						continue
					}
					if card.odir == 2 && (oc_rid == -1 || oc_rid == rind) {
						continue
					}
					if card.do(d, rind, oc_id) != 0 {
						continue
					}
					break
				}
			case 2:
				for {
					or_id := d.oi[rind].SelCardObjRole(action)
					if or_id == -1 {
						continue sel
					}
					if card.odir == 1 && or_id != rind {
						continue
					}
					if card.odir == 2 && or_id == rind {
						continue
					}
					if card.do(d, rind, or_id) != 0 {
						continue
					}
					break
				}
			case 3:
				for {
					oloc := d.oi[rind].SelCardObjLoc(action)
					if oloc == -1 {
						continue sel
					}
					if card.do(d, rind, oloc) != 0 {
						continue
					}
					break
				}
			case 4:
				var sel_rset *dsg.Set = dsg.InitSet(d.nrole)
				for i := 0; i < d.nrole; i++ {
					if card.odir == 1 && i != rind {
						continue
					}
					if card.odir == 2 && i == rind {
						continue
					}
					sel_rset.SetLabel(i, true)
				}
			sel1:
				for {
					st, sloc := d.oi[rind].SelCardObjAny(action)
					switch st {
					case -1:
						continue sel
					case 0:
						if sloc < 0 || sloc >= d.nrole {
							continue
						}
						if !sel_rset.GetLabel(sloc) {
							continue
						}
						for {
							pl := d.roles[sloc].mos.GetAllLabel()
							sp_id := d.oi[rind].SelCardObjPeople(action, pl)
							if sp_id == -1 {
								continue sel1
							}
							if card.do(d, rind, pl[sp_id]) != 0 {
								continue
							}
							break
						}
					case 1:
						if sloc < 0 || sloc >= d.ncity {
							continue
						}
						c := &(d.cities[sloc])
						if card.odir == 1 && c.role != rind {
							continue
						}
						if card.odir == 2 && (c.role == -1 || c.role == rind) {
							continue
						}
						for {
							pl := c.mos.GetAllLabel()
							sp_id := d.oi[rind].SelCardObjPeople(action, pl)
							if sp_id == -1 {
								continue sel1
							}
							if card.do(d, rind, pl[sp_id]) != 0 {
								continue
							}
							break
						}
					case 2:
						if sloc < 0 || sloc >= d.ninst {
							continue
						}
						inst := &(d.insts[sloc])
						for {
							var plist []int
							for i := 0; i < d.nrole; i++ {
								if !sel_rset.GetLabel(i) {
									continue
								}
								if inst.mos[i] < 0 {
									continue
								}
								plist = append(plist, inst.mos[i])
							}
							sr_id := d.oi[rind].SelCardObjRoleAndPeople(action, plist)
							if sr_id == -1 {
								continue sel1
							}
							if card.do(d, rind, plist[sr_id]) != 0 {
								continue
							}
							break
						}
					case 3:
						if sloc < 0 || sloc >= d.ntrain {
							continue
						}
						train := &(d.trains[sloc])
						for {
							var plist []int
							for i := 0; i < d.nrole; i++ {
								if !sel_rset.GetLabel(i) {
									continue
								}
								if train.mos[i] < 0 {
									continue
								}
								plist = append(plist, train.mos[i])
							}
							sr_id := d.oi[rind].SelCardObjRoleAndPeople(action, plist)
							if sr_id == -1 {
								continue sel1
							}
							if card.do(d, rind, plist[sr_id]) != 0 {
								continue
							}
							break
						}
					}
					break
				}
			}
			r.cards[action]--
		}
	}
	if r.cst {
		d.uv.RoleContinueAction(rind)
	}
	return r.cst
}
