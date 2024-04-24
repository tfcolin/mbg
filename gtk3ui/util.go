package gtk3ui

import (
      "math"
      "fmt"
	"os"

	"gitee.com/tfcolin/mbg"
      "github.com/gotk3/gotk3/gtk"
      "github.com/gotk3/gotk3/glib"
)

func seq (num int) []int {
      res := make ([]int, num)
      for i := 0; i < num; i ++ {
            res[i] = i
      }
      return res
}

func swin_set_follow (swin * gtk.ScrolledWindow) {
      adjust := swin.GetVAdjustment()
      upper := adjust.GetUpper()
      adjust.SetValue (upper)
}

func select_setup (tv * gtk.TreeView, is_multi bool) {
      var mode gtk.SelectionMode
      if (is_multi) {
            mode = gtk.SELECTION_MULTIPLE
      } else {
            mode = gtk.SELECTION_SINGLE
      }

      sel, _ := tv.GetSelection()
      sel.SetMode (mode)
}

func buf_add_msg (buffer * gtk.TextBuffer, msg string) {
      lmsg := fmt.Sprintf ("%s\n", msg)
      text_iter := buffer.GetEndIter()
      buffer.Insert (text_iter, lmsg)
}

func onoff_widget (widget gtk.IWidget, is_on bool) {
      widget.ToWidget().SetSensitive (is_on)
}

func tv_add_column (tv * gtk.TreeView, id int, title string) {
      renderer, _ := gtk.CellRendererTextNew()
      column, _ := gtk.TreeViewColumnNewWithAttribute (title, renderer, "text", id + 1)
      column.SetResizable (true)
      column.SetSortColumnID (id + 1)
      tv.AppendColumn (column)
}

func tv_clear (tv * gtk.TreeView) {
      gls_, _ := tv.GetModel()
      gls := gls_.(* gtk.ListStore)
      gls.Clear ()
}

func tv_add_people (id int, tv * gtk.TreeView, name string, role string, job int, loc int, hpmax float32, hp float32, prop [5]float32, is_quit bool, lst int, pst int, price float32) {
      gls_, _ := tv.GetModel()
      gls := gls_.(* gtk.ListStore)
      iter := gls.Append ()

      var jobstr, locstr, isqstr string
      switch (job) {
      case -1:
            jobstr =  "无"
      case 0:
            jobstr = "市长"
      case 1:
            jobstr = "财务官"
      case 2:
            jobstr = "市民"
      case 3:
            jobstr = "研究"
      case 4:
            jobstr = "训练"
      }

      if (loc == -1) {
            locstr = "随军"
      } else {
            switch  {
            case job == -1:
                  locstr = "随军"
            case job >= 0 && job <= 2:
                  locstr = city_name[loc]
            case job == 3:
                  locstr = inst_name[loc]
            case job == 4:
                  locstr = train_name[loc]
            }
      }

      if (is_quit) {
            isqstr = "是"
      } else {
            isqstr = "否"
      }

      columns := seq(16)

      if (price >= 0) {
            gls.Set (iter, columns, []interface{}{
                  id,
                  name,
                  role,
                  jobstr, 
                  locstr,
                  int(math.Ceil(float64(hp))),
                  int(math.Ceil(float64(hpmax))),
                  int(math.Ceil(float64(prop[0]))), 
                  int(math.Ceil(float64(prop[1]))), 
                  int(math.Ceil(float64(prop[2]))),
                  int(math.Ceil(float64(prop[3]))),
                  int(math.Ceil(float64(prop[4]))),
                  isqstr,
                  lst, 
                  pst,
                  int(math.Ceil(float64(price))),
            })
      } else {
            gls.Set (iter, columns[:15], []interface{}{
                  id,
                  name,
                  role,
                  jobstr, 
                  locstr,
                  int(math.Ceil(float64(hp))),
                  int(math.Ceil(float64(hpmax))),
                  int(math.Ceil(float64(prop[0]))), 
                  int(math.Ceil(float64(prop[1]))), 
                  int(math.Ceil(float64(prop[2]))),
                  int(math.Ceil(float64(prop[3]))),
                  int(math.Ceil(float64(prop[4]))),
                  isqstr,
                  lst,
                  pst,
            })
      }
}

func tv_add_city (id int, tv * gtk.TreeView,   name string, scope int, hpmax float32, hp float32, role string, nmos int, mayor string, treasurer string, fengshui float32) {
      gls_, _ := tv.GetModel()
      gls := gls_.(* gtk.ListStore)
      iter := gls.Append ()

      var sstr string
      switch (scope) {
      case 0:
            sstr = "小"
      case 1:
            sstr = "中"
      case 2:
            sstr = "大"
      }

      gls.Set (iter, seq(10), []interface{}{
            id,
            name,
            sstr,
            int (math.Ceil(float64(hp))),
            int (math.Ceil(float64(hpmax))),
            role,
            nmos,
            mayor,
            treasurer,
            int (math.Ceil(float64(fengshui * 100))),
      })
}

func tv_add_tech (id int, tv * gtk.TreeView,   name string, study float32, scond int) {
      gls_, _ := tv.GetModel()
      gls := gls_.(* gtk.ListStore)
      iter := gls.Append ()

      var scond_str string
      switch (scond) {
      case 0:
            scond_str = "任意"
      case 1:
            scond_str = "回合胜利"
      case 2:
            scond_str = "回合失败"
      case 3:
            scond_str = "战斗胜利"
      }

      gls.Set (iter, seq(4), []interface{}{
            id,
            name,
            int (math.Ceil(float64(study))),
            scond_str,
      })

}

func labels_clear (labels [] (* gtk.Label)) {
      for _, l := range labels {
            l.SetSensitive (false)
      }
}

func label_set_color_blue (label * gtk.Label) {
	str := label.GetLabel()
	/*
	mark_up := fmt.Sprintf ("<span foreground=\"blue\">%s</span>", str)
      label.SetMarkup (mark_up)
	*/
	n := len(str)
	if str[n - 1] != '+' {
		str += "+"
	}
	label.SetLabel(str)
}

func label_set_color_red (label * gtk.Label) {
	/*
	str := label.GetLabel()
	mark_up := fmt.Sprintf ("<span foreground=\"red\">%s</span>", str)
      label.SetMarkup (mark_up)
	*/
	label_set_color_blue (label)
}

func label_set_color_no (label * gtk.Label) {
	/*
	str := label.GetLabel()
	mark_up := fmt.Sprintf ("<span foreground=\"black\">%s</span>", str)
      label.SetMarkup (mark_up)
	*/
	str := label.GetLabel()
	n := len(str)
	if str[n - 1] == '+' {
		str = str[:n-1]
	}
	label.SetLabel(str)
}

func pb_set_value (pb * gtk.ProgressBar, value float32, max_value float32) {
      var frac float32
      if max_value == 0 {
            frac = 0
      } else {
            frac = value / max_value
      }

      msg := fmt.Sprintf ("%.0f / %.0f", value, max_value)
      pb.SetFraction (float64(frac))
      pb.SetText (msg)
}


func buf_clear (buf * gtk.TextBuffer) {
      buf.SetText ("")
}

func battle_set_turn (turn int) {
      title := fmt.Sprintf ("第 %d / %d 回合", turn + 1, max_turn)
      label_battle.SetLabel (title)
}

/* mode: 0: OK. 1: YES/NO; 
 * return 0: OK. -1: CANCEL; 1: CHECK
 */
func msg_box (msg string, mode int) int {

      if (st == GAME_OVER) {return -1}
      if (mode == 1) {
            onoff_widget (button_main_yes, true)
            onoff_widget (button_main_no, true)
      } else {
            onoff_widget (button_main_yes, true)
            onoff_widget (button_main_no, false)
      }
      if (len(msg) > 0) {
            buf_add_msg (buf_main_msg, msg)
            swin_set_follow (swin_main_msg)
      }

      msg_res = -1
      if (mode == 1) {
            st = WAIT_MSG_ANSWER
      } else {
            st = WAIT_MSG_CONFIRM
      }

      gtk.Main()
      if (st == GAME_OVER) {return -1}

      onoff_widget (button_main_yes, false)
      onoff_widget (button_main_no, false)

      return msg_res

}

func msg_map (msg string) {
      if (st == GAME_OVER) {return}
      buf_add_msg (buf_main_msg, msg)
      swin_set_follow (swin_main_msg)
}

