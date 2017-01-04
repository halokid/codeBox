package org.r00txx.inject;

import java.io.*;
import java.util.*;

public class PropertiesUtil {

  private static InputStream in = PropertiesUtil.class.getClassLoadwer()
                                  .getResourceAsStream(my.properties);

  public PropertiesUtil() {

  }


  private static PropertiesUtil props = new PropertiesUtil();

  static {
    try {
      props.load(in);
    }
    catch (FileNotFoundException e) {
      e.printStackTrace();
    }
    catch (IOException e) {
      e.printStackTrace();
    }
  }


  public static String getValue(String key) {
    return props.getProperty(key);
  }

  public static void updateProperties(String key, String value) {
    props.setProperty(key, value);
  }


}
