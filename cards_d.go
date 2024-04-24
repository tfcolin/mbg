package mbg

type Card struct {
      name string
      odir int // 对象阵营: 0: 任意, 1: 己方 2: 他方
      otype int // 对象类型: 0: 无对象(自己), 1: 城市, 2: 角色, 3: 地图, 4: 人员;
      do func (d *Driver, sub int, obj int) int // 操作函数 obj: 操作对象: 城市(*City), 角色(*Role), 武将(*Office), 地图(loc int). 返回: 0: 成功. -1: 失败
}

var cards [CARD_COUNT]Card

func GetCardInfo (card int) (name string, odir int, otype int) {
      if card < 0 || card >= CARD_COUNT {return}
      name = cards[card].name
      odir = cards[card].odir
      otype = cards[card].otype
      return
}