func LoadUI () {
      var i int

      gtk.Init (nil)
      builder, _ := gtk.BuilderNewFromFile ("gtk3ui.glade")

      win_battle_, _ := builder.GetObject ("win_battle")
      win_battle = win_battle_.(* gtk.Window)
      label_battle_, _ := builder.GetObject ("battle_label")
      label_battle = label_battle_.(* gtk.Label)

      var iobj glib.IObject

      iobj, _ = builder.GetObject ("label_battle_role_0")
      label_battle_role[0] = iobj.(* gtk.Label)
      iobj, _ = builder.GetObject ("label_battle_role_1")
      label_battle_role[1] = iobj.(* gtk.Label)
      iobj, _ = builder.GetObject ("pb_battle_power_0")
      pb_battle_power[0] = iobj.(* gtk.ProgressBar)
      iobj, _ = builder.GetObject ("pb_battle_power_1")
      pb_battle_power[1] = iobj.(* gtk.ProgressBar)
      iobj, _ = builder.GetObject ("pb_battle_winp_0")
      pb_battle_winp[0] = iobj.(* gtk.ProgressBar)
      iobj, _ = builder.GetObject ("pb_battle_winp_1")
      pb_battle_winp[1] = iobj.(* gtk.ProgressBar)
      iobj, _ = builder.GetObject ("pb_battle_celve_0")
      pb_battle_celve[0] = iobj.(* gtk.ProgressBar)
      iobj, _ = builder.GetObject ("pb_battle_celve_1")
      pb_battle_celve[1] = iobj.(* gtk.ProgressBar)
      iobj, _ = builder.GetObject ("pb_battle_hp_0")
      pb_battle_hp[0] = iobj.(* gtk.ProgressBar)
      iobj, _ = builder.GetObject ("pb_battle_hp_1")
      pb_battle_hp[1] = iobj.(* gtk.ProgressBar)
      iobj, _ = builder.GetObject ("pb_battle_def_0")
      pb_battle_def[0] = iobj.(* gtk.ProgressBar)
      iobj, _ = builder.GetObject ("pb_battle_def_1")
      pb_battle_def[1] = iobj.(* gtk.ProgressBar)

      iobj, _ = builder.GetObject ("label_battle_tech_0_0")
      label_battle_tech[0][0] = iobj.(* gtk.Label)
      iobj, _ = builder.GetObject ("label_battle_tech_0_1")
      label_battle_tech[0][1] = iobj.(* gtk.Label)
      iobj, _ = builder.GetObject ("label_battle_tech_0_2")
      label_battle_tech[0][2] = iobj.(* gtk.Label)
      iobj, _ = builder.GetObject ("label_battle_tech_0_3")
      label_battle_tech[0][3] = iobj.(* gtk.Label)
      iobj, _ = builder.GetObject ("label_battle_tech_0_4")
      label_battle_tech[0][4] = iobj.(* gtk.Label)
      iobj, _ = builder.GetObject ("label_battle_tech_0_5")
      label_battle_tech[0][5] = iobj.(* gtk.Label)
      iobj, _ = builder.GetObject ("label_battle_tech_0_6")
      label_battle_tech[0][6] = iobj.(* gtk.Label)
      iobj, _ = builder.GetObject ("label_battle_tech_0_7")
      label_battle_tech[0][7] = iobj.(* gtk.Label)
      iobj, _ = builder.GetObject ("label_battle_tech_0_8")
      label_battle_tech[0][8] = iobj.(* gtk.Label)
      iobj, _ = builder.GetObject ("label_battle_tech_0_9")
      label_battle_tech[0][9] = iobj.(* gtk.Label)
      iobj, _ = builder.GetObject ("label_battle_tech_0_10")
      label_battle_tech[0][10] = iobj.(* gtk.Label)
      iobj, _ = builder.GetObject ("label_battle_tech_0_11")
      label_battle_tech[0][11] = iobj.(* gtk.Label)
      iobj, _ = builder.GetObject ("label_battle_tech_1_0")
      label_battle_tech[1][0] = iobj.(* gtk.Label)
      iobj, _ = builder.GetObject ("label_battle_tech_1_1")
      label_battle_tech[1][1] = iobj.(* gtk.Label)
      iobj, _ = builder.GetObject ("label_battle_tech_1_2")
      label_battle_tech[1][2] = iobj.(* gtk.Label)
      iobj, _ = builder.GetObject ("label_battle_tech_1_3")
      label_battle_tech[1][3] = iobj.(* gtk.Label)
      iobj, _ = builder.GetObject ("label_battle_tech_1_4")
      label_battle_tech[1][4] = iobj.(* gtk.Label)
      iobj, _ = builder.GetObject ("label_battle_tech_1_5")
      label_battle_tech[1][5] = iobj.(* gtk.Label)
      iobj, _ = builder.GetObject ("label_battle_tech_1_6")
      label_battle_tech[1][6] = iobj.(* gtk.Label)
      iobj, _ = builder.GetObject ("label_battle_tech_1_7")
      label_battle_tech[1][7] = iobj.(* gtk.Label)
      iobj, _ = builder.GetObject ("label_battle_tech_1_8")
      label_battle_tech[1][8] = iobj.(* gtk.Label)
      iobj, _ = builder.GetObject ("label_battle_tech_1_9")
      label_battle_tech[1][9] = iobj.(* gtk.Label)
      iobj, _ = builder.GetObject ("label_battle_tech_1_10")
      label_battle_tech[1][10] = iobj.(* gtk.Label)
      iobj, _ = builder.GetObject ("label_battle_tech_1_11")
      label_battle_tech[1][11] = iobj.(* gtk.Label)

      iobj, _ = builder.GetObject ("label_battle_st_0_0")
      label_battle_st[0][0] = iobj.(* gtk.Label)
      iobj, _ = builder.GetObject ("label_battle_st_0_1")
      label_battle_st[0][1] = iobj.(* gtk.Label)
      iobj, _ = builder.GetObject ("label_battle_st_0_2")
      label_battle_st[0][2] = iobj.(* gtk.Label)
      iobj, _ = builder.GetObject ("label_battle_st_1_0")
      label_battle_st[1][0] = iobj.(* gtk.Label)
      iobj, _ = builder.GetObject ("label_battle_st_1_1")
      label_battle_st[1][1] = iobj.(* gtk.Label)
      iobj, _ = builder.GetObject ("label_battle_st_1_2")
      label_battle_st[1][2] = iobj.(* gtk.Label)

      iobj, _ = builder.GetObject ("tv_battle_msg")
      tv_battle_msg = iobj.(* gtk.TextView)

      iobj, _ = builder.GetObject ("button_battle_next")
      button_battle_next = iobj.(* gtk.Button)
      iobj, _ = builder.GetObject ("tb_battle_auto")
      tb_battle_auto = iobj.(* gtk.ToggleButton)

      iobj, _ = builder.GetObject ("win_people")
      win_people = iobj.(* gtk.Window)
      iobj, _ = builder.GetObject ("label_people")
      label_people = iobj.(* gtk.Label)
      iobj, _ = builder.GetObject ("tv_people")
      tv_people = iobj.(* gtk.TreeView)
      iobj, _ = builder.GetObject ("button_people")
      button_people = iobj.(* gtk.Button)
      iobj, _ = builder.GetObject ("button_people_main")
      button_people_main = iobj.(* gtk.Button)
      iobj, _ = builder.GetObject ("button_people_vice")
      button_people_vice = iobj.(* gtk.Button)
      iobj, _ = builder.GetObject ("label_people_main")
      label_people_main = iobj.(* gtk.Label)
      iobj, _ = builder.GetObject ("label_people_vice")
      label_people_vice = iobj.(* gtk.Label)

      iobj, _ = builder.GetObject ("win_city")
      win_city = iobj.(* gtk.Window)
      iobj, _ = builder.GetObject ("tv_city")
      tv_city = iobj.(* gtk.TreeView)
      iobj, _ = builder.GetObject ("button_city")
      button_city = iobj.(* gtk.Button)

      iobj, _ = builder.GetObject ("win_card")
      win_card = iobj.(* gtk.Window)
      iobj, _ = builder.GetObject ("tv_card")
      tv_card = iobj.(* gtk.TreeView)
      iobj, _ = builder.GetObject ("button_card")
      button_card = iobj.(* gtk.Button)

      iobj, _ = builder.GetObject ("win_inst")
      win_inst = iobj.(* gtk.Window)
      iobj, _ = builder.GetObject ("tv_inst")
      tv_inst = iobj.(* gtk.TreeView)
      iobj, _ = builder.GetObject ("tv_inst_mos")
      tv_inst_mos = iobj.(* gtk.TreeView)
      iobj, _ = builder.GetObject ("tv_inst_tech")
      tv_inst_tech = iobj.(* gtk.TreeView)
      iobj, _ = builder.GetObject ("button_inst")
      button_inst = iobj.(* gtk.Button)

      iobj, _ = builder.GetObject ("win_train")
      win_train = iobj.(* gtk.Window)
      iobj, _ = builder.GetObject ("tv_train")
      tv_train = iobj.(* gtk.TreeView)
      iobj, _ = builder.GetObject ("tv_train_mos")
      tv_train_mos = iobj.(* gtk.TreeView)
      iobj, _ = builder.GetObject ("button_train")
      button_train = iobj.(* gtk.Button)

      iobj, _ = builder.GetObject ("win_exchange")
      win_exchange = iobj.(* gtk.Window)
      iobj, _ = builder.GetObject ("label_exchange")
      label_exchange = iobj.(* gtk.Label )
      iobj, _ = builder.GetObject ("tv_exchange_in")
      tv_exchange_in = iobj.(* gtk.TreeView)
      iobj, _ = builder.GetObject ("tv_exchange_out")
      tv_exchange_out = iobj.(* gtk.TreeView)
      iobj, _ = builder.GetObject ("button_exchange")
      button_exchange = iobj.(* gtk.Button)

      iobj, _ = builder.GetObject ("win_study")
      win_study = iobj.(* gtk.Window) 
      iobj, _ = builder.GetObject ("label_study")
      label_study = iobj.(* gtk.Label)
      iobj, _ = builder.GetObject ("tv_study_tech")
      tv_study_tech = iobj.(* gtk.TreeView) 
      iobj, _ = builder.GetObject ("tv_study_people")
      tv_study_people = iobj.(* gtk.TreeView) 
      iobj, _ = builder.GetObject ("button_study")
      button_study = iobj.(* gtk.Button) 

      iobj, _ = builder.GetObject ("win_role")
      win_role = iobj.(* gtk.Window) 
      iobj, _ = builder.GetObject ("tv_role")
      tv_role = iobj.(* gtk.TreeView) 
      iobj, _ = builder.GetObject ("tv_role_mos")
      tv_role_mos = iobj.(* gtk.TreeView) 
      iobj, _ = builder.GetObject ("tv_role_mcs")
      tv_role_mcs = iobj.(* gtk.TreeView) 
      iobj, _ = builder.GetObject ("tv_role_sos")
      tv_role_sos = iobj.(* gtk.TreeView) 
      iobj, _ = builder.GetObject ("tv_role_tos")
      tv_role_tos = iobj.(* gtk.TreeView) 
      iobj, _ = builder.GetObject ("tv_role_cards")
      tv_role_cards = iobj.(* gtk.TreeView) 

      iobj, _ = builder.GetObject ("label_role_tech_0")
      label_role_tech[0] = iobj.(* gtk.Label)
      iobj, _ = builder.GetObject ("label_role_tech_1")
      label_role_tech[1] = iobj.(* gtk.Label)
      iobj, _ = builder.GetObject ("label_role_tech_2")
      label_role_tech[2] = iobj.(* gtk.Label)
      iobj, _ = builder.GetObject ("label_role_tech_3")
      label_role_tech[3] = iobj.(* gtk.Label)
      iobj, _ = builder.GetObject ("label_role_tech_4")
      label_role_tech[4] = iobj.(* gtk.Label)
      iobj, _ = builder.GetObject ("label_role_tech_5")
      label_role_tech[5] = iobj.(* gtk.Label)
      iobj, _ = builder.GetObject ("label_role_tech_6")
      label_role_tech[6] = iobj.(* gtk.Label)
      iobj, _ = builder.GetObject ("label_role_tech_7")
      label_role_tech[7] = iobj.(* gtk.Label)
      iobj, _ = builder.GetObject ("label_role_tech_8")
      label_role_tech[8] = iobj.(* gtk.Label)
      iobj, _ = builder.GetObject ("label_role_tech_9")
      label_role_tech[9] = iobj.(* gtk.Label)
      iobj, _ = builder.GetObject ("label_role_tech_10")
      label_role_tech[10] = iobj.(* gtk.Label)
      iobj, _ = builder.GetObject ("label_role_tech_11")
      label_role_tech[11] = iobj.(* gtk.Label)

      iobj, _ = builder.GetObject ("button_role")
      button_role = iobj.(* gtk.Button) 
      iobj, _ = builder.GetObject ("button_role_mcs")
      button_role_mcs = iobj.(* gtk.Button) 

      iobj, _ = builder.GetObject ("win_main")
      win_main = iobj.(* gtk.Window) 
      iobj, _ = builder.GetObject ("layout_main_map")
      layout_main_map = iobj.(* gtk.Layout)
      iobj, _ = builder.GetObject ("button_main_move")
      button_main_move = iobj.(* gtk.Button) 
      iobj, _ = builder.GetObject ("button_main_card")
      button_main_card = iobj.(* gtk.Button) 
      iobj, _ = builder.GetObject ("button_main_save")
      button_main_save = iobj.(* gtk.Button)
      iobj, _ = builder.GetObject ("button_main_role")
      button_main_role = iobj.(* gtk.Button) 
      iobj, _ = builder.GetObject ("button_main_people")
      button_main_people = iobj.(* gtk.Button) 
      iobj, _ = builder.GetObject ("button_main_rp")
      button_main_rp = iobj.(* gtk.Button) 
      iobj, _ = builder.GetObject ("button_main_city")
      button_main_city = iobj.(* gtk.Button) 
      iobj, _ = builder.GetObject ("button_main_inst")
      button_main_inst = iobj.(* gtk.Button) 
      iobj, _ = builder.GetObject ("button_main_train")
      button_main_train = iobj.(* gtk.Button) 
      iobj, _ = builder.GetObject ("tv_main_msg")
      tv_main_msg = iobj.(* gtk.TextView) 
      iobj, _ = builder.GetObject ("button_main_yes")
      button_main_yes = iobj.(* gtk.Button) 
      iobj, _ = builder.GetObject ("button_main_no")
      button_main_no = iobj.(* gtk.Button) 
      iobj, _ = builder.GetObject ("button_main_small")
      button_main_small = iobj.(* gtk.Button) 
      iobj, _ = builder.GetObject ("button_main_medium")
      button_main_medium = iobj.(* gtk.Button) 
      iobj, _ = builder.GetObject ("button_main_large")
      button_main_large = iobj.(* gtk.Button) 

      iobj, _ = builder.GetObject ("buf_main_msg")
      buf_main_msg = iobj.(* gtk.TextBuffer) 
      iobj, _ = builder.GetObject ("buf_battle_msg")
      buf_battle_msg = iobj.(* gtk.TextBuffer) 
      iobj, _ = builder.GetObject ("swin_main_msg")
      swin_main_msg = iobj.(* gtk.ScrolledWindow) 
      iobj, _ = builder.GetObject ("swin_battle_msg")
      swin_battle_msg = iobj.(* gtk.ScrolledWindow) 

      button_main_role.Connect ( "clicked", cbf_main_role_check)
      button_main_people.Connect ( "clicked", cbf_main_people_check)
      button_main_move.Connect ( "clicked", cbf_main_move_click)
      button_main_card.Connect ( "clicked", cbf_main_card_click)
	button_main_save.Connect ( "clicked", cbf_main_save_click)
      button_main_city.Connect ( "clicked", cbf_main_city_check)
      button_main_inst.Connect ( "clicked", cbf_main_inst_check)
      button_main_train.Connect ( "clicked", cbf_main_train_check)
      button_main_rp.Connect ( "clicked", cbf_main_rp_check)
      button_main_yes.Connect ( "clicked", cbf_main_yes_click)
      button_main_no.Connect ( "clicked", cbf_main_no_click)
      button_main_small.Connect ( "clicked", cbf_main_small_click)
      button_main_medium.Connect ( "clicked", cbf_main_medium_click)
      button_main_large.Connect ( "clicked", cbf_main_large_click)

      button_battle_next.Connect ( "clicked", cbf_battle_next_click)
      button_people.Connect ( "clicked", cbf_people_click)
      button_people_main.Connect ( "clicked", cbf_people_main_click)
      button_people_vice.Connect ( "clicked", cbf_people_vice_click)
      button_city.Connect ( "clicked", cbf_city_click)
      button_card.Connect ( "clicked", cbf_card_click)
      button_inst.Connect ( "clicked", cbf_inst_click)
      button_train.Connect ( "clicked", cbf_train_click)
      button_exchange.Connect ( "clicked", cbf_exchange_click)
      button_study.Connect ( "clicked", cbf_study_click)
      button_role.Connect ( "clicked", cbf_role_click)
      button_role_mcs.Connect ( "clicked", cbf_role_mcs_click)

      win_main.Connect ( "delete-event", cbf_main_close)
      win_card.Connect ( "delete-event", cbf_card_close)
      win_battle.Connect ( "delete-event", cbf_battle_close)
      win_role.Connect ( "delete-event", cbf_role_close)
      win_people.Connect ( "delete-event", cbf_people_close)
      win_city.Connect ( "delete-event", cbf_city_close)
      win_inst.Connect ( "delete-event", cbf_inst_close)
      win_train.Connect ( "delete-event", cbf_train_close)
      win_study.Connect ( "delete-event", cbf_study_close)
      win_exchange.Connect ( "delete-event", cbf_exchange_close)

      // set model and parent
      
      win_card.SetTransientFor (win_main)
      win_battle.SetTransientFor (win_main)
      win_role.SetTransientFor (win_main)
      win_people.SetTransientFor (win_main)
      win_city.SetTransientFor (win_main)
      win_inst.SetTransientFor (win_main)
      win_train.SetTransientFor (win_main)
      win_study.SetTransientFor (win_main)
      win_exchange.SetTransientFor (win_main)

      win_card.SetModal (true)
      win_battle.SetModal (true)
      win_role.SetModal (true)
      win_people.SetModal (true)
      win_city.SetModal (true)
      win_inst.SetModal (true)
      win_train.SetModal (true)
      win_study.SetModal (true)
      win_exchange.SetModal (true)

      glib.TimeoutAdd (500, cbf_battle_timer)

      onoff_widget (button_main_yes, false)
      onoff_widget (button_main_no, false)
      onoff_widget (button_main_small, false)
      onoff_widget (button_main_medium, false)
      onoff_widget (button_main_large, false)
      onoff_widget (button_main_move, false)
      onoff_widget (button_main_card, false)

      // treeview initialization
      
      var gls * gtk.ListStore

      gls, _ = gtk.ListStoreNew (glib.TYPE_INT, glib.TYPE_STRING, glib.TYPE_INT, glib.TYPE_STRING, glib.TYPE_STRING)
      tv_card.SetModel (gls)

      tv_add_column (tv_card, 0, "名称")
      tv_add_column (tv_card, 1, "持有量")
      tv_add_column (tv_card, 2, "对象阵营")
      tv_add_column (tv_card, 3, "对象类型")

      gls, _ = gtk.ListStoreNew (glib.TYPE_INT, glib.TYPE_STRING, glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_STRING, 
      glib.TYPE_STRING, glib.TYPE_INT, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_INT)
      tv_role.SetModel (gls)

      tv_add_column (tv_role, 0, "名称")
      tv_add_column (tv_role, 1, "位置")
      tv_add_column (tv_role, 2, "金币")
      tv_add_column (tv_role, 3, "移动方向")
      tv_add_column (tv_role, 4, "行动状态")
      tv_add_column (tv_role, 5, "行动状态剩余")
      tv_add_column (tv_role, 6, "连续行动")
      tv_add_column (tv_role, 7, "联盟")
      tv_add_column (tv_role, 8, "联盟剩余")

      gls, _ = gtk.ListStoreNew (glib.TYPE_INT, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING, 
      glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_INT, 
      glib.TYPE_STRING, glib.TYPE_INT, glib.TYPE_INT)
      tv_role_mos.SetModel (gls)

      tv_add_column (tv_role_mos, 0, "名称")
      tv_add_column (tv_role_mos, 1, "阵营")
      tv_add_column (tv_role_mos, 2, "职业")
      tv_add_column (tv_role_mos, 3, "位置")
      tv_add_column (tv_role_mos, 4, "兵力")
      tv_add_column (tv_role_mos, 5, "最大兵力")
      for i = 0; i < 5; i ++ {
            tv_add_column (tv_role_mos, 6 + i, prop_name[i])
      }
      tv_add_column (tv_role_mos, 11, "准备离职")
      tv_add_column (tv_role_mos, 12, "流言状态剩余")
      tv_add_column (tv_role_mos, 13, "中毒状态剩余")

      gls, _ = gtk.ListStoreNew (glib.TYPE_INT, glib.TYPE_STRING, glib.TYPE_STRING, 
      glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_STRING, glib.TYPE_INT,
      glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_INT)
      tv_role_mcs.SetModel (gls)

      tv_add_column (tv_role_mcs, 0, "名称")
      tv_add_column (tv_role_mcs, 1, "规模")
      tv_add_column (tv_role_mcs, 2, "防御")
      tv_add_column (tv_role_mcs, 3, "最大防御")
      tv_add_column (tv_role_mcs, 4, "所属阵营")
      tv_add_column (tv_role_mcs, 5, "驻扎官员数")
      tv_add_column (tv_role_mcs, 6, "市长")
      tv_add_column (tv_role_mcs, 7, "财务官")
      tv_add_column (tv_role_mcs, 8, "风水")

      gls, _ = gtk.ListStoreNew (glib.TYPE_INT, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING,
      glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_INT, 
      glib.TYPE_STRING, glib.TYPE_INT, glib.TYPE_INT)
      tv_role_sos.SetModel (gls)

      tv_add_column (tv_role_sos, 0, "名称")
      tv_add_column (tv_role_sos, 1, "阵营")
      tv_add_column (tv_role_sos, 2, "职业")
      tv_add_column (tv_role_sos, 3, "位置")
      tv_add_column (tv_role_sos, 4, "兵力")
      tv_add_column (tv_role_sos, 5, "最大兵力")
      for i = 0; i < 5; i ++ {
            tv_add_column (tv_role_sos, 6 + i, prop_name[i])
      }
      tv_add_column (tv_role_sos, 11, "准备离职")
      tv_add_column (tv_role_sos, 12, "流言状态剩余")
      tv_add_column (tv_role_sos, 13, "中毒状态剩余")

      gls, _ = gtk.ListStoreNew (glib.TYPE_INT, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING, 
      glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_INT, 
      glib.TYPE_STRING, glib.TYPE_INT, glib.TYPE_INT)
      tv_role_tos.SetModel (gls)

      tv_add_column (tv_role_tos, 0, "名称")
      tv_add_column (tv_role_tos, 1, "阵营")
      tv_add_column (tv_role_tos, 2, "职业")
      tv_add_column (tv_role_tos, 3, "位置")
      tv_add_column (tv_role_tos, 4, "兵力")
      tv_add_column (tv_role_tos, 5, "最大兵力")
      for i = 0; i < 5; i ++ {
            tv_add_column (tv_role_tos, 6 + i, prop_name[i])
      }
      tv_add_column (tv_role_tos, 11, "准备离职")
      tv_add_column (tv_role_tos, 12, "流言状态剩余")
      tv_add_column (tv_role_tos, 13, "中毒状态剩余")

      gls, _ = gtk.ListStoreNew (glib.TYPE_INT, glib.TYPE_STRING, glib.TYPE_INT)
      tv_role_cards.SetModel (gls)

      tv_add_column (tv_role_cards, 0, "名称")
      tv_add_column (tv_role_cards, 1, "数量")

      gls, _ = gtk.ListStoreNew (glib.TYPE_INT, glib.TYPE_STRING, glib.TYPE_STRING, 
            glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_STRING, glib.TYPE_INT,
            glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_INT)
      tv_city.SetModel (gls)

      tv_add_column (tv_city, 0, "名称")
      tv_add_column (tv_city, 1, "规模")
      tv_add_column (tv_city, 2, "防御")
      tv_add_column (tv_city, 3, "最大防御")
      tv_add_column (tv_city, 4, "所属阵营")
      tv_add_column (tv_city, 5, "驻扎官员数")
      tv_add_column (tv_city, 6, "市长")
      tv_add_column (tv_city, 7, "财务官")
      tv_add_column (tv_city, 8, "风水")

      gls, _ = gtk.ListStoreNew (glib.TYPE_INT, glib.TYPE_STRING, glib.TYPE_INT, 
            glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_INT, glib.TYPE_INT)
      tv_inst.SetModel (gls)

      tv_add_column (tv_inst, 0, "名称")
      tv_add_column (tv_inst, 1, "可研究策略数")
      tv_add_column (tv_inst, 2, "在研人员")
      tv_add_column (tv_inst, 3, "在研项目")
      tv_add_column (tv_inst, 4, "剩余成功点数")
      tv_add_column (tv_inst, 5, "预计剩余回合")

      gls, _ = gtk.ListStoreNew (glib.TYPE_INT, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING, 
            glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_INT, 
            glib.TYPE_STRING, glib.TYPE_INT, glib.TYPE_INT)
      tv_inst_mos.SetModel (gls)

      tv_add_column (tv_inst_mos, 0, "名称")
      tv_add_column (tv_inst_mos, 1, "阵营")
      tv_add_column (tv_inst_mos, 2, "职业")
      tv_add_column (tv_inst_mos, 3, "位置")
      tv_add_column (tv_inst_mos, 4, "兵力")
      tv_add_column (tv_inst_mos, 5, "最大兵力")
      for i = 0; i < 5; i ++ {
            tv_add_column (tv_inst_mos, 6 + i, prop_name[i])
      }
      tv_add_column (tv_inst_mos, 11, "准备离职")
      tv_add_column (tv_inst_mos, 12, "流言状态剩余")
      tv_add_column (tv_inst_mos, 13, "中毒状态剩余")

      gls, _ = gtk.ListStoreNew (glib.TYPE_INT, glib.TYPE_STRING, glib.TYPE_INT, glib.TYPE_STRING)
      tv_inst_tech.SetModel (gls)

      tv_add_column (tv_inst_tech, 0, "名称")
      tv_add_column (tv_inst_tech, 1, "研究点数")
      tv_add_column (tv_inst_tech, 2, "发动时机")

      gls, _ = gtk.ListStoreNew (glib.TYPE_INT, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_INT)
      tv_train.SetModel (gls)

      tv_add_column (tv_train, 0, "名称")
      tv_add_column (tv_train, 1, "训练项目")
      tv_add_column (tv_train, 2, "训练人员")
      tv_add_column (tv_train, 3, "剩余回合")

      gls, _ = gtk.ListStoreNew (glib.TYPE_INT, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING, 
            glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_INT, 
            glib.TYPE_STRING, glib.TYPE_INT, glib.TYPE_INT)
      tv_train_mos.SetModel (gls)

      tv_add_column (tv_train_mos, 0, "名称")
      tv_add_column (tv_train_mos, 1, "阵营")
      tv_add_column (tv_train_mos, 2, "职业")
      tv_add_column (tv_train_mos, 3, "位置")
      tv_add_column (tv_train_mos, 4, "兵力")
      tv_add_column (tv_train_mos, 5, "最大兵力")
      for i = 0; i < 5; i ++ {
            tv_add_column (tv_train_mos, 6 + i, prop_name[i])
      }
      tv_add_column (tv_train_mos, 11, "准备离职")
      tv_add_column (tv_train_mos, 12, "流言状态剩余")
      tv_add_column (tv_train_mos, 13, "中毒状态剩余")

      gls, _ = gtk.ListStoreNew (glib.TYPE_INT, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING, 
            glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_INT, 
            glib.TYPE_STRING, glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_INT)
      tv_people.SetModel (gls)

      tv_add_column (tv_people, 0, "名称")
      tv_add_column (tv_people, 1, "阵营")
      tv_add_column (tv_people, 2, "职业")
      tv_add_column (tv_people, 3, "位置")
      tv_add_column (tv_people, 4, "兵力")
      tv_add_column (tv_people, 5, "最大兵力")
      for i = 0; i < 5; i ++ {
            tv_add_column (tv_people, 6 + i, prop_name[i])
      }
      tv_add_column (tv_people, 11, "准备离职")
      tv_add_column (tv_people, 12, "流言状态剩余")
      tv_add_column (tv_people, 13, "中毒状态剩余")
      tv_add_column (tv_people, 14, "聘用费")

      gls, _ = gtk.ListStoreNew (glib.TYPE_INT, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING, 
            glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_INT, 
            glib.TYPE_STRING, glib.TYPE_INT, glib.TYPE_INT)
      tv_exchange_out.SetModel (gls)

      tv_add_column (tv_exchange_out, 0, "名称")
      tv_add_column (tv_exchange_out, 1, "阵营")
      tv_add_column (tv_exchange_out, 2, "职业")
      tv_add_column (tv_exchange_out, 3, "位置")
      tv_add_column (tv_exchange_out, 4, "兵力")
      tv_add_column (tv_exchange_out, 5, "最大兵力")
      for i = 0; i < 5; i ++ {
            tv_add_column (tv_exchange_out, 6 + i, prop_name[i])
      }
      tv_add_column (tv_exchange_out, 11, "准备离职")
      tv_add_column (tv_exchange_out, 12, "流言状态剩余")
      tv_add_column (tv_exchange_out, 13, "中毒状态剩余")

      gls, _ = gtk.ListStoreNew (glib.TYPE_INT, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING, 
            glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_INT, 
            glib.TYPE_STRING, glib.TYPE_INT, glib.TYPE_INT)
      tv_exchange_in.SetModel (gls)

      tv_add_column (tv_exchange_in, 0, "名称")
      tv_add_column (tv_exchange_in, 1, "阵营")
      tv_add_column (tv_exchange_in, 2, "职业")
      tv_add_column (tv_exchange_in, 3, "位置")
      tv_add_column (tv_exchange_in, 4, "兵力")
      tv_add_column (tv_exchange_in, 5, "最大兵力")
      for i = 0; i < 5; i ++ {
            tv_add_column (tv_exchange_in, 6 + i, prop_name[i])
      }
      tv_add_column (tv_exchange_in, 11, "准备离职")
      tv_add_column (tv_exchange_in, 12, "流言状态剩余")
      tv_add_column (tv_exchange_in, 13, "中毒状态剩余")

      gls, _ = gtk.ListStoreNew (glib.TYPE_INT, glib.TYPE_STRING, glib.TYPE_INT, glib.TYPE_STRING)
      tv_study_tech.SetModel (gls)

      tv_add_column (tv_study_tech, 0, "名称")
      tv_add_column (tv_study_tech, 1, "研究点数")
      tv_add_column (tv_study_tech, 2, "发动时机")

      gls, _ = gtk.ListStoreNew (glib.TYPE_INT, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING, 
            glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_INT, 
            glib.TYPE_STRING, glib.TYPE_INT, glib.TYPE_INT)
      tv_study_people.SetModel (gls)

      tv_add_column (tv_study_people, 0, "名称")
      tv_add_column (tv_study_people, 1, "阵营")
      tv_add_column (tv_study_people, 2, "职业")
      tv_add_column (tv_study_people, 3, "位置")
      tv_add_column (tv_study_people, 4, "兵力")
      tv_add_column (tv_study_people, 5, "最大兵力")
      for i = 0; i < 5; i ++ {
            tv_add_column (tv_study_people, 6 + i, prop_name[i])
      }
      tv_add_column (tv_study_people, 11, "准备离职")
      tv_add_column (tv_study_people, 12, "流言状态剩余")
      tv_add_column (tv_study_people, 13, "中毒状态剩余")

      // init variables
      
      st = GAME_OVER

      obj_type = -1
      obj = -1
      action = -1
      card_action = -1

      check_type = -1
      check_role_select = -1
      check_role_mcs_select = -1
      check_city_mcs_select = -1
      check_inst_select = -1
      check_train_select = -1

      people_select_st = 0
      people_select = -1
      people_select_vice = -1
      people_select_num = 0

      battle_scale = -1
      wait_battle_tick = false
      is_battle_finish = false

      exchange_out_num = 0
      exchange_in_num = 0

      study_tech = -1
      study_people = -1

      msg_res = -1

      wp = 85
      hp = 85
      is_align = false

	// init file save dialog
	
	file_save_dialog, _ = gtk.FileChooserDialogNewWith2Buttons ("保存游戏", win_main, gtk.FILE_CHOOSER_ACTION_SAVE, 
	"保存", gtk.RESPONSE_ACCEPT, "取消", gtk.RESPONSE_CANCEL )
	pwd, _ := os.Getwd()
	file_save_dialog.SetCurrentFolder (pwd)
	file_save_dialog.SetDoOverwriteConfirmation (true)
	
}

