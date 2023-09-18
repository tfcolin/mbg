#include <gtk/gtk.h>
#include "sp.h"

#define IMG_SIZE 420 

GtkWidget *window;
GtkWidget *box1;
GtkWidget *box2;
GtkWidget *box3;
GtkWidget *oriib;
GtkWidget **slide;
GtkWidget **slide_img;
GtkWidget *grid;

GtkWidget *label_size;
GtkWidget *spin_size;
GtkWidget *label_if;
GtkWidget *fcb_if;

GtkWidget *but_start;
GtkWidget *label_step;

GdkPixbuf *image;
GError * gerr;

typedef struct
{
      int x;
      int y;
} Point;

int quit; 
int st; // 0: 游戏结束 1: 游戏进行

int selx, sely;

int msgbox (char *msg, int but); // but=0: 一个按钮, 返回 0; but=1: 两个按钮(Yes/No) 返回 0: Yes, -1: No
void main_quit (GtkWidget *widget, gpointer  data);
void but_start_cb (GtkButton * button, gpointer data);
void but_slide_cb (GtkButton * button, gpointer data);

int  (slf)  (SPDriver * sp, int * x, int * y, int * dir); 
int  (csf)  (SPDriver * sp); // -1: 终止 0: 继续
void (nwf) (SPDriver * sp, int res); // res = 0: 成功, 1: 中止
void (ncf) (SPDriver * sp, int x, int y, int cx, int cy);

void (nwf) (SPDriver * sp, int res)
{
      if (quit) return;
      char msg [2048] = "";
      if (res == 0) 
      {
            sprintf (msg, "恭喜, 完成拼图. 共用 %d 步.", sp->step);
            msgbox (msg, 0);
      }
      st = 0;
      gtk_button_set_label ((GtkButton *)(but_start), "开始");
}

int  (csf)  (SPDriver * sp)
{
      if (quit) return -1;
      int res = msgbox ("确认开始游戏?", 1);
      if (res == 0) 
      {
            st = 1;
            gtk_button_set_label ((GtkButton *)(but_start), "中止");
      }
      else 
      {
            st = 0;
            gtk_button_set_label ((GtkButton *)(but_start), "开始");
      }
      return res;
}

int  (slf)  (SPDriver * sp, int * x, int * y, int * dir)
{
      if (quit) return -1;

      char msg[2048] = "";
      sprintf (msg, "Step: %d", sp->step);
      gtk_label_set_text ((GtkLabel *)(label_step), msg);
      
      gtk_main ();
      if (quit || st != 1) return -1;
      *x = selx;
      *y = sely;
      return 0;
}

void (ncf) (SPDriver * sp, int x, int y, int cx, int cy)
{
      if (quit) return;
      int ind = y * sp->nx + x;
      int x0, y0, nx, ny;
      
      if (cx == -1)
      {
            gtk_image_clear ((GtkImage *)(slide_img[ind]));
      }
      else
      {
            nx = IMG_SIZE / sp->nx;
            x0 = cx * nx;
            ny = IMG_SIZE / sp->ny;
            y0 = cy * ny;

            GdkPixbuf * subimage = gdk_pixbuf_new_subpixbuf (image, x0, y0, nx, ny);
            gtk_image_set_from_pixbuf ((GtkImage *)(slide_img[ind]), subimage);

            g_clear_object (&subimage);
      }
}

void main_quit (GtkWidget *widget, gpointer  data)
{
      quit = 1;
      int level = gtk_main_level();
      gtk_main_quit ();
}

int msgbox (char *msg, int but) // but=0: 一个按钮, 返回 0; but=1: 两个按钮(Yes/No) 返回 0: Yes, -1: No
{
      int res, dres;
      GtkButtonsType bt = but ? GTK_BUTTONS_YES_NO : GTK_BUTTONS_CLOSE;
      
      GtkWidget * adia = gtk_message_dialog_new (GTK_WINDOW(window), GTK_DIALOG_MODAL, GTK_MESSAGE_INFO, bt, msg);
      dres = gtk_dialog_run ((GtkDialog *) adia);
      if (but) res = dres == GTK_RESPONSE_YES ? 0 : -1;
      else res = 0;

      gtk_widget_destroy (adia);

      return res;
}

