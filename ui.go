package mbg

// loc 为地图位置
// role, rind 为阵营编号, pind 为武将编号, iind 为研究所编号, tind 为训练所编号, cind 为城市编号, 在此统一说明

type UserView interface {

      StartGame (d * Driver) // 游戏开始

      SetBarrier (loc int) // 设置路障
      UnSetBarrier (loc int) // 路障生效并去除
      SetRobber (loc int) // 设置山贼
      UnSetRobber (loc int) // 山贼生效并去除
      SetSlow (rind int) // 慢行
      SetFast (rind int) // 急行
      SetStop (rind int) // 禁行
      SetContinue (rind int) // 双行
      ChangeDir (rind int) // 设置反向
      StealMoney (sub_role int, obj_role int, money float32) // 偷钱
      SelfRecover (rind int) // 自我恢复
      SetLiuYan (pind int) // 设置流言状态
      SetPoison (pind int) // 设置中毒状态
      SetQuit (pind int) // 设置离职状态
      SetAlign (sub_role int, obj_role int) // 设置结盟
      SetPE (role int) // 设置公敌

      // this section are called before game start (d is not available)
      InitMap (ngrid, nrole, npeople, ncity, ninst, ntrain, TECH_COUNT, CARD_COUNT, nx, ny int) // 初始化地图
      ShowMap (id int, class int, base int, x int, y int) // 显示地图块
      ShowCity (id int, name string, scope int, x0, y0, x1, y1 int) // 显示城市 
      ShowInstitute (id int, name string, x0, y0, x1, y1 int) // 显示策略研究所
      ShowTrainingRoom (id int, name string, item Property, x0, y0, x1, y1 int) // 显示训练所

      NoteInCity (pind int, new_city int) // 某人进驻己方城市
      NoteInInst (pind int, new_inst int) // 某人进驻研究所
      NoteInTrain (pind int, new_train int) // 某人进驻训练所
      NoteOutCity (pind int, cind int) // 某人撤离己方城市
      NoteOutInst (pind int, iind int) // 某人撤离研究所
      NoteOutTrain (pind int, tind int) // 某人撤离训练所
      NoteJoin (pind int, rind int) // 某人加入某阵营
      NoteQuit (pind int, rind int) // 某人离开某阵营
      NotePayTax (send_role int, recv_role int, recv_city int, tax float32) // 某人向某城支付租金
      NoteCollectTax (rind int, cind int, tax float32) // 某角色到达自身城市收取税收

      RoleContinueAction (rind int) // 某阵营连续行动 (在 RoleStartAction 之前调用)
      RoleStartAction (rind int) // 某阵营开始行动
      RoleForbidAction (rind int) // 某阵营被禁止行动

      NoteRoleOccupyCity (rind int, cind int) // 某阵营占领城市
      NoteRoleDropCity  (rind int, cind int) // 某阵营放弃城市

      NoteRoleStEnd (rind int, st int) // 阵营特殊行动状态终止
      NoteRoleCeEnd (rind int)         // 阵营公敌状态终止
      NoteRoleAlignEnd (r1 int, r2 int) // 阵营间结盟状态终止 (r1 < r2)

      NoteRoleWin (rind int) // 阵营胜利
      NoteRoleLoss (rind int) // 阵营破产
      NoteAllLoss () // 所有阵营同时破产 
      NoteForceQuit () // 用户强制退出

      NoteRoleMoveOneStep (rind int, sloc int, dir int) // 角色移动一步
      NoteRoleMove (rind int, sloc int, oloc int, step int) // 角色移动一个回合
      NoteGetCard (rind int, card_id int) // 角色获得卡片

      EndTurn (turn int) // 回合结束

      BattleNoGerenal (rind int, side int) // 无主将参与作战, 自动战败

      // sst: 发动方状态, ost: 被动方状态, max_bt: 最大回合数
      BattleStart (sst * BattleState, ost * BattleState, max_bt int) // 战斗开始
      BattleTurnStart (iturn int) // 战斗回合开始. iturn: 回合数 (0 起始)
      BattleAttack (ls int, ost * BattleState)
      BattleTech (ss int, tind int, sst * BattleState, ost * BattleState) // 战斗时发动技能. ss: 0|1 发动方
      BattleBurn (side int, st * BattleState) // 战斗中被烧
      BattleSelfAttack (side int, st * BattleState) // 战斗中受到内讧伤害
      BattleEnd  (ret int) // 战斗结束: ret: 战斗结果. 0: 和平结束. -1: 攻方失败. 1: 攻方取胜.
}
// 52 view function

