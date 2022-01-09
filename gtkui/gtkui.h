#ifndef gtkui_INC
#define gtkui_INC

#ifdef __cplusplus
extern "C" {
#endif

#define lstr 2048

typedef void (* SelPushFunc) (GtkTreeModel *model, GtkTreePath *path, GtkTreeIter *iter, gpointer data);
typedef enum 
{
      GAME_OVER,
      WAIT_ACTION,
      SEL_CARD_OBJ,
      WAIT_BATTLE_SCALE,
      WAIT_MSG_ANSWER,
      WAIT_MSG_CONFIRM
} Status;

typedef struct 
{
      int loc;
      int role;
} RoleLocInd;

extern char prop_name [5][64];

extern Status st;

// 返回位置: -1: 取消. -2: map. >=0: role. -3: check. (waiting)
extern int obj_type;
extern int obj;

// 0: 移动. 1: 使用卡片. 2: check(waiting). -1: 无效. 
extern int action; 
// -1: cancel. >=0: 卡片号
extern int card_action;

// 0: role. 1: people. 2: city; 3:inst. 4:trainingroom. 5: role_people
extern int check_type;
extern int check_role_select;
extern int check_role_mcs_select;
extern int check_city_mcs_select;
extern int check_inst_select;
extern int check_train_select;

extern int people_select_st; // 0: 单选一人; 1: 选主副将; 2: 选列表
extern int people_select;
extern int people_select_vice;
extern int people_select_list [2048];
extern int people_select_num;

// choose battle scale. -1: cancel -2: check (waiting)
extern int battle_scale;
extern int is_battle_finish;
extern int wait_battle_tick;

extern int exchange_out_num;
extern int exchange_out_list [2048];
extern int exchange_in_num;
extern int exchange_in_list [2048];

extern int study_tech;
extern int study_people;

/* return 0: OK. -1: CANCEL; 1: CHECK*/
extern int msg_res; // msg_button. 

extern int ngrid, nrole, npeople, ncity, ninst, ntrain, ntech, ncard, nx, ny;
extern int wp, hp;
extern int is_align;

extern char ** role_name;

extern char ** tech_name;
extern int  * tech_scond;

extern char ** card_name;
extern int  * card_odir;
extern int  * card_otype;

// 0: 空地 1:城 2:会馆 3:谋略研究所 4:修炼所 5: 机会
extern int  * map_type; // [ngrid]
extern int  * map_obj;  // [ngrid] object index

extern int * is_barrier; // [ngrid]
extern int * is_robber; // [ngrid]
extern int ** is_role; //[ngrid][nrole]

extern char ** city_name;
extern char ** inst_name;
extern char ** train_name;

extern GtkWidget * win_battle;
extern GtkWidget * label_battle; // title
extern GtkWidget * label_battle_role[2]; // label_battle_role_[01]
extern GtkWidget * pb_battle_power[2];
extern GtkWidget * pb_battle_winp[2]; // 胜率, 战法
extern GtkWidget * pb_battle_celve[2]; // 策略
extern GtkWidget * pb_battle_hp[2];
extern GtkWidget * pb_battle_def[2];
extern GtkWidget * label_battle_tech[2][12]; // tech
extern GtkWidget * label_battle_st[2][3]; // 状态: 0: burn; 1: 混乱. 2: 内讧
extern GtkWidget * tv_battle_msg; // 战斗消息
extern GtkWidget * button_battle_next;
extern GtkWidget * tb_battle_auto;
extern GtkTextBuffer * buf_battle_msg;
extern GtkWidget * swin_battle_msg;

extern GtkWidget * win_people;
extern GtkWidget * label_people;
extern GtkWidget * tv_people;
extern GtkWidget * button_people;
extern GtkWidget * button_people_main;
extern GtkWidget * button_people_vice;
extern GtkWidget * label_people_main;
extern GtkWidget * label_people_vice;

extern GtkWidget * win_city;
extern GtkWidget * tv_city;
extern GtkWidget * button_city;

extern GtkWidget * win_card;
extern GtkWidget * tv_card;
extern GtkWidget * button_card;

extern GtkWidget * win_inst;
extern GtkWidget * tv_inst;
extern GtkWidget * tv_inst_mos;
extern GtkWidget * tv_inst_tech;
extern GtkWidget * button_inst;

extern GtkWidget * win_train;
extern GtkWidget * tv_train;
extern GtkWidget * tv_train_mos;
extern GtkWidget * button_train;

extern GtkWidget * win_exchange;
extern GtkWidget * label_exchange;
extern GtkWidget * tv_exchange_in;
extern GtkWidget * tv_exchange_out;
extern GtkWidget * button_exchange;

extern GtkWidget * win_study;
extern GtkWidget * label_study;
extern GtkWidget * tv_study_tech;
extern GtkWidget * tv_study_people;
extern GtkWidget * button_study;

extern GtkWidget * win_role;
extern GtkWidget * tv_role;
extern GtkWidget * tv_role_mos;
extern GtkWidget * tv_role_mcs;
extern GtkWidget * tv_role_sos;
extern GtkWidget * tv_role_tos;
extern GtkWidget * tv_role_cards;
extern GtkWidget * label_role_tech[12];
extern GtkWidget * button_role;
extern GtkWidget * button_role_mcs;

extern GtkWidget * win_main;
extern GtkWidget * layout_main_map;
extern GtkWidget * button_main_move;
extern GtkWidget * button_main_card;
extern GtkWidget * button_main_role;
extern GtkWidget * button_main_people;
extern GtkWidget * button_main_rp;
extern GtkWidget * button_main_city;
extern GtkWidget * button_main_inst;
extern GtkWidget * button_main_train;
extern GtkWidget * tv_main_msg;
extern GtkWidget * button_main_yes;
extern GtkWidget * button_main_no;
extern GtkWidget * button_main_small;
extern GtkWidget * button_main_medium;
extern GtkWidget * button_main_large;

extern GtkWidget ** map_point; // [ngrid]
extern GtkWidget *** role_point; // [ngrid][nrole]

extern GtkTextBuffer * buf_main_msg;
extern GtkWidget * swin_main_msg;

void cbf_main_board_click (GtkWidget * widget, gpointer data);
void cbf_main_role_click (GtkWidget * widget, gpointer data);
void cbf_main_move_click (GtkWidget * widget, gpointer data);
void cbf_main_card_click (GtkWidget * widget, gpointer data);
void cbf_main_role_check (GtkWidget * widget, gpointer data);
void cbf_main_people_check (GtkWidget * widget, gpointer data);
void cbf_main_rp_check (GtkWidget * widget, gpointer data);
void cbf_main_city_check (GtkWidget * widget, gpointer data);
void cbf_main_inst_check (GtkWidget * widget, gpointer data);
void cbf_main_train_check (GtkWidget * widget, gpointer data);
void cbf_main_yes_click (GtkWidget * widget, gpointer data);
void cbf_main_no_click (GtkWidget * widget, gpointer data);
void cbf_main_small_click (GtkWidget * widget, gpointer data);
void cbf_main_medium_click (GtkWidget * widget, gpointer data);
void cbf_main_large_click (GtkWidget * widget, gpointer data);

void cbf_battle_next_click (GtkWidget * widget, gpointer data);
void cbf_people_click (GtkWidget * widget, gpointer data);
void cbf_people_main_click (GtkWidget * widget, gpointer data);
void cbf_people_vice_click (GtkWidget * widget, gpointer data);
void cbf_city_click (GtkWidget * widget, gpointer data);
void cbf_card_click (GtkWidget * widget, gpointer data);
void cbf_inst_click (GtkWidget * widget, gpointer data);
void cbf_train_click (GtkWidget * widget, gpointer data);
void cbf_exchange_click  (GtkWidget * widget, gpointer data);
void cbf_study_click (GtkWidget * widget, gpointer data);
void cbf_role_click (GtkWidget * widget, gpointer data);
void cbf_role_mcs_click (GtkWidget * widget, gpointer data);

int cbf_main_close (GtkWidget *widget, GdkEvent  *event, gpointer   user_data);
int cbf_card_close (GtkWidget *widget, GdkEvent  *event, gpointer   user_data);
int cbf_battle_close (GtkWidget *widget, GdkEvent  *event, gpointer   user_data);
int cbf_role_close (GtkWidget *widget, GdkEvent  *event, gpointer   user_data);
int cbf_people_close (GtkWidget *widget, GdkEvent  *event, gpointer   user_data);
int cbf_city_close (GtkWidget *widget, GdkEvent  *event, gpointer   user_data);
int cbf_inst_close (GtkWidget *widget, GdkEvent  *event, gpointer   user_data);
int cbf_train_close (GtkWidget *widget, GdkEvent  *event, gpointer   user_data);
int cbf_study_close (GtkWidget *widget, GdkEvent  *event, gpointer   user_data);
int cbf_exchange_close (GtkWidget *widget, GdkEvent  *event, gpointer   user_data);

int cbf_battle_timer (gpointer data);

#ifdef __cplusplus
}
#endif

#endif