void but_start_cb (GtkButton * button, gpointer data)
{
      int i, j;

      SPDriver * sp;
      int nx, ny;

      if (st == 0)
      {
            GdkPixbuf * oimage;
            char * filename = gtk_file_chooser_get_filename (GTK_FILE_CHOOSER(fcb_if));
            if (filename == NULL)
            {
                  msgbox ("请选择背景图像.", 0);
                  return;
            }
            oimage = gdk_pixbuf_new_from_file (filename, &gerr);
            if (oimage == NULL)
            {
                  msgbox ("无法载入图像.", 0);
                  return;
            }

            image = gdk_pixbuf_scale_simple (oimage, IMG_SIZE, IMG_SIZE, GDK_INTERP_BILINEAR);  
            gtk_image_set_from_pixbuf ((GtkImage *)(oriib), image);

            nx = ny = gtk_spin_button_get_value_as_int ((GtkSpinButton *)(spin_size));
            grid = gtk_grid_new ();
            slide = malloc (sizeof(GtkWidget *) * nx * ny);
            slide_img = malloc (sizeof(GtkWidget *) * nx * ny);

            Point p[nx * ny]; 
            for (i = 0; i < ny; ++ i)
            {
                  for (j = 0; j < nx; ++ j)
                  {
                        int ind = i * nx + j;
                        p[ind].x = j; p[ind].y = i;
                        slide[ind] = gtk_button_new();
                        slide_img[ind] = gtk_image_new();
                        gtk_button_set_image ((GtkButton *)(slide[ind]), slide_img[ind]);
                        g_signal_connect (slide[ind], "clicked", G_CALLBACK(but_slide_cb), p + ind);
                        gtk_widget_set_size_request (slide_img[ind], IMG_SIZE / nx, IMG_SIZE / ny);
                        gtk_grid_attach (GTK_GRID(grid), slide[ind], j, i, 1, 1);
                  }
            }

            gtk_box_pack_end (GTK_BOX(box2), grid, TRUE, FALSE, 0);
            gtk_widget_set_halign (grid, GTK_ALIGN_CENTER);
            gtk_widget_set_valign (grid, GTK_ALIGN_CENTER);

            gtk_widget_show_all (grid);

            sp = InitSP (nx, ny, nx - 1, ny - 1, slf, csf, nwf, ncf);
            RunSP (sp);
            FreeSP (sp);

            if (!quit) gtk_widget_destroy (grid);

            free (slide_img);
            free (slide);

            if (!quit) gtk_image_clear ((GtkImage *)oriib);
            g_clear_object (&image);
            g_clear_object (&oimage);

            if (quit) 
            {
                  // g_print ("Level: %d\n", gtk_main_level());
                  gtk_main_quit();
            }
      }
      else 
      {
            st = 0;
            gtk_main_quit();
      }
}

void but_slide_cb (GtkButton * button, gpointer data)
{
      if (st == 0) return;
      Point * p = (Point *)data;
      selx = p->x;
      sely = p->y;
      gtk_main_quit();
}

int main (int   argc, char *argv[])
{

      gtk_init (&argc, &argv);

      window = gtk_window_new (GTK_WINDOW_TOPLEVEL);
      gtk_window_set_title (GTK_WINDOW (window), "Slide Puzzle");

      box1 = gtk_box_new (GTK_ORIENTATION_VERTICAL, 5);
      box2 = gtk_box_new (GTK_ORIENTATION_HORIZONTAL, 5);
      box3 = gtk_box_new (GTK_ORIENTATION_HORIZONTAL, 5);

      gtk_container_add (GTK_CONTAINER(window), box1);
      gtk_box_pack_start (GTK_BOX(box1), box2, TRUE, FALSE, 0);
      gtk_box_pack_start (GTK_BOX(box1), box3, FALSE, FALSE, 0);

      gtk_widget_set_halign (box2, GTK_ALIGN_CENTER);
      gtk_widget_set_valign (box2, GTK_ALIGN_CENTER);
      gtk_widget_set_halign (box3, GTK_ALIGN_CENTER);
      gtk_widget_set_valign (box3, GTK_ALIGN_CENTER);

      gtk_box_set_homogeneous (GTK_BOX(box2), TRUE);
      gtk_widget_set_size_request (box2, 900, 500);

      oriib = gtk_image_new ();
      gtk_box_pack_end (GTK_BOX(box2), oriib, TRUE, FALSE, 0);
      gtk_widget_set_halign (oriib, GTK_ALIGN_CENTER);
      gtk_widget_set_valign (oriib, GTK_ALIGN_CENTER);

      label_size = gtk_label_new ("格数");
      spin_size = gtk_spin_button_new_with_range (2, 7, 1);
      gtk_spin_button_set_value (GTK_SPIN_BUTTON(spin_size), 5);

      label_if = gtk_label_new ("图片");
      fcb_if = gtk_file_chooser_button_new ("选择图片", GTK_FILE_CHOOSER_ACTION_OPEN);

      but_start = gtk_button_new_with_label ("开始");
      
      label_step = gtk_label_new ("Step: 0");

      gtk_box_pack_start (GTK_BOX(box3), label_size, FALSE, FALSE, 0);
      gtk_box_pack_start (GTK_BOX(box3), spin_size,  FALSE, FALSE, 0);
      gtk_box_pack_start (GTK_BOX(box3), label_if,   FALSE, FALSE, 0);
      gtk_box_pack_start (GTK_BOX(box3), fcb_if,     FALSE, FALSE, 0);
      gtk_box_pack_start (GTK_BOX(box3), but_start,  FALSE, FALSE, 0);
      gtk_box_pack_start (GTK_BOX(box3), label_step, FALSE, FALSE, 0);

      g_signal_connect (window, "destroy", G_CALLBACK (main_quit), NULL);
      g_signal_connect (but_start, "clicked", G_CALLBACK (but_start_cb), NULL);

      gtk_widget_show_all (window);

      st = 0;
      quit = 0;

      gtk_main ();

      return 0;
}