func get_file_choose () string {
	res := file_save_dialog.Run()
	var fname string
	if res == gtk.RESPONSE_ACCEPT {
		fname = file_save_dialog.GetFilename()
	} else {
		fname = ""
	}

	file_save_dialog.Hide()
	return fname
}

func init_map  (ngrid_, nrole_, npeople_, ncity_, ninst_, ntrain_, ntech_, ncard_ int) {
      var i int

      ngrid = ngrid_
      if (ngrid % 4 == 0) {
            is_align = true
      } else {
            is_align = false
      }
      nrole = nrole_
      npeople = npeople_
      ncity = ncity_
      ninst = ninst_
      ntrain = ntrain_
      ntech = ntech_
      ncard = ncard_

      nx = ngrid / 4 + 1 + nrole * 2
      ny = nx 

      layout_main_map.SetSize ( uint((nx + 2) * wp), uint((ny + 2) * hp) )

      map_point = make ([](* gtk.Button), ngrid)
      role_point = make ([][](* gtk.Button), ngrid)
      is_role = make ([][]bool, ngrid)
      for i = 0; i < ngrid; i ++ {
            role_point[i] = make ([](* gtk.Button), nrole)
            is_role[i] = make ([]bool, nrole)
      }
      map_type = make ([]int, ngrid)
      map_obj = make ([]int, ngrid)
      is_barrier = make ([]bool, ngrid)
      is_robber = make ([]bool, ngrid)
      role_name = make ([]string, nrole)
      tech_name = make ([]string, ntech)
      tech_scond = make ([]int, ntech)
      card_name = make ([]string, ncard)
      card_odir = make ([]int, ncard)
      card_otype = make ([]int, ncard)
      city_name = make ([]string, ncity)
      inst_name = make ([]string, ninst)
      train_name = make ([]string, ntrain)

      people_select_list = make ([]int, MAX_LIST_LENGTH)
      exchange_out_list = make ([]int, MAX_LIST_LENGTH)
      exchange_in_list = make ([]int, MAX_LIST_LENGTH)

}

