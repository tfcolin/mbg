#include <gtk/gtk.h>

#include "gtkui_wrapper.h"
#include "gtkui.h"

#include <stdio.h>
#include <stdlib.h>
#include <math.h>

char prop_name [5][64] = {"武术", "战术", "策略", "政治", "经济"}; 

Status st;

// 返回位置: -1: 取消. -2: map. >=0: role. -3: check. (waiting)
int obj_type;
int obj;

// 0: 移动. 1: 使用卡片. 2: check(waiting). -1: 无效. 
int action; 
// -1: cancel. >=0: 卡片号
int card_action;

// 0: role. 1: people. 2: city; 3:inst. 4:trainingroom. 5: role_people
int check_type;
int check_role_select;
int check_role_mcs_select;
int check_city_mcs_select;
int check_inst_select;
int check_train_select;

int people_select_st; // 0: 单选一人; 1: 选主副将; 2: 选列表
int people_select;
int people_select_vice;
int people_select_list [2048];
int people_select_num;

// choose battle scale. -1: cancel -2: check (waiting)
int battle_scale;
int is_battle_finish;
int wait_battle_tick;

int exchange_out_num;
int exchange_out_list [2048];
int exchange_in_num;
int exchange_in_list [2048];

int study_tech;
int study_people;

/* return 0: OK. -1: CANCEL; 1: CHECK*/
int msg_res; // msg_button. 

int ngrid, nrole, npeople, ncity, ninst, ntrain, ntech, ncard, nx, ny;
int wp, hp;
int is_align;

char ** role_name;

char ** tech_name;
int  * tech_scond;

char ** card_name;
int  * card_odir;
int  * card_otype;

// 0: 空地 1:城 2:会馆 3:谋略研究所 4:修炼所 5: 机会
int  * map_type; // [ngrid]
int  * map_obj;  // [ngrid] object index

int * is_barrier; // [ngrid]
int * is_robber; // [ngrid]
int ** is_role; //[ngrid][nrole]

char ** city_name;
char ** inst_name;
char ** train_name;

GtkWidget * win_battle; 
GtkWidget * label_battle; // title
GtkWidget * label_battle_role[2]; // label_battle_role_[01]
GtkWidget * pb_battle_power[2];
GtkWidget * pb_battle_winp[2]; // 胜率, 战法
GtkWidget * pb_battle_celve[2]; // 策略
GtkWidget * pb_battle_hp[2]; 
GtkWidget * pb_battle_def[2]; 
GtkWidget * label_battle_tech[2][12]; // tech
GtkWidget * label_battle_st[2][3]; // 状态: 0: burn; 1: 混乱. 2: 内讧
GtkWidget * tv_battle_msg; // 战斗消息
GtkWidget * button_battle_next; 
GtkWidget * tb_battle_auto; 
GtkTextBuffer * buf_battle_msg;
GtkWidget * swin_battle_msg;

GtkWidget * win_people;
GtkWidget * label_people;
GtkWidget * tv_people;
GtkWidget * button_people;
GtkWidget * button_people_main;
GtkWidget * button_people_vice;
GtkWidget * label_people_main;
GtkWidget * label_people_vice;

GtkWidget * win_city;
GtkWidget * tv_city;
GtkWidget * button_city;

GtkWidget * win_card;
GtkWidget * tv_card;
GtkWidget * button_card;

GtkWidget * win_inst;
GtkWidget * tv_inst;
GtkWidget * tv_inst_mos;
GtkWidget * tv_inst_tech;
GtkWidget * button_inst;

GtkWidget * win_train;
GtkWidget * tv_train;
GtkWidget * tv_train_mos;
GtkWidget * button_train;

GtkWidget * win_exchange;
GtkWidget * label_exchange;
GtkWidget * tv_exchange_in;
GtkWidget * tv_exchange_out;
GtkWidget * button_exchange;

GtkWidget * win_study;
GtkWidget * label_study;
GtkWidget * tv_study_tech;
GtkWidget * tv_study_people;
GtkWidget * button_study;

GtkWidget * win_role;
GtkWidget * tv_role;
GtkWidget * tv_role_mos;
GtkWidget * tv_role_mcs;
GtkWidget * tv_role_sos;
GtkWidget * tv_role_tos;
GtkWidget * tv_role_cards;
GtkWidget * label_role_tech[12];
GtkWidget * button_role;
GtkWidget * button_role_mcs;

GtkWidget * win_main;
GtkWidget * layout_main_map;
GtkWidget * button_main_move;
GtkWidget * button_main_card;
GtkWidget * button_main_role;
GtkWidget * button_main_people;
GtkWidget * button_main_rp;
GtkWidget * button_main_city;
GtkWidget * button_main_inst;
GtkWidget * button_main_train;
GtkWidget * tv_main_msg;
GtkWidget * button_main_yes;
GtkWidget * button_main_no;
GtkWidget * button_main_small;
GtkWidget * button_main_medium;
GtkWidget * button_main_large;

GtkWidget ** map_point; // [ngrid]
GtkWidget *** role_point; // [ngrid][nrole]

GtkTextBuffer * buf_main_msg;
GtkWidget * swin_main_msg;

static void swin_set_follow (GtkWidget * swin);
static void select_setup (GtkWidget * tv, int is_multi);
static void buf_add_msg (GtkTextBuffer * buffer, char * msg);
static void onoff_widget (GtkWidget * widget, int is_on);
static void tv_add_column (GtkWidget * tv, int id, char * title);
static void tv_clear (GtkWidget * tv);
static void tv_add_people (int id, GtkWidget * tv, char * name, char * role, int job, int loc, float hpmax, float hp, float * prop, int is_quit, int lst, int pst, float price);
static void tv_add_city (int id, GtkWidget * tv, char * name, int scope, float hpmax, float hp, char * role, int nmos, char * mayor, char * treasurer, float fengshui);
static void tv_add_tech (int id, GtkWidget * tv, char * name, float study, int scond);
static void labels_clear (GtkWidget ** labels, int n);

static void label_set_color_blue (GtkWidget * label);
static void label_set_color_red (GtkWidget * label);
static void label_set_color_no (GtkWidget * label);
static void pb_set_value (GtkWidget * pb, float value, float max_value);
static void buf_clear (GtkTextBuffer * buf);

static void battle_set_turn (int turn);

static int max_turn;
static int defmax = 3000;
static int hpmax[2];

static int * loc_ind;
static RoleLocInd ** rloc_ind;

void swin_set_follow (GtkWidget * swin_)
{
      GtkScrolledWindow * swin = GTK_SCROLLED_WINDOW (swin_);
      GtkAdjustment * adjust = gtk_scrolled_window_get_vadjustment (swin);
      double upper = gtk_adjustment_get_upper (adjust);
      gtk_adjustment_set_value (adjust, upper);
}

void buf_clear (GtkTextBuffer * buf)
{
      gtk_text_buffer_set_text (buf, "", -1);
}

void battle_set_turn (int turn)
{
      char title[lstr] = "";
      sprintf (title, "第 %d / %d 回合", turn + 1, max_turn);
      gtk_label_set_label (GTK_LABEL(label_battle), title);
}

void pb_set_value (GtkWidget * pb, float value, float max_value)
{
      float frac = max_value == 0 ? 0 : value / max_value;
      char msg[lstr] = "";

      sprintf (msg, "%.0f / %.0f", value, max_value);

      gtk_progress_bar_set_fraction (GTK_PROGRESS_BAR(pb), frac);
      gtk_progress_bar_set_text (GTK_PROGRESS_BAR(pb), msg);
}

void label_set_color_no (GtkWidget * label)
{
      PangoAttrList * pal = pango_attr_list_new();
      gtk_label_set_attributes (GTK_LABEL(label), pal);
      pango_attr_list_unref(pal);
}

void label_set_color_blue (GtkWidget * label)
{
      PangoAttrList * pal = pango_attr_list_new();
      PangoAttribute * attr = pango_attr_foreground_new (0x0000, 0x0000, 0xcccc);
      pango_attr_list_change (pal, attr);
      gtk_label_set_attributes (GTK_LABEL(label), pal);
      pango_attr_list_unref(pal);
}

void label_set_color_red (GtkWidget * label)
{
      PangoAttrList * pal = pango_attr_list_new();
      PangoAttribute * attr = pango_attr_foreground_new (0xcccc, 0x0000, 0x0000);
      pango_attr_list_change (pal, attr);
      gtk_label_set_attributes (GTK_LABEL(label), pal);
      pango_attr_list_unref(pal);
}

void tv_clear (GtkWidget * tv)
{
      GtkListStore * gls = GTK_LIST_STORE(gtk_tree_view_get_model (GTK_TREE_VIEW (tv)));
      gtk_list_store_clear (gls);
}

void tv_add_people (int id, GtkWidget * tv, char * name, char * role, int job, int loc, float hpmax, float hp, float * prop, int is_quit, int lst, int pst, float price)
{
      GtkListStore * gls = GTK_LIST_STORE(gtk_tree_view_get_model (GTK_TREE_VIEW (tv)));
      GtkTreeIter iter;

      char jobstr[32], locstr[32], isqstr[32];
      switch (job)
      {
      case -1:
            strcpy (jobstr, "无");
            break;
      case 0:
            strcpy (jobstr, "市长");
            break;
      case 1:
            strcpy (jobstr, "财务官");
            break;
      case 2:
            strcpy (jobstr, "市民");
            break;
      case 3:
            strcpy (jobstr, "研究");
            break;
      case 4:
            strcpy (jobstr, "训练");
            break;
      }

      if (loc == -1) 
      {
            strcpy (locstr, "随军");
      }
      else 
      {
            switch (job)
            {
            case -1:
                  strcpy (locstr, "随军");
                  break;
            case 0: case 1: case 2:
                  strcpy (locstr, city_name[loc]);
                  break;
            case 3:
                  strcpy (locstr, inst_name[loc]);
                  break;
            case 4:
                  strcpy (locstr, train_name[loc]);
                  break;
            }
      }

      if (is_quit)
            strcpy (isqstr, "是");
      else
            strcpy (isqstr, "否");

      gtk_list_store_append (gls, &iter);

      if (price >= 0)
      {
            gtk_list_store_set (gls, &iter, 
                  0, id,
                  1, name,
                  2, role, 
                  3, jobstr, 
                  4, locstr, 
                  5, (int)ceil(hp), 
                  6, (int)ceil(hpmax), 
                  7, (int)ceil(prop[0]), 
                  8, (int)ceil(prop[1]), 
                  9, (int)ceil(prop[2]),
                  10, (int)ceil(prop[3]),
                  11, (int)ceil(prop[4]),
                  12, isqstr,
                  13, lst, 
                  14, pst,
                  15, (int)ceil(price),
                  -1);
      }
      else
      {
            gtk_list_store_set (gls, &iter, 
                  0, id,
                  1, name,
                  2, role, 
                  3, jobstr, 
                  4, locstr, 
                  5, (int)ceil(hp), 
                  6, (int)ceil(hpmax), 
                  7, (int)ceil(prop[0]), 
                  8, (int)ceil(prop[1]), 
                  9, (int)ceil(prop[2]),
                  10, (int)ceil(prop[3]),
                  11, (int)ceil(prop[4]),
                  12, isqstr,
                  13, lst, 
                  14, pst,
                  -1);
      }
}

