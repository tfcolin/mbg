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
      StealMoney (sub_role int, obj_role int, money float32) // 偷钱, money < 0. 在遇到山贼且未取胜时以及使用偷盗的卡片时会被调用.
      SelfRecover (rind int) // 自我恢复
      SetLiuYan (pind int) // 设置流言状态. 
      SetPoison (pind int) // 设置中毒状态
      SetQuit (pind int) // 设置离职状态
      SetAlign (sub_role int, obj_role int) // 设置结盟
      SetPE (role int) // 设置公敌

      // this section are called before game start (driver is not available)
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
      NotePayTax (send_role int, recv_role int, recv_city int, tax float32) // 某角色向某城支付租金
      NoteCollectTax (rind int, cind int, tax float32) // 某角色到达自身城市收取税收

      RoleContinueAction (rind int) // 某阵营连续行动 (在 RoleStartAction 之前调用)
      RoleStartAction (rind int) // 某阵营开始行动
      RoleForbidAction (rind int) // 某阵营被禁止行动

      NoteRoleOccupyCity (rind int, cind int) // 某阵营占领城市
      NoteRoleDropCity  (rind int, cind int) // 某阵营放弃城市

      NoteRoleStEnd (rind int, st int) // 阵营特殊行动状态终止. st : 该角色准备终止的 mst 属性值.
      NoteRoleCeEnd (rind int)         // 阵营公敌状态终止
      NoteRoleAlignEnd (r1 int, r2 int) // 阵营间结盟状态终止 (r1 < r2)

      NoteRoleWin (rind int) // 阵营胜利
      NoteRoleLoss (rind int) // 阵营破产, 此时 rind 的所有人员, 位置都还保留, 以便于 UI 清理数据.
      NoteAllLoss () // 所有阵营同时破产 
      NoteForceQuit () // 用户强制退出

      NoteRoleMoveOneStep (rind int, sloc int, dir int) // 角色移动一步, sloc: 起始位置. dir: +1/-1: 移动方向
      NoteRoleMove (rind int, sloc int, oloc int, step int) // 角色移动一个回合. sloc: 起始位置. oloc: 终止位置, step: 移动步数 (NoteRoleMove 会在 NoteRoleMoveOneStep 之后被调用)
      NoteGetCard (rind int, card_id int) // 角色获得卡片. card_id: 卡片编号.

      EndTurn (turn int) // 回合结束. turn: 回合数.

      // 无主将参与作战, 自动战败. side: 0: 角色作为战斗发动方. 1: 角色作为战斗接受方. 如果双方均未能指定主将, 则发动方战败并收到该消息.
      BattleNoGerenal (rind int, side int) 

      // sst: 发动方状态, ost: 被动方状态, max_bt: 最大回合数
      BattleStart (sst * BattleState, ost * BattleState, max_bt int) // 战斗开始
      BattleTurnStart (iturn int) // 战斗回合开始. iturn: 回合数 (0 起始)
      BattleAttack (ls int, ost * BattleState) // 正常战斗. ls: 输方 (0: 战斗发动方, 1: 战斗接受方) ost: 输方状态.
      BattleTech (ss int, tind int, sst * BattleState, ost * BattleState) // 战斗时发动技能. ss: 0|1 (发动方|接受方)
      BattleBurn (side int, st * BattleState) // 战斗中被烧. side: 0|1 (发动方|接受方)
      BattleSelfAttack (side int, st * BattleState) // 战斗中受到内讧伤害. side: 0|1 (发动方|接受方)
      BattleEnd  (ret int) // 战斗结束: ret: 战斗结果. 0: 和平结束. -1: 发动方失败. 1: 发动方取胜.
}
// 56 view function