func new_button (x, y int) * gtk.Button {
      rx, ry := (x + 1) * wp, (y + 1) * hp
      res, _ := gtk.ButtonNew()
      res.SetSizeRequest (wp, hp)
      layout_main_map.Put (res, rx, ry)

      return res
}

func init_tech (tech int, tname string, scond int) {
      tech_name[tech] = tname
      tech_scond[tech] = scond
}

func init_card (card int,  dname string, odir, otype int) {
      card_name[card] = dname
      card_odir[card] = odir
      card_otype[card] = otype
}

func init_role (role int, rname string) {
      role_name[role] = rname
}

func set_board (loc, type_, ind, x, y int) {
      var i int
      if (!is_align) {return}

      var name string
      switch (type_) {
      case 1:
            name = city_name[ind]
      case 2:
            name = "会所"
      case 3:
            name = inst_name[ind]
      case 4:
            name = train_name[ind]
      case 5:
            name = "机会"
      }

      sec := loc / (ngrid / 4)
      resi := loc % (ngrid / 4)
      switch (sec) {
      case 0:
            x = resi
            y = 0
      case 1: 
            x = ngrid / 4
            y = resi
      case 2:
            x = ngrid / 4 - resi
            y = ngrid / 4
      case 3:
            x = 0
            y = ngrid / 4 - resi
      }

      for i = 0; i < nrole; i ++ {
            is_role[loc][i] = false
      }
      is_barrier[loc] = false
      is_robber[loc] = false

      map_point[loc] = new_button (x + nrole, y + nrole)
      map_type[loc] = type_
      map_obj[loc] = ind

      map_point[loc].SetLabel (name)
      map_point[loc].Connect ("clicked", cbf_main_board_click_maker (loc))

      var xx, yy int
      for i = 0; i < nrole; i ++ {
            switch (sec) {
            case 0:
                  xx = x + nrole
                  yy = i
            case 1:
                  xx = ngrid / 4 + nrole + 1 + i
                  yy = y + nrole
            case 2:
                  xx = x + nrole
                  yy = ngrid / 4 + nrole + 1 + i
            case 3:
                  xx = i
                  yy = y + nrole
            }
            role_point[loc][i] = new_button (xx, yy)
            role_point[loc][i].Connect ("clicked", cbf_main_role_click_maker (loc, i))
      }
}