void tv_add_city (int id, GtkWidget * tv, char * name, int scope, float hpmax, float hp, char * role, int nmos, char * mayor, char * treasurer, float fengshui)
{
      GtkListStore * gls = GTK_LIST_STORE(gtk_tree_view_get_model (GTK_TREE_VIEW (tv)));
      GtkTreeIter iter;

      char sstr[32] = "";
      switch (scope) 
      {
      case 0:
            strcpy (sstr, "小");
            break;
      case 1:
            strcpy (sstr, "中");
            break;
      case 2:
            strcpy (sstr, "大");
            break;
      }

      gtk_list_store_append (gls, &iter);
      gtk_list_store_set (gls, &iter, 
            0, id,
            1, name,
            2, sstr, 
            3, (int)ceil(hp), 
            4, (int)ceil(hpmax), 
            5, role, 
            6, nmos, 
            7, mayor, 
            8, treasurer, 
            9, (int)ceil(fengshui),
            -1);

}

void tv_add_tech (int id, GtkWidget * tv, char * name, float study, int scond)
{
      GtkListStore * gls = GTK_LIST_STORE(gtk_tree_view_get_model (GTK_TREE_VIEW (tv)));
      GtkTreeIter iter;

      char scond_str [lstr];
      switch (scond) 
      {
      case 0:
            strcpy (scond_str, "任意");
            break;
      case 1:
            strcpy (scond_str, "回合胜利");
            break;
      case 2:
            strcpy (scond_str, "回合失败");
            break;
      case 3:
            strcpy (scond_str, "战斗胜利");
            break;
      }

      gtk_list_store_append (gls, &iter);
      gtk_list_store_set (gls, &iter, 
            0, id,
            1, name,
            2, (int)ceil(study), 
            3, scond_str, 
            -1);

}

void buf_add_msg (GtkTextBuffer * buffer, char * msg)
{
      char lmsg[lstr] = "";
      sprintf (lmsg, "%s\n", msg);
      GtkTextIter text_iter;
      gtk_text_buffer_get_end_iter (buffer, &text_iter);
      gtk_text_buffer_insert (buffer, &text_iter, lmsg, -1);
}

void onoff_widget (GtkWidget * widget, int is_on)
{
      gtk_widget_set_sensitive (widget, is_on);
}

void select_setup (GtkWidget * tv, int is_multi)
{
      GtkTreeSelection * select;
      GtkSelectionMode mode;

      if (is_multi) 
            mode = GTK_SELECTION_MULTIPLE;
      else 
            mode = GTK_SELECTION_SINGLE;

      select = gtk_tree_view_get_selection (GTK_TREE_VIEW (tv));
      gtk_tree_selection_set_mode (select, mode);
}

void tv_add_column (GtkWidget * tv, int id, char * title)
{
      GtkCellRenderer *renderer;
      GtkTreeViewColumn *column;

      renderer = gtk_cell_renderer_text_new ();
      column = gtk_tree_view_column_new_with_attributes (title, renderer, 
            "text", id + 1, NULL);
      gtk_tree_view_column_set_resizable (column, TRUE);
      gtk_tree_view_column_set_sort_column_id (column, id + 1);
      gtk_tree_view_append_column (GTK_TREE_VIEW (tv), column);
}

