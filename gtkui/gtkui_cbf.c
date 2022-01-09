#include <gtk/gtk.h>

#include "gtkui_wrapper.h"
#include "gtkui.h"

#include <stdio.h>
#include <stdlib.h>

void spf_exchange_select_in  (GtkTreeModel *model, GtkTreePath *path, GtkTreeIter *iter, gpointer data);
void spf_exchange_select_out (GtkTreeModel *model, GtkTreePath *path, GtkTreeIter *iter, gpointer data);
void spf_people_select (GtkTreeModel *model, GtkTreePath *path, GtkTreeIter *iter, gpointer data);

int get_select_id   (GtkWidget * tv, char * name);
void put_select_ids (GtkWidget * tv, SelPushFunc spf);

void set_check_status ();

int get_select_id (GtkWidget * tv, char * name)
{
      GtkTreeSelection * ts = gtk_tree_view_get_selection (GTK_TREE_VIEW (tv));
      GtkTreeModel * model = gtk_tree_view_get_model (GTK_TREE_VIEW (tv));
      GtkTreeIter iter;
      int id;

      int is_sel = gtk_tree_selection_get_selected (ts, &model, &iter);

      if (is_sel)
      {
            gtk_tree_model_get (model, &iter, 0, &id, -1);
            if (name)
            {
                  char * rstr;
                  gtk_tree_model_get (model, &iter, 1, &rstr, -1);
                  strcpy (name, rstr);
                  free (rstr);
            }
      }
      else 
      {
            id = -1;
            if (name)
                  strcpy (name, "");
      }

      return id;
}

void put_select_ids (GtkWidget * tv, SelPushFunc spf)
{
      GtkTreeSelection * gts = gtk_tree_view_get_selection (GTK_TREE_VIEW (tv));
      gtk_tree_selection_selected_foreach (gts, spf, NULL);
}

void spf_people_select (GtkTreeModel *model, GtkTreePath *path, GtkTreeIter *iter, gpointer data)
{
      int id;
      gtk_tree_model_get (model, iter, 0, &id, -1);
      people_select_list [people_select_num ++] = id;
}

void spf_exchange_select_in  (GtkTreeModel *model, GtkTreePath *path, GtkTreeIter *iter, gpointer data)
{
      int id; 
      gtk_tree_model_get (model, iter, 0, &id, -1);
      exchange_in_list [exchange_in_num ++] = id;
}

void spf_exchange_select_out (GtkTreeModel *model, GtkTreePath *path, GtkTreeIter *iter, gpointer data)
{
      int id;
      gtk_tree_model_get (model, iter, 0, &id, -1);
      exchange_out_list [exchange_out_num ++] = id;
}

void cbf_main_board_click (GtkWidget * widget, gpointer data)
{
      if (st == GAME_OVER) return;
      int loc = *(int *)(data);
      if (st != SEL_CARD_OBJ) return;
      obj_type = -2;
      obj = loc;
      gtk_main_quit();
}

void cbf_main_role_click (GtkWidget * widget, gpointer data)
{
      if (st == GAME_OVER) return;
      int loc = ((RoleLocInd *)data) -> loc;
      int rind = ((RoleLocInd *)data) -> role;
      if (st != SEL_CARD_OBJ) return;
      if (!is_role[loc][rind]) return;
      obj_type = rind;
      gtk_main_quit();
}

void cbf_main_move_click (GtkWidget * widget, gpointer data)
{
      if (st == GAME_OVER) return;
      if (st != WAIT_ACTION) return;
      action = 0;
      gtk_main_quit();
}

void cbf_main_card_click (GtkWidget * widget, gpointer data)
{
      if (st == GAME_OVER) return;
      if (st != WAIT_ACTION) return;
      action = 1;
      gtk_main_quit();
}

void set_check_status ()
{
      switch (st)
      {
      case WAIT_ACTION:
            action = 2;
            break;
      case SEL_CARD_OBJ:
            obj_type = -3;
            break;
      case WAIT_BATTLE_SCALE:
            battle_scale = -2;
            break;
      case WAIT_MSG_ANSWER:
            msg_res = 1;
            break;
      }
}

void cbf_main_role_check (GtkWidget * widget, gpointer data)
{
      if (st == GAME_OVER) return;
      if (st == WAIT_MSG_CONFIRM) return;
      set_check_status();
      check_type = 0;
      gtk_main_quit();
}

