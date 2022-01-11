package gtkui

// #cgo pkg-config: gtk+-3.0
//#include <stdlib.h>
//#include "gtkui_wrapper.h"
import "C"

// bridge LoadUI and FreeUI function

func LoadUI () {
      C.load_ui ()
}
func FreeUI () {
      C.free_ui ()
}

// bridge c function to go function to use as function pointers

func add_role_mos_list_g(name *C.char, role *C.char, job C.int, loc C.int, hpmax C.float, hp C.float, prop *C.float, is_quit C.int, lst C.int, pst C.int) {
	C.add_role_mos_list(name, role, job, loc, hpmax, hp, prop, is_quit, lst, pst)
}
func add_role_sos_list_g(name *C.char, role *C.char, job C.int, loc C.int, hpmax C.float, hp C.float, prop *C.float, is_quit C.int, lst C.int, pst C.int) {
	C.add_role_sos_list(name, role, job, loc, hpmax, hp, prop, is_quit, lst, pst)
}
func add_role_tos_list_g(name *C.char, role *C.char, job C.int, loc C.int, hpmax C.float, hp C.float, prop *C.float, is_quit C.int, lst C.int, pst C.int) {
	C.add_role_tos_list(name, role, job, loc, hpmax, hp, prop, is_quit, lst, pst)
}
func add_people_list_g(name *C.char, role *C.char, job C.int, loc C.int, hpmax C.float, hp C.float, prop *C.float, is_quit C.int, lst C.int, pst C.int) {
	C.add_people_list(name, role, job, loc, hpmax, hp, prop, is_quit, lst, pst)
}
func show_inst_mos_g(name *C.char, role *C.char, job C.int, loc C.int, hpmax C.float, hp C.float, prop *C.float, is_quit C.int, lst C.int, pst C.int) {
	C.show_inst_mos(name, role, job, loc, hpmax, hp, prop, is_quit, lst, pst)
}
func show_train_mos_g(name *C.char, role *C.char, job C.int, loc C.int, hpmax C.float, hp C.float, prop *C.float, is_quit C.int, lst C.int, pst C.int) {
	C.show_train_mos(name, role, job, loc, hpmax, hp, prop, is_quit, lst, pst)
}
func add_study_people_list_g(name *C.char, role *C.char, job C.int, loc C.int, hpmax C.float, hp C.float, prop *C.float, is_quit C.int, lst C.int, pst C.int) {
	C.add_study_people_list(name, role, job, loc, hpmax, hp, prop, is_quit, lst, pst)
}

func add_role_mcs_list_g(name *C.char, scope C.int, hpmax C.float, hp C.float, role *C.char, nmos C.int, mayor *C.char, treasurer *C.char, fengshui C.float) {
	C.add_role_mcs_list(name, scope, hpmax, hp, role, nmos, mayor, treasurer, fengshui)
}
func add_city_list_g(name *C.char, scope C.int, hpmax C.float, hp C.float, role *C.char, nmos C.int, mayor *C.char, treasurer *C.char, fengshui C.float) {
	C.add_city_list(name, scope, hpmax, hp, role, nmos, mayor, treasurer, fengshui)
}

func add_study_tech_list_g(name *C.char, study C.float, scond C.int) {
	C.add_study_tech_list(name, study, scond)
}
func add_inst_tech_list_g(name *C.char, study C.float, scond C.int) {
	C.add_inst_tech_list(name, study, scond)
}
