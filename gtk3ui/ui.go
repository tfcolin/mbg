package gtk3ui

import (
      "fmt"
	"gitee.com/tfcolin/mbg"
)

func (uv *GtkUserView) StartGame (d *mbg.Driver) { // 游戏开始
	uv.d = d
	_, nrole, _, _, _, _, _, _, _ := d.GetInfo()

	for i := 0; i < mbg.TECH_COUNT; i++ {
		tname, _, scond, _, _, _, _, _ := mbg.GetTechInfo(i)
		init_tech(i, tname, scond)
	}
	for i := 0; i < mbg.CARD_COUNT; i++ {
		dname, odir, otype := mbg.GetCardInfo(i)
		init_card(i, dname, odir, otype)
	}

	for i := 0; i < nrole; i++ {
		rname, loc, _, _, _, _, _, _, _, _, _, _, _, _, _ := d.GetRoleInfo(i)
		init_role(i, rname)
		if loc != -1 {
			set_role (loc, i, false)
		}
	}
	finish_draw_map()
}

func (uv *GtkUserView) SetBarrier(loc int) { // 设置路障
	set_barrier(loc, false)
	update_role_point(loc)
}
func (uv *GtkUserView) UnSetBarrier(loc int) { // 路障生效并去除
	set_barrier(loc, true)
	update_role_point(loc)
}
func (uv *GtkUserView) SetRobber(loc int) { // 设置山贼
	set_robber(loc, false)
	update_role_point(loc)
}
func (uv *GtkUserView) UnSetRobber(loc int) { // 山贼生效并去除
	set_robber(loc, true)
	update_role_point(loc)
}

func (uv *GtkUserView) SetSlow(rind int) { // 慢行
	rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(rind)
	msg_map(fmt.Sprintf("角色 %s 速度减慢", rname))
}

func (uv *GtkUserView) SetFast(rind int) { // 急行
	rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(rind)
	msg_map (fmt.Sprintf("角色 %s 速度加快", rname))
}

func (uv *GtkUserView) SetStop(rind int) { // 禁行
	rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(rind)
	msg_map (fmt.Sprintf("角色 %s 禁止行动", rname))
}
func (uv *GtkUserView) SetContinue(rind int) { // 双行
	rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(rind)
	msg_map(fmt.Sprintf("角色 %s 连续行动", rname))
}

func (uv *GtkUserView) ChangeDir(rind int) { // 设置反向
	rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(rind)
	msg_map(fmt.Sprintf("角色 %s 掉头行动", rname))
}
func (uv *GtkUserView) StealMoney(sub_role int, obj_role int, money float32) { // 偷钱
	var msg, sub_name, obj_name string
	if sub_role == -1 {
		sub_name = "强盗"
	} else {
		sub_name, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = uv.d.GetRoleInfo(sub_role)
	}
	if obj_role == -1 {
		obj_name =  "强盗"
	} else {
		obj_name, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = uv.d.GetRoleInfo(obj_role)
	}
	msg = fmt.Sprintf("角色 %s 从角色 %s 手中偷去 %f 金币", sub_name, obj_name, money)
	msg_map(msg)
}

func (uv *GtkUserView) SelfRecover(rind int) { // 自我恢复
	rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(rind)
	msg_map(fmt.Sprintf("角色 %s 的官员得到恢复", rname))
}
func (uv *GtkUserView) SetLiuYan(pind int) { // 设置流言状态
	pname, _, _, _, _, _, _, _, _, _ := uv.d.GetPeopleInfo(pind)
	msg_map(fmt.Sprintf("官员 %s 进入流言状态, 暂时无法工作", pname))
}
func (uv *GtkUserView) SetPoison(pind int) { // 设置中毒状态
	pname, _, _, _, _, _, _, _, _, _ := uv.d.GetPeopleInfo(pind)
	msg_map(fmt.Sprintf("官员 %s 进入中毒状态, 能力属性下降", pname))
}
func (uv *GtkUserView) SetQuit(pind int) { // 设置离职状态
	pname, rind, _, _, _, _, _, _, _, _ := uv.d.GetPeopleInfo(pind)
	rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(rind)
	msg_map(fmt.Sprintf("官员 %s 接受劝说, 准备离开 %s", pname, rname))
}
func (uv *GtkUserView) SetAlign(sub_role int, obj_role int) { // 设置结盟
	sub_name, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(sub_role)
	obj_name, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(obj_role)
	msg_map(fmt.Sprintf("%s 与 %s 结盟", sub_name, obj_name))
}
func (uv *GtkUserView) SetPE(role int) { // 设置公敌
	name, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(role)
	msg_map(fmt.Sprintf("%s 阵营与所有其他阵营宣战", name))
}

func (uv *GtkUserView) InitMap(ngrid, nrole, npeople, ncity, ninst, ntrain, ntech, ncard, nx, ny int) { // 初始化地图
      init_map (ngrid, nrole, npeople, ncity, ninst, ntrain, ntech, ncard)
}

func (uv *GtkUserView) ShowMap(id int, class int, base int, x int, y int) { // 显示地图块
	set_board(id, class, base, x, y)
}

func (uv *GtkUserView) ShowCity(id int, name string, scope int, x0, y0, x1, y1 int) { // 显示城市
	set_city(id, name)
}
func (uv *GtkUserView) ShowInstitute(id int, name string, x0, y0, x1, y1 int) { // 显示策略研究所
	set_inst(id, name)
}
func (uv *GtkUserView) ShowTrainingRoom(id int, name string, item mbg.Property, x0, y0, x1, y1 int) { // 显示训练所
	set_train(id, name)
}

func (uv *GtkUserView) NoteInCity(pind int, new_city int) { // 某人进驻己方城市
	var msg string
	cname, _, _, _, _, _, _, _, _, _, _ := uv.d.GetCityInfo(new_city)
	pname, rind, _, _, _, _, _, _, _, _ := uv.d.GetPeopleInfo(pind)
	rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(rind)
	if rind != -1 {
		msg = fmt.Sprintf("官员 %s 进驻由 %s 占领的城市 %s", pname, rname, cname)
	} else {
		msg = fmt.Sprintf("官员 %s 进驻城市 %s", pname, cname)
	}
	msg_map(msg)
}