func set_city    (id int, name string) {
      city_name[id] = name
}

func set_inst    (id int, name string) {
      inst_name[id] = name
}

func set_train    (id int, name string) {
      train_name[id] = name
}

func set_barrier (loc int, is_unset bool) {
      is_barrier[loc] = !is_unset
}

func set_robber  (loc int, is_unset bool) {
      is_robber[loc] = !is_unset
}

func set_role   (loc, rind int, is_unset bool) {
      is_role[loc][rind] = !is_unset
}

func update_role_point (loc int) {
      for i := 0; i < nrole; i ++ {
            var label string
            if (is_role[loc][i]) {
                  label += role_name[i]
            }
            if (is_barrier[loc]) {
                  label += "\n障碍"
            }
            if (is_robber[loc]) {
                  label += "\n强盗"
            }
            role_point[loc][i].SetLabel (label)
      }
}

func finish_draw_map () {
      if (!is_align) {return}

      buf_clear (buf_main_msg)

      var i, j int

      for i = 0; i < ngrid; i ++ {
            update_role_point (i)
      }

      // set tech names
      for i = 0; i < 2; i ++ {
            for j = 0; j < 12; j ++ {
                  label_battle_tech[i][j].SetLabel (tech_name[j])
            }
      }

      for i = 0; i < 12; i ++ {
            label_role_tech[i].SetLabel (tech_name[i])
      }

      st = WAIT_MSG_CONFIRM
      win_main.ShowAll()
}