void cbf_main_people_check (GtkWidget * widget, gpointer data)
{
      if (st == GAME_OVER) return;
      if (st == WAIT_MSG_CONFIRM) return;
      set_check_status();
      check_type = 1;
      gtk_main_quit();
}

void cbf_main_rp_check (GtkWidget * widget, gpointer data)
{
      if (st == GAME_OVER) return;
      if (st == WAIT_MSG_CONFIRM) return;
      set_check_status();
      check_type = 5;
      gtk_main_quit();
}

void cbf_main_city_check (GtkWidget * widget, gpointer data)
{
      if (st == GAME_OVER) return;
      if (st == WAIT_MSG_CONFIRM) return;
      set_check_status();
      check_type = 2;
      gtk_main_quit();
}

void cbf_main_inst_check (GtkWidget * widget, gpointer data)
{
      if (st == GAME_OVER) return;
      if (st == WAIT_MSG_CONFIRM) return;
      set_check_status();
      check_type = 3;
      gtk_main_quit();
}

void cbf_main_train_check (GtkWidget * widget, gpointer data)
{
      if (st == GAME_OVER) return;
      if (st == WAIT_MSG_CONFIRM) return;
      set_check_status();
      check_type = 4;
      gtk_main_quit();
}

void cbf_main_yes_click (GtkWidget * widget, gpointer data)
{
      if (st != WAIT_MSG_ANSWER && st != WAIT_MSG_CONFIRM) return;
      msg_res = 0;
      gtk_main_quit();
}

void cbf_main_no_click (GtkWidget * widget, gpointer data)
{
      switch (st)
      {
      case WAIT_MSG_ANSWER:
            msg_res = -1;
            break;
      case WAIT_BATTLE_SCALE:
            battle_scale = -1;
            break;
      case SEL_CARD_OBJ:
            obj_type = -1;
            obj = -1;
            break;
      default:
            return;
      }
      gtk_main_quit();
}

void cbf_main_small_click (GtkWidget * widget, gpointer data)
{
      if (st != WAIT_BATTLE_SCALE) return;
      battle_scale = 0;
      gtk_main_quit();
}

void cbf_main_medium_click (GtkWidget * widget, gpointer data)
{
      if (st != WAIT_BATTLE_SCALE) return;
      battle_scale = 1;
      gtk_main_quit();
}

void cbf_main_large_click (GtkWidget * widget, gpointer data)
{
      if (st != WAIT_BATTLE_SCALE) return;
      battle_scale = 2;
      gtk_main_quit();
}


void cbf_battle_next_click (GtkWidget * widget, gpointer data)
{
      if (st == GAME_OVER) return;
      if (!wait_battle_tick) return;
      wait_battle_tick = 0;
      gtk_main_quit();
}

int cbf_battle_timer (gpointer data)
{
      if (st == GAME_OVER) return TRUE;
      if (!is_align) return FALSE;
      int is_auto = gtk_toggle_button_get_active(GTK_TOGGLE_BUTTON(tb_battle_auto));
      if (!is_auto) return TRUE;
      if (!wait_battle_tick) return TRUE;
      wait_battle_tick = 0;
      gtk_main_quit();
      return TRUE;
}

void cbf_people_click (GtkWidget * widget, gpointer data)
{
      if (st == GAME_OVER) return;

      switch (people_select_st)
      {
      case 0:
            people_select = get_select_id (tv_people, 0);
            break;
      case 1:
            break;
      case 2:
            put_select_ids (tv_people, spf_people_select);
      }

      gtk_widget_hide (win_people);
      gtk_main_quit();
}

void cbf_people_main_click (GtkWidget * widget, gpointer data)
{
      if (st == GAME_OVER) return;
      if (people_select_st != 1) return;
      char name[lstr];

      people_select = get_select_id (tv_people, name);
      gtk_label_set_label (GTK_LABEL(label_people_main), name);
}

void cbf_people_vice_click (GtkWidget * widget, gpointer data)
{
      if (st == GAME_OVER) return;
      if (people_select_st != 1) return;
      char name[lstr];

      people_select_vice = get_select_id (tv_people, name);
      gtk_label_set_label (GTK_LABEL(label_people_vice), name);
}

void cbf_city_click (GtkWidget * widget, gpointer data)
{
      if (st == GAME_OVER) return;
      check_city_mcs_select = get_select_id (tv_city, 0);
      gtk_main_quit();
}

