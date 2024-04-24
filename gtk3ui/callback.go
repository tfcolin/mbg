package gtk3ui

import (
      "github.com/gotk3/gotk3/gtk"
      "github.com/gotk3/gotk3/gdk"
)

func get_select_id (tv * gtk.TreeView, is_name bool) (name string, id int) {
      ts, _ := tv.GetSelection()
      model_, iter, is_sel := ts.GetSelected()
      model := model_.(*gtk.TreeModel)

      name, id = "", -1

      if (is_sel) {
            v, _ := model.GetValue (iter, 0) 
            id_, _ := v.GoValue()
            id = id_.(int)
            if (is_name) {
                  v, _ := model.GetValue (iter, 1)
                  name, _ = v.GetString ()
            }
      } 

      return 
}

func put_select_ids (tv * gtk.TreeView, spf gtk.TreeSelectionForeachFunc ) {
      gts, _ := tv.GetSelection()
      gts.SelectedForEach (spf)
}

func spf_people_select (model * gtk.TreeModel, path * gtk.TreePath, iter * gtk.TreeIter) {
      v, _ := model.GetValue (iter, 0) 
      id_, _ := v.GoValue()
      id := id_.(int)
      people_select_list [people_select_num] = id
      people_select_num ++
}

func spf_exchange_select_in  (model * gtk.TreeModel, path * gtk.TreePath, iter * gtk.TreeIter) {
      v, _ := model.GetValue (iter, 0) 
      id_, _ := v.GoValue()
      id := id_.(int)
      exchange_in_list [exchange_in_num] = id
      exchange_in_num ++
}

func spf_exchange_select_out (model * gtk.TreeModel, path * gtk.TreePath, iter * gtk.TreeIter) {
      v, _ := model.GetValue (iter, 0) 
      id_, _ := v.GoValue()
      id := id_.(int)
      exchange_out_list [exchange_out_num] = id
      exchange_out_num ++ 
}

func cbf_main_board_click_maker (loc int) func (w * gtk.Button) {
      return func (w * gtk.Button) {
            if (st == GAME_OVER) { return }
            if (st != SEL_CARD_OBJ) { return }
            obj_type = -2
            obj = loc
            gtk.MainQuit()
      }
}

func cbf_main_role_click_maker (loc, rind int) func (w * gtk.Button) {
      return func (w * gtk.Button) {
            if (st == GAME_OVER) { return }
            if (st != SEL_CARD_OBJ)  {return}
            if (!is_role[loc][rind]) {return}
            obj_type = rind
            gtk.MainQuit()
      }
}

func cbf_main_move_click (w * gtk.Button) {
      if (st == GAME_OVER) {return}
      if (st != WAIT_ACTION) {return}
      action = 0
      gtk.MainQuit()
}

func cbf_main_card_click (w * gtk.Button) {
      if (st == GAME_OVER) {return}
      if (st != WAIT_ACTION) {return}
      action = 1
      gtk.MainQuit()
}

func  cbf_main_save_click (w * gtk.Button) {
      if (st == GAME_OVER) {return}
      if (st != WAIT_ACTION) {return}
	action = 3
	gtk.MainQuit()
}

func set_check_status () {
      switch (st) {
      case WAIT_ACTION:
            action = 2
      case SEL_CARD_OBJ:
            obj_type = -3
      case WAIT_BATTLE_SCALE:
            battle_scale = -2
      case WAIT_MSG_ANSWER:
            msg_res = 1
      }
}

func cbf_main_role_check (w * gtk.Button) {
      if (st == GAME_OVER) {return}
      if (st == WAIT_MSG_CONFIRM) {return}
      set_check_status()
      check_type = 0
      gtk.MainQuit()
}

func cbf_main_people_check (w * gtk.Button) {
      if (st == GAME_OVER) {return}
      if (st == WAIT_MSG_CONFIRM) {return}
      set_check_status()
      check_type = 1
      gtk.MainQuit()
}

func cbf_main_rp_check (w * gtk.Button) {
      if (st == GAME_OVER) {return}
      if (st == WAIT_MSG_CONFIRM) {return}
      set_check_status()
      check_type = 5
      gtk.MainQuit()
}

func cbf_main_city_check (w * gtk.Button) {
      if (st == GAME_OVER) {return}
      if (st == WAIT_MSG_CONFIRM) {return}
      set_check_status()
      check_type = 2
      gtk.MainQuit()
}

func cbf_main_inst_check (w * gtk.Button) {
      if (st == GAME_OVER) {return}
      if (st == WAIT_MSG_CONFIRM) {return}
      set_check_status()
      check_type = 3
      gtk.MainQuit()
}

func cbf_main_train_check (w * gtk.Button) {
      if (st == GAME_OVER) {return}
      if (st == WAIT_MSG_CONFIRM) {return}
      set_check_status()
      check_type = 4
      gtk.MainQuit()
}

func cbf_main_yes_click (w * gtk.Button) {
      if (st != WAIT_MSG_ANSWER && st != WAIT_MSG_CONFIRM) {return}
      msg_res = 0
      gtk.MainQuit()
}

func cbf_main_no_click (w * gtk.Button) {
      switch (st) {
      case WAIT_MSG_ANSWER:
            msg_res = -1
      case WAIT_BATTLE_SCALE:
            battle_scale = -1
      case SEL_CARD_OBJ:
            obj_type = -1
            obj = -1
      default:
            return
      }
      gtk.MainQuit()
}

func cbf_main_small_click (w * gtk.Button) {
      if (st != WAIT_BATTLE_SCALE) {return}
      battle_scale = 0
      gtk.MainQuit()
}

