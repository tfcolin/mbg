package mbg

import "fmt"

type TestUserView struct {
      d * Driver
}

func (uv * TestUserView) StartGame (d * Driver) { // 游戏开始
      uv.d = d
      fmt.Print ("Game Start\n")
}

func (uv * TestUserView) SetBarrier (loc int) {// 设置路障
      fmt.Printf ("Set Barrier at %d\n", loc)
}
func (uv * TestUserView) UnSetBarrier (loc int) {// 路障生效并去除
      fmt.Printf ("UnSet Barrier at %d\n", loc)
}
func (uv * TestUserView) SetRobber (loc int) { // 设置山贼
      fmt.Printf ("Set Robber at %d\n", loc)
}
func (uv * TestUserView) UnSetRobber (loc int){ // 山贼生效并去除
      fmt.Printf ("UnSet Robber at %d\n", loc)
}
func (uv * TestUserView) SetSlow (rind int) {// 慢行
      fmt.Printf ("Set Slow for Role %s\n", uv.d.roles[rind].name)
}
func (uv * TestUserView) SetFast (rind int) {// 急行
      fmt.Printf ("Set Fast for Role %s\n", uv.d.roles[rind].name)
}
func (uv * TestUserView) SetStop (rind int) {// 禁行
      fmt.Printf ("Set Stop for Role %s\n", uv.d.roles[rind].name)
}
func (uv * TestUserView) SetContinue (rind int) {// 双行
      fmt.Printf ("Set Continue for Role %s\n", uv.d.roles[rind].name)
}
func (uv * TestUserView) ChangeDir (rind int) {// 设置反向
      fmt.Printf ("change dir for Role %s\n", uv.d.roles[rind].name)
}
func (uv * TestUserView) StealMoney (sub_role int, obj_role int, money float32) {// 偷钱
      fmt.Printf ("Role %d steal Role %d money : %f\n", sub_role, obj_role, money)
}
func (uv * TestUserView) SelfRecover (rind int) {// 自我恢复
      fmt.Printf ("Role %d finish recovery\n", rind)
}
func (uv * TestUserView) SetLiuYan (pind int) { // 设置流言状态
      fmt.Printf ("People %d enter LiuYan state\n", pind)
}
func (uv * TestUserView) SetPoison (pind int) {// 设置中毒状态
      fmt.Printf ("People %d enter Poison state\n", pind)
}
func (uv * TestUserView) SetQuit (pind int) {// 设置离职状态
      fmt.Printf ("People %d prepare to leave %d\n", pind, uv.d.people[pind].role)
}
func (uv * TestUserView) SetAlign (sub_role int, obj_role int) {// 设置结盟
      fmt.Printf ("Role %d set align with Role %d\n", sub_role, obj_role)
}
func (uv * TestUserView) SetPE (role int) {// 设置公敌
      fmt.Printf ("Role %d declare war with all other roles\n", role)
}

func (uv * TestUserView)     InitMap (ngrid, nrole, npeople, ncity, ninst, ntrain, ntech, ncard, nx, ny int) {// 初始化地图
      fmt.Printf ("Init Map. Size = %d %d\n", nx, ny)
}
func (uv * TestUserView)     ShowMap (id int, class int, base int, x int, y int) {// 显示地图块
      fmt.Printf ("Set Map %d (%d %d) : class = %d, base = %d\n", id, x, y, class, base)
}
func (uv * TestUserView)     ShowCity (id int, name string, scope int, x0, y0, x1, y1 int) {// 显示城市 
      fmt.Printf ("Show City %d (%d %d %d %d): name = %s, scope = %d\n", id, x0, y0, x1, y1, name, scope)
}
func (uv * TestUserView)     ShowInstitute (id int, name string, x0, y0, x1, y1 int) {// 显示策略研究所
      fmt.Printf ("Show Institute %d (%d %d %d %d): name = %s\n", id, x0, y0, x1, y1, name)
}
func (uv * TestUserView)     ShowTrainingRoom (id int, name string, item Property, x0, y0, x1, y1 int) {// 显示训练所
      fmt.Printf ("Show TrainingRoom %d (%d %d %d %d): name = %s, Item = %d\n", id, x0, y0, x1, y1, name, item)
}