void load_ui()
{
      int i;

      gtk_init (NULL, NULL);

      GtkBuilder * builder = gtk_builder_new_from_file ("gtkui.glade");

      win_battle = (GtkWidget *)gtk_builder_get_object (builder, "win_battle");
      label_battle = (GtkWidget *)gtk_builder_get_object (builder, "label_battle");
      label_battle_role[0] = (GtkWidget *)gtk_builder_get_object (builder, "label_battle_role_0");
      label_battle_role[1] = (GtkWidget *)gtk_builder_get_object (builder, "label_battle_role_1");
      pb_battle_power[0] = (GtkWidget *)gtk_builder_get_object (builder, "pb_battle_power_0");
      pb_battle_power[1] = (GtkWidget *)gtk_builder_get_object (builder, "pb_battle_power_1");
      pb_battle_winp[0] = (GtkWidget *)gtk_builder_get_object (builder, "pb_battle_winp_0");
      pb_battle_winp[1] = (GtkWidget *)gtk_builder_get_object (builder, "pb_battle_winp_1");
      pb_battle_celve[0] = (GtkWidget *)gtk_builder_get_object (builder, "pb_battle_celve_0");
      pb_battle_celve[1] = (GtkWidget *)gtk_builder_get_object (builder, "pb_battle_celve_1");
      pb_battle_hp[0] = (GtkWidget *)gtk_builder_get_object (builder, "pb_battle_hp_0");
      pb_battle_hp[1] = (GtkWidget *)gtk_builder_get_object (builder, "pb_battle_hp_1");
      pb_battle_def[0] = (GtkWidget *)gtk_builder_get_object (builder, "pb_battle_def_0");
      pb_battle_def[1] = (GtkWidget *)gtk_builder_get_object (builder, "pb_battle_def_1");

      label_battle_tech[0][0] = (GtkWidget *)gtk_builder_get_object (builder, "label_battle_tech_0_0");
      label_battle_tech[0][1] = (GtkWidget *)gtk_builder_get_object (builder, "label_battle_tech_0_1");
      label_battle_tech[0][2] = (GtkWidget *)gtk_builder_get_object (builder, "label_battle_tech_0_2");
      label_battle_tech[0][3] = (GtkWidget *)gtk_builder_get_object (builder, "label_battle_tech_0_3");
      label_battle_tech[0][4] = (GtkWidget *)gtk_builder_get_object (builder, "label_battle_tech_0_4");
      label_battle_tech[0][5] = (GtkWidget *)gtk_builder_get_object (builder, "label_battle_tech_0_5");
      label_battle_tech[0][6] = (GtkWidget *)gtk_builder_get_object (builder, "label_battle_tech_0_6");
      label_battle_tech[0][7] = (GtkWidget *)gtk_builder_get_object (builder, "label_battle_tech_0_7");
      label_battle_tech[0][8] = (GtkWidget *)gtk_builder_get_object (builder, "label_battle_tech_0_8");
      label_battle_tech[0][9] = (GtkWidget *)gtk_builder_get_object (builder, "label_battle_tech_0_9");
      label_battle_tech[0][10] = (GtkWidget *)gtk_builder_get_object (builder, "label_battle_tech_0_10");
      label_battle_tech[0][11] = (GtkWidget *)gtk_builder_get_object (builder, "label_battle_tech_0_11");
      label_battle_tech[1][0] = (GtkWidget *)gtk_builder_get_object (builder, "label_battle_tech_1_0");
      label_battle_tech[1][1] = (GtkWidget *)gtk_builder_get_object (builder, "label_battle_tech_1_1");
      label_battle_tech[1][2] = (GtkWidget *)gtk_builder_get_object (builder, "label_battle_tech_1_2");
      label_battle_tech[1][3] = (GtkWidget *)gtk_builder_get_object (builder, "label_battle_tech_1_3");
      label_battle_tech[1][4] = (GtkWidget *)gtk_builder_get_object (builder, "label_battle_tech_1_4");
      label_battle_tech[1][5] = (GtkWidget *)gtk_builder_get_object (builder, "label_battle_tech_1_5");
      label_battle_tech[1][6] = (GtkWidget *)gtk_builder_get_object (builder, "label_battle_tech_1_6");
      label_battle_tech[1][7] = (GtkWidget *)gtk_builder_get_object (builder, "label_battle_tech_1_7");
      label_battle_tech[1][8] = (GtkWidget *)gtk_builder_get_object (builder, "label_battle_tech_1_8");
      label_battle_tech[1][9] = (GtkWidget *)gtk_builder_get_object (builder, "label_battle_tech_1_9");
      label_battle_tech[1][10] = (GtkWidget *)gtk_builder_get_object (builder, "label_battle_tech_1_10");
      label_battle_tech[1][11] = (GtkWidget *)gtk_builder_get_object (builder, "label_battle_tech_1_11");

      label_battle_st[0][0] = (GtkWidget *)gtk_builder_get_object (builder, "label_battle_st_0_0");
      label_battle_st[0][1] = (GtkWidget *)gtk_builder_get_object (builder, "label_battle_st_0_1");
      label_battle_st[0][2] = (GtkWidget *)gtk_builder_get_object (builder, "label_battle_st_0_2");
      label_battle_st[1][0] = (GtkWidget *)gtk_builder_get_object (builder, "label_battle_st_1_0");
      label_battle_st[1][1] = (GtkWidget *)gtk_builder_get_object (builder, "label_battle_st_1_1");
      label_battle_st[1][2] = (GtkWidget *)gtk_builder_get_object (builder, "label_battle_st_1_2");

      tv_battle_msg = (GtkWidget *)gtk_builder_get_object (builder, "tv_battle_msg");

      button_battle_next = (GtkWidget *)gtk_builder_get_object (builder, "button_battle_next");
      tb_battle_auto = (GtkWidget *)gtk_builder_get_object (builder, "tb_battle_auto");

      win_people = (GtkWidget *)gtk_builder_get_object (builder, "win_people");
      label_people = (GtkWidget *)gtk_builder_get_object (builder, "label_people");
      tv_people = (GtkWidget *)gtk_builder_get_object (builder, "tv_people");
      button_people = (GtkWidget *)gtk_builder_get_object (builder, "button_people");
      button_people_main = (GtkWidget *)gtk_builder_get_object (builder, "button_people_main");
      button_people_vice = (GtkWidget *)gtk_builder_get_object (builder, "button_people_vice");
      label_people_main = (GtkWidget *)gtk_builder_get_object (builder, "label_people_main");
      label_people_vice = (GtkWidget *)gtk_builder_get_object (builder, "label_people_vice");

      win_city = (GtkWidget *)gtk_builder_get_object (builder, "win_city");
      tv_city = (GtkWidget *)gtk_builder_get_object (builder, "tv_city");
      button_city = (GtkWidget *)gtk_builder_get_object (builder, "button_city");

      win_card = (GtkWidget *)gtk_builder_get_object (builder, "win_card");
      tv_card = (GtkWidget *)gtk_builder_get_object (builder, "tv_card");
      button_card = (GtkWidget *)gtk_builder_get_object (builder, "button_card");

      win_inst = (GtkWidget *)gtk_builder_get_object (builder, "win_inst");
      tv_inst = (GtkWidget *)gtk_builder_get_object (builder, "tv_inst");
      tv_inst_mos = (GtkWidget *)gtk_builder_get_object (builder, "tv_inst_mos");
      tv_inst_tech = (GtkWidget *)gtk_builder_get_object (builder, "tv_inst_tech");
      button_inst = (GtkWidget *)gtk_builder_get_object (builder, "button_inst");

      win_train = (GtkWidget *)gtk_builder_get_object (builder, "win_train");
      tv_train = (GtkWidget *)gtk_builder_get_object (builder, "tv_train");
      tv_train_mos = (GtkWidget *)gtk_builder_get_object (builder, "tv_train_mos");
      button_train = (GtkWidget *)gtk_builder_get_object (builder, "button_train");

      win_exchange = (GtkWidget *)gtk_builder_get_object (builder, "win_exchange");
      label_exchange = (GtkWidget *)gtk_builder_get_object (builder, "label_exchange");
      tv_exchange_in = (GtkWidget *)gtk_builder_get_object (builder, "tv_exchange_in");
      tv_exchange_out = (GtkWidget *)gtk_builder_get_object (builder, "tv_exchange_out");
      button_exchange = (GtkWidget *)gtk_builder_get_object (builder, "button_exchange");

      win_study = (GtkWidget *)gtk_builder_get_object (builder, "win_study");
      label_study = (GtkWidget *)gtk_builder_get_object (builder, "label_study");
      tv_study_tech = (GtkWidget *)gtk_builder_get_object (builder, "tv_study_tech");
      tv_study_people = (GtkWidget *)gtk_builder_get_object (builder, "tv_study_people");
      button_study = (GtkWidget *)gtk_builder_get_object (builder, "button_study");

      win_role = (GtkWidget *)gtk_builder_get_object (builder, "win_role");
      tv_role = (GtkWidget *)gtk_builder_get_object (builder, "tv_role");
      tv_role_mos = (GtkWidget *)gtk_builder_get_object (builder, "tv_role_mos");
      tv_role_mcs = (GtkWidget *)gtk_builder_get_object (builder, "tv_role_mcs");
      tv_role_sos = (GtkWidget *)gtk_builder_get_object (builder, "tv_role_sos");
      tv_role_tos = (GtkWidget *)gtk_builder_get_object (builder, "tv_role_tos");
      tv_role_cards = (GtkWidget *)gtk_builder_get_object (builder, "tv_role_cards");

      label_role_tech[0] = (GtkWidget *)gtk_builder_get_object (builder, "label_role_tech_0");
      label_role_tech[1] = (GtkWidget *)gtk_builder_get_object (builder, "label_role_tech_1");
      label_role_tech[2] = (GtkWidget *)gtk_builder_get_object (builder, "label_role_tech_2");
      label_role_tech[3] = (GtkWidget *)gtk_builder_get_object (builder, "label_role_tech_3");
      label_role_tech[4] = (GtkWidget *)gtk_builder_get_object (builder, "label_role_tech_4");
      label_role_tech[5] = (GtkWidget *)gtk_builder_get_object (builder, "label_role_tech_5");
      label_role_tech[6] = (GtkWidget *)gtk_builder_get_object (builder, "label_role_tech_6");
      label_role_tech[7] = (GtkWidget *)gtk_builder_get_object (builder, "label_role_tech_7");
      label_role_tech[8] = (GtkWidget *)gtk_builder_get_object (builder, "label_role_tech_8");
      label_role_tech[9] = (GtkWidget *)gtk_builder_get_object (builder, "label_role_tech_9");
      label_role_tech[10] = (GtkWidget *)gtk_builder_get_object (builder, "label_role_tech_10");
      label_role_tech[11] = (GtkWidget *)gtk_builder_get_object (builder, "label_role_tech_11");

      button_role = (GtkWidget *)gtk_builder_get_object (builder, "button_role");
      button_role_mcs = (GtkWidget *)gtk_builder_get_object (builder, "button_role_mcs");

      win_main = (GtkWidget *)gtk_builder_get_object (builder, "win_main");
      layout_main_map = (GtkWidget *)gtk_builder_get_object (builder, "layout_main_map");
      button_main_move = (GtkWidget *)gtk_builder_get_object (builder, "button_main_move");
      button_main_card = (GtkWidget *)gtk_builder_get_object (builder, "button_main_card");
      button_main_role = (GtkWidget *)gtk_builder_get_object (builder, "button_main_role");
      button_main_people = (GtkWidget *)gtk_builder_get_object (builder, "button_main_people");
      button_main_rp = (GtkWidget *)gtk_builder_get_object (builder, "button_main_rp");
      button_main_city = (GtkWidget *)gtk_builder_get_object (builder, "button_main_city");
      button_main_inst = (GtkWidget *)gtk_builder_get_object (builder, "button_main_inst");
      button_main_train = (GtkWidget *)gtk_builder_get_object (builder, "button_main_train");
      tv_main_msg = (GtkWidget *)gtk_builder_get_object (builder, "tv_main_msg");
      button_main_yes = (GtkWidget *)gtk_builder_get_object (builder, "button_main_yes");
      button_main_no = (GtkWidget *)gtk_builder_get_object (builder, "button_main_no");
      button_main_small = (GtkWidget *)gtk_builder_get_object (builder, "button_main_small");
      button_main_medium = (GtkWidget *)gtk_builder_get_object (builder, "button_main_medium");
      button_main_large = (GtkWidget *)gtk_builder_get_object (builder, "button_main_large");

      buf_main_msg = (GtkTextBuffer *)gtk_builder_get_object (builder, "buf_main_msg");
      buf_battle_msg = (GtkTextBuffer *)gtk_builder_get_object (builder, "buf_battle_msg");
      swin_main_msg = (GtkWidget *)gtk_builder_get_object (builder,   "swin_main_msg");
      swin_battle_msg = (GtkWidget *)gtk_builder_get_object (builder, "swin_battle_msg");

      g_signal_connect(button_main_role, "clicked", G_CALLBACK(cbf_main_role_check), NULL); 
      g_signal_connect(button_main_people, "clicked", G_CALLBACK(cbf_main_people_check), NULL); 
      g_signal_connect(button_main_move, "clicked", G_CALLBACK(cbf_main_move_click), NULL); 
      g_signal_connect(button_main_card, "clicked", G_CALLBACK(cbf_main_card_click), NULL); 
      g_signal_connect(button_main_city, "clicked", G_CALLBACK(cbf_main_city_check), NULL); 
      g_signal_connect(button_main_inst, "clicked", G_CALLBACK(cbf_main_inst_check), NULL); 
      g_signal_connect(button_main_train, "clicked", G_CALLBACK(cbf_main_train_check), NULL); 
      g_signal_connect(button_main_rp, "clicked", G_CALLBACK(cbf_main_rp_check), NULL); 
      g_signal_connect(button_main_yes, "clicked", G_CALLBACK(cbf_main_yes_click), NULL); 
      g_signal_connect(button_main_no, "clicked", G_CALLBACK(cbf_main_no_click), NULL); 
      g_signal_connect(button_main_small, "clicked", G_CALLBACK(cbf_main_small_click), NULL); 
      g_signal_connect(button_main_medium, "clicked", G_CALLBACK(cbf_main_medium_click), NULL); 
      g_signal_connect(button_main_large, "clicked", G_CALLBACK(cbf_main_large_click), NULL); 

      g_signal_connect (button_battle_next, "clicked", G_CALLBACK(cbf_battle_next_click), NULL); 
      g_signal_connect (button_people, "clicked", G_CALLBACK (cbf_people_click), NULL);
      g_signal_connect (button_people_main, "clicked", G_CALLBACK (cbf_people_main_click), NULL);
      g_signal_connect (button_people_vice, "clicked", G_CALLBACK (cbf_people_vice_click), NULL);
      g_signal_connect (button_city, "clicked", G_CALLBACK(cbf_city_click), NULL); 
      g_signal_connect (button_card, "clicked", G_CALLBACK (cbf_card_click), NULL);
      g_signal_connect (button_inst, "clicked", G_CALLBACK (cbf_inst_click), NULL);
      g_signal_connect (button_train, "clicked", G_CALLBACK (cbf_train_click), NULL);
      g_signal_connect (button_exchange, "clicked", G_CALLBACK(cbf_exchange_click), NULL); 
      g_signal_connect (button_study, "clicked", G_CALLBACK (cbf_study_click), NULL);
      g_signal_connect (button_role, "clicked", G_CALLBACK (cbf_role_click), NULL);
      g_signal_connect (button_role_mcs, "clicked", G_CALLBACK (cbf_role_mcs_click), NULL);

      g_signal_connect (win_main, "delete-event", G_CALLBACK (cbf_main_close), NULL);
      g_signal_connect (win_card, "delete-event", G_CALLBACK (cbf_card_close), NULL);
      g_signal_connect (win_battle, "delete-event", G_CALLBACK (cbf_battle_close), NULL);
      g_signal_connect (win_role, "delete-event", G_CALLBACK (cbf_role_close), NULL);
      g_signal_connect (win_people, "delete-event", G_CALLBACK (cbf_people_close), NULL);
      g_signal_connect (win_city, "delete-event", G_CALLBACK (cbf_city_close), NULL);
      g_signal_connect (win_inst, "delete-event", G_CALLBACK (cbf_inst_close), NULL);
      g_signal_connect (win_train, "delete-event", G_CALLBACK (cbf_train_close), NULL);
      g_signal_connect (win_study, "delete-event", G_CALLBACK (cbf_study_close), NULL);
      g_signal_connect (win_exchange, "delete-event", G_CALLBACK (cbf_exchange_close), NULL);

      // set model and parent
      
      gtk_window_set_transient_for (GTK_WINDOW(win_card), GTK_WINDOW(win_main));
      gtk_window_set_transient_for (GTK_WINDOW(win_battle), GTK_WINDOW(win_main));
      gtk_window_set_transient_for (GTK_WINDOW(win_role), GTK_WINDOW(win_main));
      gtk_window_set_transient_for (GTK_WINDOW(win_people), GTK_WINDOW(win_main));
      gtk_window_set_transient_for (GTK_WINDOW(win_city), GTK_WINDOW(win_main));
      gtk_window_set_transient_for (GTK_WINDOW(win_inst), GTK_WINDOW(win_main));
      gtk_window_set_transient_for (GTK_WINDOW(win_train), GTK_WINDOW(win_main));
      gtk_window_set_transient_for (GTK_WINDOW(win_study), GTK_WINDOW(win_main));
      gtk_window_set_transient_for (GTK_WINDOW(win_exchange), GTK_WINDOW(win_main));

      gtk_window_set_modal (GTK_WINDOW(win_card), TRUE);
      gtk_window_set_modal (GTK_WINDOW(win_battle), TRUE);
      gtk_window_set_modal (GTK_WINDOW(win_role), TRUE);
      gtk_window_set_modal (GTK_WINDOW(win_people), TRUE);
      gtk_window_set_modal (GTK_WINDOW(win_city), TRUE);
      gtk_window_set_modal (GTK_WINDOW(win_inst), TRUE);
      gtk_window_set_modal (GTK_WINDOW(win_train), TRUE);
      gtk_window_set_modal (GTK_WINDOW(win_study), TRUE);
      gtk_window_set_modal (GTK_WINDOW(win_exchange), TRUE);

      g_timeout_add (500, cbf_battle_timer, NULL);

      onoff_widget (button_main_yes, 0);
      onoff_widget (button_main_no, 0);
      onoff_widget (button_main_small, 0);
      onoff_widget (button_main_medium, 0);
      onoff_widget (button_main_large, 0);
      onoff_widget (button_main_move, 0);
      onoff_widget (button_main_card, 0);

      // treeview initialization
      
      GtkListStore * gls; 

      gls = gtk_list_store_new (5, G_TYPE_INT, G_TYPE_STRING, G_TYPE_INT, G_TYPE_STRING, G_TYPE_STRING);
      gtk_tree_view_set_model (GTK_TREE_VIEW(tv_card), GTK_TREE_MODEL(gls));

      tv_add_column (tv_card, 0, "名称");
      tv_add_column (tv_card, 1, "持有量");
      tv_add_column (tv_card, 2, "对象阵营");
      tv_add_column (tv_card, 3, "对象类型");

      gls = gtk_list_store_new (10, G_TYPE_INT, G_TYPE_STRING, G_TYPE_INT, G_TYPE_INT, G_TYPE_STRING,
            G_TYPE_STRING, G_TYPE_INT, G_TYPE_STRING, G_TYPE_STRING, G_TYPE_INT);
      gtk_tree_view_set_model (GTK_TREE_VIEW(tv_role), GTK_TREE_MODEL(gls));

      tv_add_column (tv_role, 0, "名称");
      tv_add_column (tv_role, 1, "位置");
      tv_add_column (tv_role, 2, "金币");
      tv_add_column (tv_role, 3, "移动方向");
      tv_add_column (tv_role, 4, "行动状态");
      tv_add_column (tv_role, 5, "行动状态剩余");
      tv_add_column (tv_role, 6, "连续行动");
      tv_add_column (tv_role, 7, "联盟");
      tv_add_column (tv_role, 8, "联盟剩余");

      gls = gtk_list_store_new (15, G_TYPE_INT, G_TYPE_STRING, G_TYPE_STRING, G_TYPE_STRING, G_TYPE_STRING, 
            G_TYPE_INT, G_TYPE_INT, G_TYPE_INT, G_TYPE_INT, G_TYPE_INT, G_TYPE_INT, G_TYPE_INT, 
            G_TYPE_STRING, G_TYPE_INT, G_TYPE_INT);
      gtk_tree_view_set_model (GTK_TREE_VIEW(tv_role_mos), GTK_TREE_MODEL(gls));

      tv_add_column (tv_role_mos, 0, "名称");
      tv_add_column (tv_role_mos, 1, "阵营");
      tv_add_column (tv_role_mos, 2, "职业");
      tv_add_column (tv_role_mos, 3, "位置");
      tv_add_column (tv_role_mos, 4, "兵力");
      tv_add_column (tv_role_mos, 5, "最大兵力");
      for (i = 0; i < 5; ++ i)
            tv_add_column (tv_role_mos, 6 + i, prop_name[i]);
      tv_add_column (tv_role_mos, 11, "准备离职");
      tv_add_column (tv_role_mos, 12, "流言状态剩余");
      tv_add_column (tv_role_mos, 13, "中毒状态剩余");

      gls = gtk_list_store_new (10, G_TYPE_INT, G_TYPE_STRING, G_TYPE_STRING, 
            G_TYPE_INT, G_TYPE_INT, G_TYPE_STRING, G_TYPE_INT,
            G_TYPE_STRING, G_TYPE_STRING, G_TYPE_INT);
      gtk_tree_view_set_model (GTK_TREE_VIEW(tv_role_mcs), GTK_TREE_MODEL(gls));

      tv_add_column (tv_role_mcs, 0, "名称");
      tv_add_column (tv_role_mcs, 1, "规模");
      tv_add_column (tv_role_mcs, 2, "防御");
      tv_add_column (tv_role_mcs, 3, "最大防御");
      tv_add_column (tv_role_mcs, 4, "所属阵营");
      tv_add_column (tv_role_mcs, 5, "驻扎官员数");
      tv_add_column (tv_role_mcs, 6, "市长");
      tv_add_column (tv_role_mcs, 7, "财务官");
      tv_add_column (tv_role_mcs, 8, "风水");

      gls = gtk_list_store_new (15, G_TYPE_INT, G_TYPE_STRING, G_TYPE_STRING, G_TYPE_STRING, G_TYPE_STRING, 
            G_TYPE_INT, G_TYPE_INT, G_TYPE_INT, G_TYPE_INT, G_TYPE_INT, G_TYPE_INT, G_TYPE_INT, 
            G_TYPE_STRING, G_TYPE_INT, G_TYPE_INT);
      gtk_tree_view_set_model (GTK_TREE_VIEW(tv_role_sos), GTK_TREE_MODEL(gls));

      tv_add_column (tv_role_sos, 0, "名称");
      tv_add_column (tv_role_sos, 1, "阵营");
      tv_add_column (tv_role_sos, 2, "职业");
      tv_add_column (tv_role_sos, 3, "位置");
      tv_add_column (tv_role_sos, 4, "兵力");
      tv_add_column (tv_role_sos, 5, "最大兵力");
      for (i = 0; i < 5; ++ i)
            tv_add_column (tv_role_sos, 6 + i, prop_name[i]);
      tv_add_column (tv_role_sos, 11, "准备离职");
      tv_add_column (tv_role_sos, 12, "流言状态剩余");
      tv_add_column (tv_role_sos, 13, "中毒状态剩余");

      gls = gtk_list_store_new (15, G_TYPE_INT, G_TYPE_STRING, G_TYPE_STRING, G_TYPE_STRING, G_TYPE_STRING, 
            G_TYPE_INT, G_TYPE_INT, G_TYPE_INT, G_TYPE_INT, G_TYPE_INT, G_TYPE_INT, G_TYPE_INT, 
            G_TYPE_STRING, G_TYPE_INT, G_TYPE_INT);
      gtk_tree_view_set_model (GTK_TREE_VIEW(tv_role_tos), GTK_TREE_MODEL(gls));

      tv_add_column (tv_role_tos, 0, "名称");
      tv_add_column (tv_role_tos, 1, "阵营");
      tv_add_column (tv_role_tos, 2, "职业");
      tv_add_column (tv_role_tos, 3, "位置");
      tv_add_column (tv_role_tos, 4, "兵力");
      tv_add_column (tv_role_tos, 5, "最大兵力");
      for (i = 0; i < 5; ++ i)
            tv_add_column (tv_role_tos, 6 + i, prop_name[i]);
      tv_add_column (tv_role_tos, 11, "准备离职");
      tv_add_column (tv_role_tos, 12, "流言状态剩余");
      tv_add_column (tv_role_tos, 13, "中毒状态剩余");

      gls = gtk_list_store_new (3, G_TYPE_INT, G_TYPE_STRING, G_TYPE_INT);
      gtk_tree_view_set_model (GTK_TREE_VIEW(tv_role_cards), GTK_TREE_MODEL(gls));

      tv_add_column (tv_role_cards, 0, "名称");
      tv_add_column (tv_role_cards, 1, "数量");

      gls = gtk_list_store_new (10, G_TYPE_INT, G_TYPE_STRING, G_TYPE_STRING, 
            G_TYPE_INT, G_TYPE_INT, G_TYPE_STRING, G_TYPE_INT,
            G_TYPE_STRING, G_TYPE_STRING, G_TYPE_INT);
      gtk_tree_view_set_model (GTK_TREE_VIEW(tv_city), GTK_TREE_MODEL(gls));

      tv_add_column (tv_city, 0, "名称");
      tv_add_column (tv_city, 1, "规模");
      tv_add_column (tv_city, 2, "防御");
      tv_add_column (tv_city, 3, "最大防御");
      tv_add_column (tv_city, 4, "所属阵营");
      tv_add_column (tv_city, 5, "驻扎官员数");
      tv_add_column (tv_city, 6, "市长");
      tv_add_column (tv_city, 7, "财务官");
      tv_add_column (tv_city, 8, "风水");

      gls = gtk_list_store_new (7, G_TYPE_INT, G_TYPE_STRING, G_TYPE_INT, 
            G_TYPE_STRING, G_TYPE_STRING, G_TYPE_INT, G_TYPE_INT);
      gtk_tree_view_set_model (GTK_TREE_VIEW(tv_inst), GTK_TREE_MODEL(gls));

      tv_add_column (tv_inst, 0, "名称");
      tv_add_column (tv_inst, 1, "可研究策略数");
      tv_add_column (tv_inst, 2, "在研人员");
      tv_add_column (tv_inst, 3, "在研项目");
      tv_add_column (tv_inst, 4, "剩余成功点数");
      tv_add_column (tv_inst, 5, "预计剩余回合");

      gls = gtk_list_store_new (15, G_TYPE_INT, G_TYPE_STRING, G_TYPE_STRING, G_TYPE_STRING, G_TYPE_STRING, 
            G_TYPE_INT, G_TYPE_INT, G_TYPE_INT, G_TYPE_INT, G_TYPE_INT, G_TYPE_INT, G_TYPE_INT, 
            G_TYPE_STRING, G_TYPE_INT, G_TYPE_INT);
      gtk_tree_view_set_model (GTK_TREE_VIEW(tv_inst_mos), GTK_TREE_MODEL(gls));

      tv_add_column (tv_inst_mos, 0, "名称");
      tv_add_column (tv_inst_mos, 1, "阵营");
      tv_add_column (tv_inst_mos, 2, "职业");
      tv_add_column (tv_inst_mos, 3, "位置");
      tv_add_column (tv_inst_mos, 4, "兵力");
      tv_add_column (tv_inst_mos, 5, "最大兵力");
      for (i = 0; i < 5; ++ i)
            tv_add_column (tv_inst_mos, 6 + i, prop_name[i]);
      tv_add_column (tv_inst_mos, 11, "准备离职");
      tv_add_column (tv_inst_mos, 12, "流言状态剩余");
      tv_add_column (tv_inst_mos, 13, "中毒状态剩余");

      gls = gtk_list_store_new (4, G_TYPE_INT, G_TYPE_STRING, G_TYPE_INT, G_TYPE_STRING);
      gtk_tree_view_set_model (GTK_TREE_VIEW(tv_inst_tech), GTK_TREE_MODEL(gls));

      tv_add_column (tv_inst_tech, 0, "名称");
      tv_add_column (tv_inst_tech, 1, "研究点数");
      tv_add_column (tv_inst_tech, 2, "发动时机");

      gls = gtk_list_store_new (5, G_TYPE_INT, G_TYPE_STRING, G_TYPE_STRING, G_TYPE_STRING, G_TYPE_INT);
      gtk_tree_view_set_model (GTK_TREE_VIEW(tv_train), GTK_TREE_MODEL(gls));

      tv_add_column (tv_train, 0, "名称");
      tv_add_column (tv_train, 1, "训练项目");
      tv_add_column (tv_train, 2, "训练人员");
      tv_add_column (tv_train, 3, "剩余回合");

      gls = gtk_list_store_new (15, G_TYPE_INT, G_TYPE_STRING, G_TYPE_STRING, G_TYPE_STRING, G_TYPE_STRING, 
            G_TYPE_INT, G_TYPE_INT, G_TYPE_INT, G_TYPE_INT, G_TYPE_INT, G_TYPE_INT, G_TYPE_INT, 
            G_TYPE_STRING, G_TYPE_INT, G_TYPE_INT);
      gtk_tree_view_set_model (GTK_TREE_VIEW(tv_train_mos), GTK_TREE_MODEL(gls));

      tv_add_column (tv_train_mos, 0, "名称");
      tv_add_column (tv_train_mos, 1, "阵营");
      tv_add_column (tv_train_mos, 2, "职业");
      tv_add_column (tv_train_mos, 3, "位置");
      tv_add_column (tv_train_mos, 4, "兵力");
      tv_add_column (tv_train_mos, 5, "最大兵力");
      for (i = 0; i < 5; ++ i)
            tv_add_column (tv_train_mos, 6 + i, prop_name[i]);
      tv_add_column (tv_train_mos, 11, "准备离职");
      tv_add_column (tv_train_mos, 12, "流言状态剩余");
      tv_add_column (tv_train_mos, 13, "中毒状态剩余");

      gls = gtk_list_store_new (16, G_TYPE_INT, G_TYPE_STRING, G_TYPE_STRING, G_TYPE_STRING, G_TYPE_STRING, 
            G_TYPE_INT, G_TYPE_INT, G_TYPE_INT, G_TYPE_INT, G_TYPE_INT, G_TYPE_INT, G_TYPE_INT, 
            G_TYPE_STRING, G_TYPE_INT, G_TYPE_INT, G_TYPE_INT);
      gtk_tree_view_set_model (GTK_TREE_VIEW(tv_people), GTK_TREE_MODEL(gls));

      tv_add_column (tv_people, 0, "名称");
      tv_add_column (tv_people, 1, "阵营");
      tv_add_column (tv_people, 2, "职业");
      tv_add_column (tv_people, 3, "位置");
      tv_add_column (tv_people, 4, "兵力");
      tv_add_column (tv_people, 5, "最大兵力");
      for (i = 0; i < 5; ++ i)
            tv_add_column (tv_people, 6 + i, prop_name[i]);
      tv_add_column (tv_people, 11, "准备离职");
      tv_add_column (tv_people, 12, "流言状态剩余");
      tv_add_column (tv_people, 13, "中毒状态剩余");
      tv_add_column (tv_people, 14, "聘用费");

      gls = gtk_list_store_new (15, G_TYPE_INT, G_TYPE_STRING, G_TYPE_STRING, G_TYPE_STRING, G_TYPE_STRING, 
            G_TYPE_INT, G_TYPE_INT, G_TYPE_INT, G_TYPE_INT, G_TYPE_INT, G_TYPE_INT, G_TYPE_INT, 
            G_TYPE_STRING, G_TYPE_INT, G_TYPE_INT);
      gtk_tree_view_set_model (GTK_TREE_VIEW(tv_exchange_out), GTK_TREE_MODEL(gls));

      tv_add_column (tv_exchange_out, 0, "名称");
      tv_add_column (tv_exchange_out, 1, "阵营");
      tv_add_column (tv_exchange_out, 2, "职业");
      tv_add_column (tv_exchange_out, 3, "位置");
      tv_add_column (tv_exchange_out, 4, "兵力");
      tv_add_column (tv_exchange_out, 5, "最大兵力");
      for (i = 0; i < 5; ++ i)
            tv_add_column (tv_exchange_out, 6 + i, prop_name[i]);
      tv_add_column (tv_exchange_out, 11, "准备离职");
      tv_add_column (tv_exchange_out, 12, "流言状态剩余");
      tv_add_column (tv_exchange_out, 13, "中毒状态剩余");

      gls = gtk_list_store_new (15, G_TYPE_INT, G_TYPE_STRING, G_TYPE_STRING, G_TYPE_STRING, G_TYPE_STRING, 
            G_TYPE_INT, G_TYPE_INT, G_TYPE_INT, G_TYPE_INT, G_TYPE_INT, G_TYPE_INT, G_TYPE_INT, 
            G_TYPE_STRING, G_TYPE_INT, G_TYPE_INT);
      gtk_tree_view_set_model (GTK_TREE_VIEW(tv_exchange_in), GTK_TREE_MODEL(gls));

      tv_add_column (tv_exchange_in, 0, "名称");
      tv_add_column (tv_exchange_in, 1, "阵营");
      tv_add_column (tv_exchange_in, 2, "职业");
      tv_add_column (tv_exchange_in, 3, "位置");
      tv_add_column (tv_exchange_in, 4, "兵力");
      tv_add_column (tv_exchange_in, 5, "最大兵力");
      for (i = 0; i < 5; ++ i)
            tv_add_column (tv_exchange_in, 6 + i, prop_name[i]);
      tv_add_column (tv_exchange_in, 11, "准备离职");
      tv_add_column (tv_exchange_in, 12, "流言状态剩余");
      tv_add_column (tv_exchange_in, 13, "中毒状态剩余");

      gls = gtk_list_store_new (4, G_TYPE_INT, G_TYPE_STRING, G_TYPE_INT, G_TYPE_STRING);
      gtk_tree_view_set_model (GTK_TREE_VIEW(tv_study_tech), GTK_TREE_MODEL(gls));

      tv_add_column (tv_study_tech, 0, "名称");
      tv_add_column (tv_study_tech, 1, "研究点数");
      tv_add_column (tv_study_tech, 2, "发动时机");

      gls = gtk_list_store_new (15, G_TYPE_INT, G_TYPE_STRING, G_TYPE_STRING, G_TYPE_STRING, G_TYPE_STRING, 
            G_TYPE_INT, G_TYPE_INT, G_TYPE_INT, G_TYPE_INT, G_TYPE_INT, G_TYPE_INT, G_TYPE_INT, 
            G_TYPE_STRING, G_TYPE_INT, G_TYPE_INT);
      gtk_tree_view_set_model (GTK_TREE_VIEW(tv_study_people), GTK_TREE_MODEL(gls));

      tv_add_column (tv_study_people, 0, "名称");
      tv_add_column (tv_study_people, 1, "阵营");
      tv_add_column (tv_study_people, 2, "职业");
      tv_add_column (tv_study_people, 3, "位置");
      tv_add_column (tv_study_people, 4, "兵力");
      tv_add_column (tv_study_people, 5, "最大兵力");
      for (i = 0; i < 5; ++ i)
            tv_add_column (tv_study_people, 6 + i, prop_name[i]);
      tv_add_column (tv_study_people, 11, "准备离职");
      tv_add_column (tv_study_people, 12, "流言状态剩余");
      tv_add_column (tv_study_people, 13, "中毒状态剩余");

      // init variables
      
      st = GAME_OVER;

      obj_type = -1;
      obj = -1;
      action = -1;
      card_action = -1;

      check_type = -1;
      check_role_select = -1;
      check_role_mcs_select = -1;
      check_city_mcs_select = -1;
      check_inst_select = -1;
      check_train_select = -1;

      people_select_st = 0;
      people_select = -1;
      people_select_vice = -1;
      people_select_num = 0;

      battle_scale = -1;
      wait_battle_tick = 0;
      is_battle_finish = 0;

      exchange_out_num = 0;
      exchange_in_num = 0;

      study_tech = -1;
      study_people = -1;

      msg_res = -1;

      wp = 85;
      hp = 85;
      is_align = 0;
}

