#ifndef sp_cs_INC
#define sp_cs_INC

#ifdef __cplusplus
extern "C" {
#endif

void cs_init (int * nx, int * ny);
void cs_end ();
void pexit (char * msg);
void redraw_win (SPDriver * sp, int x, int y);

int  (slf)  (SPDriver * sp, int * x, int * y, int * dir);
int  (csf)  (SPDriver * sp); 
void (nwf) (SPDriver * sp, int res); 
void (ncf) (SPDriver * sp, int x, int y, int cx, int cy);

#ifdef __cplusplus
}
#endif

#endif