func (uv * TestUserView)     NoteInCity (pind int, new_city int) {// 某人进驻己方城市
      fmt.Printf ("People %d enter city %d for role %d\n", pind, new_city, uv.d.people[pind].role)
}
func (uv * TestUserView)     NoteInInst (pind int, new_inst int) {// 某人进驻研究所
      fmt.Printf ("People %d enter inst %d for role %d\n", pind, new_inst, uv.d.people[pind].role)
}
func (uv * TestUserView)     NoteInTrain (pind int, new_train int) {// 某人进驻训练所
      fmt.Printf ("People %d enter train %d for role %d\n", pind, new_train, uv.d.people[pind].role)
}
func (uv * TestUserView)     NoteOutCity (pind int, cind int) {// 某人撤离己方城市
      fmt.Printf ("People %d exit city %d\n", pind, cind)
}
func (uv * TestUserView)     NoteOutInst (pind int, iind int) {// 某人撤离研究所
      fmt.Printf ("People %d exit inst %d\n", pind, iind)
}
func (uv * TestUserView)     NoteOutTrain (pind int, tind int) {// 某人撤离训练所
      fmt.Printf ("People %d exit train %d\n", pind, tind)
}
func (uv * TestUserView)     NoteJoin (pind int, rind int) {// 某人加入某阵营
      fmt.Printf ("people %d join role %d\n", pind, rind)
}
func (uv * TestUserView)     NoteQuit (pind int, rind int) {// 某人离开某阵营
      fmt.Printf ("people %d quit role %d\n", pind, rind)
}
func (uv * TestUserView)     NotePayTax (send_role int, recv_role int, recv_city int, tax float32) {// 某人向某城支付租金
      fmt.Printf ("role %d pay role %d city %d tax: %f\n", send_role, recv_role, recv_city, tax)
}
func (uv * TestUserView)     NoteCollectTax (rind int, cind int, tax float32) {// 某角色到达自身城市收取税收
      fmt.Printf ("role %d collect tax from %d city: %f\n", rind, cind, tax)
}

func (uv * TestUserView)     RoleContinueAction (rind int) {// 某阵营连续行动 (在 RoleStartAction 之前调用)
      fmt.Printf ("role %d continue action.\n", rind)
}
func (uv * TestUserView)     RoleStartAction (rind int) {// 某阵营开始行动
      fmt.Printf ("role %d start action.\n", rind)
}
func (uv * TestUserView)     RoleForbidAction (rind int) {// 某阵营被禁止行动
      fmt.Printf ("role %d forbid action.\n", rind)
}

func (uv * TestUserView) NoteRoleOccupyCity (rind int, cind int) {// 某阵营占领城市
      fmt.Printf ("role %d occupy city %d.\n", rind, cind)
}

func (uv * TestUserView) NoteRoleDropCity  (rind int, cind int) {// 某阵营占领城市
      fmt.Printf ("role %d drop city %d.\n", rind, cind)
}

func (uv * TestUserView)     NoteRoleStEnd (rind int, st int) {// 阵营特殊行动状态终止
      fmt.Printf ("role %d end special action %d state.\n", rind, st)
}
func (uv * TestUserView)     NoteRoleCeEnd (rind int)         {// 阵营公敌状态终止
      fmt.Printf ("role %d end declare-war state.\n", rind)
}
func (uv * TestUserView)     NoteRoleAlignEnd (r1 int, r2 int) {// 阵营间结盟状态终止 (r1 < r2)
      fmt.Printf ("role %d end align relation with role %d.\n", r1, r2)
}

func      PrintMoney (d * Driver) {
      for i := 0; i < d.nrole; i ++ {
            fmt.Printf ("     role %d: money = %f\n", i, d.roles[i].money)
      }
      fmt.Printf ("\n")
}

func (uv * TestUserView)     NoteRoleWin (rind int) {// 阵营胜利
      fmt.Printf ("role %d win the game.\n", rind)
      PrintMoney (uv.d)
}

func (uv * TestUserView)     NoteRoleLoss (rind int) {// 阵营破产
      fmt.Printf ("role %d loss the game.\n", rind)
      PrintMoney (uv.d)
}

