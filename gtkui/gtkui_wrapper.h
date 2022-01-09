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
long msg_box (char * msg, long mode);
void msg_map (char * msg);

void init_map  (long ngrid, long nrole, long npeople, long ncity, long ninst, long ntrain, long ntech, long ncard, long nx_, long ny_);
void init_tech (long tech , char *tname, long scond);
void init_card (long card, char * dname, long odir, long otype);
void init_role (long role, char * rname);

void set_board   (long loc, long type, long ind, long x, long y); 
void set_city    (long id, char * name);
void set_inst    (long id, char * name);
void set_train   (long id, char * name);
void set_barrier (long loc, long is_unset);
void set_robber  (long loc, long is_unset);
void set_role    (long loc, long rind, long is_unset);

void finish_draw_map ();
void update_role_point (long loc);

// 0: 移动. 1: 使用卡片. 2: check(waiting). -1: 中途退出游戏
long get_action (long rind);

// 0: role. 1: people. 2: city; 3:inst. 4:trainingroom. 5: role_people
long get_check_type ();

// called after call get_action () (return 1) to show cards
void begin_cards_list ();
void add_cards_list (long cind, long num);
void end_cards_list ();
// -1: cancel
long get_card_action ();

/* return: 0-nrole-1: select role to check. -1: exit. (waiting)
 * check_mcs: [out] mcs to check. -1: show role info only
 */
long get_role_select (); 

/* called immediatly after each get_role_select() 
 * return -1: show base role info  
 * >=0: city id. show mcs city info 
 * if >=0: call begin_people_list ... get_people_exit in subsequence 
 */
long get_role_mcs_select ();

// check role info
void begin_role_list ();
void add_role_list (char * name, long loc, float money, long dir, long mst, long mtime, long cst, long ast, long atime);
void end_role_list ();
void begin_role_tech_list ();
void add_role_tech_list (long hind);
void end_role_tech_list ();
void begin_role_mos_list ();
void add_role_mos_list (char * name, char * role, long job, long loc, float hpmax, float hp, float * prop, long is_quit, long lst, long pst);
void end_role_mos_list ();
void begin_role_mcs_list ();
void add_role_mcs_list (char * name, long scope, float hpmax, float hp, char * role, long nmos, char * mayor, char * treasurer, float fengshui);
void end_role_mcs_list ();
void begin_role_sos_list ();
void add_role_sos_list (char * name, char * role, long job, long loc, float hpmax, float hp, float * prop, long is_quit, long lst, long pst);
void end_role_sos_list ();
void begin_role_tos_list ();
void add_role_tos_list (char * name, char * role, long job, long loc, float hpmax, float hp, float * prop, long is_quit, long lst, long pst);
void end_role_tos_list ();
void begin_role_cards_list ();
void add_role_cards_list (long dind, long num);
void end_role_cards_list ();

// check people info
void begin_people_list(char * title);
void add_people_list(char * name, char * role, long job, long loc, float hpmax, float hp, float * prop, long is_quit, long lst, long pst);
// topof 0: main; 1: role 2: city
void end_people_list(long topof);
/* return to exit check people info (waiting) */
void get_people_exit ();

/* return one select people. -1: cancel (waiting)*/
long get_people_one (); 
/* return two selected people. main == -1: cancel. vice == -1: only main (waiting) */
long get_people_two_main (); 
long get_people_two_vice ();
/* return a selected list of people. empty: cancel (waiting)*/
long get_people_list_num (); 
long get_people_list (long ind);

void begin_city_list();
void add_city_list(char * name, long scope, float hpmax, float hp, char * role, long nmos, char * mayor, char * treasurer, float fengshui);
void end_city_list();

/* return -1: exit;  (waiting)
 *        >=0: id in mcs list
 * if >=0: call begin_people_list ... get_people_exit in subsequence 
 */
long get_city_mcs_select ();

void begin_inst_list();
void add_inst_list(char * name, long ntech, char * mos, char * on_study, float point, long left_round);  
void end_inst_list();
// -1: exit. >=0: inst index (waiting)
long get_inst_select ();
void show_inst_mos (char * name, char * role, long job, long loc, float hpmax, float hp, float * prop, long is_quit, long lst, long pst);
void begin_inst_tech_list ();
void add_inst_tech_list (char * name, float study, long scond);
void end_inst_tech_list ();

void begin_train_list();
void add_train_list(char * name, long item, char * mos, long round);  
void end_train_list();
// -1: exit. >=0: train index (waiting)
long get_train_select ();
void show_train_mos (char * name, char * role, long job, long loc, float hpmax, float hp, float * prop, long is_quit, long lst, long pst);

void begin_saloon_list();
void add_saloon_list (char * name, char * role, long job, long loc, float hpmax, float hp, float * prop, long is_quit, long lst, long pst, float price);
void end_saloon_list();
long get_saloon_select ();

long sel_battle_scale (char * msg); // choose battle scale. -1: cancel -2: check (waiting)
void battle_start (long max_turn); // open battle field.
void battle_msg (char *msg);
void battle_set (long side, char * role, float power, float winp, float celve, float hpmax, float hp, float def);
void battle_tech_set (long side , long tech);
void battle_change_hp (long side, float hp, float def);
// is_quit: 0: false 1:true
void battle_change_st (long side, long bst, long is_quit, long btime, long fst);
void battle_active_tech (long side , long tech , float shp, float ohp, float sdef, float odef);
void battle_turn_start (long turn); // waiting for battle end (timing)
void battle_end (); // close battle field.

// side = 0: role. side = 1: city
// side = 0 to open the select window.
void begin_exchange_people_list(long side, char * title);
void add_exchange_people_list(long side, char * name, char * role, long job, long loc, float hpmax, float hp, float * prop, long is_quit, long lst, long pst);
void end_exchange_people_list(long side);
// side = 0: 0->1. 1: 1->0
// side = 1: return immediatly. side = 0: return after select complete(waiting).
long get_exchange_people_list_num (long side);
long get_exchange_people_list (long side, long isel);

void begin_study_tech_list (char * name);
void add_study_tech_list (char * name, float study, long scond);
void end_study_tech_list ();
void begin_study_people_list ();
void add_study_people_list (char * name, char * role, long job, long loc, float hpmax, float hp, float * prop, long is_quit, long lst, long pst);
void end_study_people_list ();
// return people list num. -1: cancel. (waiting)
// tech [out] tech list num.
long get_study_select (long * tech);

// 返回位置: -1: 取消. -2: map. >=0: role. -3: check. (waiting)
long get_obj_type ();
// return immediatly (loc, for get_obj_type == -2)
long get_obj ();

// res: -1: 平局. >=0 取胜方 -2: 用户强制退出
void game_end (long res);

#ifdef __cplusplus                                              
}                                                               
#endif

#endif
