package gtkui

// #cgo pkg-config: gtk+-3.0
//#include <stdlib.h>
//#include "gtkui_wrapper.h"
import "C"

import (
	"fmt"
	"unsafe"

	"github.com/tfcolin/mbg"
)

const max_cstr = 1024

type GtkUserView struct {
	d *mbg.Driver
}

func (uv *GtkUserView) StartGame(d *mbg.Driver) { // 游戏开始
	uv.d = d
	_, nrole, _, _, _, _, _, _, _ := d.GetInfo()

	for i := 0; i < mbg.TECH_COUNT; i++ {
		tname, _, scond, _, _, _, _, _ := mbg.GetTechInfo(i)
		cstr := C.CString(tname)
		C.init_tech(C.int(i), cstr, C.int(scond))
		C.free(unsafe.Pointer(cstr))
	}
	for i := 0; i < mbg.CARD_COUNT; i++ {
		dname, odir, otype := mbg.GetCardInfo(i)
		cstr := C.CString(dname)
		C.init_card(C.int(i), cstr, C.int(odir), C.int(otype))
		C.free(unsafe.Pointer(cstr))
	}

	for i := 0; i < nrole; i++ {
		rname, loc, _, _, _, _, _, _, _, _, _, _, _, _, _ := d.GetRoleInfo(i)
		cstr := C.CString(rname)
		C.init_role(C.int(i), cstr)
		C.free(unsafe.Pointer(cstr))
		C.set_role(C.int(loc), C.int(i), 0)
	}
	C.finish_draw_map()
}

func (uv *GtkUserView) SetBarrier(loc int) { // 设置路障
	C.set_barrier(C.int(loc), 0)
	C.update_role_point(C.int(loc))
}
func (uv *GtkUserView) UnSetBarrier(loc int) { // 路障生效并去除
	C.set_barrier(C.int(loc), 1)
	C.update_role_point(C.int(loc))
}
func (uv *GtkUserView) SetRobber(loc int) { // 设置山贼
	C.set_robber(C.int(loc), 0)
	C.update_role_point(C.int(loc))
}
func (uv *GtkUserView) UnSetRobber(loc int) { // 山贼生效并去除
	C.set_robber(C.int(loc), 1)
	C.update_role_point(C.int(loc))
}

func (uv *GtkUserView) SetSlow(rind int) { // 慢行
	rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(rind)
	var msg string
	msg = fmt.Sprintf("角色 %s 速度减慢", rname)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}
func (uv *GtkUserView) SetFast(rind int) { // 急行
	rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(rind)
	var msg string
	msg = fmt.Sprintf("角色 %s 速度加快", rname)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}
func (uv *GtkUserView) SetStop(rind int) { // 禁行
	rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(rind)
	var msg string
	msg = fmt.Sprintf("角色 %s 禁止行动", rname)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}
func (uv *GtkUserView) SetContinue(rind int) { // 双行
	rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(rind)
	var msg string
	msg = fmt.Sprintf("角色 %s 连续行动", rname)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}
func (uv *GtkUserView) ChangeDir(rind int) { // 设置反向
	rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(rind)
	var msg string
	msg = fmt.Sprintf("角色 %s 掉头行动", rname)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}
func (uv *GtkUserView) StealMoney(sub_role int, obj_role int, money float32) { // 偷钱
	var msg string
	sub_name, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(sub_role)
	obj_name, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(obj_role)
	msg = fmt.Sprintf("角色 %s 从角色 %s 手中偷去 %f 金币", sub_name, obj_name, money)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}
func (uv *GtkUserView) SelfRecover(rind int) { // 自我恢复
	var msg string
	rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(rind)
	msg = fmt.Sprintf("角色 %s 的官员得到恢复", rname)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}
func (uv *GtkUserView) SetLiuYan(pind int) { // 设置流言状态
	var msg string
	pname, _, _, _, _, _, _, _, _, _ := uv.d.GetPeopleInfo(pind)
	msg = fmt.Sprintf("官员 %s 进入流言状态, 暂时无法工作", pname)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}
func (uv *GtkUserView) SetPoison(pind int) { // 设置中毒状态
	var msg string
	pname, _, _, _, _, _, _, _, _, _ := uv.d.GetPeopleInfo(pind)
	msg = fmt.Sprintf("官员 %s 进入中毒状态, 能力属性下降", pname)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}
func (uv *GtkUserView) SetQuit(pind int) { // 设置离职状态
	var msg string
	pname, rind, _, _, _, _, _, _, _, _ := uv.d.GetPeopleInfo(pind)
	rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(rind)
	msg = fmt.Sprintf("官员 %s 接受劝说, 准备离开 %s", pname, rname)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}
func (uv *GtkUserView) SetAlign(sub_role int, obj_role int) { // 设置结盟
	sub_name, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(sub_role)
	obj_name, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(obj_role)
	var msg string
	msg = fmt.Sprintf("%s 与 %s 结盟", sub_name, obj_name)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}
func (uv *GtkUserView) SetPE(role int) { // 设置公敌
	var msg string
	name, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(role)
	msg = fmt.Sprintf("%s 阵营与所有其他阵营宣战", name)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}

func (uv *GtkUserView) InitMap(ngrid, nrole, npeople, ncity, ninst, ntrain, ntech, ncard, nx, ny int) { // 初始化地图
	C.init_map(C.int(ngrid), C.int(nrole), C.int(npeople), C.int(ncity), C.int(ninst), C.int(ntrain), C.int(ntech), C.int(ncard), C.int(nx), C.int(ny))
}

func (uv *GtkUserView) ShowMap(id int, class int, base int, x int, y int) { // 显示地图块
	C.set_board(C.int(id), C.int(class), C.int(base), C.int(x), C.int(y))
}