func (uv * TestUserView)     NoteAllLoss () {// 所有阵营同时破产 
      fmt.Printf ("all roles loss the game.\n")
      PrintMoney (uv.d)
}

func (uv * TestUserView) NoteForceQuit () {// 用户强制退出
      fmt.Printf ("user force quit.\n")
}

func (uv * TestUserView)     NoteRoleMoveOneStep (rind int, sloc int, dir int) {// 角色移动一步
      fmt.Printf ("role %d move from %d direct to %d\n", rind, sloc, dir)
}
func (uv * TestUserView)     NoteRoleMove (rind int, sloc int, oloc int, step int) {// 角色移动一个回合
      fmt.Printf ("role %d finish move: from %d to %d. step = %d\n", rind, sloc, oloc, step)
}

func (uv * TestUserView)     NoteGetCard (rind int, card_id int) {// 角色获得卡片
      fmt.Printf ("role %d get card %d\n", rind, card_id)
}

func (uv * TestUserView)     EndTurn (turn int) {// 回合结束
      fmt.Printf ("turn %d end\n", turn)
}

func (uv * TestUserView) BattleNoGerenal (rind int, side int) {// 无主将参与作战, 自动战败
      fmt.Printf ("Role %d at side %d select no general (surrender).\n", rind, side)
}

// sst: 发动方状态, ost: 被动方状态, max_bt: 最大回合数
func (uv * TestUserView)     BattleStart (sst * BattleState, ost * BattleState, max_bt int) {// 战斗开始
      fmt.Printf ("Battle Start:\n Max number of turns = %d\n", max_bt)
      fmt.Printf ("     Sub Role: %d\n", sst.role)
      fmt.Printf ("           power=%f winp=%f celve=%f hpmax=%f hp=%f def=%f\n",
      sst.power, sst.winp, sst.celve, sst.hpmax, sst.hp, sst.def)
      fmt.Printf ("     Obj Role: %d\n", ost.role)
      fmt.Printf ("           power=%f winp=%f celve=%f hpmax=%f hp=%f def=%f\n",
      ost.power, ost.winp, ost.celve, ost.hpmax, ost.hp, ost.def)
}

func (uv * TestUserView)     BattleTurnStart (iturn int) {// 战斗回合开始. iturn: 回合数 (0 起始)
      fmt.Printf ("BattleField: turn %d.\n", iturn)
}
func (uv * TestUserView)     BattleAttack (ls int, ost * BattleState) {
      fmt.Printf ("BattleField: side %d loss %f def and %f hp. hp = %f, def = %f\n", ls, ost.last_ddef, ost.last_dhp, ost.hp, ost.def)
}
func (uv * TestUserView)     BattleTech (ss int, tind int, sst * BattleState, ost * BattleState) {// 战斗时发动技能. ss: 0|1 发动方
      fmt.Printf ("BattleField: side %d launch %d tech:\n",  ss, tind)
      fmt.Printf ("     Sub side: dhp = %f, ddef = %f, hp = %f, hpmax = %f, def = %f\n", sst.last_dhp, sst.last_ddef, sst.hp, sst.hpmax, sst.def)
      fmt.Printf ("     Obj side: dhp = %f, ddef = %f, hp = %f, hpmax = %f, def = %f\n", ost.last_dhp, ost.last_ddef, ost.hp, ost.hpmax, ost.def)
}
func (uv * TestUserView)     BattleBurn (side int, st * BattleState) {// 战斗中被烧
      fmt.Printf ("side %d burned: dhp = %f, ddef = %f, hp = %f, hpmax = %f, def = %f\n", side, st.last_dhp, st.last_ddef, st.hp, st.hpmax, st.def)
}
func (uv * TestUserView)     BattleSelfAttack (side int, st * BattleState) {// 战斗中受到内讧伤害
      fmt.Printf ("side %d attack itself : dhp = %f, ddef = %f, hp = %f, hpmax = %f, def = %f\n", side, st.last_dhp, st.last_ddef, st.hp, st.hpmax, st.def)
}
func (uv * TestUserView)     BattleEnd  (ret int) {// 战斗结束: ret: 战斗结果. 0: 和平结束. -1: 攻方失败. 1: 攻方取胜.
      fmt.Printf ("Battle End: result = %d\n", ret)
}