func (uv *GtkUserView) NoteInInst(pind int, new_inst int) { // 某人进驻研究所
	pname, rind, _, _, _, _, _, _, _, _ := uv.d.GetPeopleInfo(pind)
	if rind != -1 {
		iname, _, _, on_study, _, _ := uv.d.GetInstInfo(new_inst, rind)
		rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(rind)
		hname, _, _, _, _, _, _, _ := mbg.GetTechInfo(on_study)
		msg_map(fmt.Sprintf("官员 %s 进入 %s 为 %s 研究策略 %s", pname, iname, rname, hname))
	}
}
func (uv *GtkUserView) NoteInTrain(pind int, new_train int) { // 某人进驻训练所
	pname, role, _, _, _, _, _, _, _, _ := uv.d.GetPeopleInfo(pind)
	tname, _, _, _ := uv.d.GetTrainInfo(new_train, role)
	msg_map(fmt.Sprintf("官员 %s 进入 %s 进行训练", pname, tname))
}
func (uv *GtkUserView) NoteOutCity(pind int, cind int) { // 某人撤离己方城市
	pname, _, _, _, _, _, _, _, _, _ := uv.d.GetPeopleInfo(pind)
	cname, _, _, _, _, _, _, _, _, _, _ := uv.d.GetCityInfo(cind)
	msg_map(fmt.Sprintf("官员 %s 离开城市 %s", pname, cname))
}
func (uv *GtkUserView) NoteOutInst(pind int, iind int) { // 某人撤离研究所
	pname, _, _, _, _, _, _, _, _, _ := uv.d.GetPeopleInfo(pind)
	iname, _, _, _, _, _ := uv.d.GetInstInfo(iind, 0)
	msg_map(fmt.Sprintf("官员 %s 离开 %s", pname, iname))
}
func (uv *GtkUserView) NoteOutTrain(pind int, tind int) { // 某人撤离训练所
	pname, _, _, _, _, _, _, _, _, _ := uv.d.GetPeopleInfo(pind)
	tname, _, _, _ := uv.d.GetTrainInfo(tind, 0)
	msg_map(fmt.Sprintf("官员 %s 离开 %s", pname, tname))
}
func (uv *GtkUserView) NoteJoin(pind int, rind int) { // 某人加入某阵营
	pname, _, _, _, _, _, _, _, _, _ := uv.d.GetPeopleInfo(pind)
	rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(rind)
	msg_map(fmt.Sprintf("官员 %s 加入 %s 阵营", pname, rname))
}
func (uv *GtkUserView) NoteQuit(pind int, rind int) { // 某人离开某阵营
	pname, _, _, _, _, _, _, _, _, _ := uv.d.GetPeopleInfo(pind)
	rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(rind)
	msg_map(fmt.Sprintf("官员 %s 离开 %s 阵营", pname, rname))
}
func (uv *GtkUserView) NotePayTax(send_role int, recv_role int, recv_city int, tax float32) { // 某人向某城支付租金
	send_name, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(send_role)
	recv_name, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(recv_role)
	cname, _, _, _, _, _, _, _, _, _, _ := uv.d.GetCityInfo(recv_city)
	msg_map(fmt.Sprintf("角色 %s 途径城市 %s, 向角色 %s 支付税金 %.0f", send_name, cname, recv_name, tax))
}
func (uv *GtkUserView) NoteCollectTax(rind int, cind int, tax float32) { // 某角色到达自身城市收取税收
	rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(rind)
	cname, _, _, _, _, _, _, _, _, _, _ := uv.d.GetCityInfo(cind)
	msg_map(fmt.Sprintf("角色 %s 途径已占领城市 %s, 额外收取税金 %.0f", rname, cname, tax))
}

func (uv *GtkUserView) RoleContinueAction(rind int) { // 某阵营连续行动 (在 RoleStartAction 之前调用)
	rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(rind)
	msg_map(fmt.Sprintf("角色 %s 连续行动", rname))
}
func (uv *GtkUserView) RoleStartAction(rind int) { // 某阵营开始行动
	rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(rind)
	msg_map(fmt.Sprintf("角色 %s 开始行动", rname))
}
func (uv *GtkUserView) RoleForbidAction(rind int) { // 某阵营被禁止行动
	name, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(rind)
	msg_map(fmt.Sprintf("角色 %s 本回合无法行动", name))
}

func (uv *GtkUserView) NoteRoleOccupyCity(rind int, cind int) { // 某阵营占领城市
	rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(rind)
	cname, _, _, _, _, _, _, _, _, _, _ := uv.d.GetCityInfo(cind)
	msg_map(fmt.Sprintf("%s 阵营占领城市 %s", rname, cname))
}

func (uv *GtkUserView) NoteRoleDropCity(rind int, cind int) { // 某阵营占领城市
	rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(rind)
	cname, _, _, _, _, _, _, _, _, _, _ := uv.d.GetCityInfo(cind)
	msg_map(fmt.Sprintf("%s 阵营放弃城市 %s", rname, cname))
}

func (uv *GtkUserView) NoteRoleStEnd(rind int, st int) { // 阵营特殊行动状态终止
	rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(rind)
	var stname string
	switch st {
	case 1:
		stname = "慢行"
	case 2:
		stname = "急行"
	case 3:
		stname = "禁行"
	}

	msg_map(fmt.Sprintf("角色 %s 结束 %s 状态", rname, stname))
}
func (uv *GtkUserView) NoteRoleCeEnd(rind int) { // 阵营公敌状态终止
	name, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(rind)
	msg_map(fmt.Sprintf("%s 阵营结束宣战状态", name))
}
func (uv *GtkUserView) NoteRoleAlignEnd(r1 int, r2 int) { // 阵营间结盟状态终止 (r1 < r2)
	r1n, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(r1)
	r2n, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(r2)
	msg_map(fmt.Sprintf("%s 阵营结束与 %s 阵营的盟约", r1n, r2n))
}

