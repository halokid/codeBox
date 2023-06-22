package org.halokid;

import java.text.SimpleDateFormat;
import java.util.Date;

public class TimeTool {
  private static final SimpleDateFormat fmt = new SimpleDateFormat("HH:mm:sss.SSS");

  public interface Task {
    void execute();
  }

  public static void check(String title, Task task) {
    if (task == null) return;
    title = (title == null) ? "" : ("【"+ title + "】");
    System.out.println(title);
    System.out.println("Start -->>> " + fmt.format(new Date()));

    long begin = System.currentTimeMillis();
    task.execute();
    long end = System.currentTimeMillis();
    System.out.println("Done -->>> " + fmt.format(new Date()));
    double delta = (end - begin) / 1000.0;
    System.out.println("Elapsed -->>> " + delta + " seconds");
    System.out.println("----------------------------------------------");
  }
}



