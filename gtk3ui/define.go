package gtk3ui

import (
	"gitee.com/tfcolin/mbg"
      "github.com/gotk3/gotk3/gtk"
)

type GtkUserView struct {
	d *mbg.Driver
}

type GtkOperationInterface struct {
	d    *mbg.Driver
	rind int
}

const (
      GAME_OVER = 0
      WAIT_ACTION = 1
      SEL_CARD_OBJ = 2
      WAIT_BATTLE_SCALE = 3
      WAIT_MSG_ANSWER = 4 
      WAIT_MSG_CONFIRM = 5
)
type Status int

const (
      MAX_LIST_LENGTH = 2048
)

var (
      prop_name = []string {"武术", "战术", "策略", "政治", "经济"}
      st Status

      // 返回位置: -1: 取消. -2: map. >=0: role. -3: check. (waiting)
      obj_type, obj int
      // 0: 移动. 1: 使用卡片. 2: check(waiting). -1: 无效. 
      action int
      // -1: cancel. >=0: 卡片号
      card_action int
      // 0: role. 1: people. 2: city; 3:inst. 4:trainingroom. 5: role_people
      check_type int 
      check_role_select, check_role_mcs_select, check_city_mcs_select, check_inst_select, check_train_select int

      // 0: 单选一人; 1: 选主副将; 2: 选列表
      people_select_st, people_select, people_select_vice int
      people_select_list []int // [MAX_LIST_LENGTH]
      people_select_num int
      // choose battle scale. -1: cancel -2: check (waiting)
      battle_scale int 
	// 是否在等待一个战斗回合的结束.
      wait_battle_tick bool
	// 战斗窗口已被关闭或战斗已结束
      is_battle_finish bool
      exchange_out_list, exchange_in_list []int // [MAX_LIST_LENGTH]
      exchange_out_num, exchange_in_num int
      study_tech, study_people int
      /* return 0: OK. -1: CANCEL; 1: CHECK*/
      msg_res int
      ngrid, nrole, npeople, ncity, ninst, ntrain, ntech, ncard, nx, ny int
      wp, hp int
      is_align bool
      role_name, tech_name, card_name []string // [nrole], [ntech], [ncard]
      tech_scond, card_odir, card_otype []int // [ntech], [ncard], [ncard]
      // 0: 空地 1:城 2:会馆 3:谋略研究所 4:修炼所 5: 机会
      map_type, map_obj []int // [ngrid]
      city_name, inst_name, train_name []string // [ncity], [ninst], [ntrain]
      is_barrier, is_robber []bool // [ngrid]
      is_role [][]bool // [ngrid][nrole]

      win_battle * gtk.Window
      label_battle * gtk.Label // title
      label_battle_role [2](* gtk.Label) // label_battle_role_[01]
      pb_battle_power [2](* gtk.ProgressBar)
      pb_battle_winp [2](* gtk.ProgressBar) // 胜率, 战法
      pb_battle_celve [2](* gtk.ProgressBar) // 策略
      pb_battle_hp [2](* gtk.ProgressBar)
      pb_battle_def [2](* gtk.ProgressBar)
      label_battle_tech [2][12] (* gtk.Label) // tech
      label_battle_st [2][3] (* gtk.Label) // 状态: 0: burn; 1: 混乱. 2: 内讧
      tv_battle_msg * gtk.TextView // 战斗消息
      button_battle_next * gtk.Button
      tb_battle_auto * gtk.ToggleButton
      buf_battle_msg * gtk.TextBuffer
      swin_battle_msg * gtk.ScrolledWindow

      win_people * gtk.Window
      label_people * gtk.Label
      tv_people * gtk.TreeView
      button_people * gtk.Button
      button_people_main * gtk.Button
      button_people_vice * gtk.Button
      label_people_main * gtk.Label
      label_people_vice * gtk.Label

      win_city * gtk.Window
      tv_city * gtk.TreeView
      button_city * gtk.Button

      win_card * gtk.Window
      tv_card * gtk.TreeView
      button_card * gtk.Button

      win_inst * gtk.Window
      tv_inst * gtk.TreeView
      tv_inst_mos * gtk.TreeView
      tv_inst_tech * gtk.TreeView
      button_inst * gtk.Button

      win_train * gtk.Window
      tv_train * gtk.TreeView
      tv_train_mos * gtk.TreeView
      button_train * gtk.Button

      win_exchange * gtk.Window
      label_exchange * gtk.Label
      tv_exchange_in * gtk.TreeView
      tv_exchange_out * gtk.TreeView
      button_exchange * gtk.Button

      win_study * gtk.Window
      label_study * gtk.Label
      tv_study_tech * gtk.TreeView
      tv_study_people * gtk.TreeView
      button_study * gtk.Button

      win_role * gtk.Window
      tv_role * gtk.TreeView
      tv_role_mos * gtk.TreeView
      tv_role_mcs * gtk.TreeView
      tv_role_sos * gtk.TreeView
      tv_role_tos * gtk.TreeView
      tv_role_cards * gtk.TreeView
      label_role_tech [12](* gtk.Label)
      button_role * gtk.Button
      button_role_mcs * gtk.Button

      win_main * gtk.Window
      layout_main_map * gtk.Layout
      button_main_move * gtk.Button
      button_main_card * gtk.Button
      button_main_save * gtk.Button
      button_main_role * gtk.Button
      button_main_people * gtk.Button
      button_main_rp * gtk.Button
      button_main_city * gtk.Button
      button_main_inst * gtk.Button
      button_main_train * gtk.Button
      tv_main_msg * gtk.TextView
      button_main_yes * gtk.Button
      button_main_no * gtk.Button
      button_main_small * gtk.Button
      button_main_medium * gtk.Button
      button_main_large * gtk.Button

      map_point [] (* gtk.Button) // [ngrid]
      role_point [][] (* gtk.Button) // [ngrid][nrole]

      buf_main_msg * gtk.TextBuffer
      swin_main_msg * gtk.ScrolledWindow

	file_save_dialog * gtk.FileChooserDialog

      max_turn int
      defmax float32 = 3000
      hpmax [2]float32

)