func (uv *GtkUserView) ShowCity(id int, name string, scope int, x0, y0, x1, y1 int) { // 显示城市
	cstr := C.CString(name)
	C.set_city(C.int(id), cstr)
	C.free(unsafe.Pointer(cstr))
}
func (uv *GtkUserView) ShowInstitute(id int, name string, x0, y0, x1, y1 int) { // 显示策略研究所
	cstr := C.CString(name)
	C.set_inst(C.int(id), cstr)
	C.free(unsafe.Pointer(cstr))
}
func (uv *GtkUserView) ShowTrainingRoom(id int, name string, item mbg.Property, x0, y0, x1, y1 int) { // 显示训练所
	cstr := C.CString(name)
	C.set_train(C.int(id), cstr)
	C.free(unsafe.Pointer(cstr))
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
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}
func (uv *GtkUserView) NoteInInst(pind int, new_inst int) { // 某人进驻研究所
	pname, rind, _, _, _, _, _, _, _, _ := uv.d.GetPeopleInfo(pind)
	if rind != -1 {
		iname, _, _, on_study, _, _ := uv.d.GetInstInfo(new_inst, rind)
		rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(rind)
		hname, _, _, _, _, _, _, _ := mbg.GetTechInfo(on_study)
		var msg string
		msg = fmt.Sprintf("官员 %s 进入 %s 为 %s 研究策略 %s", pname, iname, rname, hname)
		cstr := C.CString(msg)
		C.msg_map(cstr)
		C.free(unsafe.Pointer(cstr))
	}
}
func (uv *GtkUserView) NoteInTrain(pind int, new_train int) { // 某人进驻训练所
	pname, role, _, _, _, _, _, _, _, _ := uv.d.GetPeopleInfo(pind)
	tname, _, _, _ := uv.d.GetTrainInfo(new_train, role)
	var msg string
	msg = fmt.Sprintf("官员 %s 进入 %s 进行训练", pname, tname)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}
func (uv *GtkUserView) NoteOutCity(pind int, cind int) { // 某人撤离己方城市
	pname, _, _, _, _, _, _, _, _, _ := uv.d.GetPeopleInfo(pind)
	cname, _, _, _, _, _, _, _, _, _, _ := uv.d.GetCityInfo(cind)
	var msg string
	msg = fmt.Sprintf("官员 %s 离开城市 %s", pname, cname)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}
func (uv *GtkUserView) NoteOutInst(pind int, iind int) { // 某人撤离研究所
	pname, _, _, _, _, _, _, _, _, _ := uv.d.GetPeopleInfo(pind)
	iname, _, _, _, _, _ := uv.d.GetInstInfo(iind, 0)
	var msg string
	msg = fmt.Sprintf("官员 %s 离开 %s", pname, iname)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}
func (uv *GtkUserView) NoteOutTrain(pind int, tind int) { // 某人撤离训练所
	pname, _, _, _, _, _, _, _, _, _ := uv.d.GetPeopleInfo(pind)
	tname, _, _, _ := uv.d.GetTrainInfo(tind, 0)
	var msg string
	msg = fmt.Sprintf("官员 %s 离开 %s", pname, tname)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}
func (uv *GtkUserView) NoteJoin(pind int, rind int) { // 某人加入某阵营
	pname, _, _, _, _, _, _, _, _, _ := uv.d.GetPeopleInfo(pind)
	rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(rind)
	var msg string
	msg = fmt.Sprintf("官员 %s 加入 %s 阵营", pname, rname)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}
func (uv *GtkUserView) NoteQuit(pind int, rind int) { // 某人离开某阵营
	pname, _, _, _, _, _, _, _, _, _ := uv.d.GetPeopleInfo(pind)
	rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(rind)
	var msg string
	msg = fmt.Sprintf("官员 %s 离开 %s 阵营", pname, rname)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}
func (uv *GtkUserView) NotePayTax(send_role int, recv_role int, recv_city int, tax float32) { // 某人向某城支付租金
	send_name, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(send_role)
	recv_name, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(recv_role)
	cname, _, _, _, _, _, _, _, _, _, _ := uv.d.GetCityInfo(recv_city)
	var msg string
	msg = fmt.Sprintf("角色 %s 途径城市 %s, 向角色 %s 支付税金 %.0f", send_name, cname, recv_name, tax)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}
func (uv *GtkUserView) NoteCollectTax(rind int, cind int, tax float32) { // 某角色到达自身城市收取税收
	rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(rind)
	cname, _, _, _, _, _, _, _, _, _, _ := uv.d.GetCityInfo(cind)
	var msg string
	msg = fmt.Sprintf("角色 %s 途径已占领城市 %s, 额外收取税金 %.0f", rname, cname, tax)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}

func (uv *GtkUserView) RoleContinueAction(rind int) { // 某阵营连续行动 (在 RoleStartAction 之前调用)
	rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(rind)
	var msg string
	msg = fmt.Sprintf("角色 %s 连续行动", rname)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}
func (uv *GtkUserView) RoleStartAction(rind int) { // 某阵营开始行动
	rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(rind)
	var msg string
	msg = fmt.Sprintf("角色 %s 开始行动", rname)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}
func (uv *GtkUserView) RoleForbidAction(rind int) { // 某阵营被禁止行动
	name, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(rind)
	var msg string
	msg = fmt.Sprintf("角色 %s 本回合无法行动", name)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}

func (uv *GtkUserView) NoteRoleOccupyCity(rind int, cind int) { // 某阵营占领城市
	rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(rind)
	cname, _, _, _, _, _, _, _, _, _, _ := uv.d.GetCityInfo(cind)
	var msg string
	msg = fmt.Sprintf("%s 阵营占领城市 %s", rname, cname)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}

