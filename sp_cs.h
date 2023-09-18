#ifndef sp_cs_INC
#define sp_cs_INC

#ifdef __cplusplus
extern "C" {
#endif

void cs_init (int * nx_, int * ny_);
void cs_end ();

void redraw_win (int c, int x, int y, int nx, int cx, int cy);
void msgbox (char * msg); 
char inputbox (char * msg);

#ifdef __cplusplus
}
#endif

#endif
