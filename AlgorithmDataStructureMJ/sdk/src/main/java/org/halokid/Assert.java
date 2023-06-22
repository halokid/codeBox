package org.halokid;

public class Assert {
  public static void test(boolean value) {
    try {
      if (!value) throw new Exception("TEST no pass");
    } catch (Exception e) {
      e.printStackTrace();
    }
  }
}