func (uv *GtkUserView) NoteRoleWin (rind int) { // 阵营胜利
	rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(rind)
	msg_box (fmt.Sprintf("角色 %s 赢得游戏, 游戏结束", rname), 0)
	game_end(rind)
}

func (uv *GtkUserView) NoteRoleLoss(rind int) { // 阵营破产
	rname, loc, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(rind)
	set_role (loc, rind, true)
	update_role_point (loc)
	msg_box (fmt.Sprintf("角色 %s 破产, 离开游戏", rname), 0)
}

func (uv *GtkUserView) NoteAllLoss() { // 所有阵营同时破产
	msg_box(fmt.Sprintf("游戏结束, 所有阵营同时破产, 无人赢得游戏"), 0)
	game_end(-1)
}

func (uv *GtkUserView) NoteForceQuit() { // 用户强制退出
	game_end(-2)
}

func (uv *GtkUserView) NoteRoleMoveOneStep(rind int, sloc int, dir int) { // 角色移动一步
}

func (uv *GtkUserView) NoteRoleMove(rind int, sloc int, oloc int, step int) { // 角色移动一个回合
	set_role(sloc, rind, true)
	set_role(oloc, rind, false)
	update_role_point(sloc)
	update_role_point(oloc)
}

func (uv *GtkUserView) NoteGetCard(rind int, card_id int) { // 角色获得卡片
	rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(rind)
	dname, _, _ := mbg.GetCardInfo(card_id)
	msg_map(fmt.Sprintf("角色 %s 获得卡片 %s", rname, dname))
}

func (uv *GtkUserView) EndTurn(turn int) { // 回合结束
	msg_map(fmt.Sprintf("回合 %d 结束", turn))
	msg_map("资金情况:") 
	_, nrole, _, _, _, _, _, _, _  := uv.d.GetInfo()
	for i := 0; i < nrole; i ++ {
		rname, _, _, _, _, _, _, money, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(i)
		msg := fmt.Sprintf ("  阵营 %s 总资金 %.0f", rname, money)
		msg_map(msg)
	}
}

func (uv *GtkUserView) BattleNoGerenal(rind int, side int) { // 无主将参与作战, 自动战败
	name, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(rind)
	var msg string
	if side == 0 {
		msg = fmt.Sprintf("角色 %s 未选择出战将领, 无法开战", name)
	} else {
		msg = fmt.Sprintf("角色 %s 无将领出战, 自动战败", name)
	}
	msg_map(msg)
}

// sst: 发动方状态, ost: 被动方状态, max_bt: 最大回合数
func (uv *GtkUserView) BattleStart(sst *mbg.BattleState, ost *mbg.BattleState, max_bt int) { // 战斗开始
	battle_start((max_bt))
	var raname, rdname string
	var tlist []int

	role, power, winp, celve, hpmax, hp, def, _, _, _, _, _, _, _, _, tlist := sst.GetInfo()
	if role == -1 {
		raname = "山贼"
	} else {
		raname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = uv.d.GetRoleInfo(role)
	}
	battle_set (0, raname, power, winp, celve, hpmax, hp, def)
	for _, i := range tlist {
		battle_tech_set (0, i)
	}

	role, power, winp, celve, hpmax, hp, def, _, _, _, _, _, _, _, _, tlist = ost.GetInfo()
	if role == -1 {
		rdname = "山贼"
	} else {
		rdname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = uv.d.GetRoleInfo(role)
	}
	battle_set (1, rdname, power, winp, celve, hpmax, hp, def)
	for _, i := range tlist {
		battle_tech_set(1, i)
	}
	battle_msg (fmt.Sprintf("战斗开始. 攻方: %s, 守方 %s", raname, rdname))
}

func (uv *GtkUserView) BattleTurnStart(iturn int) { // 战斗回合开始. iturn: 回合数 (0 起始)
	battle_turn_start(iturn)
	// battle_msg (fmt.Sprintf("第 %d 回合", iturn))
}
func (uv *GtkUserView) BattleAttack(ls int, ost *mbg.BattleState) {
	_, _, _, _, _, hp, def, _, ddef, dhp, _, _, _, _, _, _ := ost.GetInfo()
	var sss string
	if ls == 0 {
		sss = "攻方"
	} else {
		sss = "守方"
	}
	battle_change_hp (ls, hp, def)
	battle_msg(fmt.Sprintf("%s 损失 %6.0f 城防, %6.0f 兵力.", sss, -ddef, -dhp))
}

func (uv *GtkUserView) BattleTech(ss int, tind int, sst *mbg.BattleState, ost *mbg.BattleState) { // 战斗时发动技能. ss: 0|1 发动方
	tname, _, _, _, _, _, _, _ := mbg.GetTechInfo(tind)
	_, _, _, _, _, shp, sdef, _, _, sdhp, _, sbst, sisq, sbtime, sfst, _ := sst.GetInfo()
	_, _, _, _, _, ohp, odef, _, oddef, odhp, _, obst, oisq, obtime, ofst, _ := ost.GetInfo()

	var sss, msg string
	if ss == 0 {
		sss = "攻方"
	} else {
		sss = "守方"
	}
	msg = sss + " 发动策略 " + tname
	if sdhp > 0 {
		msg += fmt.Sprintf(", 恢复 %8.0f", sdhp)
	}
	if odhp < 0 {
		msg += fmt.Sprintf(", 造成伤害 %8.0f", -odhp)
	}
	if oddef < 0 {
		msg += fmt.Sprintf(", 造成城防破坏 %8.0f", -oddef)
	}

	battle_active_tech(ss, tind, shp, ohp, sdef, odef)
	battle_msg(msg)
	battle_change_st(ss, sbst, sisq, sbtime, sfst)
	battle_change_st(1-ss, obst, oisq, obtime, ofst)
}