func (uv *GtkUserView) NoteRoleDropCity(rind int, cind int) { // 某阵营占领城市
	rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(rind)
	cname, _, _, _, _, _, _, _, _, _, _ := uv.d.GetCityInfo(cind)
	var msg string
	msg = fmt.Sprintf("%s 阵营放弃城市 %s", rname, cname)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
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

	var msg string
	msg = fmt.Sprintf("角色 %s 结束 %s 状态", rname, stname)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}
func (uv *GtkUserView) NoteRoleCeEnd(rind int) { // 阵营公敌状态终止
	name, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(rind)
	var msg string
	msg = fmt.Sprintf("%s 阵营结束宣战状态", name)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}
func (uv *GtkUserView) NoteRoleAlignEnd(r1 int, r2 int) { // 阵营间结盟状态终止 (r1 < r2)
	r1n, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(r1)
	r2n, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(r2)
	var msg string
	msg = fmt.Sprintf("%s 阵营结束与 %s 阵营的盟约", r1n, r2n)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}

func (uv *GtkUserView) NoteRoleWin(rind int) { // 阵营胜利
	rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(rind)
	var msg string
	msg = fmt.Sprintf("角色 %s 赢得游戏, 游戏结束", rname)
	cstr := C.CString(msg)
	C.msg_box(cstr, C.int(0))
	C.free(unsafe.Pointer(cstr))
	C.game_end(C.int(rind))
}

func (uv *GtkUserView) NoteRoleLoss(rind int) { // 阵营破产
	rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(rind)
	var msg string
	msg = fmt.Sprintf("角色 %s 破产, 离开游戏", rname)
	cstr := C.CString(msg)
	C.msg_box(cstr, C.int(0))
	C.free(unsafe.Pointer(cstr))
}

func (uv *GtkUserView) NoteAllLoss() { // 所有阵营同时破产
	var msg string
	msg = fmt.Sprintf("游戏结束, 所有阵营同时破产, 无人赢得游戏")
	cstr := C.CString(msg)
	C.msg_box(cstr, C.int(0))
	C.free(unsafe.Pointer(cstr))
	C.game_end(-1)
}

func (uv *GtkUserView) NoteForceQuit() { // 用户强制退出
	C.game_end(-2)
}

func (uv *GtkUserView) NoteRoleMoveOneStep(rind int, sloc int, dir int) { // 角色移动一步
}

func (uv *GtkUserView) NoteRoleMove(rind int, sloc int, oloc int, step int) { // 角色移动一个回合
	C.set_role(C.int(sloc), C.int(rind), 1)
	C.set_role(C.int(oloc), C.int(rind), 0)
	C.update_role_point(C.int(sloc))
	C.update_role_point(C.int(oloc))
}

func (uv *GtkUserView) NoteGetCard(rind int, card_id int) { // 角色获得卡片
	rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(rind)
	dname, _, _ := mbg.GetCardInfo(card_id)
	var msg string
	msg = fmt.Sprintf("角色 %s 获得卡片 %s", rname, dname)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}

func (uv *GtkUserView) EndTurn(turn int) { // 回合结束
	var msg string
	msg = fmt.Sprintf("回合 %d 结束", turn)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}

func (uv *GtkUserView) BattleNoGerenal(rind int, side int) { // 无主将参与作战, 自动战败
	name, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := uv.d.GetRoleInfo(rind)
	var msg string
	if side == 0 {
		msg = fmt.Sprintf("角色 %s 未选择出战将领, 无法开战", name)
	} else {
		msg = fmt.Sprintf("角色 %s 无将领出战, 自动战败", name)
	}
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}

// sst: 发动方状态, ost: 被动方状态, max_bt: 最大回合数
func (uv *GtkUserView) BattleStart(sst *mbg.BattleState, ost *mbg.BattleState, max_bt int) { // 战斗开始
	C.battle_start(C.int(max_bt))
	var raname, rdname string
	var tlist []int

	role, power, winp, celve, hpmax, hp, def, _, _, _, _, _, _, _, _, tlist := sst.GetInfo()
	if role == -1 {
		raname = "山贼"
	} else {
		raname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = uv.d.GetRoleInfo(role)
	}
	cstr := C.CString(raname)
	C.battle_set(0, cstr, C.float(power), C.float(winp), C.float(celve), C.float(hpmax), C.float(hp), C.float(def))
	C.free(unsafe.Pointer(cstr))
	for _, i := range tlist {
		C.battle_tech_set(0, C.int(i))
	}

	role, power, winp, celve, hpmax, hp, def, _, _, _, _, _, _, _, _, tlist = ost.GetInfo()
	if role == -1 {
		rdname = "山贼"
	} else {
		rdname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = uv.d.GetRoleInfo(role)
	}
	cstr = C.CString(rdname)
	C.battle_set(1, cstr, C.float(power), C.float(winp), C.float(celve), C.float(hpmax), C.float(hp), C.float(def))
	C.free(unsafe.Pointer(cstr))
	for _, i := range tlist {
		C.battle_tech_set(1, C.int(i))
	}
	cstr = C.CString(fmt.Sprintf("战斗开始. 攻方: %s, 守方 %s", raname, rdname))
	C.battle_msg(cstr)
	C.free(unsafe.Pointer(cstr))
}

func (uv *GtkUserView) BattleTurnStart(iturn int) { // 战斗回合开始. iturn: 回合数 (0 起始)
	C.battle_turn_start(C.int(iturn))
	// cstr := C.CString(fmt.Sprintf("第 %d 回合", iturn))
	// C.battle_msg (cstr)
	// C.free (unsafe.Pointer(cstr))
}
func (uv *GtkUserView) BattleAttack(ls int, ost *mbg.BattleState) {
	_, _, _, _, _, hp, def, _, ddef, dhp, _, _, _, _, _, _ := ost.GetInfo()
	var sss string
	if ls == 0 {
		sss = "攻方"
	} else {
		sss = "守方"
	}
	C.battle_change_hp(C.int(ls), C.float(hp), C.float(def))
	cstr := C.CString(fmt.Sprintf("%s 损失 %6.0f 城防, %6.0f 兵力.", sss, -ddef, -dhp))
	C.battle_msg(cstr)
	C.free(unsafe.Pointer(cstr))
}