void labels_clear (GtkWidget ** labels, int n)
{
      int i;
      for (i = 0; i < n; ++ i)
      {
            gtk_widget_set_sensitive (labels[i], 0);
      }
}

int msg_box (char * msg, int mode)
{
      if (st == GAME_OVER) return -1;
      if (mode == 1) {
            onoff_widget (button_main_yes, 1);
            onoff_widget (button_main_no, 1);
      } else {
            onoff_widget (button_main_yes, 1);
            onoff_widget (button_main_no, 0);
      }
      if (strlen (msg) > 0)
      {
            buf_add_msg (buf_main_msg, msg);
            swin_set_follow (swin_main_msg);
      }

      msg_res = -1;
      if (mode == 1)
            st = WAIT_MSG_ANSWER;
      else 
            st = WAIT_MSG_CONFIRM;

      gtk_main();
      if (st == GAME_OVER) return -1;

      onoff_widget (button_main_yes, 0);
      onoff_widget (button_main_no, 0);

      return msg_res;

}

void msg_map (char * msg)
{
      if (st == GAME_OVER) return;
      char msg_line[lstr] = "";
      buf_add_msg (buf_main_msg, msg);
      swin_set_follow (swin_main_msg);
}

GtkWidget * new_button (int x, int y) 
{
      int rx = (x + 1) * wp, ry = (y + 1) * hp;

      GtkWidget * res = gtk_button_new();
      gtk_widget_set_size_request (res, wp, hp);

      gtk_layout_put ((GtkLayout*)layout_main_map, res, rx, ry);

      return res;
}