type TestOperationInterface struct {
      d * Driver
      rind int
}

func (oi * TestOperationInterface)  StartGame (d * Driver, rind int) {// 游戏开始
      oi.d = d
      oi.rind = rind
      fmt.Printf ("game operational start for role %d\n", rind)
}

func (oi * TestOperationInterface) PreOperation () {
      fmt.Printf ("Role %d :", oi.rind)
}

func (oi * TestOperationInterface) Check () {
      d := oi.d

      var id int
      var class int

      sel: for {
            fmt.Printf ("Check type (0: role; 1: people; 2: city; 3: inst; 4: train, 5: map, 6: freepeople : exit; ):\n")
            fmt.Scan (&class)
            switch class {
            case 0:
                  fmt.Printf ("Role (0 - %d) id:\n", d.nrole - 1)
                  fmt.Scan (&id)
                  if id < 0 || id >= d.nrole {continue}
                  name , loc  , tech , mos  , mcs  , sos  , tos  , money, dir  , mst  , mtime, cst  , ast  , atime, cards := d.GetRoleInfo (id)
                  fmt.Printf ("    Name: %s \n", name)
                  fmt.Printf ("     Loc: %d \n", loc)
                  fmt.Printf ("    Tech: %v \n", tech)
                  fmt.Printf ("     mos: %v \n", mos)
                  fmt.Printf ("     mcs: %v \n", mcs)
                  fmt.Printf ("     sos: %v \n", sos)
                  fmt.Printf ("     tos: %v \n", tos)
                  fmt.Printf ("   money: %v \n", money)
                  fmt.Printf ("     dir: %v \n", dir)
                  fmt.Printf ("     mst: %v \n", mst)
                  fmt.Printf ("   mtime: %v \n", mtime)
                  fmt.Printf ("     cst: %v \n", cst)
                  fmt.Printf ("     ast: %v \n", ast)
                  fmt.Printf ("   atime: %v \n", atime)
                  fmt.Printf ("   cards: %v \n", cards)
            case 1:
                  fmt.Printf ("People (0 - %d) id:\n", d.npeople - 1)
                  fmt.Scan (&id)
                  if id < 0 || id >= d.npeople {continue}
                  name, role, job, loc, hpmax, hp, prop, is_quit, lst, pst := d.GetPeopleInfo(id)
                  fmt.Printf ("     name:      %v \n", name)
                  fmt.Printf ("     role:      %v \n", role)
                  fmt.Printf ("     job:       %v \n", job)
                  fmt.Printf ("     loc:       %v \n", loc)
                  fmt.Printf ("     hpmax:     %v \n", hpmax)
                  fmt.Printf ("     hp:        %v \n", hp)
                  fmt.Printf ("     prop:      %v \n", prop)
                  fmt.Printf ("     is_quit:   %v \n", is_quit)
                  fmt.Printf ("     lst:       %v \n", lst)
                  fmt.Printf ("     pst:       %v \n", pst)
            case 2:
                  fmt.Printf ("City (0 - %d) id:\n", d.ncity - 1)
                  fmt.Scan (&id)
                  if id < 0 || id >= d.ncity {continue}
                  name , scope , hpmax, hp , role , mos , mayor , treasurer , fengshui , hpplus , earn := d.GetCityInfo(id)
                  fmt.Printf ("   name:          %v \n", name)
                  fmt.Printf ("   scope:         %v \n", scope)
                  fmt.Printf ("   hpmax:         %v \n", hpmax)
                  fmt.Printf ("   hp:            %v \n", hp)
                  fmt.Printf ("   role:          %v \n", role)
                  fmt.Printf ("   mos:           %v \n", mos)
                  fmt.Printf ("   mayor:         %v \n", mayor)
                  fmt.Printf ("   treasurer:     %v \n", treasurer)
                  fmt.Printf ("   fengshui:      %v \n", fengshui)
                  fmt.Printf ("   hpplus:        %v \n", hpplus)
                  fmt.Printf ("   earn:          %v \n", earn)
            case 3:
                  fmt.Printf ("Inst (0 - %d) id:\n", d.ninst - 1)
                  fmt.Scan (&id)
                  if id < 0 || id >= d.ninst {continue}
                  name , tech , mos , on_study , point , round := d.GetInstInfo (id, oi.rind)
                  fmt.Printf ("   name:       %v \n",  name)
                  fmt.Printf ("   tech:       %v \n",  tech)
                  fmt.Printf ("   mos:        %v \n",  mos)
                  fmt.Printf ("   on_study:   %v \n",  on_study)
                  fmt.Printf ("   point:      %v \n",  point)
                  fmt.Printf ("   round:      %v \n",  round)
            case 4:
                  fmt.Printf ("Train (0 - %d) id:\n", d.ntrain - 1)
                  fmt.Scan (&id)
                  if id < 0 || id >= d.ntrain {continue}
                  name, item, mos, round := d.GetTrainInfo (id, oi.rind)
                  fmt.Printf ("   name,   %v \n",  name)
                  fmt.Printf ("   item,   %v \n",  item)
                  fmt.Printf ("   mos,    %v \n",  mos)
                  fmt.Printf ("   round,  %v \n",  round)
            case 5:
                  fmt.Printf ("Map: \n")
                  fmt.Printf ("  Class: ")
                  for i := 0; i < d.ngrid; i ++ {
                        fmt.Printf ("%v ", d.board[i].class)
                  }
                  fmt.Printf ("\n")

                  fmt.Printf ("   Base: ")
                  for i := 0; i < d.ngrid; i ++ {
                        fmt.Printf ("%v ", d.board[i].base)
                  }
                  fmt.Printf ("\n")

                  fmt.Printf ("  Roles: \n")
                  for j := 0; j < d.nrole; j ++ {
                        fmt.Printf ("      Id %d: ", j)
                        for i := 0; i < d.ngrid; i ++ {
                              fmt.Printf ("%v ", d.board[i].roles[j])
                        }
                        fmt.Printf ("\n")
                  }

                  fmt.Printf ("Barrier: \n")
                  for i := 0; i < d.ngrid; i ++ {
                        fmt.Printf ("%v ", d.board[i].barrier)
                  }
                  fmt.Printf ("\n")

                  fmt.Printf (" Robber: \n")
                  for i := 0; i < d.ngrid; i ++ {
                        fmt.Printf ("%v ", d.board[i].robber)
                  }
                  fmt.Printf ("\n")
            case 6:
                  fmt.Printf ("Free People: \n")
                  for i := 0; i < 3; i ++ {
                        fmt.Printf ("    Level %d: %v", i, d.GetFreePeopleList(i))
                  }
                  fmt.Printf ("\n")
            default:
                  break sel
            }
      }
}