func btoi(value bool) int {
	if value {
		return 1
	} else {
		return 0
	}
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

	C.battle_active_tech(C.int(ss), C.int(tind), C.float(shp), C.float(ohp), C.float(sdef), C.float(odef))
	cstr := C.CString(msg)
	C.battle_msg(cstr)
	C.free(unsafe.Pointer(cstr))
	C.battle_change_st(C.int(ss), C.int(sbst), C.int(btoi(sisq)), C.int(sbtime), C.int(sfst))
	C.battle_change_st(C.int(1-ss), C.int(obst), C.int(btoi(oisq)), C.int(obtime), C.int(ofst))
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
	C.battle_change_hp(C.int(side), C.float(hp), C.float(def))
	cstr := C.CString(msg)
	C.battle_msg(cstr)
	C.free(unsafe.Pointer(cstr))
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
	C.battle_change_hp(C.int(side), C.float(hp), C.float(def))
	cstr := C.CString(msg)
	C.battle_msg(cstr)
	C.free(unsafe.Pointer(cstr))
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
	cstr := C.CString(msg)
	C.battle_msg(cstr)
	C.free(unsafe.Pointer(cstr))
	C.battle_end()
}

type GtkOperationInterface struct {
	d    *mbg.Driver
	rind int
}

func (oi *GtkOperationInterface) StartGame(d *mbg.Driver, rind int) { // 游戏开始
	oi.d = d
	oi.rind = rind
}

func ShowPeopleInfo(d *mbg.Driver, pind int,
	ui_show_func func(name *C.char, role *C.char, job C.int, loc C.int, hpmax C.float, hp C.float, prop *C.float, is_quit C.int, lst C.int, pst C.int),
) {
	pname, role, job, loc, hpmax, hp, prop, is_quit, lst, pst := d.GetPeopleInfo(pind)
	var rname string
	if role != -1 {
		rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = d.GetRoleInfo(role)
	} else {
		rname = "在野"
	}
	cs1 := C.CString(pname)
	cs2 := C.CString(rname)
	var prop_c [5]C.float
	for i := 0; i < 5; i++ {
		prop_c[i] = C.float(prop[i])
	}
	ui_show_func(cs1, cs2, C.int(job), C.int(loc), C.float(hpmax), C.float(hp), &(prop_c[0]), C.int(btoi(is_quit)), C.int(lst), C.int(pst))
	C.free(unsafe.Pointer(cs1))
	C.free(unsafe.Pointer(cs2))
}

func ShowCityInfo(d *mbg.Driver, cind int,
	ui_show_func func(name *C.char, scope C.int, hpmax C.float, hp C.float, role *C.char, nmos C.int, mayor *C.char, treasurer *C.char, fengshui C.float),
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
	cs1 := C.CString(cname)
	cs2 := C.CString(rname)
	cs3 := C.CString(mname)
	cs4 := C.CString(trea_name)
	ui_show_func(cs1, C.int(scope), C.float(hpmax), C.float(hp), cs2, C.int(len(mos)), cs3, cs4, C.float(fengshui))
	C.free(unsafe.Pointer(cs1))
	C.free(unsafe.Pointer(cs2))
	C.free(unsafe.Pointer(cs3))
	C.free(unsafe.Pointer(cs4))
}

func ShowTechInfo(hind int,
	ui_show_tech func(name *C.char, study C.float, scond C.int),
) {
	tname, study, scond, _, _, _, _, _ := mbg.GetTechInfo(hind)
	cstr := C.CString(tname)
	ui_show_tech(cstr, C.float(study), C.int(scond))
	C.free(unsafe.Pointer(cstr))
}