void init_map  (int ngrid_, int nrole_, int npeople_, int ncity_, int ninst_, int ntrain_, int ntech_, int ncard_, int nx_, int ny_)
{
      int i, j;

      ngrid = ngrid_;
      if (ngrid % 4 == 0) is_align = 1;
      else is_align = 0;
      nrole = nrole_;
      npeople = npeople_;
      ncity = ncity_;
      ninst = ninst_;
      ntrain = ntrain_;
      ntech = ntech_;
      ncard = ncard_;

      nx = ny = ngrid / 4 + 1 + nrole * 2;

      gtk_layout_set_size ((GtkLayout*)layout_main_map, (nx + 2) * wp, (ny + 2) * hp);

      map_point = malloc (ngrid * sizeof(GtkWidget *));
      role_point = malloc (ngrid * sizeof(GtkWidget **));
      is_role = malloc (ngrid * sizeof(int *));
      loc_ind = malloc (ngrid * sizeof(int));
      rloc_ind = malloc (ngrid * sizeof(RoleLocInd *));
      for (i = 0; i < ngrid; ++ i)
      {
            role_point[i] = malloc (nrole * sizeof(GtkWidget *));
            is_role[i] = malloc (nrole * sizeof(int));
            loc_ind[i] = i;
            rloc_ind[i] = malloc (nrole * sizeof(RoleLocInd));
            for (j = 0; j < nrole; ++ j)
            {
                  rloc_ind[i][j].loc = i;
                  rloc_ind[i][j].role = j;
            }
      }

      map_type = malloc (ngrid * sizeof(int));
      map_obj  = malloc (ngrid * sizeof(int));
      is_barrier = malloc (ngrid * sizeof(int));
      is_robber = malloc (ngrid * sizeof(int));

      role_name = malloc (nrole * sizeof(char *));
      for (i = 0; i < nrole; ++ i)
      {
            role_name[i] = malloc (lstr * sizeof(char));
      }

      tech_name = malloc (ntech * sizeof(char *));
      for (i = 0; i < ntech; ++ i)
      {
            tech_name[i] = malloc (lstr * sizeof(char));
      }
      tech_scond = malloc (ntech * sizeof(int));

      card_name = malloc (ncard * sizeof(char *));
      for (i = 0; i < ncard; ++ i)
      {
            card_name[i] = malloc (lstr * sizeof(char));
      }
      card_odir = malloc (ncard * sizeof(int));
      card_otype = malloc (ncard * sizeof(int));

      city_name = malloc (ncity * sizeof(char *));
      for (i = 0; i < ncity; ++ i)
      {
            city_name[i] = malloc (lstr * sizeof(char));
      }

      inst_name = malloc (ninst * sizeof(char *));
      for (i = 0; i < ninst; ++ i)
      {
            inst_name[i] = malloc (lstr * sizeof(char));
      }

      train_name = malloc (ntrain * sizeof(char *));
      for (i = 0; i < ntrain; ++ i)
      {
            train_name[i] = malloc (lstr * sizeof(char));
      }

}