func (uv *GtkUserView) BattleBurn(side int, st *mbg.BattleState) { // 战斗中被烧
	var msg, sss string
	if side == 0 {
		sss = "攻方"
	} else {
		sss = "守方"
	}
	_, _, _, _, _, hp, def, _, ddef, dhp, _, _, _, _, _, _ := st.GetInfo()
	msg = fmt.Sprintf("%s 遭到火烧, 损失 %6.0f 防御, %6.0f 兵力", sss, -ddef, -dhp)
	battle_change_hp((side), (hp), (def))
	battle_msg(msg)
}
func (uv *GtkUserView) BattleSelfAttack(side int, st *mbg.BattleState) { // 战斗中受到内讧伤害
	var msg, sss string
	if side == 0 {
		sss = "攻方"
	} else {
		sss = "守方"
	}
	_, _, _, _, _, hp, def, _, ddef, dhp, _, _, _, _, _, _ := st.GetInfo()
	msg = fmt.Sprintf("%s 内讧, 损失 %6.0f 防御, %6.0f 兵力", sss, -ddef, -dhp)
	battle_change_hp((side), (hp), (def))
	battle_msg(msg)
}
func (uv *GtkUserView) BattleEnd(ret int) { // 战斗结束: ret: 战斗结果. 0: 和平结束. -1: 攻方失败. 1: 攻方取胜.
	var msg string
	switch ret {
	case 0:
		msg = "双方收兵, 战斗结束"
	case -1:
		msg = "攻方战败, 战斗结束"
	case 1:
		msg = "攻方胜利, 战斗结束"
	}
	battle_msg(msg)
	battle_end()
}

func (oi *GtkOperationInterface) StartGame(d *mbg.Driver, rind int) { // 游戏开始
	oi.d = d
	oi.rind = rind
}

func ShowPeopleInfo (d *mbg.Driver, pind int,
	ui_show_func func (name string, role string, job int, loc int, hpmax , hp float32, prop [5]float32, is_quit bool, lst , pst int),
) {
	pname, role, job, loc, hpmax, hp, prop, is_quit, lst, pst := d.GetPeopleInfo(pind)
	var rname string
	if role != -1 {
		rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = d.GetRoleInfo(role)
	} else {
		rname = "在野"
	}
	ui_show_func (pname, rname, job, loc, hpmax, hp, prop, is_quit, lst, pst)
}

func ShowCityInfo (d *mbg.Driver, cind int,
	ui_show_func func (name string, scope int, hpmax , hp float32, role string, nmos int, mayor, treasurer string, fengshui float32),
) {
	cname, scope, hpmax, hp, role, mos, mayor, treasurer, fengshui, _, _ := d.GetCityInfo(cind)
	var rname, mname, trea_name string
	if role != -1 {
		rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = d.GetRoleInfo(role)
	} else {
		rname = "无"
	}
	if mayor != -1 {
		mname, _, _, _, _, _, _, _, _, _ = d.GetPeopleInfo(mayor)
	} else {
		mname = "无"
	}
	if treasurer != -1 {
		trea_name, _, _, _, _, _, _, _, _, _ = d.GetPeopleInfo(treasurer)
	} else {
		trea_name = "无"
	}

	ui_show_func (cname, scope, hpmax, hp, rname, len(mos), mname, trea_name, fengshui)
}

func ShowTechInfo (hind int,
	ui_show_tech func (name string, study float32, scond int),
) {
	tname, study, scond, _, _, _, _, _ := mbg.GetTechInfo(hind)
	ui_show_tech (tname, study, scond)
}