func (oi * TestOperationInterface) AlignFail (obj_role int) {// 设置结盟失败
      oi.PreOperation()
      fmt.Printf ("set align fail\n")
}
func (oi * TestOperationInterface)     PEFail (rind int) {// 设置公敌失败 
      oi.PreOperation()
      fmt.Printf ("declare war fail\n")
}

func (oi * TestOperationInterface)     NoteAsMayer    (pind int) {// 某人就职市长
      oi.PreOperation()
      fmt.Printf ("people %d as mayer\n", pind)
}
func (oi * TestOperationInterface)     NoteAsTreasurer (pind int) {// 某人就职财务官
      oi.PreOperation()
      fmt.Printf ("people %d as treasurer\n", pind)
}
func (oi * TestOperationInterface)     NoteAsCitizen  (pind int) {// 某人撤销市政职务
      oi.PreOperation()
      fmt.Printf ("people %d as normal citizen\n", pind)
}

func (oi * TestOperationInterface)     NoteFinishStudy (iind int, pind int, tech int) {// 研发完成
      oi.PreOperation()
      fmt.Printf ("people %d in inst %d finish study. learn tech %d for role %d\n",
      pind, iind, tech, oi.d.people[pind].role)
}
func (oi * TestOperationInterface)     NoteLYStudy (iind int, pind int) {// 研发因流言无法继续
      oi.PreOperation()
      fmt.Printf ("people %d in inst %d cannot study due to LiuYan.\n", pind, iind)
}
func (oi * TestOperationInterface)     NoteFinishTrain (pind int, item Property, value float32) {// 训练完成
      oi.PreOperation()
      fmt.Printf ("people %d finish train %d, value = %f\n", pind, item, value)
}
func (oi * TestOperationInterface)     NoteLYTrain (tind int, pind int) {// 训练因流言无法继续
      oi.PreOperation()
      fmt.Printf ("people %d in training_room %d cannot train due to LiuYan.\n", pind, tind)
}
func (oi * TestOperationInterface)     NoteLYMayor (pind, cind int) {// 因流言无法履行市长职责
      oi.PreOperation()
      fmt.Printf ("people %d in city %d cannot adequate to mayor due to LiuYan.\n", pind, cind)
}
func (oi * TestOperationInterface)     NoteLYTreasurer (pind, cind int) {// 因流言无法履行财务官职责
      oi.PreOperation()
      fmt.Printf ("people %d in city %d cannot adequate to treasurer due to LiuYan.\n", pind, cind)
}
func (oi * TestOperationInterface)     NoteCityRecover (cind int, hplus float32) {// 城市因市长恢复所属人员 HP
      oi.PreOperation()
      fmt.Printf ("city %d recover because of mayor. hplus = %f \n", cind, hplus)
}
func (oi * TestOperationInterface)     NoteCityEarn (cind int, mplus float32) {// 城市获得收入
      oi.PreOperation()
      fmt.Printf ("city %d earn money for role %d because of treasurer. earn = %f \n", cind, oi.d.cities[cind].role, mplus)
}

