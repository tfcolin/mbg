package mbg

// res: 0: 正常; 1: 减为零; 2: 加满
func (o * Officer) ChangeHP (dhp float32) (real_dhp float32, res int) {
      php := o.hp + dhp
      if php < 1 {
            real_dhp = - o.hp
            res = 1
      } else if php >= o.hpmax {
            real_dhp = o.hpmax - o.hp
            res = 2
      } else {
            real_dhp = dhp
            res = 0
      }
      o.hp += real_dhp
      if o.hp < 1 {o.hp = 0}
      return
}

func (d * Driver) PJoin (pind int, rind int, is_note bool) bool {
      p := &(d.people[pind])
      if p.role != -1 {return false}
      p.role = rind
      d.roles[rind].mos.SetLabel (pind, true)
      d.lpset[p.level].SetLabel (pind, false)
      if is_note {
            d.uv.NoteJoin (pind, rind)
      }
      return true
}

func (d * Driver) PQuit (pind int, is_note bool) bool {
      p := &(d.people[pind])
      if p.role == -1 {
            return false
      }

      switch p.job {
      case 0:
            d.PAsCitizen (pind, is_note)
            d.POutCity (pind, is_note)
      case 1:
            d.PAsCitizen (pind, is_note)
            d.POutCity (pind, is_note)
      case 2:
            d.POutCity (pind, is_note)
      case 3:
            d.POutInst (pind, is_note)
      case 4:
            d.POutTrain (pind, is_note)
      }

      rind := p.role
      p.role = -1
      d.roles[rind].mos.SetLabel (pind, false)
      d.lpset[p.level].SetLabel (pind, true)

      if is_note {
            d.uv.NoteQuit (pind, rind)
      }

      return true
}

func (d * Driver) PInCity (pind int, new_city int, is_note bool) bool {
      p := &(d.people[pind])
      if p.role == -1 || p.loc != -1 {
            return false
      }
      c := &(d.cities[new_city])
      if p.role != c.role {
            return false
      }
      p.job = 2
      p.loc = new_city
      c.mos.SetLabel (pind, true)
      r := &(d.roles[p.role])
      r.mos.SetLabel (pind, false)

      if (is_note) {
            d.uv.NoteInCity (pind, new_city)
      }
      return true
}

func (d * Driver) PInInst (pind int, new_inst int, techid int, is_note bool) bool {
      p := &(d.people[pind])
      if p.role == -1 || p.loc != -1 {
            return false
      }
      if d.roles[p.role].tech.GetLabel(techid) {
            return false
      }
      inst := &(d.insts[new_inst])
      if inst.mos[p.role] != -1 {
            return false
      }

      p.job = 3
      p.loc = new_inst
      inst.mos[p.role] = pind
      inst.on_study[p.role] = techid
      inst.point[p.role] = techs[techid].study

      r := &(d.roles[p.role])
      r.sos.SetLabel (pind, true)
      r.mos.SetLabel (pind, false)

      if (is_note) {
            d.uv.NoteInInst (pind, new_inst)
      }
      return true
}

func (d * Driver) PInTrain (pind int, new_train int, is_note bool) bool {
      p := &(d.people[pind])
      if p.role == -1 || p.loc != -1 {
            return false
      }
      train := &(d.trains[new_train])
      if train.mos[p.role] != -1 {
            return false
      }

      p.job = 4
      p.loc = new_train
      train.mos[p.role] = pind
      train.round[p.role] = TRAIN_ROUND

      r := &(d.roles[p.role])
      r.tos.SetLabel (pind, true)
      r.mos.SetLabel (pind, false)

      if (is_note) {
            d.uv.NoteInTrain (pind, new_train)
      }
      return true
}

func (d * Driver) POutCity (pind int, is_note bool) bool {
      p := &(d.people[pind])
      if p.loc == -1 {
            return false
      }
      if p.job < 0 || p.job > 2 {
            return false
      }
      if p.job == 0 || p.job == 1 {
            d.PAsCitizen(pind, false)
      }
      c := &(d.cities[p.loc])
      cind := p.loc

      p.job = -1
      p.loc = -1
      c.mos.SetLabel (pind, false)

      if c.role != -1 {
            r := &(d.roles[c.role])
            r.mos.SetLabel (pind, true)
      }

      if (is_note) {
            d.uv.NoteOutCity (pind, cind)
      }

      if (c.mos.IsEmpty()) {
            rind := c.role
            c.role = -1
            if rind != -1 {
                  d.roles[rind].mcs.SetLabel (cind, false)
                  if (is_note) {
                        d.uv.NoteRoleDropCity (rind, cind)
                  }
            }
      }
      return true
}

func (d * Driver) POutInst (pind int, is_note bool) bool {
      p := &(d.people[pind])
      if p.role == -1 || p.job != 3 {
            return false
      }
      inst := &(d.insts[p.loc])
      iind := p.loc

      p.job = -1
      p.loc = -1
      inst.mos[p.role] = -1

      r := &(d.roles[p.role])
      r.sos.SetLabel (pind, false)
      r.mos.SetLabel (pind, true)

      if is_note {
            d.uv.NoteOutInst (pind, iind)
      }
      return true
}

func (d * Driver) POutTrain (pind int, is_note bool) bool {
      p := &(d.people[pind])
      if p.role == -1 || p.job != 4 {
            return false
      }
      train := &(d.trains[p.loc])
      tind := p.loc

      p.job = -1
      p.loc = -1
      train.mos[p.role] = -1

      r := &(d.roles[p.role])
      r.tos.SetLabel (pind, false)
      r.mos.SetLabel (pind, true)

      if is_note {
            d.uv.NoteOutTrain (pind, tind)
      }
      return true
}

func (d * Driver) PAsMayer (pind int, is_note bool) bool {
      o := &(d.people[pind])
      if o.job < 1 || o.job > 2 {
            return false
      }
      c := &(d.cities[o.loc])
      if c.mayor != -1 {
            d.people[c.mayor].job = 2
      }
      c.mayor = pind
      if o.job == 1 {
            c.treasurer = -1
      }
      o.job = 0
      if is_note && o.role != -1 {
            d.oi[o.role].NoteAsMayer (pind)
      }
      return true
}

func (d * Driver) PAsTreasurer (pind int, is_note bool) bool {
      o := &(d.people[pind])
      if o.job != 0 && o.job != 2 {
            if is_note {
                  return false
            }
      }
      c := &(d.cities[o.loc])
      if c.treasurer != -1 {
            d.people[c.treasurer].job = 2
      }
      c.treasurer = pind
      if o.job == 0 {
            c.mayor = -1
      }
      o.job = 1
      if is_note && o.role != -1 {
            d.oi[o.role].NoteAsTreasurer (pind)
      }
      return true
}

func (d * Driver) PAsCitizen (pind int, is_note bool) bool {
      o := &(d.people[pind])
      if o.job < 0 || o.job > 1 {
            return false
      }
      c := &(d.cities[o.loc])
      if o.job == 0 {
            c.mayor = -1
      }
      if o.job == 1 {
            c.treasurer = -1
      }
      o.job = 2
      if is_note && o.role != -1 {
            d.oi[o.role].NoteAsCitizen (pind)
      }
      return true
}

func (d * Driver) RoundPeople (pind int) {
      p := &(d.people[pind])
      if p.lst > 0 {
            p.lst --
      }
      if p.pst > 0 {
            p.pst --
      }
      if p.is_quit {
            p.is_quit = false
            d.PQuit (pind, true)
      }
}
