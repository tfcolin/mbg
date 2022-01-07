package mbg

import "math/rand"

// res: 0: 正常; 1: 减为零; 2: 加满
func (c * City) ChangeHP (dhp float32) (real_dhp float32, res int) {
      php := c.hp + dhp
      if php < 1 {
            real_dhp = - c.hp
            res = 1
      } else if php >= c.hpmax {
            real_dhp = c.hpmax - c.hp
            res = 2
      } else {
            real_dhp = dhp
            res = 0
      }
      c.hp += real_dhp
      if c.hp < 1 {c.hp = 0}
      return
}

func (c * City) GetHPPlus (d * Driver) float32 {
      var hpplus float32
      if c.mayor == -1 {
            hpplus = 0
      } else {
            hpplus = d.people[c.mayor].GetProp(Zheng) * HPPLUS_SCALE
      }
      return hpplus
}

func (c * City) GetEarn (d * Driver) float32 {
      var mplus, scale, value float32
      if c.treasurer == -1 {
             value = MIN_EARN
      } else {
            if d.people[c.treasurer].lst != 0 {
                  value = MIN_EARN
            } else {
                  value = d.people[c.treasurer].GetProp(Jing)
                  if value < MIN_EARN {
                        value = MIN_EARN
                  }
            }
      }
      scale = 1.5 + float32(c.scope) * 0.5 + c.fengshui
      mplus = value * scale
      return mplus
}

// 研发回合结算
func (d * Driver) RoundInst (iind int) {
      inst := &(d.insts[iind])
      for i := 0; i < d.nrole; i ++ {
            r := &(d.roles[i])
            if inst.mos[i] != -1 {
                  pind := inst.mos[i]
                  study_item := inst.on_study[i]
                  p := &(d.people[pind])
                  value := p.GetProp (Mou)
                  if p.lst == 0 {
                        old_ip := inst.point[i]
                        inst.point[i] -= value * STUDY_SCALE
                        if inst.point[i] <= 0 {
                              r.tech.SetLabel (inst.on_study[i], true)
                              inst.point[i] = 0
                              inst.on_study[i] = -1
                              inst.mos[i] = -1
                              d.oi[i].NoteFinishStudy (iind, pind, study_item)
                              d.POutInst (pind, true)
                        }
                        real_dm, _ := r.ChangeMoney (inst.point[i] - old_ip)
                        d.oi[i].NotePayInst (real_dm)
                  } else {
                        d.oi[i].NoteLYStudy (iind, pind)
                  }
            }
      }
}

// 训练回合结算
func (d * Driver) RoundTrain (tind int) {
      train := &(d.trains[tind])
      for i := 0; i < d.nrole; i ++ {
            if train.mos[i] != -1 {
                  pind := train.mos[i]
                  p := &(d.people[pind])
                  if p.lst == 0 {
                        train.round[i] --
                        if train.round[i] <= 0 {
                              train.mos[i] = -1
                              train.round[i] = 0
                              value := float32(rand.Intn (2) + 1)
                              if p.prop[train.item] + value > 100 {value = 100 - p.prop[train.item]}
                              p.prop[train.item] += value
                              d.oi[i].NoteFinishTrain (pind, train.item, value)
                              d.POutTrain (pind, true)
                        }
                  } else {
                        d.oi[i].NoteLYTrain (tind, pind)
                  }
            }
      }
}

// 城市回合结算
func (d * Driver) RoundCity (cind int) {
      c := &(d.cities[cind])

      c.ChangeHP (c.hpmax * 0.1)

      if c.mayor != -1 {
            if d.people[c.mayor].lst != 0 {
                  if c.role != -1 { d.oi[c.role].NoteLYMayor (c.mayor, cind) }
            } else {
                  hplus := d.people[c.mayor].GetProp(Zheng)
                  for _, pind := range c.mos.GetAllLabel() {
                        d.people[pind].ChangeHP (hplus)
                  }
                  if c.role != -1 { d.oi[c.role].NoteCityRecover(cind, hplus) }
            }
      }

      if c.role != -1 {
            if c.treasurer != -1 && d.people[c.treasurer].lst != 0 {
                  d.oi[c.role].NoteLYTreasurer (c.treasurer, cind)
            }
            mplus := c.GetEarn (d)
            d.roles[c.role].ChangeMoney (mplus)
            d.oi[c.role].NoteCityEarn (cind, mplus)
      }

      // 重置风水
      if d.turn % FENGSHUI_ROUND == 0 {
            c.fengshui = rand.Float32() * 0.5
      }
}