func (oi * TestOperationInterface) NotePayInst (dm float32) {// 支付研究经费
      fmt.Printf ("role %d pay for study %f.\n", oi.rind, -dm)
}

func (oi * TestOperationInterface)     SkipBattleByAlign (robj int) {// 因为结盟跳过战斗
      oi.PreOperation()
      fmt.Printf ("role %d skip battle by align\n", robj)
}
func (oi * TestOperationInterface)     ForceBattle (robj int) {// 因为公愤进行强制战斗
      oi.PreOperation()
      fmt.Printf ("role %d force battle by war-declarence\n", robj)
}

func (oi * TestOperationInterface)     SelRoleAction (clist []int) int {// 选择阵营行动: -1: 移动, >=0: 使用卡片
      oi.PreOperation()
      var action int
      for {
            fmt.Printf ("Select action:\n")
            fmt.Scan (&action)
            if action == -2 {
                  oi.Check()
                  continue
            } else { break }
      }
      if action < -1 {action = -1}
      return action
}

func (oi * TestOperationInterface)     IsOccuCity (cind int) bool {// 选择是否占领空白城市
      oi.PreOperation()
      var res bool
      fmt.Printf ("Is occupy empty city %d:\n", cind)
      fmt.Scan (&res)
      return res
}

func (oi * TestOperationInterface)     OccuCityFail (cind int) {// 占领城市失败
      oi.PreOperation()
      fmt.Printf ("occupy city fail\n", cind)
}

func (oi * TestOperationInterface) SelListLessThan (max int) []int {
      var ind int
      var selres []int
      for {
            fmt.Scan (&ind)
            if ind == -2 {
                  oi.Check()
                  continue
            }
            if ind == -1 {break}

            if ind < 0 || ind >= max {continue}

            selres = append(selres, ind)
      }
      return selres
}

func (oi * TestOperationInterface) SelOneLessThan (max int) int {
      var res int
      for {
            fmt.Scan (&res)
            if res == -2 {
                  oi.Check()
                  continue
            }
            if res < 0 || res >= max {res = -1}
            break
      }
      return res
}

func (oi * TestOperationInterface)     SelCityMos (splist []int, cind int) []int {// 选择占领的人员
      oi.PreOperation()
      fmt.Printf ("select people to occupy city %d:\n", cind)
      fmt.Printf ("(select list: %v)\n", splist)
      return oi.SelListLessThan (len(splist))
}

func (oi * TestOperationInterface)     SelMayor (cind int, splist []int, old_mayor int, old_treasurer int) (mayor int) {// 选择市长
      fmt.Printf ("select mayor for %d (-1: cancel):\n", cind)
      fmt.Printf ("List: %v\n", splist)
      fmt.Printf ("current mayor: %v; current treasurer %v\n", old_mayor, old_treasurer)
      return oi.SelOneLessThan(len(splist))
}

