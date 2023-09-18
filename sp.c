#include "sp.h"
#include <math.h>
#include <time.h>
#include <stdlib.h>

int irand (int n)
{
      int res;
      res = floor ((double)(rand()) * n / RAND_MAX);
      return res;
}

SPDriver * InitSP (int nx, int ny, int ex, int ey, SelLocFunc slf, ConfStartFunc csf, NoteWinFunc nwf, NoteChangeFunc ncf)
{
      int i;

      srand (time(0));

      SPDriver * sp = malloc (sizeof(SPDriver));
      sp->nx = nx;
      sp->ny = ny;
      sp->step = 0;
      sp->slf = slf;
      sp->csf = csf;
      sp->nwf = nwf;
      sp->ncf = ncf;

      sp->eloc = ey * nx + ex;
      sp->celoc = sp->eloc;
      sp->loc = malloc (sizeof(int) * nx * ny);
      for (i = 0; i < ny * nx; ++ i) 
      {
            if (i == sp->eloc) sp->loc[i] = -1;
            else sp->loc [i] = i;
      }

      int nstep = 100 + irand (40);

      if (nstep % 2 == 1) ++ nstep;

      for (i = 0; i < nstep; ++ i)
      {
            int j1, j2, temp;
            j1 = sp->eloc;
            while (j1 == sp->eloc) j1 = irand (nx * ny);
            j2 = j1;
            while (j2 == j1 || j2 == sp->eloc) j2 = irand (nx * ny);

            temp = sp->loc[j1];
            sp->loc[j1] = sp->loc[j2];
            sp->loc[j2] = temp;
      }

      return sp;
}

void FreeSP (SPDriver * sp)
{
      free (sp->loc);
      free (sp);
}

int IsWin (SPDriver * sp)
{
      int i;
      for (i = 0; i < sp->nx * sp->ny; ++ i)
      {
            if (sp->loc[i] == -1)
            {
                  if (i != sp->eloc) return 0;
            }
            else
            {
                  if (sp->loc[i] != i) return 0;
            }
      }
      return 1;
}

int OptLoc (SPDriver * sp, int x0, int y0)
{
      int i, j;
      int nx = sp->nx, ny = sp->ny;
      int x1, y1, loc1, loc0, x, y, l, xx, yy, ll;

      if (x0 < 0 || x0 >= nx) return -1;
      if (y0 < 0 || y0 >= ny) return -1;

      loc0 = y0 * nx + x0; 
      loc1 = sp->celoc;
      x1 = loc1 % nx;
      y1 = loc1 / nx;

      if (x1 == x0)
      {
            if (y1 == y0) return -1;
            x = x0;
            int s = y0 > y1 ? 1 : -1;
            for (y = y1; (y0 - y) * s > 0; y += s)
            {
                  l = y * nx + x;
                  xx = x;
                  yy = y + s;
                  ll = yy * nx + xx;

                  sp->loc[l] = sp->loc[ll];
                  sp->ncf (sp, x, y, sp->loc[l] % nx, sp->loc[l] / nx);
            }
            sp->loc [loc0] = -1;
            sp->celoc = loc0;
            sp->ncf (sp, x0, y0, -1, -1);
            sp->step ++;
            return 0;
      }
      if (y1 == y0)
      {
            if (x1 == x0) return -1;
            y = y0;
            int s = x0 > x1 ? 1 : -1;
            for (x = x1; (x0 - x) * s > 0; x += s)
            {
                  l = y * nx + x;
                  yy = y;
                  xx = x + s;
                  ll = yy * nx + xx;

                  sp->loc[l] = sp->loc[ll];
                  sp->ncf (sp, x, y, sp->loc[l] % nx, sp->loc[l] / nx);
            }
            sp->loc [loc0] = -1;
            sp->celoc = loc0;
            sp->ncf (sp, x0, y0, -1, -1);
            sp->step ++;
            return 0;
      }

      return -1;
}

int OptDir (SPDriver * sp, int dir)
{
      int nx = sp->nx, ny = sp->ny;
      int x, y, loc, x1, y1, loc1;

      loc = sp->celoc;
      x = sp->celoc % nx;
      y = sp->celoc / nx;
      x1 = x;
      y1 = y;

      switch (dir)
      {
      case 0: -- x1; break;
      case 1: ++ x1; break;
      case 2: -- y1; break;
      case 3: ++ y1; break;
      }

      if (x1 < 0  || x1 >= nx) return -1;
      if (y1 < 0  || y1 >= ny) return -1;

      loc1 = y1 * nx + x1;

      sp->loc[loc] = sp->loc[loc1];
      sp->loc[loc1] = -1;
      sp->celoc = loc1;
      sp->step ++;

      sp->ncf (sp, x, y, sp->loc[loc] % nx, sp->loc[loc] / nx);
      sp->ncf (sp, x1, y1, -1, -1);
      return 0;
}

void GetLoc (SPDriver * sp, int x, int y, int * ox, int * oy)
{
      int nx = sp->nx, ny = sp->ny;
      int loc = sp->loc[y * nx + x];
      *ox = loc % nx;
      *oy = loc / nx;
}

void RunSP (SPDriver * sp)
{
      int i;
      int nx = sp->nx, ny = sp->ny;
      int res = 0, x, y, dir, res_slf;

      for (y = 0; y < ny; ++ y)
      {
            for (x = 0; x < nx; ++ x)
            {
                  int cx; 
                  int cy; 
                  i = y * nx + x;

                  if (sp->loc[i] == -1) cx = cy = -1;
                  else
                  {
                        cx = sp->loc[i] % nx;
                        cy = sp->loc[i] / nx;
                  }

                  sp->ncf (sp, x, y, cx, cy);
            }
      }

      if (sp->csf (sp)) 
      {
            sp->nwf (sp, 1);
            return;
      }

      while (1)
      {
            int opt;

            res_slf = sp->slf (sp, &x, &y, &dir);
            if (res_slf == -1) 
            {
                  res = -1;
                  break;
            }
            else if (res_slf == 1)
            {
                  opt = OptDir (sp, dir);
            }
            else
            {
                  opt = OptLoc (sp, x, y);
            }

            if (!opt) 
            {
                  if (IsWin (sp)) break;
            }
      }

      if (res) sp->nwf (sp, 1);
      else sp->nwf (sp, 0);
}