func (oi *GtkOperationInterface) Check() {
	d := oi.d

	_, nrole, npeople, ncity, ninst, ntrain, _, _, _ := d.GetInfo()
      class := get_check_type()

	switch class {
	case 0:
		begin_role_list()
		for i := 0; i < nrole; i++ {
			name, loc, _, _, _, _, _, money, dir, mst, mtime, cst, ast, atime, _ := d.GetRoleInfo(i)
			add_role_list (name, loc, money, dir, mst, mtime, cst, ast, atime)
		}
		end_role_list()

		for {
			rind := get_role_select()
			if rind == -1 {
				break
			}
			cind := get_role_mcs_select()
			if cind == -1 {
				_, _, tech, mos, mcs, sos, tos, _, _, _, _, _, _, _, cards := d.GetRoleInfo(rind)

				begin_role_tech_list()
				for _, i := range tech {
					add_role_tech_list(i)
				}
				end_role_tech_list()

				begin_role_mos_list()
				for _, i := range mos {
					ShowPeopleInfo (d, i, add_role_mos_list)
				}
				end_role_mos_list()

				begin_role_mcs_list()
				for _, i := range mcs {
					ShowCityInfo(d, i, add_role_mcs_list)
				}
				end_role_mcs_list()

				begin_role_sos_list()
				for _, i := range sos {
					ShowPeopleInfo(d, i, add_role_sos_list)
				}
				end_role_sos_list()

				begin_role_tos_list()
				for _, i := range tos {
					ShowPeopleInfo(d, i, add_role_tos_list)
				}
				end_role_tos_list()

				begin_role_cards_list()
				for i := 0; i < mbg.CARD_COUNT; i++ {
					add_role_cards_list(i, cards[i])
				}
				end_role_cards_list()
			} else {
				_, _, _, _, mcs, _, _, _, _, _, _, _, _, _, _ := d.GetRoleInfo(rind)
				cname, _, _, _, _, mos, _, _, _, _, _ := d.GetCityInfo(mcs[cind])
				begin_people_list (cname + "官员")
				for _, i := range mos {
					ShowPeopleInfo (d, i, add_people_list)
				}
				end_people_list(1)
				get_people_exit()
			}
		}
	case 1:
		begin_people_list("全体人员")
		for i := 0; i < npeople; i++ {
			ShowPeopleInfo(d, i, add_people_list)
		}
		end_people_list(0)
		get_people_exit()
	case 2:
		begin_city_list()
		for i := 0; i < ncity; i++ {
			ShowCityInfo(d, i, add_city_list)
		}
		end_city_list()

		for {
			cind := get_city_mcs_select()
			if cind == -1 {
				break
			}
			cname, _, _, _, _, mos, _, _, _, _, _ := d.GetCityInfo(cind)
			begin_people_list(cname + "官员")
			for _, i := range mos {
				ShowPeopleInfo(d, i, add_people_list)
			}
			end_people_list(2)
			get_people_exit()
		}
	case 3:
		begin_inst_list()
		for i := 0; i < ninst; i++ {
			iname, tech, mos, on_study, point, left_round := d.GetInstInfo(i, oi.rind)
			var mos_name, study_name string
			if mos == -1 {
				mos_name = "无"
			} else {
				mos_name, _, _, _, _, _, _, _, _, _ = d.GetPeopleInfo(mos)
			}
			if on_study == -1 {
				study_name = "无"
			} else {
				study_name, _, _, _, _, _, _, _ = mbg.GetTechInfo(on_study)
			}
			add_inst_list(iname, len(tech), mos_name, study_name, point, left_round)
		}
		end_inst_list()

		for {
			iind := get_inst_select()
			if iind == -1 {
				break
			}
			iname, tech, mos, _, _, _ := d.GetInstInfo(iind, oi.rind)
			if mos != -1 {
				ShowPeopleInfo (d, mos, show_inst_mos)
			}
			begin_inst_tech_list(iname)
			for _, i := range tech {
				ShowTechInfo (i, add_inst_tech_list)
			}
			end_inst_tech_list()
		}
	case 4:
		begin_train_list()
		for i := 0; i < ntrain; i++ {
			tname, item, mos, round := d.GetTrainInfo(i, oi.rind)
                  var pname string
			if mos != -1 {
				pname, _, _, _, _, _, _, _, _, _ = d.GetPeopleInfo(mos)
			} else {
				pname = "无"
			}
			add_train_list (tname, item, pname, round)
		}
		end_train_list()

		for {
			tind := get_train_select()
			if tind == -1 {
				break
			}
			_, _, mos, _ := d.GetTrainInfo(tind, oi.rind)
			if mos != -1 {
				ShowPeopleInfo (d, mos, show_train_mos)
			}
		}
	case 5:
		begin_people_list("人员查询")
		for i := 0; i < npeople; i++ {
			_, role, _, _, _, _, _, _, _, _ := d.GetPeopleInfo(i)
			if role == oi.rind {
				ShowPeopleInfo(d, i, add_people_list)
			}
		}

		end_people_list(0)
		get_people_exit()
	}
}

func (oi *GtkOperationInterface) AlignFail(obj_role int) { // 设置结盟失败
	msg_map("无法结盟")
}
func (oi *GtkOperationInterface) PEFail(rind int) { // 设置公敌失败
	msg_map("无法宣战")
}

func (oi *GtkOperationInterface) NoteAsMayer(pind int) { // 某人就职市长
	var msg string
	pname, _, _, _, _, _, _, _, _, _ := oi.d.GetPeopleInfo(pind)
	msg = fmt.Sprintf("%s 就任市长", pname)
	msg_map(msg)
}
func (oi *GtkOperationInterface) NoteAsTreasurer(pind int) { // 某人就职财务官
	pname, _, _, _, _, _, _, _, _, _ := oi.d.GetPeopleInfo(pind)
	msg_map(fmt.Sprintf("%s 就任财务官", pname))
}
func (oi *GtkOperationInterface) NoteAsCitizen(pind int) { // 某人撤销市政职务
	pname, _, _, _, _, _, _, _, _, _ := oi.d.GetPeopleInfo(pind)
	msg_map(fmt.Sprintf("%s 被撤销市政职务", pname))
}

func (oi *GtkOperationInterface) NoteFinishStudy(iind int, pind int, tech int) { // 研发完成
	rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := oi.d.GetRoleInfo(oi.rind)
	pname, _, _, _, _, _, _, _, _, _ := oi.d.GetPeopleInfo(pind)
	hname, _, _, _, _, _, _, _ := mbg.GetTechInfo(tech)
	msg_map(fmt.Sprintf("%s 为 %s 研究策略 %s 完成.", pname, rname, hname))
}
func (oi *GtkOperationInterface) NoteLYStudy(iind int, pind int) { // 研发因流言无法继续
	iname, _, _, _, _, _ := oi.d.GetInstInfo(iind, 0)
	pname, _, _, _, _, _, _, _, _, _ := oi.d.GetPeopleInfo(pind)
	msg_map(fmt.Sprintf("%s 处于流言状态, 暂停位于 %s 的研究工作", pname, iname))
}

func GetPropName(item mbg.Property) string {
	switch item {
	case mbg.Wu:
		return "武术"
	case mbg.Zhan:
		return "战术"
	case mbg.Mou:
		return "策略"
	case mbg.Zheng:
		return "政治"
	case mbg.Jing:
		return "经济"
	}
	return ""
}