func cbf_main_medium_click (w * gtk.Button) {
      if (st != WAIT_BATTLE_SCALE) {return}
      battle_scale = 1
      gtk.MainQuit()
}

func cbf_main_large_click (w * gtk.Button) {
      if (st != WAIT_BATTLE_SCALE) {return}
      battle_scale = 2
      gtk.MainQuit()
}

func cbf_battle_next_click (w * gtk.Button) {
      if (st == GAME_OVER) {return}
      if (!wait_battle_tick) {return}
      wait_battle_tick = false
      gtk.MainQuit()
}

func cbf_battle_timer () bool {
      if (st == GAME_OVER) {return true}
      if (!is_align) {return false}
      is_auto := tb_battle_auto.GetActive()
      if (!is_auto) {return true}
      if (!wait_battle_tick) {return true}
      wait_battle_tick = false
      gtk.MainQuit()
      return true
}

func cbf_people_click (widget * gtk.Button) {
      if (st == GAME_OVER) {return}

      switch (people_select_st) {
      case 0:
            _, people_select = get_select_id (tv_people, false)
      case 2:
            put_select_ids (tv_people, spf_people_select)
      }

      win_people.Hide()
      gtk.MainQuit()
}

func cbf_people_main_click (widget * gtk.Button) {
      if (st == GAME_OVER) {return}
      if (people_select_st != 1) {return}

      var name string
      name, people_select = get_select_id (tv_people, true)
      label_people_main.SetLabel (name)
}

func cbf_people_vice_click (widget * gtk.Button) {
      if (st == GAME_OVER) { return }
      if (people_select_st != 1) { return }

      var name string
      name, people_select_vice = get_select_id (tv_people, true)
      label_people_vice.SetLabel (name)
}

func cbf_city_click (widget * gtk.Button) {
      if (st == GAME_OVER){ return}
      _, check_city_mcs_select = get_select_id (tv_city, false)
      gtk.MainQuit()
}

func cbf_card_click (widget * gtk.Button) {
      if (st == GAME_OVER){ return}
      _, card_action = get_select_id (tv_card, false)
      win_card.Hide()
      gtk.MainQuit()
}

func cbf_inst_click (widget * gtk.Button) {
      if (st == GAME_OVER){ return}
      _, check_inst_select = get_select_id (tv_inst, false)
      gtk.MainQuit()
}

func cbf_train_click (widget * gtk.Button) {
      if (st == GAME_OVER){ return}
      _, check_train_select = get_select_id(tv_train, false)
      gtk.MainQuit()
}

func cbf_exchange_click  (widget * gtk.Button) {
      if (st == GAME_OVER) {return}
      put_select_ids (tv_exchange_in, spf_exchange_select_in)
      put_select_ids (tv_exchange_out, spf_exchange_select_out)
      win_exchange.Hide()
      gtk.MainQuit()
}

func cbf_study_click (widget * gtk.Button) {
      if (st == GAME_OVER) {return}
      _, study_tech = get_select_id (tv_study_tech, false)
      _, study_people = get_select_id (tv_study_people, false)
      if (study_tech == -1 || study_people == -1) {return}
      win_study.Hide()
      gtk.MainQuit()
}

func cbf_role_click (widget * gtk.Button) {
      if (st == GAME_OVER) {return}
      _, check_role_select = get_select_id (tv_role, false)
      check_role_mcs_select = -1
      gtk.MainQuit()
      return
}

func cbf_role_mcs_click (widget * gtk.Button) {
      if (st == GAME_OVER) {return}
      _, check_role_select = get_select_id (tv_role, false)
      _, check_role_mcs_select = get_select_id (tv_role_mcs, false)
      gtk.MainQuit()
      return
}

func cbf_main_close (widget * gtk.Window, event * gdk.Event) bool {
      st = GAME_OVER
      widget.Hide()
      gtk.MainQuit()
      return false
}

func cbf_card_close (widget * gtk.Window, event * gdk.Event) bool {
      card_action = -1
      widget.Hide()
      gtk.MainQuit()
      return true
}

func cbf_battle_close (widget * gtk.Window, event * gdk.Event) bool {
      wait_battle_tick = false
      is_battle_finish = true
      widget.Hide()
      gtk.MainQuit()
      return true
}

func cbf_role_close (widget * gtk.Window, event * gdk.Event) bool {
      check_role_select = -1
      check_role_mcs_select = -1
      widget.Hide()
      gtk.MainQuit()
      return true
}

func cbf_people_close (widget * gtk.Window, event * gdk.Event) bool {
      people_select = -1
      people_select_vice = -1
      people_select_num = 0
      widget.Hide()
      gtk.MainQuit()
      return true
}

func cbf_city_close (widget * gtk.Window, event * gdk.Event) bool {
      check_city_mcs_select = -1
      widget.Hide()
      gtk.MainQuit()
      return true
}

func cbf_inst_close (widget * gtk.Window, event * gdk.Event) bool {
      check_inst_select = -1
      widget.Hide()
      gtk.MainQuit()
      return true
}

func cbf_train_close (widget * gtk.Window, event * gdk.Event) bool {
      check_train_select = -1
      widget.Hide()
      gtk.MainQuit()
      return true
}

func cbf_study_close (widget * gtk.Window, event * gdk.Event) bool {
      study_tech, study_people = -1, -1
      widget.Hide()
      gtk.MainQuit()
      return true
}

func cbf_exchange_close (widget * gtk.Window, event * gdk.Event) bool {
      exchange_in_num, exchange_out_num = 0, 0
      widget.Hide()
      gtk.MainQuit()
      return true
}