/* 每个角色有属于自己的 OperationInterface */
type OperationInterface interface {

      StartGame (d * Driver, rind int) // 游戏开始, rind 为本 OperationInterface 所代表的角色.

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
      NotePayInst (dm float32) // 支付研究经费, dm < 0: 支付的经费数量.

      SkipBattleByAlign (robj int) // 因为结盟跳过战斗, robj: 结盟阵营.
      ForceBattle (robj int) // 因为公愤进行强制战斗, robj: 发生战斗的对象阵营.

      IsOccuCity (cind int) bool // 选择是否占领空白城市

      /* 选择阵营行动: -1: 移动, >=0: 使用卡片, -2: 中途退出游戏. -3: 保存游戏进度.
      * 引擎会在每个回合每个角色行动前调用本函数, 然后根据用户选择调用后续函数 */ 
      SelRoleAction (clist []int) int 
      GetFileName () string // 返回游戏进度文件名 (返回空串表示取消) ( SelRoleAction = -3 时会紧接着被调用)
	SaveReport (is_success bool) // 保存文件通知 (is_success: 是否保存成功)
      OccuCityFail (cind int) // 占领城市失败
      // 选择占领的人员. splist 为可选择人员的列表. 返回值为由 splist 的下标所构成的列表, 以表示从 splist 中选择了一个子列表.
      SelCityMos (splist []int, cind int) []int 
      /* 选择市长. 
      * splist: 为可选择人员的列表. 
      * old_mayor: 该城市之前的市长.
      * old_treasurer: 该城市之前的财务官.
      * 返回值: splist 的一个下标, 表示玩家从中选择的市长人选.
      */
      SelMayor (cind int, splist []int, old_mayor int, old_treasurer int) int 
      SelTreasurer (cind int, splist []int, old_mayor int, old_treasurer int) int // 选择财务官. 返回值为 splist 的一个下标. 参数含义与选择市长相同.
      IsAttackCity (cind int) int // 选择是否攻城 -1: 不攻城 0-2: 三种攻城规模
      SaloonSelPeople (pind []int , price []float32) int // 选择招募人员 -1: 放弃, 返回 pind 及 price 中的下标.
      RecruitFail (pind int) // 因无法支付招募失败
      SelStudy (iind int, plist []int, tlist []int) (sp int, st int) // 选择研究人员和项目, 返回 plist 和 tlist 中所选择的下标索引, (-1, -1): 放弃研究
      SelTrain (tind int, plist []int, item Property) int // 选择训练人员, 返回从 plist 中选择的下标. -1: 放弃.
      BattleWithRole (obj_role int) int // 是否打遭遇战. 返回: -1: 不打. 0-2: 战斗规模

      ConfirmCard (card_id int) int // 确认是否使用卡片 (针对无操作对象的卡片). 返回: 0: 确定; -1:取消
      SelCardObjCity (card_id int) int // 返回卡片作用对象城市. -1 取消
      // 对于作用于人员的卡片, 会先调用 SelCardObjAny, 再调用 SelCardObjPeople.
      SelCardObjAny (card_id int) (st int, obj int) // 返回卡片作用对象建筑或角色 (用于从某个场所或角色中选择一个雇员的第一步). st: -1: 取消. 0: 角色 1: 城市: 2: 研究所 3: 修炼所
      SelCardObjPeople (card_id int, plist []int) int // 返回卡片作用人员 (plist 中的下标). -1: 取消
      SelCardObjRole (card_id int) int // 返回选择卡片作用阵营 -1: 取消
      SelCardObjLoc (card_id int) int // 选择卡片作用地图位置 -1: 取消
      /* 在调度回合, 会首先调用 StartAllocate, 然后反复调用 SelAllocObj, 并根据 st 选择性的调用
      * ExchangeCityPeople, IsCancelStudy, IsCancelTrain, 直到 st == -1, 则终止本角色的本轮调度 */  
      SelAllocObj () (st int, obj int) // 选择调度一级目标. st: -1: 调度完成. 0: 城市: 1: 研究所 2: 修炼所. obj: 目标编号.
      /* 调度随身人员与城市内驻扎人员. 
      * splist: 调度前角色身边人员
      * oplist: 调度前城市驻扎人员
      * in_list: 从角色身边进入城市驻扎的人员, 取值为 splist 的下标.
      * out_list: 从城市中撤回到角色身边的人员, 取值为 oplist 的下标.
      */
      ExchangeCityPeople (cind int, splist []int, oplist []int) (in_list []int, out_list []int) 
      /* 是否终止研究
      * on_study: 正在研究的技术编号
      * left_point: 正在研究的剩余研究点数
      * 返回: 是否在本调度回合终止研究
      */
      IsCancelStudy (pind int, on_study int, left_point float32) bool 
      /* 是否终止训练
      * item: 正在训练的人员属性.
      * left_round: 距离训练完成的剩余回合数.
      * 返回: 是否在本调度回合终止训练.
      */
      IsCancelTrain (pind int, item Property, left_round int) bool 
	// 开始人员调度, 通知 UI 哪些位置是可用的 (av_loc [ngrid]bool: 位置是否可被调度)
      StartAllocate (av_loc []bool) 

      SelGeneral (plist []int) (main, vice int) // 选择战争将领: main: 主将. vice: 副将. 取值为 plist 中的下标. -1: 表示放弃选择. 
}

// 43 operational function