func (oi * TestOperationInterface)     SelTreasurer (cind int, splist []int, old_mayor int, old_treasurer int) (treasurer int) {// 选择财务官
      fmt.Printf ("select treasurer for %d (-1: cancel):\n", cind)
      fmt.Printf ("List: %v\n", splist)
      fmt.Printf ("current mayor: %v; current treasurer %v\n", old_mayor, old_treasurer)
      return oi.SelOneLessThan(len(splist))
}

func (oi * TestOperationInterface)     IsAttackCity (cind int) int {// 选择是否攻城 -1: 不攻城 0-2: 三种攻城规模
      oi.PreOperation()
      fmt.Printf ("Is attack city? 0-2: attack scale \n")
      return oi.SelOneLessThan (3)
}

func (oi * TestOperationInterface)     SaloonSelPeople (pind []int , price []float32) int {// 选择招募人员 : -1: 放弃
      oi.PreOperation()
      n := len (pind)
      fmt.Printf ("select people to join you:\n")
      fmt.Printf ("People list: index price \n")
      for i := 0; i < n; i ++ {
            fmt.Printf ("%8d: %10.2d%10.2f\n", i, pind[i], price[i])
      }
      return oi.SelOneLessThan (n)
}

func (oi * TestOperationInterface)     RecruitFail (pind int) {// 因无法支付招募失败
      oi.PreOperation()
      fmt.Printf ("Not enough money. People join fail.\n")
}

func (oi * TestOperationInterface)     SelStudy (iind int, plist []int, tlist []int) (sp int, st int) {// 选择研究人员和项目, 返回 plist 和 tlist 中所选择的编号(不是索引), (-1, -1): 放弃研究
      oi.PreOperation()
      fmt.Printf ("enter inst %d", iind)
      fmt.Printf ("select people to study: %v\n", plist)
      sp = oi.SelOneLessThan (len(plist))
      fmt.Printf ("select tech to study: %v\n", tlist)
      st = oi.SelOneLessThan (len(tlist))
      if sp == -1 || st == -1 {
            sp = -1
            st = -1
      }
      return
}

func (oi * TestOperationInterface)     SelTrain (tind int, plist []int, item Property) int {// 选择训练人员, 返回人员在 plist 中的索引号, -1: 放弃
      oi.PreOperation()
      fmt.Printf ("enter train %d", tind)
      fmt.Printf ("select people to train %d.:\n", item)
      fmt.Printf ("(%v)\n", plist)
      return oi.SelOneLessThan (len(plist))
}

func (oi * TestOperationInterface)     BattleWithRole (obj_role int) int {// 是否打遭遇战. 返回: -1: 不打. 0-2: 战斗规模
      oi.PreOperation()
      fmt.Printf ("Battle with role %d ? (0-2: scale. -1: do not attack)\n", obj_role)
      return oi.SelOneLessThan (3)
}

func (oi * TestOperationInterface)     ConfirmCard (card_id int) int {// 确认是否使用卡片: 无对象卡片 返回: 0: 确定; -1:取消
      oi.PreOperation()
      fmt.Printf ("Confirm use card %d (no obj): 0: yes; other: no\n", card_id)
      return oi.SelOneLessThan (1)
}

func (oi * TestOperationInterface)     SelCardObjCity (card_id int) int {// 返回卡片作用对象城市. -1 取消
      oi.PreOperation()
      fmt.Printf ("select the object city of card %d\n", card_id)
      return oi.SelOneLessThan (oi.d.ncity)
}

func (oi * TestOperationInterface)     SelCardObjAny (card_id int) (st int, obj int) {// 返回卡片作用对象建筑或角色. st: -1: 取消. 0: 角色 1: 城市: 2: 研究所 3: 修炼所
      oi.PreOperation()
      fmt.Printf ("select card %d object class: (-1: cancel; 0: role; 1: city; 2: inst; 3:trainingroom) \n", card_id)
      fmt.Scan (&st)
      if st < 0 || st > 3 {st = -1}
      if st == -1 {return}

      fmt.Printf ("select card %d object: \n", card_id)
      fmt.Scan (&obj)

      if obj < 0 {obj = -1}
      switch st {
      case 0:
            if obj >= oi.d.nrole {obj = -1}
      case 1:
            if obj >= oi.d.ncity {obj = -1}
      case 2:
            if obj >= oi.d.ninst {obj = -1}
      case 3:
            if obj >= oi.d.ntrain {obj = -1}
      }

      return
}