func (d * Driver) CityOccupy (rind int, cind int) {
      c := &(d.cities[cind])
      if c.role != -1 {return}
      splist := d.roles[rind].mos.GetAllLabel()
      mos := d.oi[rind].SelCityMos(splist, cind)
      if len (mos) == 0 {
            d.oi[rind].OccuCityFail (cind)
      } else {
            c.role = rind
            d.roles[rind].mcs.SetLabel (cind, true)
            d.uv.NoteRoleOccupyCity (rind, cind)
            for _, i := range mos {
                  d.PInCity (splist[i], cind, false)
            }
            splist := c.mos.GetAllLabel()
            mayor := d.oi[rind].SelMayor (cind, splist, c.mayor, c.treasurer)
            if mayor != -1 {
                  d.PAsMayer(splist[mayor], true)
            }
            splist = c.mos.GetAllLabel()
            trea := d.oi[rind].SelTreasurer (cind, splist, c.mayor, c.treasurer)
            if trea != -1 {
                  d.PAsTreasurer(splist[trea], true)
            }
      }
}

func (d * Driver) CityDrop (rind int, cind int) {
      r := &(d.roles[rind])
      c := &(d.cities[cind])
      if c.role != rind {return}
      for _, pind := range c.mos.GetAllLabel() {
            d.POutCity (pind, false)
      }
      r.mcs.SetLabel (cind, false)
      c.role = -1
      d.uv.NoteRoleDropCity (rind, cind)
}
func (d * Driver) CityAction (rind int, cind int) {
      r := &(d.roles[rind])
      c := &(d.cities[cind])
      if c.role == rind {
            tax := c.GetEarn(d) * OWN_TAX_SCALE
            r.ChangeMoney (tax)
            d.uv.NoteCollectTax (rind, cind, tax)
            return
      }
      if c.role == -1 {
            if d.oi[rind].IsOccuCity(cind) { d.CityOccupy (rind, cind) }
      } else {
            tax := c.GetEarn(d) * TAX_SCALE
            r.ChangeMoney (- tax)
            d.uv.NotePayTax (rind, c.role, cind, tax)
            if r.ast == c.role {
                  d.oi[rind].SkipBattleByAlign (c.role)
                  return
            }
            bs := d.oi[rind].IsAttackCity(cind)
            if bs == -1 {
                  if r.ast == -2 {
                        d.oi[rind].ForceBattle (c.role)
                        bs = 0
                  }
                  if d.roles[c.role].ast == -2 {
                        d.oi[c.role].ForceBattle (rind)
                        bs = 0
                  }
            }
            if bs >= 0 {
                  bres := d.Battle (rind, -1, c.role, cind, bs)
                  if bres == 1 {
                        d.CityDrop (c.role, cind)
                        d.CityOccupy (rind, cind)
                  }
            }
      }
}

func (d * Driver) SaloonAction (rind int) {
      r := &(d.roles[rind])

      var pind []int
      var price []float32

      for k := 0; k < 3; k ++ {
            nl := d.lpset[k].GetNLabel()
            if nl == 0 {continue}
            i := d.lpset[k].GetAllLabel()[rand.Intn (nl)]
            pind = append (pind, i)
            price = append (price, d.people[i].GetPrice())
      }

      for {
            sel_ind := d.oi[rind].SaloonSelPeople (pind[:], price[:])
            if sel_ind == -1 {break}
            if r.money > price[sel_ind] {
                  r.ChangeMoney (- price[sel_ind])
                  d.PJoin(pind[sel_ind], rind, true)
                  break
            } else { d.oi[rind].RecruitFail (pind[sel_ind]) }
      }
}

func (d * Driver) InstAction (rind int, iind int) {
      r := &(d.roles[rind])
      inst := &(d.insts[iind])
      if inst.mos[rind] == -1 {
            var tl []int
            for _, tind := range inst.tech {
                  if !r.tech.GetLabel (tind) {
                        tl = append (tl, tind)
                  }
            }
            pl := r.mos.GetAllLabel()
            sp, st := d.oi[rind].SelStudy (iind, pl, tl)
            if sp != -1 && st != -1 {
                  d.PInInst (pl[sp], iind, tl[st], true)
            }
      }
}

func (d * Driver) TrainAction (rind int, tind int) {
      r := &(d.roles[rind])
      train := &(d.trains[tind])
      if train.mos[rind] == -1 {
            pl := r.mos.GetAllLabel()
            sp := d.oi[rind].SelTrain (tind, pl, train.item)
            if sp != -1 { d.PInTrain (pl[sp], tind, true) }
      }
}

func (d * Driver) ChanceAction (rind int) {
      r := &(d.roles[rind])
      nc := len (cards)
      sc_id := rand.Intn (nc)
      r.cards[sc_id] ++
      d.uv.NoteGetCard (rind, sc_id)
}

