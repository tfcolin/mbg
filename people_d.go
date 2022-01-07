package mbg

import "math/rand"
import "io"
import "fmt"

type Property int
const (
      Wu    Property = iota
      Zhan  Property = iota
      Mou   Property = iota
      Zheng Property = iota
      Jing  Property = iota
)

type Officer struct {
      name string
      role int  // 所属阵营. -1: 下野 0-: 玩家阵营                                            
      job  int  // -1: 行军或下野 0:mayor 1:treasurer 2:citizen 3: institute 4: training      
      loc int   // -1: 行军或下野 0-n: city (institute, training_room)                        
      hpmax float32
      hp float32
      prop [5]float32 // 属性:  0-4: 武术, 战术, 谋略, 政治, 经济

      is_quit bool // 是否离职

      lst   int    // 流言状态: 剩余回合数. 0: 正常
      pst   int    // 中毒状态: 剩余回合数. 0: 正常

      level int    // 级别
}

func (d * Driver) GetPeopleInfo (pind int) (name string, role, job, loc int, hpmax, hp float32, prop [5]float32, is_quit bool, lst, pst int) {
      if pind == -1 {return}
      p := &(d.people[pind])
      name = p.name
      role = p.role
      job = p.job
      loc = p.loc
      hpmax = p.hpmax
      hp = p.hp
      for i := 0; i < 5; i ++ {
            prop[i] = p.GetProp(Property(i))
      }
      is_quit = p.is_quit
      lst = p.lst
      pst = p.pst
      return
}

func (o * Officer) GetProp (ind Property) float32 {
      value := o.prop[ind]
      if o.pst != 0 {
            value *= 0.7
      }
      return value
}

func (o * Officer) GetLevel () (level int) {
      var ave float32 = 0
      var np95, np85, np80, np60 int
      for i := 0; i < 5; i ++ {
            pv := o.prop[i]
            if pv > 95 {np95 ++}
            if pv > 85 {np85 ++}
            if pv > 80 {np80 ++}
            if pv > 60 {np60 ++}
            ave += pv
      }
      ave /= 5

      switch {
      case (np95 >= 1 && ave >= 50) || (np80 >= 2 && ave >= 70) :
            level = 2
      case (np85 >= 1 && ave >= 50) || (ave >= 60) :
            level = 1
      default:
            level = 0
      }

      return
}

func (o * Officer) GetPrice () float32 {
      var sum float32 = 0
      var max float32 = 0
      for i := 0; i < 5; i ++ {
            sum += o.prop[i]
            if o.prop[i] > max {max = o.prop[i]}
      }

      return ((10 * sum + 50 * max) * (0.9 + 0.2 * rand.Float32()))
}

func (p * Officer) Save (fout io.Writer) {
      fmt.Fprintf (fout, "%s %d %d %d %f %f ", p.name, p.role, p.job, p.loc, p.hpmax, p.hp)
      for i := 0; i < 5; i ++ {
            fmt.Fprintf (fout, "%f ", p.prop[i])
      }
      fmt.Fprintf (fout, "%t %d %d %d \n", p.is_quit, p.lst, p.pst, p.level)
}

func (p * Officer) Load (fin io.Reader) {
      fmt.Fscan (fin, &p.name, &p.role, &p.job, &p.loc, &p.hpmax, &p.hp)
      for i := 0; i < 5; i ++ {
            fmt.Fscan (fin, &p.prop[i])
      }
      fmt.Fscan (fin, &p.is_quit, &p.lst, &p.pst, &p.level)
}