// 0: 移动. 1: 使用卡片. 2: check(waiting). 3: 保存游戏. -1: 中途退出游戏
func get_action (rind int) int {

      if (!is_align || st == GAME_OVER) {return -1}

      st = WAIT_ACTION
      onoff_widget (button_main_move, true)
      onoff_widget (button_main_card, true)
	onoff_widget (button_main_save, true)

      action = -1
      for (action == -1) {
            gtk.Main ()
            if (st == GAME_OVER) {return -1}
      }

      onoff_widget (button_main_move, false)
      onoff_widget (button_main_card, false)
	onoff_widget (button_main_save, false)

      return action
}

// 0: role. 1: people. 2: city; 3:inst. 4:trainingroom. 5: role_people
func get_check_type () int {
      return check_type
}

var cards_list_num int

// called after call get_action () (return 1) to show cards
func begin_cards_list () {
      tv_clear (tv_card)
      cards_list_num = 0
}

func add_cards_list (cind, num int) {

      gls_, _ := tv_card.GetModel()
      gls := gls_.(* gtk.ListStore)
      iter := gls.Append ()

      var odir_str, otype_str string

      switch (card_odir[cind]) {
      case 0:
            odir_str = "任意"
      case 1:
            odir_str = "己方"
      case 2:
            odir_str = "他方"
      }

      switch (card_otype[cind]) {
      case 0:
            otype_str = "自己"
      case 1:
            otype_str = "城市"
      case 2:
            otype_str = "角色"
      case 3:
            otype_str = "地图"
      case 4:
            otype_str = "武将"
      }

      gls.Set (iter, seq(5), []interface{}{
            cards_list_num, 
            card_name[cind], 
            num, 
            odir_str, 
            otype_str, 
      })

      cards_list_num ++ 
}

func end_cards_list () {
      win_card.SetTransientFor (win_main)
      select_setup (tv_card, false)
}

func show_top_window (win * gtk.Window) {
	win.ShowAll()
	win.SetKeepAbove(true)
	win.Present()
}

// -1: cancel
func get_card_action () int {
      if (st == GAME_OVER) {return -1}
      card_action = -1
	show_top_window (win_card)
      gtk.Main()
      if (st == GAME_OVER) {return -1}
      return card_action
}

var role_list_num int

// check role info
func begin_role_list () {
      tv_clear (tv_role)
      tv_clear (tv_role_mos)
      tv_clear (tv_role_mcs)
      tv_clear (tv_role_sos)
      tv_clear (tv_role_tos)
      tv_clear (tv_role_cards)
      labels_clear (label_role_tech[:])
      role_list_num = 0
}

func add_role_list (name string, loc int, money float32, dir, mst, mtime int, cst bool, ast, atime int) {

      gls_, _ := tv_role.GetModel()
      gls := gls_.(* gtk.ListStore)
      iter := gls.Append ()

      var dirstr, mststr, cststr, aststr string

      if (dir > 0) { 
            dirstr = "正向"
      } else {
            dirstr = "反向"
      }

      switch (mst) {
      case 0:
            mststr = "正常"
      case 1:
            mststr = "慢行"
      case 2:
            mststr = "急行"
      case 3:
            mststr = "禁行"
      }

      if (cst) {
            cststr = "是"
      } else {
            cststr = "否"
      }

      switch (ast) {
      case -1:
            aststr = "无"
      case -2:
            aststr = "公敌"
      default:
            aststr = role_name[ast]
      }

      gls.Set (iter, seq (10), []interface{}{
            role_list_num,
            name,
            loc,
            int(math.Ceil(float64(money))),
            dirstr,
            mststr,
            mtime,
            cststr,
            aststr,
            atime,
      })

      role_list_num ++ 
}

func end_role_list () {
      win_role.SetTransientFor (win_main)
      select_setup (tv_role, false)
}

/* return: 0-nrole-1: select role to check. -1: exit. (waiting)
 * check_mcs: [out] mcs to check. -1: show role info only
 */
func get_role_select () int {
      if (st == GAME_OVER) {return -1}
      check_role_select = -1
      check_role_mcs_select = -1
      show_top_window(win_role)
      gtk.Main()
      if (st == GAME_OVER) {return -1}
      if (check_role_select < 0) {
            return -1
      } else {
            return check_role_select
      }
}

func begin_role_tech_list () {
      labels_clear (label_role_tech[:])
}

func add_role_tech_list (hind int) {
      label_role_tech[hind].SetSensitive (true)
}

func end_role_tech_list () {
}

var role_mos_list_num int

func begin_role_mos_list () {
      tv_clear(tv_role_mos)
      role_mos_list_num = 0
}

func add_role_mos_list (name, role string, job, loc int, hpmax, hp float32, prop [5]float32, is_quit bool, lst, pst int) {
      tv_add_people (role_mos_list_num, tv_role_mos, name, role, job, loc, hpmax, hp, prop, is_quit, lst, pst, -1)
      role_mos_list_num ++
}

func end_role_mos_list () {
      select_setup (tv_role_mos, false)
}

var role_mcs_list_num int
func begin_role_mcs_list () {
      tv_clear (tv_role_mcs)
      role_mcs_list_num = 0
}

func add_role_mcs_list (name string, scope int, hpmax, hp float32, role string, nmos int, mayor, treasurer string, fengshui float32) {
      tv_add_city (role_mcs_list_num, tv_role_mcs, name, scope, hpmax, hp, role, nmos, mayor, treasurer, fengshui)
      role_mcs_list_num ++ 
}

func end_role_mcs_list () {
      select_setup (tv_role_mcs, false)
}