func (oi *GtkOperationInterface) Check() {
	d := oi.d

	var class int

	_, nrole, npeople, ncity, ninst, ntrain, _, _, _ := d.GetInfo()

	class = int(C.get_check_type())
	switch class {
	case 0:
		C.begin_role_list()
		for i := 0; i < nrole; i++ {
			name, loc, _, _, _, _, _, money, dir, mst, mtime, cst, ast, atime, _ := d.GetRoleInfo(i)
			cstr := C.CString(name)
			C.add_role_list(cstr, C.int(loc), C.float(money), C.int(dir), C.int(mst), C.int(mtime), C.int(btoi(cst)), C.int(ast), C.int(atime))
			C.free(unsafe.Pointer(cstr))
		}
		C.end_role_list()

		for {
			rind := int(C.get_role_select())
			if rind == -1 {
				break
			}
			cind := int(C.get_role_mcs_select())
			if cind == -1 {
				_, _, tech, mos, mcs, sos, tos, _, _, _, _, _, _, _, cards := d.GetRoleInfo(rind)

				C.begin_role_tech_list()
				for _, i := range tech {
					C.add_role_tech_list(C.int(i))
				}
				C.end_role_tech_list()

				C.begin_role_mos_list()
				for _, i := range mos {
					ShowPeopleInfo(d, i, add_role_mos_list_g)
				}
				C.end_role_mos_list()

				C.begin_role_mcs_list()
				for _, i := range mcs {
					ShowCityInfo(d, i, add_role_mcs_list_g)
				}
				C.end_role_mcs_list()

				C.begin_role_sos_list()
				for _, i := range sos {
					ShowPeopleInfo(d, i, add_role_sos_list_g)
				}
				C.end_role_sos_list()

				C.begin_role_tos_list()
				for _, i := range tos {
					ShowPeopleInfo(d, i, add_role_tos_list_g)
				}
				C.end_role_tos_list()

				C.begin_role_cards_list()
				for i := 0; i < mbg.CARD_COUNT; i++ {
					C.add_role_cards_list(C.int(i), C.int(cards[i]))
				}
				C.end_role_cards_list()
			} else {
				_, _, _, _, mcs, _, _, _, _, _, _, _, _, _, _ := d.GetRoleInfo(rind)
				cname, _, _, _, _, mos, _, _, _, _, _ := d.GetCityInfo(mcs[cind])
				cstr := C.CString(cname + "官员")
				C.begin_people_list(cstr)
				C.free(unsafe.Pointer(cstr))
				for _, i := range mos {
					ShowPeopleInfo(d, i, add_people_list_g)
				}
				C.end_people_list(1)
				C.get_people_exit()
			}
		}
	case 1:
		cstr := C.CString("全体人员")
		C.begin_people_list(cstr)
		C.free(unsafe.Pointer(cstr))
		for i := 0; i < npeople; i++ {
			ShowPeopleInfo(d, i, add_people_list_g)
		}
		C.end_people_list(0)
		C.get_people_exit()
	case 2:
		C.begin_city_list()
		for i := 0; i < ncity; i++ {
			ShowCityInfo(d, i, add_city_list_g)
		}
		C.end_city_list()

		for {
			cind := int(C.get_city_mcs_select())
			if cind == -1 {
				break
			}
			cname, _, _, _, _, mos, _, _, _, _, _ := d.GetCityInfo(cind)
			cstr := C.CString(cname + "官员")
			C.begin_people_list(cstr)
			C.free(unsafe.Pointer(cstr))
			for _, i := range mos {
				ShowPeopleInfo(d, i, add_people_list_g)
			}
			C.end_people_list(2)
			C.get_people_exit()
		}
	case 3:
		C.begin_inst_list()
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
			cs1 := C.CString(iname)
			cs2 := C.CString(mos_name)
			cs3 := C.CString(study_name)
			C.add_inst_list(cs1, C.int(len(tech)), cs2, cs3, C.float(point), C.int(left_round))
			C.free(unsafe.Pointer(cs1))
			C.free(unsafe.Pointer(cs2))
			C.free(unsafe.Pointer(cs3))
		}
		C.end_inst_list()

		for {
			iind := int(C.get_inst_select())
			if iind == -1 {
				break
			}
			iname, tech, mos, _, _, _ := d.GetInstInfo(iind, oi.rind)
			if mos != -1 {
				ShowPeopleInfo(d, mos, show_inst_mos_g)
			}
			cstr := C.CString(iname)
			C.begin_inst_tech_list(cstr)
			C.free(unsafe.Pointer(cstr))
			for _, i := range tech {
				ShowTechInfo(i, add_inst_tech_list_g)
			}
			C.end_inst_tech_list()
		}
	case 4:
		C.begin_train_list()
		for i := 0; i < ntrain; i++ {
			tname, item, mos, round := d.GetTrainInfo(i, oi.rind)
			var cs1, cs2 *C.char
			cs1 = C.CString(tname)
			if mos != -1 {
				pname, _, _, _, _, _, _, _, _, _ := d.GetPeopleInfo(mos)
				cs2 = C.CString(pname)
			} else {
				cs2 = C.CString("无")
			}
			C.add_train_list(cs1, C.int(item), cs2, C.int(round))
			C.free(unsafe.Pointer(cs1))
			C.free(unsafe.Pointer(cs2))
		}
		C.end_train_list()
		for {
			tind := int(C.get_train_select())
			if tind == -1 {
				break
			}
			_, _, mos, _ := d.GetTrainInfo(tind, oi.rind)
			if mos != -1 {
				ShowPeopleInfo(d, mos, show_train_mos_g)
			}
		}
	case 5:
		msg := C.CString("人员查询")
		C.begin_people_list(msg)
		C.free(unsafe.Pointer(msg))
		for i := 0; i < npeople; i++ {
			_, role, _, _, _, _, _, _, _, _ := d.GetPeopleInfo(i)
			if role == oi.rind {
				ShowPeopleInfo(d, i, add_people_list_g)
			}
		}

		C.end_people_list(0)
		C.get_people_exit()
	}
}

func (oi *GtkOperationInterface) AlignFail(obj_role int) { // 设置结盟失败
	cstr := C.CString("无法结盟")
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}
func (oi *GtkOperationInterface) PEFail(rind int) { // 设置公敌失败
	cstr := C.CString("无法宣战")
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}

func (oi *GtkOperationInterface) NoteAsMayer(pind int) { // 某人就职市长
	var msg string
	pname, _, _, _, _, _, _, _, _, _ := oi.d.GetPeopleInfo(pind)
	msg = fmt.Sprintf("%s 就任市长", pname)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}
func (oi *GtkOperationInterface) NoteAsTreasurer(pind int) { // 某人就职财务官
	var msg string
	pname, _, _, _, _, _, _, _, _, _ := oi.d.GetPeopleInfo(pind)
	msg = fmt.Sprintf("%s 就任财务官", pname)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}
func (oi *GtkOperationInterface) NoteAsCitizen(pind int) { // 某人撤销市政职务
	var msg string
	pname, _, _, _, _, _, _, _, _, _ := oi.d.GetPeopleInfo(pind)
	msg = fmt.Sprintf("%s 被撤销市政职务", pname)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}

func (oi *GtkOperationInterface) NoteFinishStudy(iind int, pind int, tech int) { // 研发完成
	rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := oi.d.GetRoleInfo(oi.rind)
	pname, _, _, _, _, _, _, _, _, _ := oi.d.GetPeopleInfo(pind)
	hname, _, _, _, _, _, _, _ := mbg.GetTechInfo(tech)
	var msg string
	msg = fmt.Sprintf("%s 为 %s 研究策略 %s 完成.", pname, rname, hname)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}