func (oi *GtkOperationInterface) NoteFinishTrain (pind int, item mbg.Property, value float32) { // 训练完成
	pname, _, _, _, _, _, _, _, _, _ := oi.d.GetPeopleInfo(pind)
	msg_map(fmt.Sprintf("%s 完成 %s 的训练, 能力提升 %6.0f", pname, GetPropName(item), value))
}
func (oi *GtkOperationInterface) NoteLYTrain(tind int, pind int) { // 训练因流言无法继续
	pname, _, _, _, _, _, _, _, _, _ := oi.d.GetPeopleInfo(pind)
	_, item, _, _ := oi.d.GetTrainInfo(tind, oi.rind)
	msg_map(fmt.Sprintf("%s 关于 %s 的训练因流言暂停", pname, GetPropName(item)))
}
func (oi *GtkOperationInterface) NoteLYMayor(pind, cind int) { // 因流言无法履行市长职责
	pname, _, _, _, _, _, _, _, _, _ := oi.d.GetPeopleInfo(pind)
	cname, _, _, _, _, _, _, _, _, _, _ := oi.d.GetCityInfo(cind)
	msg_map(fmt.Sprintf("%s 因流言无法履行 %s 的市长职责", pname, cname))
}
func (oi *GtkOperationInterface) NoteLYTreasurer(pind, cind int) { // 因流言无法履行财务官职责
	pname, _, _, _, _, _, _, _, _, _ := oi.d.GetPeopleInfo(pind)
	cname, _, _, _, _, _, _, _, _, _, _ := oi.d.GetCityInfo(cind)
	msg_map(fmt.Sprintf("%s 因流言无法履行 %s 的财务官职责", pname, cname))
}
func (oi *GtkOperationInterface) NoteCityRecover(cind int, hplus float32) { // 城市因市长恢复所属人员 HP
	cname, _, _, _, _, _, _, _, _, _, _ := oi.d.GetCityInfo(cind)
	msg_map(fmt.Sprintf("城市 %s 恢复兵力 %6.0f", cname, hplus))
}
func (oi *GtkOperationInterface) NoteCityEarn(cind int, mplus float32) { // 城市获得收入
	cname, _, _, _, _, _, _, _, _, _, _ := oi.d.GetCityInfo(cind)
	msg_map(fmt.Sprintf("城市 %s 获得收入 %6.0f", cname, mplus))
}

func (oi *GtkOperationInterface) NotePayInst(dm float32) { // 支付研究经费
	rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := oi.d.GetRoleInfo(oi.rind)
	msg_map(fmt.Sprintf("%s 支付研究经费 %6.0f", rname, -dm))
}

func (oi *GtkOperationInterface) SkipBattleByAlign(robj int) { // 因为结盟跳过战斗
	srname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := oi.d.GetRoleInfo(oi.rind)
	orname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := oi.d.GetRoleInfo(robj)
	msg_map(fmt.Sprintf("%s 与 %s 因结盟无法战斗", srname, orname))
}
func (oi *GtkOperationInterface) ForceBattle(robj int) { // 因为公愤进行强制战斗
	srname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := oi.d.GetRoleInfo(oi.rind)
	orname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := oi.d.GetRoleInfo(robj)
	msg_map(fmt.Sprintf("%s 与 %s 因宣战状态强制发生战斗", srname, orname))
}

func (oi *GtkOperationInterface) SelRoleAction (clist []int) int { // 选择阵营行动: -1: 移动, >=0: 使用卡片
	for {
		action := get_action(oi.rind)
		switch action {
		case -1:
			return -2
		case 0:
			return -1
		case 3:
			return -3
		case 1:
			_, _, _, _, _, _, _, _, _, _, _, _, _, _, cards := oi.d.GetRoleInfo(oi.rind)
			begin_cards_list()
			for i := 0; i < mbg.CARD_COUNT; i++ {
				add_cards_list(i, cards[i])
			}
			end_cards_list()
			sc := (get_card_action())
			if sc < 0 {
				continue
			}
			return sc
		case 2:
			oi.Check()
		}
	}
}

func (oi *GtkOperationInterface) GetFileName () string { // 返回游戏进度文件名 ( SelRoleAction = -3 时会紧接着被调用), 返回空串表示取消
	return get_file_choose()
}

func (oi *GtkOperationInterface) SaveReport (is_success bool) { // 保存文件通知
	var msg string
	if is_success {
		rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := oi.d.GetRoleInfo(oi.rind)
		msg = fmt.Sprintf ("保存文件成功, 当前行动角色为 %s", rname)
	} else {
		msg = "保存文件失败"
	}

	msg_map (msg)
}

func (oi *GtkOperationInterface) IsOccuCity (cind int) bool { // 选择是否占领空白城市
	rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := oi.d.GetRoleInfo(oi.rind)
	cname, scope, _, _, _, _, _, _, _, _, _ := oi.d.GetCityInfo(cind)
	var msg, scope_str string
	switch scope {
	case 0:
		scope_str = "小"
	case 1:
		scope_str = "中"
	case 2:
		scope_str = "大"
	}
	msg = fmt.Sprintf("%s 是否占领%s型城市 %s", rname, scope_str, cname)
	for {
		res := msg_box (msg, 1)
		if res == 1 {
			oi.Check()
			msg = ""
			continue
		}
		if res == 0 {
			return true
		} else {
			return false
		}
	}
}

func (oi *GtkOperationInterface) OccuCityFail(cind int) { // 占领城市失败
	rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := oi.d.GetRoleInfo(oi.rind)
	cname, _, _, _, _, _, _, _, _, _, _ := oi.d.GetCityInfo(cind)
	msg_map(fmt.Sprintf("%s 占领城市 %s 失败", rname, cname))
}

func (oi *GtkOperationInterface) SelCityMos(splist []int, cind int) []int { // 选择占领的人员
	begin_people_list("选择占领人员")
	for _, i := range splist {
		ShowPeopleInfo(oi.d, i, add_people_list)
	}
	end_people_list(0)
	nsel := get_people_list_num()
	var res []int = make([]int, nsel)
	for i := 0; i < nsel; i++ {
		res[i] = get_people_list(i)
	}
	return res
}

func (oi *GtkOperationInterface) SelMayor (cind int, splist []int, old_mayor int, old_treasurer int) (mayor int) { // 选择市长
	cname, _, _, _, _, _, _, _, _, _, _ := oi.d.GetCityInfo(cind)
	begin_people_list(fmt.Sprintf("为 %s 选择市长", cname))
	for _, i := range splist {
		ShowPeopleInfo(oi.d, i, add_people_list)
	}
	end_people_list(0)
	return get_people_one()
}

func (oi *GtkOperationInterface) SelTreasurer(cind int, splist []int, old_mayor int, old_treasurer int) (treasurer int) { // 选择财务官
	cname, _, _, _, _, _, _, _, _, _, _ := oi.d.GetCityInfo(cind)
	begin_people_list(fmt.Sprintf("为 %s 选择财务官", cname))
	for _, i := range splist {
		ShowPeopleInfo(oi.d, i, add_people_list)
	}
	end_people_list(0)
	return get_people_one()
}

