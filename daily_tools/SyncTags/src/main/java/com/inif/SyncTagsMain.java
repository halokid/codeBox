package com.inif;


import java.util.ArrayList;
import java.util.List;

/**
 * Created by r00xx on 2017/10/9.
 */
public class SyncTagsMain {

  public static void  main(String[] args) {
    Ums ums = new Ums();
//    List<String> projects = new ArrayList<String>();
    List<String> projects  = ums.getProjects();
    System.out.println(projects);
  }

}