type OperationInterface interface {

      StartGame (d * Driver, rind int) // 游戏开始

      AlignFail(obj_role int) // 设置结盟失败
      PEFail (rind int) // 设置公敌失败 

      NoteAsMayer    (pind int) // 某人就职市长
      NoteAsTreasurer (pind int) // 某人就职财务官
      NoteAsCitizen  (pind int) // 某人撤销市政职务

      NoteFinishStudy (iind int, pind int, tech int) // 研发完成
      NoteLYStudy (iind int, pind int) // 研发因流言无法继续
      NoteFinishTrain (pind int, item Property, value float32) // 训练完成
      NoteLYTrain (tind int, pind int) // 训练因流言无法继续
      NoteLYMayor (pind, cind int) // 因流言无法履行市长职责
      NoteLYTreasurer (pind, cind int) // 因流言无法履行财务官职责
      NoteCityRecover (cind int, hplus float32) // 城市因市长恢复所属人员 HP
      NoteCityEarn (cind int, mplus float32) // 城市获得收入
      NotePayInst (dm float32) // 支付研究经费

      SkipBattleByAlign (robj int) // 因为结盟跳过战斗
      ForceBattle (robj int) // 因为公愤进行强制战斗

      SelRoleAction (clist []int) int // 选择阵营行动: -1: 移动, >=0: 使用卡片, -2: 中途退出游戏.
      IsOccuCity (cind int) bool // 选择是否占领空白城市
      OccuCityFail (cind int) // 占领城市失败
      SelCityMos (splist []int, cind int) []int // 选择占领的人员
      SelMayor (cind int, splist []int, old_mayor int, old_treasurer int) int // 选择市长
      SelTreasurer (cind int, splist []int, old_mayor int, old_treasurer int) int // 选择财务官
      IsAttackCity (cind int) int // 选择是否攻城 -1: 不攻城 0-2: 三种攻城规模
      SaloonSelPeople (pind []int , price []float32) int // 选择招募人员 -1: 放弃, 返回 price 中的编号
      RecruitFail (pind int) // 因无法支付招募失败
      SelStudy (iind int, plist []int, tlist []int) (sp int, st int) // 选择研究人员和项目, 返回 plist 和 tlist 中所选择的编号索引, (-1, -1): 放弃研究
      SelTrain (tind int, plist []int, item Property) int // 选择训练人员, 返回人员索引, -1: 放弃
      BattleWithRole (obj_role int) int // 是否打遭遇战. 返回: -1: 不打. 0-2: 战斗规模

      ConfirmCard (card_id int) int // 确认是否使用卡片: 无对象卡片 返回: 0: 确定; -1:取消
      SelCardObjCity (card_id int) int // 返回卡片作用对象城市. -1 取消
      SelCardObjAny (card_id int) (st int, obj int) // 返回卡片作用对象建筑或角色. st: -1: 取消. 0: 角色 1: 城市: 2: 研究所 3: 修炼所
      SelCardObjPeople (card_id int, plist []int) int // 返回卡片作用人员. -1: 取消
      SelCardObjRoleAndPeople (card_id int, rplist []int) (sr_id int) // 返回卡片作用阵营和人员 (每阵营一人供选择) -1:取消
      SelCardObjRole (card_id int) int // 返回选择卡片作用阵营 -1: 取消
      SelCardObjLoc (card_id int) int // 选择卡片作用地图位置 -1: 取消
      SelAllocObj () (st int, obj int) // 选择调度一级目标. st: -1: 取消. 0: 城市: 1: 研究所 2: 修炼所
      ExchangeCityPeople (cind int, splist []int, oplist []int) (in_list []int, out_list []int) // 调度随身人员与目标内驻扎人员. 
      IsCancelStudy (pind int, on_study int, left_point float32) bool // 是否终止研究
      IsCancelTrain (pind int, item Property, left_round int) bool // 是否终止训练
      StartAllocate () // 开始人员调度

      SelGeneral (plist []int) (main, vice int) // 选择战争将领: main: 主将. vice: 副将 -1: 无副将
}

// 39 operational function