void free_ui ()
{
      int i;

      free (map_point);
      for (i = 0; i < ngrid; ++ i)
      {
            free (role_point[i]);
            free (is_role[i]);
            free (rloc_ind[i]);
      }
      free (role_point);
      free (is_role);
      free (loc_ind);
      free (rloc_ind);

      free (map_type);
      free (map_obj);
      free (is_barrier);
      free (is_robber);

      for (i = 0; i < nrole; ++ i)
      {
            free (role_name[i]);
      }
      free (role_name);

      for (i = 0; i < ntech; ++ i)
      {
            free (tech_name[i]);
      }
      free (tech_name);
      free (tech_scond);

      for (i = 0; i < ncard; ++ i)
      {
            free (card_name[i]);
      }
      free (card_name);
      free (card_odir);
      free (card_otype);

      for (i = 0; i < ncity; ++ i)
      {
            free (city_name[i]);
      }
      free (city_name);

      for (i = 0; i < ninst; ++ i)
      {
            free (inst_name[i]);
      }
      free (inst_name);

      for (i = 0; i < ntrain; ++ i)
      {
            free (train_name[i]);
      }
      free (train_name);
}

void init_tech (int tech, char *tname, int scond)
{
      strcpy (tech_name[tech], tname);
      tech_scond[tech] = scond;
}

void init_card (int card, char * dname, int odir, int otype)
{
      strcpy (card_name[card], dname);
      card_odir[card] = odir;
      card_otype[card] = otype;
}

void init_role (int role, char * rname)
{
      strcpy (role_name[role], rname);
}

void set_board (int loc, int type, int ind, int x, int y)
{
      int i;

      if (!is_align) return;
      char name [lstr];
      switch (type)
      {
      case 0:
            strcpy (name, "");
            break;
      case 1:
            strcpy (name, city_name[ind]);
            break;
      case 2:
            strcpy (name, "会所");
            break;
      case 3:
            strcpy (name, inst_name[ind]);
            break;
      case 4:
            strcpy (name, train_name[ind]);
            break;
      case 5:
            strcpy (name, "机会");
            break;
      }

      int sec = loc / (ngrid / 4); 
      int resi = loc % (ngrid / 4);
      switch (sec) 
      {
      case 0:
            x = resi;
            y = 0;
            break;
      case 1: 
            x = ngrid / 4;
            y = resi;
            break;
      case 2:
            x = ngrid / 4 - resi;
            y = ngrid / 4;
            break;
      case 3:
            x = 0;
            y = ngrid / 4 - resi;
            break;
      }

      for (i = 0; i < nrole; ++ i)
            is_role[loc][i] = 0;
      is_barrier[loc] = 0;
      is_robber[loc] = 0;

      map_point[loc] = new_button (x + nrole, y + nrole);
      map_type[loc] = type;
      map_obj[loc] = ind;

      gtk_button_set_label (GTK_BUTTON(map_point[loc]), name);

      g_signal_connect (map_point[loc], "clicked", G_CALLBACK(cbf_main_board_click), &loc_ind[loc]);

      int xx, yy;
      for (i = 0; i < nrole; ++ i)
      {
            switch (sec) 
            {
            case 0:
                  xx = x + nrole;
                  yy = i;
                  break;
            case 1:
                  xx = ngrid / 4 + nrole + 1 + i;
                  yy = y + nrole;
                  break;
            case 2:
                  xx = x + nrole;
                  yy = ngrid / 4 + nrole + 1 + i;
                  break;
            case 3:
                  xx = i;
                  yy = y + nrole;
                  break;
            }
            role_point[loc][i] = new_button (xx, yy);
            g_signal_connect (role_point[loc][i], "clicked", G_CALLBACK(cbf_main_role_click), &rloc_ind[loc][i]);
      }
}

void set_city    (int id, char * name)
{
      strcpy (city_name[id], name);
}

void set_inst    (int id, char * name)
{
      strcpy (inst_name[id], name);
}

void set_train    (int id, char * name)
{
      strcpy (train_name[id], name);
}

void set_barrier (int loc, int is_unset)
{
      if (is_unset)
            is_barrier[loc] = 0;
      else
            is_barrier[loc] = 1;
}

void set_robber  (int loc, int is_unset)
{
      if (is_unset)
            is_robber[loc] = 0;
      else
            is_robber[loc] = 1;
}

void set_role    (int loc, int rind, int is_unset)
{
      if (is_unset)
            is_role[loc][rind] = 0;
      else
            is_role[loc][rind] = 1;
}

void update_role_point (int loc)
{
      int i;
      for (i = 0; i < nrole; ++ i)
      {
            char label[lstr] = "";
            if (is_role[loc][i]) 
                  strcat (label, role_name[i]);
            if (is_barrier[loc])
                  strcat (label, "\n障碍");
            if (is_robber[loc])
                  strcat (label, "\n强盗");
            gtk_button_set_label ((GtkButton *)role_point[loc][i], label);
      }
}

void finish_draw_map ()
{
      if (!is_align) return;

      buf_clear (buf_main_msg);

      int i, j;

      for (i = 0; i < ngrid; ++ i)
            update_role_point (i);

      // set tech names
      for (i = 0; i < 2; ++ i)
      {
            for (j = 0; j < 12; ++ j)
            {
                  gtk_label_set_label (GTK_LABEL(label_battle_tech[i][j]), tech_name[j]);
            }
      }

      for (i = 0; i < 12; ++ i)
      {
            gtk_label_set_label (GTK_LABEL(label_role_tech[i]), tech_name[i]);
      }

      st = WAIT_MSG_CONFIRM;
      gtk_widget_show_all (win_main);
}

int get_action (int rind)
{
      if (!is_align || st == GAME_OVER) return -1;

      st = WAIT_ACTION;
      onoff_widget (button_main_move, 1);
      onoff_widget (button_main_card, 1);

      action = -1;
      while (action == -1)
      {
            gtk_main ();
            if (st == GAME_OVER) return -1;
      }
      onoff_widget (button_main_move, 0);
      onoff_widget (button_main_card, 0);
      return action;
}

int get_check_type () 
{
      return check_type;
}

static int cards_list_num;
void begin_cards_list () 
{
      tv_clear (tv_card);
      cards_list_num = 0;
}

void add_cards_list (int cind, int num)
{
      GtkListStore * gls = GTK_LIST_STORE(gtk_tree_view_get_model (GTK_TREE_VIEW (tv_card)));
      GtkTreeIter iter;

      char odir_str[64] = "";
      char otype_str[64] = "";

      switch (card_odir[cind]) 
      {
      case 0:
            strcpy (odir_str, "任意");
            break;
      case 1:
            strcpy (odir_str, "己方");
            break;
      case 2:
            strcpy (odir_str, "他方");
            break;
      }

      switch (card_otype[cind])
      {
      case 0:
            strcpy (otype_str, "自己");
            break;
      case 1:
            strcpy (otype_str, "城市");
            break;
      case 2:
            strcpy (otype_str, "角色");
            break;
      case 3:
            strcpy (otype_str, "地图");
            break;
      case 4:
            strcpy (otype_str, "武将");
            break;
      }

      gtk_list_store_append (gls, &iter);
      gtk_list_store_set (gls, &iter, 
            0, cards_list_num, 
            1, card_name[cind], 
            2, num, 
            3, odir_str, 
            4, otype_str, 
            -1);

      ++ cards_list_num;
}

void end_cards_list ()
{
      gtk_window_set_transient_for (GTK_WINDOW(win_card), GTK_WINDOW(win_main));
      select_setup (tv_card, 0);
}

int get_card_action () 
{
      if (st == GAME_OVER) return -1;
      card_action = -1;
      gtk_widget_show_all (win_card);
      gtk_main();
      if (st == GAME_OVER) return -1;
      return card_action;
}

static int role_list_num;
void begin_role_list ()
{
      tv_clear (tv_role);
      tv_clear (tv_role_mos);
      tv_clear (tv_role_mcs);
      tv_clear (tv_role_sos);
      tv_clear (tv_role_tos);
      tv_clear (tv_role_cards);
      labels_clear (label_role_tech, 12);

      role_list_num = 0;
}

