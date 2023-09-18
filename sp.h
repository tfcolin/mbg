#ifndef sp_INC
#define sp_INC

/* sliding puzzle driver */

#ifdef __cpluscplus
extern "C" {
#endif

struct strSPDriver;
typedef struct strSPDriver SPDriver;

// -1: 中止 0: 返回选择位置到 (x, y) 1: 返回位置到 dir: 0: x+ 1: x- 2: y+ 3: y-; 
typedef int  (* SelLocFunc)  (SPDriver * sp, int * x, int * y, int * dir); 
typedef int  (* ConfStartFunc)  (SPDriver * sp); // -1: 终止 0: 继续
/* x, y: 实际位置, cx, cy: 位于该位置的图块的原始坐标. cx == -1: 表示空位置.
 */
typedef void (* NoteChangeFunc) (SPDriver * sp, int x, int y, int cx, int cy); 
typedef void (* NoteWinFunc) (SPDriver * sp, int res); // res = 0: 成功, 1: 中止

struct strSPDriver
{
      int nx;
      int ny;

      int eloc;
      int celoc;
      int * loc; // -1: 空位

      int step;

      SelLocFunc slf;
      ConfStartFunc csf;
      NoteWinFunc nwf;
      NoteChangeFunc ncf;

};

SPDriver * InitSP (int nx, int ny, int ex, int ey, SelLocFunc slf, ConfStartFunc csf, NoteWinFunc nwf, NoteChangeFunc ncf);
void FreeSP (SPDriver * sp);

int OptLoc (SPDriver * sp, int x, int y); // 0: 成功, -1: 失败
int OptDir (SPDriver * sp, int dir); // 0: 成功, -1: 失败
void GetLoc (SPDriver * sp, int x, int y, int * ox, int * oy);
int IsWin (SPDriver * sp); // 1: 胜利, 0: 未胜利

void RunSP (SPDriver * sp);

#ifdef __cpluscplus
}
#endif

#endif