func (oi *GtkOperationInterface) IsAttackCity(cind int) int { // 选择是否攻城 -1: 不攻城 0-2: 三种攻城规模
	cname, _, _, _, _, _, _, _, _, _, _ := oi.d.GetCityInfo(cind)
	var msg string
	msg = fmt.Sprintf("是否进攻城市 %s? 请选择发动战斗的规模.", cname)
	for {
		res := sel_battle_scale (msg)
		if res == -2 {
			oi.Check()
			msg = ""
			continue
		}
		return res
	}
}

func add_saloon_list_functor (price float32) (
	res_func func(name, role string, job , loc int, hpmax , hp float32, prop [5]float32, is_quit bool, lst , pst int),
) {
	return func(name, role string, job , loc int, hpmax , hp float32, prop [5]float32, is_quit bool, lst , pst int) {
		add_saloon_list (name, role, job, loc, hpmax, hp, prop, is_quit, lst, pst, price)
	}
}

func (oi *GtkOperationInterface) SaloonSelPeople (pind []int, price []float32) int { // 选择招募人员 : -1: 放弃
	begin_saloon_list()
	for i := 0; i < len(pind); i++ {
		ShowPeopleInfo (oi.d, pind[i], add_saloon_list_functor (price[i]))
	}
	end_saloon_list()

	return get_saloon_select()
}

func (oi *GtkOperationInterface) RecruitFail(pind int) { // 因无法支付招募失败
	msg_map ("金币不足, 招募失败")
}

func (oi *GtkOperationInterface) SelStudy (iind int, plist []int, tlist []int) (sp int, st int) { // 选择研究人员和项目, (-1, -1): 放弃研究
	iname, _, _, _, _, _ := oi.d.GetInstInfo(iind, oi.rind)
	begin_study_tech_list(iname)
	for _, i := range tlist {
		ShowTechInfo(i, add_study_tech_list)
	}
	end_study_tech_list()
	begin_study_people_list()
	for _, i := range plist {
		ShowPeopleInfo(oi.d, i, add_study_people_list)
	}
	end_study_people_list()

	st, sp = get_study_select()
	if sp == -1 || st == -1 {
		sp = -1
		st = -1
	}
	return
}

func (oi *GtkOperationInterface) SelTrain(tind int, plist []int, item mbg.Property) int { // 选择训练人员, 返回人员编号, -1: 放弃
	tname, _, _, _ := oi.d.GetTrainInfo(tind, oi.rind)
	begin_people_list(fmt.Sprintf("%s: 选择训练人员. 训练项目: %s", tname, GetPropName(item)))
	for _, i := range plist {
		ShowPeopleInfo(oi.d, i, add_people_list)
	}
	end_people_list(0)

	return get_people_one()
}

func (oi *GtkOperationInterface) BattleWithRole(obj_role int) int { // 是否打遭遇战. 返回: -1: 不打. 0-2: 战斗规模
	var msg string
	rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := oi.d.GetRoleInfo(obj_role)
	msg = fmt.Sprintf("是否与 %s 发生战斗? 请选择战斗规模", rname)
	for {
		res := sel_battle_scale (msg)
		if res == -2 {
			oi.Check()
			msg = ""
			continue
		}
		return res
	}
}

func (oi *GtkOperationInterface) ConfirmCard(card_id int) int { // 确认是否使用卡片: 无对象卡片 返回: 0: 确定; -1:取消
	var msg string
	dname, _, _ := mbg.GetCardInfo(card_id)
	msg = fmt.Sprintf("是否确认使用卡片 %s", dname)
	for {
		res := msg_box(msg, 1)
		if res == 1 {
			oi.Check()
			msg = ""
			continue
		}
		if res != 0 {
			res = -1
		}
		return res
	}
}

func (oi *GtkOperationInterface) SelCardObjCity(card_id int) int { // 返回卡片作用对象城市. -1 取消
	dname, _, _ := mbg.GetCardInfo(card_id)
	msg_map (fmt.Sprintf("请选择卡片 %s 的作用城市", dname))

	for {
		obj_type := get_obj_type()
		if obj_type == -1 {
			return -1
		}
		if obj_type == -3 {
			oi.Check()
			continue
		}
		if obj_type != -2 {
			continue
		}
		loc := get_obj()

		class, base, _, _, _ := oi.d.GetBoardInfo(loc)
		if class != 1 {
			continue
		}
		return base
	}
}

func (oi *GtkOperationInterface) SelCardObjAny (card_id int) (st int, obj int) { // 返回卡片作用对象建筑或角色. st: -1: 取消. 0: 角色 1: 城市: 2: 研究所 3: 修炼所
	dname, _, _ := mbg.GetCardInfo(card_id)
	msg_map(fmt.Sprintf("请选择卡片 %s 的作用官员", dname))

	for {
		obj_type := get_obj_type()
		switch obj_type {
		case -1:
			return -1, -1
		case -2:
			loc := get_obj()
			class, base, _, _, _ := oi.d.GetBoardInfo(loc)
			switch class {
			case 1:
				return 1, base
			case 3:
				return 2, base
			case 4:
				return 3, base
			}
		case -3:
			oi.Check()
		default:
			return 0, obj_type
		}
	}
}

func (oi *GtkOperationInterface) SelCardObjPeople(card_id int, plist []int) int { // 返回卡片作用人员. -1: 取消
	dname, _, _ := mbg.GetCardInfo(card_id)
	begin_people_list (fmt.Sprintf("请选择卡片 %s 的作用官员", dname))
	for _, i := range plist {
		ShowPeopleInfo(oi.d, i, add_people_list)
	}
	end_people_list(0)
	return get_people_one()
}

