package mbg


func InitCard () {

      cards[0] = Card{
            name: "障碍",
            odir: 0,
            otype: 3,
      }

      cards[0].do = func (d *Driver, sub int, obj int) int {
            d.board[obj].barrier = true
            d.uv.SetBarrier (obj)
            return 0
      }

      cards[1] = Card{
            name: "慢行",
            odir: 0,
            otype: 2,
      }

      cards[1].do = func (d *Driver, sub int, obj int) int {
            d.roles[obj].mst = 1
            d.roles[obj].mtime = 3
            d.uv.SetSlow (obj)
            return 0
      }

      cards[2] = Card{
            name: "急行",
            odir: 0,
            otype: 2,
      }

      cards[2].do = func (d *Driver, sub int, obj int) int {
            d.roles[obj].mst = 2
            d.roles[obj].mtime = 3
            d.uv.SetFast (obj)
            return 0
      }


      cards[3] = Card{
            name: "停止",
            odir: 0,
            otype: 2,
      }

      cards[3].do = func (d *Driver, sub int, obj int) int {
            d.roles[obj].mst = 3
            d.roles[obj].mtime = 3
            d.uv.SetStop (obj)
            return 0
      }


      cards[4] = Card{
            name: "双行",
            odir: 1,
            otype: 0,
      }

      cards[4].do = func (d *Driver, sub int, obj int) int {
            d.roles[sub].cst = true
            d.uv.SetContinue (sub)
            return 0
      }

      cards[5] = Card{
            name: "偷盗",
            odir: 2,
            otype: 2,
      }

      cards[5].do = func (d *Driver, sub int, obj int) int {
            real_dm, _ := d.roles[obj].ChangeMoney(-CARD_STEAL_MONEY)
            d.roles[sub].ChangeMoney (-real_dm)
            d.uv.StealMoney (sub, obj, real_dm)
            return 0
      }

      cards[6] = Card{
            name: "恢复",
            odir: 1,
            otype: 0,
      }

      cards[6].do = func (d *Driver, sub int, obj int) int {
		for _, i := range d.roles[sub].mos.GetAllLabel() {
                  d.people[i].ChangeHP (d.people[i].hpmax * CARD_HP_RECOVER_SCALE)
            }
            d.uv.SelfRecover(sub)
            return 0
      }

      cards[7] = Card{
            name: "流言",
            odir: 2,
            otype: 4,
      }

      cards[7].do = func (d *Driver, sub int, obj int) int {
            d.people[obj].lst = 4
            d.uv.SetLiuYan(obj)
            return 0
      }

      cards[8] = Card{
            name: "毒药",
            odir: 2,
            otype: 4,
      }

      cards[8].do = func (d *Driver, sub int, obj int) int {
            d.people[obj].pst = 8
            d.uv.SetPoison(obj)
            return 0
      }

      cards[9] = Card{
            name: "离间",
            odir: 2,
            otype: 4,
      }

      cards[9].do = func (d *Driver, sub int, obj int) int {
            d.people[obj].is_quit = true
            d.uv.SetQuit (obj)
            return 0
      }

      cards[10] = Card{
            name: "山贼",
            odir: 0,
            otype: 3,
      }

      cards[10].do = func (d *Driver, sub int, obj int) int {
            d.board[obj].robber = true
            d.uv.SetRobber(obj)
            return 0
      }

      cards[11] = Card{
            name: "联盟",
            odir: 2,
            otype: 2,
      }

      cards[11].do = func (d *Driver, sub int, obj int) int {
            a1 := d.roles[sub].ast
            a2 := d.roles[obj].ast
            if a1 != -1 || a2 != -1 {
                  d.oi[sub].AlignFail(obj)
                  return -1
            } else {
                  d.roles[obj].ast = sub
                  d.roles[obj].atime = 7
                  d.roles[sub].ast = obj
                  d.roles[sub].atime = 7
                  d.uv.SetAlign (sub, obj)
                  return 0
            }
      }

      cards[12] = Card{
            name: "公愤",
            odir: 2,
            otype: 2,
      }

      cards[12].do = func (d *Driver, sub int, obj int) int {
            ast := d.roles[obj].ast
            if ast != -1 {
                  d.oi[sub].PEFail(obj)
                  return -1
            } else {
                  d.roles[obj].ast = -2
                  d.roles[obj].atime = 5
                  d.uv.SetPE(obj)
                  return 0
            }
      }

      cards[13] = Card{
            name: "反向",
            odir: 0,
            otype: 2,
      }

      cards[13].do = func (d *Driver, sub int, obj int) int {
            d.roles[obj].dir = - d.roles[obj].dir
            d.uv.ChangeDir (obj)
            return 0
      }

}
