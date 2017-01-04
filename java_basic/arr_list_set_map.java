

package org.r00txx.tese;


import com.sun.javafx.collections.MappingChange;

import java.awt.*;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.Map;
import java.util.Set;

public class ConvertorTest {

  public static void main (String[] args) {

    testList2Array();
    testArra2List();

    testSet2List();
    testList2Set();

    testSet2Array();
    testArray2Set();

    testMap2Set();
    testMap2List();
  }



  private static void testMap2List() {
    Map<String, String> map = new HashMap<String, String>();
    map.put("A", "ABC");
    map.put("K", "KK");
    map.put("L", "LV");

    //将 map key 转化为 list
    List<String> mapKeyList = new ArrayList<String>(map.keySet());
    System.out.println("map key list: " + mapKeyList);

    //将 map value 转化为  list
    List<String> mapValueList = new ArrayList<String>(map.values());
    System.out.println("map value list: " + mapValueList);

  }



  private static void testMap2Set () {

    Map<String, String> map = new HashMap<String, String>();
    map.put("A", "ABC");
    map.put("K", "KK");
    map.put("L", "LV");

    //讲 map 的键转化为 set
    Set<String> mapKeySet = map.keySet();


  }

}

