void add_role_list (char * name, int loc, float money, int dir, int mst, int mtime, int cst, int ast, int atime)
{
      GtkListStore * gls = GTK_LIST_STORE(gtk_tree_view_get_model (GTK_TREE_VIEW (tv_role)));
      GtkTreeIter iter;

      char dirstr[32], mststr[32], cststr[32], aststr[32];

      if (dir > 0) strcpy (dirstr, "正向");
      else strcpy (dirstr, "反向");
      switch (mst) 
      {
      case 0:
            strcpy (mststr, "正常");
            break;
      case 1:
            strcpy (mststr, "慢行");
            break;
      case 2:
            strcpy (mststr, "急行");
            break;
      case 3:
            strcpy (mststr, "禁行");
            break;
      }
      if (cst > 0) strcpy (cststr, "是");
      else strcpy (cststr, "否");
      if (ast == -1) 
            strcpy (aststr, "无");
      else if (ast == -2)
            strcpy (aststr, "公敌");
      else 
            strcpy (aststr, role_name[ast]);

      gtk_list_store_append (gls, &iter);
      gtk_list_store_set (gls, &iter, 
            0, role_list_num,
            1, name,
            2, loc, 
            3, (int)ceil(money), 
            4, dirstr, 
            5, mststr, 
            6, mtime, 
            7, cststr, 
            8, aststr, 
            9, atime,
            -1);

      ++ role_list_num;
}

void end_role_list ()
{
      gtk_window_set_transient_for (GTK_WINDOW(win_role), GTK_WINDOW(win_main));
      select_setup (tv_role, 0);
}

int get_role_select ()
{
      if (st == GAME_OVER) return -1;
      check_role_select = -1;
      check_role_mcs_select = -1;
      gtk_widget_show_all (win_role);
      gtk_main();
      if (st == GAME_OVER) return -1;
      if (check_role_select < 0) return -1;
      else return check_role_select;
}

void begin_role_tech_list ()
{
      labels_clear (label_role_tech, 12);
}

void add_role_tech_list (int hind)
{
      gtk_widget_set_sensitive (label_role_tech[hind], 1);
}

void end_role_tech_list ()
{
}

static int role_mos_list_num;
void begin_role_mos_list ()
{
      tv_clear(tv_role_mos);
      role_mos_list_num = 0;
}

void add_role_mos_list (char * name, char * role, int job, int loc, float hpmax, float hp, float * prop, int is_quit, int lst, int pst)
{
      tv_add_people (role_mos_list_num, tv_role_mos, name, role, job, loc, hpmax, hp, prop, is_quit, lst, pst, -1);
      ++ role_mos_list_num;
}

void end_role_mos_list ()
{
      select_setup (tv_role_mos, 0);
}

static int role_mcs_list_num;
void begin_role_mcs_list ()
{
      tv_clear (tv_role_mcs);
      role_mcs_list_num = 0;
}

void add_role_mcs_list (char * name, int scope, float hpmax, float hp, char * role, int nmos, char * mayor, char * treasurer, float fengshui)
{
      tv_add_city (role_mcs_list_num, tv_role_mcs, name, scope, hpmax, hp, role, nmos, mayor, treasurer, fengshui);
      ++ role_mcs_list_num;
}

void end_role_mcs_list ()
{
      select_setup (tv_role_mcs, 0);
}

static int role_sos_list_num;
void begin_role_sos_list ()
{
      tv_clear (tv_role_sos);
      role_sos_list_num = 0;
}

void add_role_sos_list (char * name, char * role, int job, int loc, float hpmax, float hp, float * prop, int is_quit, int lst, int pst)
{
      tv_add_people (role_sos_list_num, tv_role_sos, name, role, job, loc, hpmax, hp, prop, is_quit, lst, pst, -1);
      ++ role_sos_list_num;
}

void end_role_sos_list ()
{
      select_setup (tv_role_sos, 0);
}

static int role_tos_list_num;
void begin_role_tos_list ()
{
      tv_clear (tv_role_tos);
      role_tos_list_num = 0;
}

void add_role_tos_list (char * name, char * role, int job, int loc, float hpmax, float hp, float * prop, int is_quit, int lst, int pst)
{
      tv_add_people (role_tos_list_num, tv_role_tos, name, role, job, loc, hpmax, hp, prop, is_quit, lst, pst, -1);
      ++ role_tos_list_num;
}

void end_role_tos_list ()
{
      select_setup (tv_role_tos, 0);
}

static int role_cards_list_num;
void begin_role_cards_list ()
{
      tv_clear (tv_role_cards);
      role_cards_list_num  = 0;
}

void add_role_cards_list (int dind, int num)
{
      GtkListStore * gls = GTK_LIST_STORE(gtk_tree_view_get_model (GTK_TREE_VIEW (tv_role_cards)));
      GtkTreeIter iter;
      gtk_list_store_append (gls, &iter);

      gtk_list_store_set (gls, &iter, 
            0, role_cards_list_num, 
            1, card_name[dind],
            2, num, 
            -1);

      ++ role_cards_list_num;
}

void end_role_cards_list ()
{
      select_setup (tv_role_cards, 0);
}

int get_role_mcs_select ()
{
      return check_role_mcs_select;
}

static int people_list_num ;
void begin_people_list(char * title)
{
      onoff_widget (button_people, 0);
      onoff_widget (button_people_main, 0);
      onoff_widget (button_people_vice, 0);
      gtk_label_set_label (GTK_LABEL(label_people_main), "无");
      gtk_label_set_label (GTK_LABEL(label_people_vice), "无");

      gtk_label_set_label ((GtkLabel *)label_people, title);
      tv_clear (tv_people);
      people_list_num  = 0;
}

void add_people_list(char * name, char * role, int job, int loc, float hpmax, float hp, float * prop, int is_quit, int lst, int pst)
{
      tv_add_people (people_list_num, tv_people, name, role, job, loc, hpmax, hp, prop, is_quit, lst, pst, 0);
      ++ people_list_num;
}

void end_people_list(int topof)
{
      switch (topof)
      {
      case 0:
            gtk_window_set_transient_for (GTK_WINDOW(win_people), GTK_WINDOW(win_main));
            break;
      case 1:
            gtk_window_set_transient_for (GTK_WINDOW(win_people), GTK_WINDOW(win_role));
            break;
      case 2:
            gtk_window_set_transient_for (GTK_WINDOW(win_people), GTK_WINDOW(win_city));
            break;
      }
}

void get_people_exit ()
{
      if (st == GAME_OVER) return;
      select_setup (tv_people, 0);
      people_select_st = 0;
      gtk_widget_show_all (win_people);
      gtk_main();
}

int get_people_one ()
{
      if (st == GAME_OVER) return -1;
      select_setup (tv_people, 0);
      onoff_widget (button_people, 1);
      people_select = -1;
      people_select_st = 0;
      gtk_widget_show_all (win_people);
      gtk_main ();
      if (st == GAME_OVER) return -1;
      return people_select;
}

int get_people_two_main () 
{
      if (st == GAME_OVER) return -1;
      select_setup (tv_people, 0);
      onoff_widget (button_people, 1);
      onoff_widget (button_people_main, 1);
      onoff_widget (button_people_vice, 1);
      people_select = -1;
      people_select_vice = -1;
      people_select_st = 1;
      gtk_widget_show_all (win_people);
      gtk_main ();
      if (st == GAME_OVER) return -1;
      return people_select;
}

int get_people_two_vice ()
{
      return people_select_vice;
}

int get_people_list_num ()
{
      if (st == GAME_OVER) return 0;
      select_setup (tv_people, 1);
      onoff_widget (button_people, 1);
      people_select_num = 0;
      people_select_st = 2;
      gtk_widget_show_all (win_people);
      gtk_main ();
      if (st == GAME_OVER) return 0;
      return people_select_num;
}

int get_people_list (int ind)
{
      return people_select_list[ind];
}

static int city_list_num;
void begin_city_list()
{
      tv_clear(tv_city);
      city_list_num = 0;
}

void add_city_list(char * name, int scope, float hpmax, float hp, char * role, int nmos, char * mayor, char * treasurer, float fengshui)
{
      tv_add_city (city_list_num, tv_city, name, scope, hpmax, hp, role, nmos, mayor, treasurer, fengshui);
      ++ city_list_num;
}

void end_city_list()
{
      gtk_window_set_transient_for (GTK_WINDOW(win_city), GTK_WINDOW(win_main));
      select_setup (tv_city, 0);
}

int get_city_mcs_select ()
{
      if (st == GAME_OVER) return -1;
      check_city_mcs_select = -1;
      gtk_widget_show_all(win_city);
      gtk_main ();
      if (st == GAME_OVER) return -1;
      return check_city_mcs_select;
}

static int inst_list_num ;
void begin_inst_list()
{
      tv_clear (tv_inst);
      tv_clear (tv_inst_mos);
      tv_clear (tv_inst_tech);
      inst_list_num = 0;
}

void add_inst_list(char * name, int ntech, char * mos, char * on_study, float point, int left_round)
{
      GtkListStore * gls = GTK_LIST_STORE(gtk_tree_view_get_model (GTK_TREE_VIEW (tv_inst)));
      GtkTreeIter iter;

      gtk_list_store_append (gls, &iter);
      gtk_list_store_set (gls, &iter, 
            0, inst_list_num,
            1, name,
            2, ntech, 
            3, mos, 
            4, on_study, 
            5, (int)ceil(point), 
            6, left_round, 
            -1);
      ++ inst_list_num;

}

void end_inst_list()
{
      gtk_window_set_transient_for (GTK_WINDOW(win_inst), GTK_WINDOW(win_main));
      select_setup (tv_inst, 0);
}

int get_inst_select ()
{
      if (st == GAME_OVER) return -1;
      check_inst_select = -1;
      gtk_widget_show_all (win_inst);
      gtk_main ();
      if (st == GAME_OVER) return -1;
      return check_inst_select;
}

static int inst_tech_list_num;
void begin_inst_tech_list (char * title)
{
      tv_clear (tv_inst_tech);
      inst_tech_list_num = 0;
}

void add_inst_tech_list (char * name, float study, int scond)
{
      tv_add_tech (inst_tech_list_num, tv_inst_tech, name, study, scond);
      ++ inst_tech_list_num;
}

void end_inst_tech_list ()
{
      select_setup (tv_inst_tech, 0);
}