func (oi * TestOperationInterface)     SelCardObjPeople (card_id int, plist []int) int {// 返回卡片作用人员. -1: 取消
      oi.PreOperation()
      fmt.Printf ("select person as card %d object:\n", card_id)
      fmt.Printf ("List: %v\n", plist)
      return oi.SelOneLessThan (len(plist))
}

func (oi * TestOperationInterface)     SelCardObjRoleAndPeople (card_id int, rplist []int) (sr_id int) {// 返回卡片作用人员 (每阵营一人供选择) -1:取消
      oi.PreOperation()
      fmt.Printf ("select person as card %d object: ", card_id)
      for _, i := range rplist {
            fmt.Printf (" %d ", i)
      }
      fmt.Printf ("\n")
      fmt.Scan (&sr_id)
      if sr_id < 0 || sr_id >= oi.d.nrole {sr_id = -1}
      return
}

func (oi * TestOperationInterface)     SelCardObjRole (card_id int) int {// 返回选择卡片作用阵营 -1: 取消
      oi.PreOperation()
      fmt.Printf ("select object role of card %d\n", card_id)
      return oi.SelOneLessThan (oi.d.nrole)
}

func (oi * TestOperationInterface)     SelCardObjLoc (card_id int) int {// 选择卡片作用地图位置 -1: 取消
      oi.PreOperation()
      fmt.Printf ("select object location of card %d\n", card_id)
      return oi.SelOneLessThan (oi.d.ngrid)
}

func (oi * TestOperationInterface)     SelAllocObj () (st int, obj int) {// 选择调度一级目标. st: -1: 取消. 0: 城市: 1: 研究所 2: 修炼所
      oi.PreOperation()
      fmt.Printf ("select allocate object: (-1: cancel, 0: city; 1: inst; 2: train\n")
      st = oi.SelOneLessThan (3)
      switch st {
      case 0:
            obj = oi.SelOneLessThan (oi.d.ncity)
      case 1:
            obj = oi.SelOneLessThan (oi.d.ninst)
      case 2:
            obj = oi.SelOneLessThan (oi.d.ntrain)
      default:
            obj = -1
      }
      return
}

func (oi * TestOperationInterface) ExchangeCityPeople (cind int, splist []int, oplist []int) (in_list []int, out_list []int) {// 调度随身人员与目标内驻扎人员. 
      oi.PreOperation()
      fmt.Printf ("exchange people for role %d and city %d.\n", oi.rind, cind)
      fmt.Printf ("Role people: %v\n", splist)
      fmt.Printf ("City people: %v\n", oplist)
      in_list  = oi.SelListLessThan (len(splist))
      out_list = oi.SelListLessThan (len(oplist))

      return
}

func (oi * TestOperationInterface)     IsCancelStudy (pind int, on_study int, left_point float32) bool {// 是否终止研究
      oi.PreOperation()
      fmt.Printf ("If cancel study: people = %d, tech = %d, left_point = %f?\n", pind, on_study, left_point)
      var res bool
      fmt.Scan (&res)
      return res
}

func (oi * TestOperationInterface)     IsCancelTrain (pind int, item Property, left_round int) bool {// 是否终止训练
      oi.PreOperation()
      fmt.Printf ("If cancel training: people = %d, item = %d, left_round = %d?\n", pind, item, left_round)
      var res bool
      fmt.Scan (&res)
      return res
}
func (oi * TestOperationInterface)     StartAllocate () {// 开始人员调度
      oi.PreOperation()
      fmt.Printf ("Start allocating.\n")
}

func (oi * TestOperationInterface)     SelGeneral (plist []int) (main, vice int) {// 选择战争将领: main: 主将. vice: 副将 -1: 无副将
      oi.PreOperation()
      fmt.Printf ("Select generals for the battle.\n")
      fmt.Printf ("List: %v\n", plist)
      main = oi.SelOneLessThan (len(plist))
      vice = oi.SelOneLessThan (len(plist))
      return
}