void cbf_card_click (GtkWidget * widget, gpointer data)
{
      if (st == GAME_OVER) return;
      card_action = get_select_id (tv_card, 0);
      gtk_widget_hide (win_card);
      gtk_main_quit();
}

void cbf_inst_click (GtkWidget * widget, gpointer data)
{
      if (st == GAME_OVER) return;
      check_inst_select = get_select_id (tv_inst, 0);
      gtk_main_quit();
}

void cbf_train_click (GtkWidget * widget, gpointer data)
{
      if (st == GAME_OVER) return;
      check_train_select = get_select_id(tv_train, 0);
      gtk_main_quit();
}

void cbf_exchange_click  (GtkWidget * widget, gpointer data)
{
      if (st == GAME_OVER) return;
      put_select_ids (tv_exchange_in, spf_exchange_select_in);
      put_select_ids (tv_exchange_out, spf_exchange_select_out);
      gtk_widget_hide (win_exchange);
      gtk_main_quit();
}

void cbf_study_click (GtkWidget * widget, gpointer data)
{
      if (st == GAME_OVER) return;
      study_tech = get_select_id (tv_study_tech, 0);
      study_people = get_select_id (tv_study_people, 0);
      if (study_tech == -1 || study_people == -1) return;
      gtk_widget_hide (win_study);
      gtk_main_quit();
}

void cbf_role_click (GtkWidget * widget, gpointer data)
{
      if (st == GAME_OVER) return;
      check_role_select = get_select_id (tv_role, 0);
      check_role_mcs_select = -1;
      gtk_main_quit();
      return;
}

void cbf_role_mcs_click (GtkWidget * widget, gpointer data)
{
      if (st == GAME_OVER) return;
      check_role_select = get_select_id (tv_role, 0);
      check_role_mcs_select = get_select_id (tv_role_mcs, 0);
      gtk_main_quit();
      return;
}

int cbf_main_close (GtkWidget *widget, GdkEvent  *event, gpointer   user_data)
{
      st = GAME_OVER;
      gtk_widget_hide(widget);
      gtk_main_quit();
      return FALSE;
}

int cbf_card_close (GtkWidget *widget, GdkEvent  *event, gpointer   user_data)
{
      card_action = -1;
      gtk_widget_hide(widget);
      gtk_main_quit();

      return TRUE;
}

int cbf_battle_close (GtkWidget *widget, GdkEvent  *event, gpointer   user_data)
{
      wait_battle_tick = 0;
      is_battle_finish = 1;
      gtk_widget_hide (widget);
      gtk_main_quit();

      return TRUE;
}

int cbf_role_close (GtkWidget *widget, GdkEvent  *event, gpointer   user_data)
{
      check_role_select = -1;
      check_role_mcs_select = -1;
      gtk_widget_hide (widget);
      gtk_main_quit();
      return TRUE;
}

int cbf_people_close (GtkWidget *widget, GdkEvent  *event, gpointer   user_data)
{
      people_select = -1;
      people_select_vice = -1;
      people_select_num = 0;
      gtk_widget_hide (widget);
      gtk_main_quit();
      return TRUE;
}

int cbf_city_close (GtkWidget *widget, GdkEvent  *event, gpointer   user_data)
{
      check_city_mcs_select = -1;
      gtk_widget_hide (widget);
      gtk_main_quit();
      return TRUE;
}

int cbf_inst_close (GtkWidget *widget, GdkEvent  *event, gpointer   user_data)
{
      check_inst_select = -1;
      gtk_widget_hide (widget);
      gtk_main_quit();
      return TRUE;
}

int cbf_train_close (GtkWidget *widget, GdkEvent  *event, gpointer   user_data)
{
      check_train_select = -1;
      gtk_widget_hide (widget);
      gtk_main_quit();
      return TRUE;
}

int cbf_study_close (GtkWidget *widget, GdkEvent  *event, gpointer   user_data)
{
      study_tech = study_people = -1;
      gtk_widget_hide (widget);
      gtk_main_quit();
      return TRUE;
}

int cbf_exchange_close (GtkWidget *widget, GdkEvent  *event, gpointer   user_data)
{
      exchange_in_num = exchange_out_num = 0;
      gtk_widget_hide (widget);
      gtk_main_quit();
      return TRUE;
}