func (oi *GtkOperationInterface) NoteLYStudy(iind int, pind int) { // 研发因流言无法继续
	iname, _, _, _, _, _ := oi.d.GetInstInfo(iind, 0)
	pname, _, _, _, _, _, _, _, _, _ := oi.d.GetPeopleInfo(pind)
	var msg string
	msg = fmt.Sprintf("%s 处于流言状态, 暂停位于 %s 的研究工作", pname, iname)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
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

func (oi *GtkOperationInterface) NoteFinishTrain(pind int, item mbg.Property, value float32) { // 训练完成
	pname, _, _, _, _, _, _, _, _, _ := oi.d.GetPeopleInfo(pind)
	var msg string
	msg = fmt.Sprintf("%s 完成 %s 的训练, 能力提升 %6.0f", pname, GetPropName(item), value)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}
func (oi *GtkOperationInterface) NoteLYTrain(tind int, pind int) { // 训练因流言无法继续
	pname, _, _, _, _, _, _, _, _, _ := oi.d.GetPeopleInfo(pind)
	_, item, _, _ := oi.d.GetTrainInfo(tind, oi.rind)
	var msg string
	msg = fmt.Sprintf("%s 关于 %s 的训练因流言暂停", pname, GetPropName(item))
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}
func (oi *GtkOperationInterface) NoteLYMayor(pind, cind int) { // 因流言无法履行市长职责
	pname, _, _, _, _, _, _, _, _, _ := oi.d.GetPeopleInfo(pind)
	cname, _, _, _, _, _, _, _, _, _, _ := oi.d.GetCityInfo(cind)
	var msg string
	msg = fmt.Sprintf("%s 因流言无法履行 %s 的市长职责", pname, cname)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}
func (oi *GtkOperationInterface) NoteLYTreasurer(pind, cind int) { // 因流言无法履行财务官职责
	pname, _, _, _, _, _, _, _, _, _ := oi.d.GetPeopleInfo(pind)
	cname, _, _, _, _, _, _, _, _, _, _ := oi.d.GetCityInfo(cind)
	var msg string
	msg = fmt.Sprintf("%s 因流言无法履行 %s 的财务官职责", pname, cname)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}
func (oi *GtkOperationInterface) NoteCityRecover(cind int, hplus float32) { // 城市因市长恢复所属人员 HP
	cname, _, _, _, _, _, _, _, _, _, _ := oi.d.GetCityInfo(cind)
	var msg string
	msg = fmt.Sprintf("城市 %s 恢复兵力 %6.0f", cname, hplus)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}
func (oi *GtkOperationInterface) NoteCityEarn(cind int, mplus float32) { // 城市获得收入
	cname, _, _, _, _, _, _, _, _, _, _ := oi.d.GetCityInfo(cind)
	var msg string
	msg = fmt.Sprintf("城市 %s 获得收入 %6.0f", cname, mplus)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}

func (oi *GtkOperationInterface) NotePayInst(dm float32) { // 支付研究经费
	rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := oi.d.GetRoleInfo(oi.rind)
	var msg string
	msg = fmt.Sprintf("%s 支付研究经费 %6.0f", rname, -dm)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}

func (oi *GtkOperationInterface) SkipBattleByAlign(robj int) { // 因为结盟跳过战斗
	srname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := oi.d.GetRoleInfo(oi.rind)
	orname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := oi.d.GetRoleInfo(robj)
	var msg string
	msg = fmt.Sprintf("%s 与 %s 因结盟无法战斗", srname, orname)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}
func (oi *GtkOperationInterface) ForceBattle(robj int) { // 因为公愤进行强制战斗
	srname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := oi.d.GetRoleInfo(oi.rind)
	orname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := oi.d.GetRoleInfo(robj)
	var msg string
	msg = fmt.Sprintf("%s 与 %s 因宣战状态强制发生战斗", srname, orname)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}

func (oi *GtkOperationInterface) SelRoleAction(clist []int) int { // 选择阵营行动: -1: 移动, >=0: 使用卡片
	for {
		action := int(C.get_action(C.int(oi.rind)))
		switch action {
		case -1:
			return -2
		case 0:
			return -1
		case 1:
			_, _, _, _, _, _, _, _, _, _, _, _, _, _, cards := oi.d.GetRoleInfo(oi.rind)
			C.begin_cards_list()
			for i := 0; i < mbg.CARD_COUNT; i++ {
				C.add_cards_list(C.int(i), C.int(cards[i]))
			}
			C.end_cards_list()
			sc := int(C.get_card_action())
			if sc < 0 {
				continue
			}
			return sc
		case 2:
			oi.Check()
		}
	}
}

func (oi *GtkOperationInterface) IsOccuCity(cind int) bool { // 选择是否占领空白城市
	rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := oi.d.GetRoleInfo(oi.rind)
	cname, _, _, _, _, _, _, _, _, _, _ := oi.d.GetCityInfo(cind)
	var msg string
	msg = fmt.Sprintf("%s 是否占领城市 %s", rname, cname)
	for {
		cstr := C.CString(msg)
		res := int(C.msg_box(cstr, C.int(1)))
		C.free(unsafe.Pointer(cstr))
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
	var msg string
	msg = fmt.Sprintf("%s 占领城市 %s 失败", rname, cname)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}

func (oi *GtkOperationInterface) SelCityMos(splist []int, cind int) []int { // 选择占领的人员
	cstr := C.CString("选择占领人员")
	C.begin_people_list(cstr)
	C.free(unsafe.Pointer(cstr))
	for _, i := range splist {
		ShowPeopleInfo(oi.d, i, add_people_list_g)
	}
	C.end_people_list(0)
	nsel := int(C.get_people_list_num())
	var res []int = make([]int, nsel)
	for i := 0; i < nsel; i++ {
		res[i] = int(C.get_people_list(C.int(i)))
	}
	return res
}

func (oi *GtkOperationInterface) SelMayor(cind int, splist []int, old_mayor int, old_treasurer int) (mayor int) { // 选择市长
	cname, _, _, _, _, _, _, _, _, _, _ := oi.d.GetCityInfo(cind)
	cstr := C.CString(fmt.Sprintf("为 %s 选择市长", cname))
	C.begin_people_list(cstr)
	C.free(unsafe.Pointer(cstr))
	for _, i := range splist {
		ShowPeopleInfo(oi.d, i, add_people_list_g)
	}
	C.end_people_list(0)
	return int(C.get_people_one())
}

func (oi *GtkOperationInterface) SelTreasurer(cind int, splist []int, old_mayor int, old_treasurer int) (treasurer int) { // 选择财务官
	cname, _, _, _, _, _, _, _, _, _, _ := oi.d.GetCityInfo(cind)
	cstr := C.CString(fmt.Sprintf("为 %s 选择财务官", cname))
	C.begin_people_list(cstr)
	C.free(unsafe.Pointer(cstr))
	for _, i := range splist {
		ShowPeopleInfo(oi.d, i, add_people_list_g)
	}
	C.end_people_list(0)
	return int(C.get_people_one())
}

func (oi *GtkOperationInterface) IsAttackCity(cind int) int { // 选择是否攻城 -1: 不攻城 0-2: 三种攻城规模
	cname, _, _, _, _, _, _, _, _, _, _ := oi.d.GetCityInfo(cind)
	var msg string
	msg = fmt.Sprintf("是否进攻城市 %s? 请选择发动战斗的规模.", cname)
	for {
		cstr := C.CString(msg)
		res := int(C.sel_battle_scale(cstr))
		C.free(unsafe.Pointer(cstr))
		if res == -2 {
			oi.Check()
			msg = ""
			continue
		}
		return res
	}
}

func add_saloon_list_functor(price float32) (
	res_func func(name *C.char, role *C.char, job C.int, loc C.int, hpmax C.float, hp C.float, prop *C.float, is_quit C.int, lst C.int, pst C.int),
) {
	return func(name *C.char, role *C.char, job C.int, loc C.int, hpmax C.float, hp C.float, prop *C.float, is_quit C.int, lst C.int, pst C.int) {
		C.add_saloon_list(name, role, job, loc, hpmax, hp, prop, is_quit, lst, pst, C.float(price))
	}
}

func (oi *GtkOperationInterface) SaloonSelPeople(pind []int, price []float32) int { // 选择招募人员 : -1: 放弃
	C.begin_saloon_list()
	for i := 0; i < len(pind); i++ {
		ShowPeopleInfo(oi.d, pind[i], add_saloon_list_functor(price[i]))
	}
	C.end_saloon_list()

	return int(C.get_saloon_select())
}

func (oi *GtkOperationInterface) RecruitFail(pind int) { // 因无法支付招募失败
	var msg string
	msg = fmt.Sprintf("金币不足, 招募失败")
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}

func (oi *GtkOperationInterface) SelStudy(iind int, plist []int, tlist []int) (sp int, st int) { // 选择研究人员和项目, (-1, -1): 放弃研究
	iname, _, _, _, _, _ := oi.d.GetInstInfo(iind, oi.rind)
	cstr := C.CString(iname)
	C.begin_study_tech_list(cstr)
	C.free(unsafe.Pointer(cstr))
	for _, i := range tlist {
		ShowTechInfo(i, add_study_tech_list_g)
	}
	C.end_study_tech_list()
	C.begin_study_people_list()
	for _, i := range plist {
		ShowPeopleInfo(oi.d, i, add_study_people_list_g)
	}
	C.end_study_people_list()

	st_c := C.int(st)
	sp = int(C.get_study_select(&st_c))
	st = int(st_c)
	if sp == -1 || st == -1 {
		sp = -1
		st = -1
	}
	return
}

func (oi *GtkOperationInterface) SelTrain(tind int, plist []int, item mbg.Property) int { // 选择训练人员, 返回人员编号, -1: 放弃
	tname, _, _, _ := oi.d.GetTrainInfo(tind, oi.rind)
	var msg string = fmt.Sprintf("%s: 选择训练人员. 训练项目: %s", tname, GetPropName(item))
	cstr := C.CString(msg)
	C.begin_people_list(cstr)
	C.free(unsafe.Pointer(cstr))
	for _, i := range plist {
		ShowPeopleInfo(oi.d, i, add_people_list_g)
	}
	C.end_people_list(0)

	return int(C.get_people_one())
}

func (oi *GtkOperationInterface) BattleWithRole(obj_role int) int { // 是否打遭遇战. 返回: -1: 不打. 0-2: 战斗规模
	var msg string
	rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := oi.d.GetRoleInfo(obj_role)
	msg = fmt.Sprintf("是否与 %s 发生战斗? 请选择战斗规模", rname)
	for {
		cstr := C.CString(msg)
		res := int(C.sel_battle_scale(cstr))
		C.free(unsafe.Pointer(cstr))
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
		cstr := C.CString(msg)
		res := int(C.msg_box(cstr, C.int(1)))
		C.free(unsafe.Pointer(cstr))
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
	msg := fmt.Sprintf("请选择卡片 %s 的作用城市", dname)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))

	for {
		obj_type := int(C.get_obj_type())
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
		loc := int(C.get_obj())

		class, base, _, _, _ := oi.d.GetBoardInfo(loc)
		if class != 1 {
			continue
		}
		return base
	}
}

func (oi *GtkOperationInterface) SelCardObjAny(card_id int) (st int, obj int) { // 返回卡片作用对象建筑或角色. st: -1: 取消. 0: 角色 1: 城市: 2: 研究所 3: 修炼所
	var msg string
	dname, _, _ := mbg.GetCardInfo(card_id)
	msg = fmt.Sprintf("请选择卡片 %s 的作用官员", dname)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))

	for {
		obj_type := int(C.get_obj_type())
		switch obj_type {
		case -1:
			return -1, -1
		case -2:
			loc := int(C.get_obj())
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
	var msg string
	dname, _, _ := mbg.GetCardInfo(card_id)
	msg = fmt.Sprintf("请选择卡片 %s 的作用官员", dname)
	cstr := C.CString(msg)
	C.begin_people_list(cstr)
	C.free(unsafe.Pointer(cstr))
	for _, i := range plist {
		ShowPeopleInfo(oi.d, i, add_people_list_g)
	}
	C.end_people_list(0)
	return int(C.get_people_one())
}

func (oi *GtkOperationInterface) SelCardObjRoleAndPeople(card_id int, rplist []int) (sr_id int) { // 返回卡片作用人员  -1:取消

	var msg string
	dname, _, _ := mbg.GetCardInfo(card_id)
	msg = fmt.Sprintf("请选择卡片 %s 的作用官员", dname)
	cstr := C.CString(msg)
	C.begin_people_list(cstr)
	C.free(unsafe.Pointer(cstr))
	for _, i := range rplist {
		ShowPeopleInfo(oi.d, i, add_people_list_g)
	}
	C.end_people_list(0)
	return int(C.get_people_one())
}

func (oi *GtkOperationInterface) SelCardObjRole(card_id int) int { // 返回选择卡片作用阵营 -1: 取消
	var msg string
	dname, _, _ := mbg.GetCardInfo(card_id)
	msg = fmt.Sprintf("请选择卡片 %s 的作用阵营", dname)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))

	for {
		obj_type := int(C.get_obj_type())
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
	var msg string
	dname, _, _ := mbg.GetCardInfo(card_id)
	msg = fmt.Sprintf("请选择卡片 %s 的作用地图位置", dname)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))

	for {
		obj_type := int(C.get_obj_type())
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
		return int(C.get_obj())
	}
}

