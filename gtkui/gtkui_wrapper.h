#ifndef gtkui_wrapper_INC
#define gtkui_wrapper_INC

#ifdef __cplusplus
extern "C" {
#endif

void load_ui ();
void free_ui ();

/* mode: 0: OK. 1: YES/NO; 
 * return 0: OK. -1: CANCEL; 1: CHECK
 */
int msg_box (char * msg, int mode);
void msg_map (char * msg);

void init_map  (int ngrid, int nrole, int npeople, int ncity, int ninst, int ntrain, int ntech, int ncard, int nx_, int ny_);
void init_tech (int tech , char *tname, int scond);
void init_card (int card, char * dname, int odir, int otype);
void init_role (int role, char * rname);

void set_board   (int loc, int type, int ind, int x, int y); 
void set_city    (int id, char * name);
void set_inst    (int id, char * name);
void set_train   (int id, char * name);
void set_barrier (int loc, int is_unset);
void set_robber  (int loc, int is_unset);
void set_role    (int loc, int rind, int is_unset);

void finish_draw_map ();
void update_role_point (int loc);

// 0: 移动. 1: 使用卡片. 2: check(waiting). -1: 中途退出游戏
int get_action (int rind);

// 0: role. 1: people. 2: city; 3:inst. 4:trainingroom. 5: role_people
int get_check_type ();

// called after call get_action () (return 1) to show cards
void begin_cards_list ();
void add_cards_list (int cind, int num);
void end_cards_list ();
// -1: cancel
int get_card_action ();

/* return: 0-nrole-1: select role to check. -1: exit. (waiting)
 * check_mcs: [out] mcs to check. -1: show role info only
 */
int get_role_select (); 

/* called immediatly after each get_role_select() 
 * return -1: show base role info  
 * >=0: city id. show mcs city info 
 * if >=0: call begin_people_list ... get_people_exit in subsequence 
 */
int get_role_mcs_select ();

// check role info
void begin_role_list ();
void add_role_list (char * name, int loc, float money, int dir, int mst, int mtime, int cst, int ast, int atime);
void end_role_list ();
void begin_role_tech_list ();
void add_role_tech_list (int hind);
void end_role_tech_list ();
void begin_role_mos_list ();
void add_role_mos_list (char * name, char * role, int job, int loc, float hpmax, float hp, float * prop, int is_quit, int lst, int pst);
void end_role_mos_list ();
void begin_role_mcs_list ();
void add_role_mcs_list (char * name, int scope, float hpmax, float hp, char * role, int nmos, char * mayor, char * treasurer, float fengshui);
void end_role_mcs_list ();
void begin_role_sos_list ();
void add_role_sos_list (char * name, char * role, int job, int loc, float hpmax, float hp, float * prop, int is_quit, int lst, int pst);
void end_role_sos_list ();
void begin_role_tos_list ();
void add_role_tos_list (char * name, char * role, int job, int loc, float hpmax, float hp, float * prop, int is_quit, int lst, int pst);
void end_role_tos_list ();
void begin_role_cards_list ();
void add_role_cards_list (int dind, int num);
void end_role_cards_list ();

// check people info
void begin_people_list(char * title);
void add_people_list(char * name, char * role, int job, int loc, float hpmax, float hp, float * prop, int is_quit, int lst, int pst);
// topof 0: main; 1: role 2: city
void end_people_list(int topof);
/* return to exit check people info (waiting) */
void get_people_exit ();

/* return one select people. -1: cancel (waiting)*/
int get_people_one (); 
/* return two selected people. main == -1: cancel. vice == -1: only main (waiting) */
int get_people_two_main (); 
int get_people_two_vice ();
/* return a selected list of people. empty: cancel (waiting)*/
int get_people_list_num (); 
int get_people_list (int ind);

void begin_city_list();
void add_city_list(char * name, int scope, float hpmax, float hp, char * role, int nmos, char * mayor, char * treasurer, float fengshui);
void end_city_list();

/* return -1: exit;  (waiting)
 *        >=0: id in mcs list
 * if >=0: call begin_people_list ... get_people_exit in subsequence 
 */
int get_city_mcs_select ();

void begin_inst_list();
void add_inst_list(char * name, int ntech, char * mos, char * on_study, float point, int left_round);  
void end_inst_list();
// -1: exit. >=0: inst index (waiting)
int get_inst_select ();
void show_inst_mos (char * name, char * role, int job, int loc, float hpmax, float hp, float * prop, int is_quit, int lst, int pst);
void begin_inst_tech_list (char * title);
void add_inst_tech_list (char * name, float study, int scond);
void end_inst_tech_list ();

void begin_train_list();
void add_train_list(char * name, int item, char * mos, int round);  
void end_train_list();
// -1: exit. >=0: train index (waiting)
int get_train_select ();
void show_train_mos (char * name, char * role, int job, int loc, float hpmax, float hp, float * prop, int is_quit, int lst, int pst);

void begin_saloon_list();
void add_saloon_list (char * name, char * role, int job, int loc, float hpmax, float hp, float * prop, int is_quit, int lst, int pst, float price);
void end_saloon_list();
int get_saloon_select ();

int sel_battle_scale (char * msg); // choose battle scale. -1: cancel -2: check (waiting)
void battle_start (int max_turn); // open battle field.
void battle_msg (char *msg);
void battle_set (int side, char * role, float power, float winp, float celve, float hpmax, float hp, float def);
void battle_tech_set (int side , int tech);
void battle_change_hp (int side, float hp, float def);
// is_quit: 0: false 1:true
void battle_change_st (int side, int bst, int is_quit, int btime, int fst);
void battle_active_tech (int side , int tech , float shp, float ohp, float sdef, float odef);
void battle_turn_start (int turn); // waiting for battle end (timing)
void battle_end (); // close battle field.

// side = 0: role. side = 1: city
// side = 0 to open the select window.
void begin_exchange_people_list(int side, char * title);
void add_exchange_people_list(int side, char * name, char * role, int job, int loc, float hpmax, float hp, float * prop, int is_quit, int lst, int pst);
void end_exchange_people_list(int side);
// side = 0: 0->1. 1: 1->0
// side = 1: return immediatly. side = 0: return after select complete(waiting).
int get_exchange_people_list_num (int side);
int get_exchange_people_list (int side, int isel);

void begin_study_tech_list (char * name);
void add_study_tech_list (char * name, float study, int scond);
void end_study_tech_list ();
void begin_study_people_list ();
void add_study_people_list (char * name, char * role, int job, int loc, float hpmax, float hp, float * prop, int is_quit, int lst, int pst);
void end_study_people_list ();
// return people list num. -1: cancel. (waiting)
// tech [out] tech list num.
int get_study_select (int * tech);

// 返回位置: -1: 取消. -2: map. >=0: role. -3: check. (waiting)
int get_obj_type ();
// return immediatly (loc, for get_obj_type == -2)
int get_obj ();

// res: -1: 平局. >=0 取胜方 -2: 用户强制退出
void game_end (int res);

#ifdef __cplusplus                                              
}                                                               
#endif

#endif
