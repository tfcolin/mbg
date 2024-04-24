package mbg

import "gitee.com/tfcolin/dsg"

const (
)

type Tech struct {
	name    string
	study   float32 // 所需研究点数/资金数
	scond   int     // 发动时机 0: 任意 1: 回合胜利 2: 回合失败 3:战斗胜利
	lprob   float32 // 发动概率低限
	hprob   float32 // 发动概率高限
	lvalue  float32 // 效果低限
	hvalue  float32 // 效果高限
	latency int     // 发动后的延迟等待时间
	// dhp: 原始伤害值(负数) value: 技能 value 值; srole: 发动技能的角色, orole: 对方角色;
	do func(sst *BattleState, ost *BattleState, dhp float32, value float32) (sdhp, odhp float32) // 操作函数
}

type BattleState struct {

	role  int // 0: 攻方; 1: 守方
	power float32 // 力量
	winp  float32 // 胜率 
	celve float32 // 谋略
	hpmax float32 // 最大 hp
	hp    float32 // 当前 hp
	def   float32 // 当前城防

	is_die    bool    // 战斗结果: false: 正常. true: 战败
	last_ddef float32 // 上次 def 增值
	last_dhp  float32 // 上次 hp 增值

	latency [TECH_COUNT]int // 策略发动延迟

	bst     int  // 战斗状态: 0: 正常, 1: 混乱, 2: 内讧
	is_quit bool // 战后是否离职
	btime   int  // 战斗状态剩余时间
	fst     int  // 燃烧状态: 剩余回合数. 0: 正常

	tech *dsg.Set // 掌握技能
} // 用于描述战斗的一方状态

var techs [TECH_COUNT]Tech

func GetTechInfo(tech int) (
	name string,
	study float32, // 所需研究点数/资金数
	scond int, // 发动时机 0: 任意 1: 回合胜利 2: 回合失败 3:战斗胜利
	lprob float32, // 发动概率低限
	hprob float32, // 发动概率高限
	lvalue float32, // 效果低限
	hvalue float32, // 效果高限
	latency int, // 发动后的延迟等待时间
) {
	if tech < 0 || tech >= TECH_COUNT {
		return
	}
	name = techs[tech].name
	study = techs[tech].study
	scond = techs[tech].scond
	lprob = techs[tech].lprob
	hprob = techs[tech].hprob
	lvalue = techs[tech].lvalue
	hvalue = techs[tech].hvalue
	latency = techs[tech].latency
	return
}

func (bs *BattleState) GetInfo() (
	role int,
	power float32,
	winp float32,
	celve float32,
	hpmax float32,
	hp float32,
	def float32,

	is_die bool, // 战斗结果: false: 正常. true: 战败
	last_ddef float32, // 上次 def 增值
	last_dhp float32, // 上次 hp 增值

	latency [TECH_COUNT]int, // 策略发动延迟

	bst int, // 战斗状态: 0: 正常, 1: 混乱, 2: 内讧
	is_quit bool, // 战后是否离职
	btime int, // 战斗状态剩余时间
	fst int, // 燃烧状态: 剩余回合数. 0: 正常

	tech []int, // 掌握技能
) {
	role = bs.role
	power = bs.power
	winp = bs.winp
	celve = bs.celve
	hpmax = bs.hpmax
	hp = bs.hp
	def = bs.def
	is_die = bs.is_die
	last_ddef = bs.last_ddef
	last_dhp = bs.last_dhp
	latency = bs.latency
	bst = bs.bst
	is_quit = bs.is_quit
	btime = bs.btime
	fst = bs.fst
	tech = bs.tech.GetAllLabel()
	return
}