func (oi *GtkOperationInterface) SelAllocObj() (st int, obj int) { // 选择调度一级目标. st: -1: 取消. 0: 城市: 1: 研究所 2: 修炼所
	var msg string
	msg = fmt.Sprintf("请在地图上选择调度对象 (城市, 研究所, 训练所)")
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))

	for {
		obj_type := int(C.get_obj_type())
		if obj_type == -1 {
			return -1, -1
		}
		if obj_type == -3 {
			oi.Check()
			continue
		}
		if obj_type >= 0 {
			continue
		}
		loc := int(C.get_obj())
		class, base, _, _, _ := oi.d.GetBoardInfo(loc)
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

func make_add_exchange_func(side int) func(name *C.char, role *C.char, job C.int, loc C.int, hpmax C.float, hp C.float, prop *C.float, is_quit C.int, lst C.int, pst C.int) {
	return func(name *C.char, role *C.char, job C.int, loc C.int, hpmax C.float, hp C.float, prop *C.float, is_quit C.int, lst C.int, pst C.int) {
		C.add_exchange_people_list(C.int(side), name, role, job, loc, hpmax, hp, prop, is_quit, lst, pst)
	}
}
func (oi *GtkOperationInterface) ExchangeCityPeople(cind int, splist []int, oplist []int) (in_list []int, out_list []int) { // 调度随身人员与目标内驻扎人员.
	rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := oi.d.GetRoleInfo(oi.rind)
	cname, _, _, _, _, _, _, _, _, _, _ := oi.d.GetCityInfo(cind)
	msg := fmt.Sprintf("%s 阵营调度城市 %s 的官员", rname, cname)

	for k := 0; k < 2; k++ {
		cstr := C.CString(msg)
		C.begin_exchange_people_list(C.int(k), cstr)
		C.free(unsafe.Pointer(cstr))
		var plist []int
		if k == 0 {
			plist = splist
		} else {
			plist = oplist
		}
		for _, i := range plist {
			ShowPeopleInfo(oi.d, i, make_add_exchange_func(k))
		}
		C.end_exchange_people_list(C.int(k))
	}

	var ret_list [2]([]int)
	for k := 0; k < 2; k++ {
		nsel := int(C.get_exchange_people_list_num(C.int(k)))
		ret_list[k] = make([]int, nsel)
		for i := 0; i < nsel; i++ {
			ret_list[k][i] = int(C.get_exchange_people_list(C.int(k), C.int(i)))
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
		cstr := C.CString(msg)
		res := int(C.msg_box(cstr, C.int(1)))
		C.free(unsafe.Pointer(cstr))
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
		cstr := C.CString(msg)
		res := int(C.msg_box(cstr, C.int(1)))
		C.free(unsafe.Pointer(cstr))
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

func (oi *GtkOperationInterface) StartAllocate() { // 开始人员调度
	var msg string
	rname, _, _, _, _, _, _, _, _, _, _, _, _, _, _ := oi.d.GetRoleInfo(oi.rind)
	msg = fmt.Sprintf("%s 开始人员调度回合", rname)
	cstr := C.CString(msg)
	C.msg_map(cstr)
	C.free(unsafe.Pointer(cstr))
}

func (oi *GtkOperationInterface) SelGeneral(plist []int) (main, vice int) { // 选择战争将领: main: 主将. vice: 副将 -1: 无副将, main == vice means vice = -1
	var msg string
	msg = fmt.Sprintf("选择出战主将与参军")

	cstr := C.CString(msg)
	C.begin_people_list(cstr)
	C.free(unsafe.Pointer(cstr))
	for _, i := range plist {
		ShowPeopleInfo(oi.d, i, add_people_list_g)
	}
	C.end_people_list(0)

	main = int(C.get_people_two_main())
	vice = int(C.get_people_two_vice())
	return
}
