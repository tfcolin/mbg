#include "sp.h"

#include <curses.h>
#include <stdlib.h>
#include <math.h>

#include "sp_cs.h"

int nx, ny;
WINDOW ** win ;
WINDOW * msgwin;
int cx, cy;

int  (slf)  (SPDriver * sp, int * x, int * y, int * dir)
{
      int c;
      *dir = -1;
      wclear (msgwin);
      wprintw (msgwin, "Step: %d", sp->step);

      int set = 0;
      while (c = wgetch(stdscr))
      {
            int ocx = cx;
            int ocy = cy;

            switch (c)
            {
            case 'q': return -1; 
            case 'j': case KEY_DOWN: ++ cy; break;
            case 'k': case KEY_UP: -- cy; break;
            case 'h': case KEY_LEFT: -- cx; break;
            case 'l': case KEY_RIGHT: ++ cx; break;
            case ' ': set = 1; break;
            }
            if (cx < 0) cx = 0;
            if (cx >= sp->nx) cx = sp->nx - 1;
            if (cy < 0) cy = 0;
            if (cy >= sp->ny) cy = sp->ny - 1;

            if (cx != ocx || cy != ocy)
            {
                  redraw_win (sp, ocx, ocy);
                  redraw_win (sp, cx, cy);
            }

            if (set)
            {
                  *x = cx;
                  *y = cy;
                  return 0;
            }
      }
      return 0;
}

int  (csf)  (SPDriver * sp)
{
      int c;
      wclear (msgwin);
      wprintw (msgwin, "Start (Y/N) ?");

      while (c = wgetch(msgwin))
      {
            switch (c)
            {
            case 'y': case 'Y': return 0;
            case 'n': case 'N': return -1;
            }
      }

      return 0;
}

void (nwf) (SPDriver * sp, int res)
{
      wclear (msgwin);
      if (res) wprintw (msgwin, "Stop. Press any key to exit");
      else wprintw (msgwin, "Congratulations. Mission is complete. Press any key to exit");
      wgetch(msgwin);
}

void (ncf) (SPDriver * sp, int x, int y, int cx, int cy)
{ redraw_win (sp, x, y); }

void pexit (char * msg)
{
      cs_end();
      printf("%s\n", msg);
      exit (1);
}

void cs_init (int * nx, int * ny)
{
      int nx_, ny_;

      initscr();
      noecho();
      cbreak();
      curs_set(0);

      keypad (stdscr,TRUE);

      wrefresh (stdscr);

      getmaxyx (stdscr, ny_, nx_);

      *nx = nx_;
      *ny = ny_;
}

void cs_end ()
{
      endwin();
}

void redraw_win (SPDriver * sp, int x, int y)
{
      int lx, ly;
      WINDOW * w = win[y * sp->nx + x];

      getmaxyx (w, ly, lx);

      int i = (ly - 2) / 2 + 1;
      int j = (lx - 2) / 2 + 1;

      wclear (w);

      if (x == cx && y == cy) 
            wborder (w, '|', '|', '-', '-', '+', '+', '+', '+');

      int c = sp->loc[y * sp->nx + x];

      if (c >= 0) mvwprintw (w, i, j, "%d", c + 1);
      else mvwprintw (w, i, j, " ");

      wrefresh (w);
}

int main ()
{
      int i, j;
      int tnx, tny, lnx, lny;
      int bx0, by0, bx, by;

      cs_init (&tnx, &tny);

      lnx = 5;
      lny = 3;
      nx = 5;
      ny = 5;
      bx0 = (tnx - nx * lnx) / 2;
      by0 = (tny - 3 - ny * lny) / 2;

      if (tnx < 8 * lnx || tny < 8 * lny + 2) pexit ("Terminal screen two small.");

      msgwin = newwin (1, COLS - 2, LINES - 2, 1); 
      wclear (msgwin);
      wprintw (msgwin, "Set number of blocks (3x3, 4x4, 5x5, 6x6, 7x7):");
      char c = wgetch(msgwin);
      if (c > '7' || c < '3') c = '5';
      nx = ny = c - '0';

      win = malloc (sizeof(WINDOW *) * nx * ny);

      for (i = 0; i < ny; ++ i)
      {
            by = by0 + i * lny;
            for (j = 0; j < nx; ++ j)
            {
                  bx = bx0 + j * lnx;
                  win [i * nx + j] = newwin (lny, lnx, by, bx);
            }
      }

      cx = nx - 1;
      cy = ny - 1;

      SPDriver * sp = InitSP (nx, ny, nx - 1, ny - 1, slf, csf, nwf, ncf);
      RunSP (sp);
      FreeSP (sp);

      for (i = 0; i < nx * ny; ++ i) delwin (win[i]);
      delwin (msgwin);
      free (win);

      cs_end();
}

