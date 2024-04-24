package mbg

import (
	"fmt"
	"io"

	"gitee.com/tfcolin/dsg"
)

type Role struct {
	name  string
	loc   int      // -1: 不出场
	tech  *dsg.Set // 掌握技术
	mos   *dsg.Set // 随身人员
	mcs   *dsg.Set // 占领城市
	sos   *dsg.Set // 研究人员
	tos   *dsg.Set // 训练人员
	money float32  // 财产
	dir   int      // 行动方向: 1: 正向. -1: 反向

	mst   int  // 行动状态: 0: 正常, 1: 慢行. 2: 急行. 3: 禁行.
	mtime int  // 行动状态剩余时间
	cst   bool // 是否连续行动
	ast   int  // 结盟状态: -1: 未结盟. -2: 公愤. 0-? 结盟对象
	atime int  // 结盟剩余时间

	cards [CARD_COUNT]int // 卡片
}

func (d *Driver) GetRoleInfo(rind int) (
	name string,
	loc int, // -1: 不出场
	tech []int, // 掌握技术
	mos []int, // 随身人员
	mcs []int, // 占领城市
	sos []int, // 研究人员
	tos []int, // 训练人员
	money float32, // 财产
	dir int, // 行动方向: 1: 正向. -1: 反向

	mst int, // 行动状态: 0: 正常, 1: 慢行. 2: 急行. 3: 禁行.
	mtime int, // 行动状态剩余时间
	cst bool, // 是否连续行动
	ast int, // 结盟状态: -1: 未结盟. -2: 公愤. 0-? 结盟对象
	atime int, // 结盟剩余时间
	cards []int,
) {
	if rind == -1 {
		return
	}
	p := &(d.roles[rind])
	name = p.name
	loc = p.loc
	tech = p.tech.GetAllLabel()
	mos = p.mos.GetAllLabel()
	mcs = p.mcs.GetAllLabel()
	sos = p.sos.GetAllLabel()
	tos = p.tos.GetAllLabel()
	money = p.money
	dir = p.dir
	mst = p.mst
	mtime = p.mtime
	cst = p.cst
	ast = p.ast
	atime = p.atime
	cards = make([]int, CARD_COUNT)
	copy(cards, p.cards[:])
	return
}

func (r *Role) Save(fout io.Writer) {
	fmt.Fprintf(fout, "%s %d %f %d %d %d %t %d %d \n",
		r.name, r.loc, r.money, r.dir, r.mst, r.mtime, r.cst, r.ast, r.atime)
	r.tech.Save(fout)
	r.mos.Save(fout)
	r.mcs.Save(fout)
	r.sos.Save(fout)
	r.tos.Save(fout)
	for i := 0; i < CARD_COUNT; i++ {
		fmt.Fprintf(fout, "%d ", r.cards[i])
	}
	fmt.Fprintf(fout, "\n")
}

func (r *Role) Load(fin io.Reader) {
	fmt.Fscan(
		fin, &r.name, &r.loc, &r.money, &r.dir, &r.mst, &r.mtime, &r.cst, &r.ast, &r.atime)
	r.tech = dsg.LoadSet(fin)
	r.mos = dsg.LoadSet(fin)
	r.mcs = dsg.LoadSet(fin)
	r.sos = dsg.LoadSet(fin)
	r.tos = dsg.LoadSet(fin)
	for i := 0; i < CARD_COUNT; i++ {
		fmt.Fscan(fin, &r.cards[i])
	}
}