var role_sos_list_num int 
func begin_role_sos_list () {
      tv_clear (tv_role_sos)
      role_sos_list_num = 0
}

func add_role_sos_list (name, role string, job, loc int, hpmax, hp float32, prop [5]float32, is_quit bool, lst, pst int) {
      tv_add_people (role_sos_list_num, tv_role_sos, name, role, job, loc, hpmax, hp, prop, is_quit, lst, pst, -1)
      role_sos_list_num ++ 
}

func end_role_sos_list () {
      select_setup (tv_role_sos, false)
}

var role_tos_list_num int 
func begin_role_tos_list () {
      tv_clear (tv_role_tos)
      role_tos_list_num = 0
}

func add_role_tos_list (name, role string, job, loc int, hpmax, hp float32, prop [5]float32, is_quit bool, lst, pst int) {
      tv_add_people (role_tos_list_num, tv_role_tos, name, role, job, loc, hpmax, hp, prop, is_quit, lst, pst, -1)
      role_tos_list_num ++ 
}

func end_role_tos_list () {
      select_setup (tv_role_tos, false)
}

var role_cards_list_num int 
func begin_role_cards_list () {
      tv_clear (tv_role_cards)
      role_cards_list_num  = 0
}

func add_role_cards_list (dind, num int) {

      gls_, _ := tv_role_cards.GetModel()
      gls := gls_.(* gtk.ListStore)
      iter := gls.Append ()

      gls.Set (iter, seq(3), []interface{} {
            role_cards_list_num, 
            card_name[dind],
            num,
      })

      role_cards_list_num ++ 
}

func end_role_cards_list () {
      select_setup (tv_role_cards, false)
}

/* called immediatly after each get_role_select() 
 * return -1: show base role info  
 * >=0: city id. show mcs city info 
 * if >=0: call begin_people_list ... get_people_exit in subsequence 
 */
func get_role_mcs_select () int {
      return check_role_mcs_select
}

var people_list_num int

// check people info
func begin_people_list (title string) {
      onoff_widget (button_people, false)
      onoff_widget (button_people_main, false)
      onoff_widget (button_people_vice, false)

      label_people_main.SetLabel ("无")
      label_people_vice.SetLabel ("无")
      label_people.SetLabel (title)

      tv_clear (tv_people)
      people_list_num  = 0
}

func add_people_list (name, role string, job, loc int, hpmax, hp float32, prop [5]float32, is_quit bool, lst, pst int) {
      tv_add_people (people_list_num, tv_people, name, role, job, loc, hpmax, hp, prop, is_quit, lst, pst, 0)
      people_list_num ++ 
}

// topof 0: main; 1: role 2: city
func end_people_list (topof int) {
      switch (topof) {
      case 0:
            win_people.SetTransientFor (win_main)
      case 1:
            win_people.SetTransientFor (win_role)
      case 2:
            win_people.SetTransientFor (win_city)
      }
}

/* return to exit check people info (waiting) */
func get_people_exit () {
      if (st == GAME_OVER) {return}
      select_setup (tv_people, false)
      people_select_st = 0
      show_top_window(win_people)
      gtk.Main()
}

/* return one select people. -1: cancel (waiting)*/
func get_people_one () int {
      if (st == GAME_OVER) {return -1}
      select_setup (tv_people, false)
      onoff_widget (button_people, true)
      people_select = -1
      people_select_st = 0
      show_top_window(win_people)
      gtk.Main ()
      if (st == GAME_OVER) {return -1}
      return people_select
}

/* return two selected people. main == -1: cancel. vice == -1: only main (waiting) */
func get_people_two_main () int {
      if (st == GAME_OVER) {return -1}
      select_setup (tv_people, false)
      onoff_widget (button_people, true)
      onoff_widget (button_people_main, true)
      onoff_widget (button_people_vice, true)
      people_select = -1
      people_select_vice = -1
      people_select_st = 1
      show_top_window(win_people)
      gtk.Main ()
      if (st == GAME_OVER) {return -1}
      return people_select
}

func get_people_two_vice () int {
      return people_select_vice
}

/* return a selected list of people. empty: cancel (waiting)*/
func get_people_list_num () int {
      if (st == GAME_OVER) {return 0}
      select_setup (tv_people, true)
      onoff_widget (button_people, true)
      people_select_num = 0
      people_select_st = 2
      show_top_window(win_people)
      gtk.Main ()
      if (st == GAME_OVER) {return 0}
      return people_select_num
}

func get_people_list (ind int) int {
      return people_select_list[ind]
}

var city_list_num int
func begin_city_list() {
      tv_clear(tv_city)
      city_list_num = 0
}

func add_city_list(name string, scope int, hpmax float32, hp float32, role string, nmos int, mayor, treasurer string, fengshui float32) {
      tv_add_city (city_list_num, tv_city, name, scope, hpmax, hp, role, nmos, mayor, treasurer, fengshui)
      city_list_num ++ 
}

func end_city_list() {
      win_city.SetTransientFor (win_main)
      select_setup (tv_city, false)
}

/* return -1: exit;  (waiting)
 *        >=0: id in mcs list
 * if >=0: call begin_people_list ... get_people_exit in subsequence 
 */
func get_city_mcs_select () int {
      if (st == GAME_OVER) {return -1}
      check_city_mcs_select = -1
      show_top_window(win_city)
      gtk.Main ()
      if (st == GAME_OVER) {return -1}
      return check_city_mcs_select
}

var inst_list_num int
func begin_inst_list() {
      tv_clear (tv_inst)
      tv_clear (tv_inst_mos)
      tv_clear (tv_inst_tech)
      inst_list_num = 0
}

func add_inst_list(name string, ntech int, mos, on_study string, point float32, left_round int) {

      gls_, _ := tv_inst.GetModel()
      gls := gls_.(* gtk.ListStore)
      iter := gls.Append ()

      gls.Set (iter, seq(7), []interface{} {
            inst_list_num,
            name,
            ntech, 
            mos, 
            on_study, 
            int(math.Ceil(float64(point))),
            left_round, 
      })

      inst_list_num ++ 

}

func end_inst_list() {
      win_inst.SetTransientFor (win_main)
      select_setup (tv_inst, false)
}

// -1: exit. >=0: inst index (waiting)
func get_inst_select () int {
      if (st == GAME_OVER) {return -1}
      check_inst_select = -1
      show_top_window(win_inst)
      gtk.Main ()
      if (st == GAME_OVER) {return -1}
      return check_inst_select
}

var inst_tech_list_num int 

func begin_inst_tech_list (title string) {
      tv_clear (tv_inst_tech)
      inst_tech_list_num = 0
}

func add_inst_tech_list (name string, study float32, scond int) {
      tv_add_tech (inst_tech_list_num, tv_inst_tech, name, study, scond)
      inst_tech_list_num ++ 
}

func end_inst_tech_list () {
      select_setup (tv_inst_tech, false)
}

func show_inst_mos (name, role string, job, loc int, hpmax, hp float32, prop [5]float32, is_quit bool, lst, pst int) {
      tv_clear (tv_inst_mos)
      tv_add_people (0, tv_inst_mos, name, role, job, loc, hpmax, hp, prop, is_quit, lst, pst, -1)
      select_setup (tv_inst_mos, false)
}

var train_list_num int 
func begin_train_list() {
      tv_clear (tv_train)
      tv_clear (tv_train_mos)
      train_list_num = 0
}

func add_train_list(name string, item mbg.Property, mos string, round int) {

      gls_, _ := tv_train.GetModel()
      gls := gls_.(* gtk.ListStore)
      iter := gls.Append ()

      gls.Set (iter, seq(5), []interface{} {
            train_list_num,
            name,
            prop_name[item], 
            mos, 
            round, 
      })
      train_list_num ++ 
}

func end_train_list() {
      win_train.SetTransientFor (win_main)
      select_setup (tv_train, false)
}

// -1: exit. >=0: train index (waiting)
func get_train_select () int {
      if (st == GAME_OVER) {return -1}
      check_train_select = -1
      show_top_window(win_train)
      gtk.Main ()
      if (st == GAME_OVER) {return -1}
      return check_train_select
}

func show_train_mos (name, role string, job, loc int, hpmax, hp float32, prop [5]float32, is_quit bool, lst, pst int) {
      tv_clear (tv_train_mos)
      tv_add_people (0, tv_train_mos, name, role, job, loc, hpmax, hp, prop, is_quit, lst, pst, -1)
      select_setup (tv_train_mos, false)
}

var saloon_list_num int 
func begin_saloon_list() {
      onoff_widget (button_people, true)
      onoff_widget (button_people_main, false)
      onoff_widget (button_people_vice, false)
      label_people_main.SetLabel ("无"    )
      label_people_vice.SetLabel ("无"    )
      label_people.SetLabel      ("请选择聘用人员")

      tv_clear (tv_people)
      saloon_list_num = 0
}