func (oi *GtkOperationInterface) SelCardObjRoleAndPeople(card_id int, rplist []int) (sr_id int) { // 返回卡片作用人员  -1:取消
	dname, _, _ := mbg.GetCardInfo(card_id)
	begin_people_list (fmt.Sprintf("请选择卡片 %s 的作用官员", dname))
	for _, i := range rplist {
		ShowPeopleInfo(oi.d, i, add_people_list)
	}
	end_people_list(0)
	return get_people_one()
}

func (oi *GtkOperationInterface) SelCardObjRole(card_id int) int { // 返回选择卡片作用阵营 -1: 取消
	dname, _, _ := mbg.GetCardInfo(card_id)
	msg_map (fmt.Sprintf("请选择卡片 %s 的作用阵营", dname))

	for {
		obj_type := get_obj_type()
		if obj_type == -1 {
			return -1
		}
		if obj_type == -2 {
			continue
		}
		if obj_type == -3 {
			oi.Check()
			continue
		}
		return obj_type
	}
}

func (oi *GtkOperationInterface) SelCardObjLoc(card_id int) int { // 选择卡片作用地图位置 -1: 取消
	dname, _, _ := mbg.GetCardInfo(card_id)
	msg_map (fmt.Sprintf("请选择卡片 %s 的作用地图位置", dname))

	for {
		obj_type := get_obj_type()
		if obj_type == -1 {
			return -1
		}
		if obj_type == -3 {
			oi.Check()
			continue
		}
		if obj_type >= 0 {
			continue
		}
		return get_obj()
	}
}

func (oi *GtkOperationInterface) ShowAllocAvailableLocation ([]bool) { // 在调度目前选择之前告诉 UI 哪些位置是可用的 
}

func (oi *GtkOperationInterface) SelAllocObj() (st int, obj int) { // 选择调度一级目标. st: -1: 取消. 0: 城市: 1: 研究所 2: 修炼所

	for {
		obj_type := get_obj_type()
		if obj_type == -1 {
			clear_map_av ()
			return -1, -1
		}
		if obj_type == -3 {
			oi.Check()
			continue
		}
		if obj_type >= 0 {
			continue
		}
		loc := get_obj()
		class, base, _, _, _ := oi.d.GetBoardInfo(loc)
		clear_map_av ()
		switch class {
		case 1:
			st = 0
			return st, base
		case 3:
			st = 1
			return st, base
		case 4:
			st = 2
			return st, base
		}
	}

}

func make_add_exchange_func (side int) func(name, role string, job , loc int, hpmax , hp float32, prop [5]float32, is_quit bool, lst , pst int) {
	return func(name , role string, job , loc int, hpmax , hp float32, prop [5]float32, is_quit bool, lst , pst int) {
		add_exchange_people_list (side, name, role, job, loc, hpmax, hp, prop, is_quit, lst, pst)
	}
}

func (oi *GtkOperationInterface) ExchangeCityPeople(cind int, splist []int, oplist []int) (in_list []int, out_list []int) { // 调度随身人员与目标内驻扎人员.
	rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := oi.d.GetRoleInfo(oi.rind)
	cname, _, _, _, _, _, _, _, _, _, _ := oi.d.GetCityInfo(cind)
	msg := fmt.Sprintf("%s 阵营调度城市 %s 的官员", rname, cname)

	for k := 0; k < 2; k++ {
		begin_exchange_people_list (k, msg)
		var plist []int
		if k == 0 {
			plist = splist
		} else {
			plist = oplist
		}
		for _, i := range plist {
			ShowPeopleInfo (oi.d, i, make_add_exchange_func (k))
		}
		end_exchange_people_list(k)
	}

	var ret_list [2]([]int)
	for k := 0; k < 2; k++ {
		nsel := get_exchange_people_list_num(k)
		ret_list[k] = make([]int, nsel)
		for i := 0; i < nsel; i++ {
			ret_list[k][i] = get_exchange_people_list(k, i)
		}
	}

	return ret_list[0], ret_list[1]
}

func (oi *GtkOperationInterface) IsCancelStudy(pind int, on_study int, left_point float32) bool { // 是否终止研究
	var msg string
	pname, _, _, _, _, _, _, _, _, _ := oi.d.GetPeopleInfo(pind)
	hname, _, _, _, _, _, _, _ := mbg.GetTechInfo(on_study)
	msg = fmt.Sprintf("是否撤销 %s 对技能 %s 的研究: 剩余研究点数: %6.0f", pname, hname, left_point)
	for {
		res := msg_box(msg, 1)
		if res == 1 {
			oi.Check()
			msg = ""
			continue
		}
		if res == 0 {
			return true
		} else {
			return false
		}
	}
}

func (oi *GtkOperationInterface) IsCancelTrain(pind int, item mbg.Property, left_round int) bool { // 是否终止训练
	var msg string
	pname, _, _, _, _, _, _, _, _, _ := oi.d.GetPeopleInfo(pind)
	msg = fmt.Sprintf("是否撤销 %s 对属性 %s 的训练: 剩余完成回合数: %d", pname, GetPropName(item), left_round)
	for {
		res := msg_box(msg, 1)
		if res == 1 {
			oi.Check()
			msg = ""
			continue
		}
		if res == 0 {
			return true
		} else {
			return false
		}
	}
}

// 开始人员调度, 通知 UI 哪些位置是可用的 (av_loc [ngrid]bool: 位置是否可被调度)
func (oi *GtkOperationInterface) StartAllocate (av_loc []bool) {
	rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := oi.d.GetRoleInfo(oi.rind)
	msg_map(fmt.Sprintf("%s 开始人员调度回合", rname))
	msg_map("请在地图上选择调度对象 (城市, 研究所, 训练所)")
	show_map_av (av_loc)
}

func (oi *GtkOperationInterface) SelGeneral(plist []int) (main, vice int) { // 选择战争将领: main: 主将. vice: 副将 -1: 无副将, main == vice means vice = -1
	begin_people_list (fmt.Sprintf("选择出战主将与参军"))
	for _, i := range plist {
		ShowPeopleInfo(oi.d, i, add_people_list)
	}
	end_people_list(0)

	main = get_people_two_main()
	vice = get_people_two_vice()
	return
}