void show_inst_mos (char * name, char * role, int job, int loc, float hpmax, float hp, float * prop, int is_quit, int lst, int pst)
{
      tv_clear (tv_inst_mos);
      tv_add_people (0, tv_inst_mos, name, role, job, loc, hpmax, hp, prop, is_quit, lst, pst, -1);
      select_setup (tv_inst_mos, 0);
}

static int train_list_num;
void begin_train_list()
{
      tv_clear (tv_train);
      tv_clear (tv_train_mos);
      train_list_num = 0;
}

void add_train_list(char * name, int item, char * mos, int round)
{
      GtkListStore * gls = GTK_LIST_STORE(gtk_tree_view_get_model (GTK_TREE_VIEW (tv_train)));
      GtkTreeIter iter;

      gtk_list_store_append (gls, &iter);
      gtk_list_store_set (gls, &iter, 
            0, train_list_num,
            1, name,
            2, prop_name[item], 
            3, mos, 
            4, round, 
            -1);
      ++ train_list_num;
}

void end_train_list()
{
      gtk_window_set_transient_for (GTK_WINDOW(win_train), GTK_WINDOW(win_main));
      select_setup (tv_train, 0);
}

int get_train_select ()
{
      if (st == GAME_OVER) return -1;
      check_train_select = -1;
      gtk_widget_show_all (win_train);
      gtk_main ();
      if (st == GAME_OVER) return -1;
      return check_train_select;
}

void show_train_mos (char * name, char * role, int job, int loc, float hpmax, float hp, float * prop, int is_quit, int lst, int pst)
{
      tv_clear (tv_train_mos);
      tv_add_people (0, tv_train_mos, name, role, job, loc, hpmax, hp, prop, is_quit, lst, pst, -1);
      select_setup (tv_train_mos, 0);
}

static int saloon_list_num;
void begin_saloon_list()
{
      onoff_widget (button_people, 1);
      onoff_widget (button_people_main, 0);
      onoff_widget (button_people_vice, 0);
      gtk_label_set_label (GTK_LABEL(label_people_main), "无");
      gtk_label_set_label (GTK_LABEL(label_people_vice), "无");
      gtk_label_set_label ((GtkLabel *)label_people, "请选择聘用人员");

      tv_clear (tv_people);
      saloon_list_num = 0;
}

void add_saloon_list (char * name, char * role, int job, int loc, float hpmax, float hp, float * prop, int is_quit, int lst, int pst, float price)
{
      tv_add_people (saloon_list_num, tv_people, name, role, job, loc, hpmax, hp, prop, is_quit, lst, pst, price);
      ++ saloon_list_num;
}

void end_saloon_list()
{
      gtk_window_set_transient_for (GTK_WINDOW(win_people), GTK_WINDOW(win_main));
      select_setup (tv_people, 0);
}

int get_saloon_select ()
{
      if (st == GAME_OVER) return -1;
      people_select_st = 0;
      people_select = -1;
      gtk_widget_show_all (win_people);
      gtk_main ();
      if (st == GAME_OVER) return -1;
      return people_select;
}

int sel_battle_scale (char * msg)
{
      if (st == GAME_OVER) return -1;
      msg_map (msg);
      onoff_widget (button_main_small, 1);
      onoff_widget (button_main_medium, 1);
      onoff_widget (button_main_large, 1);
      onoff_widget (button_main_no, 1);
      battle_scale = -1;
      st = WAIT_BATTLE_SCALE;
      gtk_main ();
      if (st == GAME_OVER) return -1;
      onoff_widget (button_main_small, 0);
      onoff_widget (button_main_medium, 0);
      onoff_widget (button_main_large, 0);
      onoff_widget (button_main_no, 0);
      return battle_scale;
}

void battle_start (int max_turn_) 
{
      int i, j;
      buf_clear (buf_battle_msg);
      for (i = 0; i < 2; ++ i)
      {
            labels_clear (label_battle_tech[i], 12);
            for (j = 0; j < 3; ++ j)
                  label_set_color_no (label_battle_st [i][j]);
      }

      gtk_toggle_button_set_active (GTK_TOGGLE_BUTTON(tb_battle_auto), 0);

      max_turn = max_turn_;
      is_battle_finish = 0;
      wait_battle_tick = 0;
      gtk_widget_show_all(win_battle);
}

void battle_msg (char *msg)
{
      buf_add_msg (buf_battle_msg, msg);
      swin_set_follow (swin_battle_msg);
}

void battle_set (int side, char * role, float power, float winp, float celve, float hpmax_, float hp, float def)
{
      gtk_label_set_label (GTK_LABEL(label_battle_role[side]), role);
      pb_set_value (pb_battle_power[side], power, 100);
      pb_set_value (pb_battle_winp[side], winp, 100);
      pb_set_value (pb_battle_celve[side], celve, 100);
      pb_set_value (pb_battle_hp[side], hp, hpmax_);
      pb_set_value (pb_battle_def[side], def, defmax);

      hpmax[side] = hpmax_;
}

void battle_tech_set (int side , int tech)
{
      onoff_widget (label_battle_tech[side][tech], 1);
}

void battle_change_hp (int side, float hp, float def)
{
      pb_set_value (pb_battle_hp [side], hp,  hpmax[side]);
      pb_set_value (pb_battle_def[side], def, defmax);
}

void battle_change_st (int side, int bst, int is_quit, int btime, int fst)
{
      if (fst)
            label_set_color_red (label_battle_st[side][0]);
      else
            label_set_color_no  (label_battle_st[side][0]);

      if (bst == 1)
            label_set_color_red (label_battle_st[side][1]);
      else
            label_set_color_no  (label_battle_st[side][1]);

      if (bst == 2)
            label_set_color_red (label_battle_st[side][2]);
      else
            label_set_color_no  (label_battle_st[side][2]);
}

void battle_active_tech (int side , int tech , float shp, float ohp, float sdef, float odef)
{
      label_set_color_blue (label_battle_tech[side][tech]);
      pb_set_value (pb_battle_hp[side], shp, hpmax[side]);
      pb_set_value (pb_battle_hp[1 - side], ohp, hpmax[1 - side]);
      pb_set_value (pb_battle_def[side], sdef, defmax);
      pb_set_value (pb_battle_def[1 - side], odef, defmax);
}

void battle_turn_start (int turn) 
{
      int i, j;
      if (st == GAME_OVER) return;
      if (is_battle_finish) return;
      battle_set_turn (turn);
      wait_battle_tick = 1;
      gtk_main ();
      for (i = 0; i < 2; ++ i)
      {
            for (j = 0; j < ntech; ++ j)
            {
                  label_set_color_no (label_battle_tech[i][j]);
            }
      }
}

void battle_end ()
{
      if (st == GAME_OVER) return;
      if (is_battle_finish) return;
      gtk_main ();
}

static int exchange_people_list_num[2];
void begin_exchange_people_list(int side, char * title)
{
      if (side)
            tv_clear (tv_exchange_out);
      else 
            tv_clear (tv_exchange_in);

      gtk_label_set_label (GTK_LABEL(label_exchange), title);
      exchange_people_list_num[side] = 0;
}

void add_exchange_people_list(int side, char * name, char * role, int job, int loc, float hpmax, float hp, float * prop, int is_quit, int lst, int pst)
{
      if (side)
            tv_add_people (exchange_people_list_num[side] ++, tv_exchange_out, name, role, job, loc, hpmax, hp, prop, is_quit, lst, pst, -1);
      else 
            tv_add_people (exchange_people_list_num[side] ++, tv_exchange_in, name, role, job, loc, hpmax, hp, prop, is_quit, lst, pst, -1);
}

void end_exchange_people_list(int side)
{
      if (side)
      {
            select_setup (tv_exchange_out, 1);
            gtk_window_set_transient_for (GTK_WINDOW(win_exchange), GTK_WINDOW(win_main));
      }
      else
      {
            select_setup (tv_exchange_in, 1);
      }
}

int get_exchange_people_list_num (int side)
{
      if (side)
      {
            return exchange_out_num;
      }
      else 
      {
            if (st == GAME_OVER) return 0;
            exchange_out_num = 0;
            exchange_in_num = 0;
            gtk_widget_show_all (win_exchange);
            gtk_main();
            if (st == GAME_OVER) return 0;
            return exchange_in_num;
      }
}

int get_exchange_people_list (int side, int isel)
{
      if (side)
      {
            return exchange_out_list [isel];
      }
      else
      {
            return exchange_in_list [isel];
      }
}

static int study_tech_list_num;
void begin_study_tech_list (char * name)
{
      gtk_label_set_label (GTK_LABEL(label_study), name);
      tv_clear (tv_study_tech);
      study_tech_list_num = 0;
}

void add_study_tech_list (char * name, float study, int scond)
{
      tv_add_tech (study_tech_list_num, tv_study_tech, name, study, scond);
      ++ study_tech_list_num;
}

void end_study_tech_list ()
{
      select_setup (tv_study_tech, 0);
}

static int study_people_list_num; 
void begin_study_people_list ()
{
      tv_clear (tv_study_people);
      study_people_list_num = 0;
}

void add_study_people_list (char * name, char * role, int job, int loc, float hpmax, float hp, float * prop, int is_quit, int lst, int pst)
{
      tv_add_people (study_people_list_num, tv_study_people, name, role, job, loc, hpmax, hp, prop, is_quit, lst, pst, -1);
      ++ study_people_list_num;
}

void end_study_people_list ()
{
      select_setup (tv_study_people, 0);
      gtk_window_set_transient_for (GTK_WINDOW(win_study), GTK_WINDOW(win_main));
}

int get_study_select (int * tech)
{
      if (st == GAME_OVER) return -1;
      study_tech = -1;
      study_people = -1;
      gtk_widget_show_all (win_study);
      gtk_main ();
      if (st == GAME_OVER) return -1;
      *tech = study_tech;
      return study_people;
}

int get_obj_type ()
{
      if (st == GAME_OVER) return -1;
      obj_type = -1;
      obj = -1;
      st = SEL_CARD_OBJ;
      onoff_widget (button_main_no, 1);
      gtk_main ();
      if (st == GAME_OVER) return -1;
      onoff_widget (button_main_no, 0);
      return obj_type;
}

int get_obj ()
{
      return obj;
}

void game_end (int res)
{
      if (st != GAME_OVER) 
      {
            st = GAME_OVER;
            if (res != -2)
                  gtk_main ();
      }

      gtk_widget_destroy (win_main);
      gtk_widget_destroy (win_card);
      gtk_widget_destroy (win_battle);
      gtk_widget_destroy (win_role);
      gtk_widget_destroy (win_people);
      gtk_widget_destroy (win_city);
      gtk_widget_destroy (win_inst);
      gtk_widget_destroy (win_train);
      gtk_widget_destroy (win_study);
      gtk_widget_destroy (win_exchange);
}