func add_saloon_list (name, role string, job, loc int, hpmax, hp float32, prop [5]float32, is_quit bool, lst, pst int, price float32) {
      tv_add_people (saloon_list_num, tv_people, name, role, job, loc, hpmax, hp, prop, is_quit, lst, pst, price)
      saloon_list_num ++ 
}

func end_saloon_list() {
      win_people.SetTransientFor (win_main)
      select_setup (tv_people, false)
}

func get_saloon_select () int {
      if (st == GAME_OVER) {return -1}
      people_select_st = 0
      people_select = -1
      show_top_window(win_people)
      gtk.Main ()
      if (st == GAME_OVER) {return -1}
      return people_select
}

// choose battle scale. -1: cancel -2: check (waiting)
func sel_battle_scale (msg string) int {
      if (st == GAME_OVER) {return -1}
      msg_map (msg)
      onoff_widget (button_main_small, true)
      onoff_widget (button_main_medium, true)
      onoff_widget (button_main_large, true)
      onoff_widget (button_main_no, true)
      battle_scale = -1
      st = WAIT_BATTLE_SCALE
      gtk.Main ()
      if (st == GAME_OVER) {return -1}
      onoff_widget (button_main_small, false)
      onoff_widget (button_main_medium, false)
      onoff_widget (button_main_large, false)
      onoff_widget (button_main_no, false)
      return battle_scale
}

// open battle field.
func battle_start (max_turn_ int) {
      var i, j int
      buf_clear (buf_battle_msg)

      for i = 0; i < 2; i ++ {
            labels_clear (label_battle_tech[i][:])
            for j = 0; j < 3; j ++ {
                  label_set_color_no (label_battle_st [i][j])
            }
      }

      tb_battle_auto.SetActive (false)

      max_turn = max_turn_
      is_battle_finish = false
      wait_battle_tick = false
      show_top_window(win_battle)
}

func battle_msg (msg string) {
      buf_add_msg (buf_battle_msg, msg)
      swin_set_follow (swin_battle_msg)
}

func battle_set (side int, role string, power, winp, celve, hpmax_, hp, def float32) {
      label_battle_role[side].SetLabel (role)

      pb_set_value (pb_battle_power[side], power, 100)
      pb_set_value (pb_battle_winp[side], winp, 100)
      pb_set_value (pb_battle_celve[side], celve, 100)
      pb_set_value (pb_battle_hp[side], hp, hpmax_)
      pb_set_value (pb_battle_def[side], def, defmax)

      hpmax[side] = hpmax_
}

func battle_tech_set (side int, tech int) {
      onoff_widget (label_battle_tech[side][tech], true)
}

func battle_change_hp (side int, hp, def float32) {
      pb_set_value (pb_battle_hp [side], hp,  hpmax[side])
      pb_set_value (pb_battle_def[side], def, defmax)
}

func battle_change_st (side, bst int, is_quit bool, btime, fst int) {
      if (fst != 0) {
            label_set_color_red (label_battle_st[side][0])
      } else {
            label_set_color_no  (label_battle_st[side][0])
      }

      if (bst == 1) {
            label_set_color_red (label_battle_st[side][1])
      } else {
            label_set_color_no  (label_battle_st[side][1])
      }

      if (bst == 2) {
            label_set_color_red (label_battle_st[side][2])
      } else {
            label_set_color_no  (label_battle_st[side][2])
      }
}

func battle_active_tech (side , tech int, shp, ohp, sdef, odef float32) {
      label_set_color_blue (label_battle_tech[side][tech])
      pb_set_value (pb_battle_hp[side], shp, hpmax[side])
      pb_set_value (pb_battle_hp[1 - side], ohp, hpmax[1 - side])
      pb_set_value (pb_battle_def[side], sdef, defmax)
      pb_set_value (pb_battle_def[1 - side], odef, defmax)
}

// waiting for battle end (timing)
func battle_turn_start (turn int) {
      var i, j int
      if (st == GAME_OVER) {return}
      if (is_battle_finish) {return}
      battle_set_turn (turn)
      wait_battle_tick = true
      gtk.Main ()
      for i = 0; i < 2; i ++ {
            for j = 0; j < ntech; j ++ {
                  label_set_color_no (label_battle_tech[i][j])
            }
      }
}

// close battle field.
func battle_end () {
      if (st == GAME_OVER) {return}
      if (is_battle_finish) {return}
      gtk.Main ()
}

var exchange_people_list_num [2]int 

// side = 0: role. side = 1: city
// side = 0 to open the select window.
func begin_exchange_people_list (side int, title string) {
      if (side != 0) {
            tv_clear (tv_exchange_out)
      } else {
            tv_clear (tv_exchange_in)
      }

      label_exchange.SetLabel (title)
      exchange_people_list_num [side] = 0
}

func add_exchange_people_list (side int,  name, role string, job, loc int, hpmax, hp float32, prop [5]float32, is_quit bool, lst, pst int) {
      if (side != 0) {
            tv_add_people (exchange_people_list_num[side], tv_exchange_out, name, role, job, loc, hpmax, hp, prop, is_quit, lst, pst, -1)
      } else {
            tv_add_people (exchange_people_list_num[side], tv_exchange_in, name, role, job, loc, hpmax, hp, prop, is_quit, lst, pst, -1)
      }
      exchange_people_list_num[side] ++
}

func end_exchange_people_list (side int) {
      if (side != 0) {
            select_setup (tv_exchange_out, true)
            win_exchange.SetTransientFor (win_main)
      } else {
            select_setup (tv_exchange_in, true)
      }
}

// side = 0: 0->1. 1: 1->0
// side = 1: return immediatly. side = 0: return after select complete(waiting).
func get_exchange_people_list_num (side int) int {
      if (side != 0) {
            return exchange_out_num
      } else {
            if (st == GAME_OVER) {return 0}
            exchange_out_num = 0
            exchange_in_num = 0
            show_top_window(win_exchange)
            gtk.Main()
            if (st == GAME_OVER) {return 0}
            return exchange_in_num
      }
}

func get_exchange_people_list (side, isel int) int {
      if (side != 0) {
            return exchange_out_list [isel]
      } else {
            return exchange_in_list [isel]
      }
}

var study_tech_list_num int 
func begin_study_tech_list (name string) {
      label_study.SetLabel (name)
      tv_clear (tv_study_tech)
      study_tech_list_num = 0
}

func add_study_tech_list (name string, study float32, scond int) {
      tv_add_tech (study_tech_list_num, tv_study_tech, name, study, scond)
      study_tech_list_num ++
}

func end_study_tech_list () {
      select_setup (tv_study_tech, false)
}

var study_people_list_num int 

func begin_study_people_list () {
      tv_clear (tv_study_people)
      study_people_list_num = 0
}

func add_study_people_list (name, role string, job, loc int, hpmax, hp float32, prop [5]float32, is_quit bool, lst, pst int) {
      tv_add_people (study_people_list_num, tv_study_people, name, role, job, loc, hpmax, hp, prop, is_quit, lst, pst, -1)
      study_people_list_num ++
}

func end_study_people_list () {
      select_setup (tv_study_people, false)
      win_study.SetTransientFor (win_main)
}

// return people list num. -1: cancel. (waiting)
// tech [out] tech list num.
func get_study_select () (tech, people int) {
      tech, people = -1, -1
      if (st == GAME_OVER) { return } 
      study_tech = -1
      study_people = -1
      show_top_window(win_study)
      gtk.Main ()
      if (st == GAME_OVER) { return } 
      tech = study_tech
      people = study_people
      return 
}

// 在地图上仅开启可用位置
func show_map_av (av_loc []bool) {
	for i := 0; i < ngrid; i ++ {
		if !av_loc[i] {
			map_point[i].SetSensitive (false) 
		}
	}
}

// 开启地图所有位置
func clear_map_av () {
	for i := 0; i < ngrid; i ++ {
		map_point[i].SetSensitive (true)
	}
}

// 返回位置: -1: 取消. -2: map. >=0: role. -3: check. (waiting)
func get_obj_type () int {
      if (st == GAME_OVER) {return -1}
      obj_type = -1
      obj = -1
      st = SEL_CARD_OBJ
      onoff_widget (button_main_no, true)
      gtk.Main ()
      if (st == GAME_OVER) {return -1}
      onoff_widget (button_main_no, false)
      return obj_type
}

// return immediatly (loc, for get_obj_type == -2)
func get_obj () int {
      return obj
}

// res: -1: 平局. >=0 取胜方 -2: 用户强制退出
func game_end (res int) {
      if (st != GAME_OVER) {
            st = GAME_OVER
            if (res != -2) {
                  gtk.Main ()
            }
      }

      win_main.Destroy()
      win_card.Destroy()
      win_battle.Destroy()
      win_role.Destroy()
      win_people.Destroy()
      win_city.Destroy()
      win_inst.Destroy()
      win_train.Destroy()
      win_study.Destroy()
      win_exchange.Destroy()
}

