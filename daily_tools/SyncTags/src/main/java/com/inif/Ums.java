package com.inif;

import java.sql.*;
import java.util.ArrayList;
import java.util.List;

/**
 * Created by r00xx on 2017/10/9.
 */
public class Ums {

  public List<String> getProjects() {

    try {
      //3: 获取连接类实例con，用con创建Statement对象类实例 sql_statement
      DbConn syncTags = new DbConn();
      Connection con = syncTags.getUmsConnection();
      Statement sqlStatement = con.createStatement();

      //4: 执行查询，用ResultSet类的对象，返回查询的结果
      String query = "select LABEL from businesstree where BPARENTID=1";
      ResultSet result = sqlStatement.executeQuery(query);


      //对获得的查询结果进行处理，对Result类的对象进行操作
      List<String> projects = new ArrayList<String>();
      while (result.next()) {
        String proName = result.getString("LABEL");
        projects.add(proName);
//        System.out.println(proName);
      }

      sqlStatement.close();
      con.close();

      return projects;

    } catch (SQLException e) {
      System.out.println("mysql operation error");
      e.printStackTrace();
    } catch (Exception e) {
      e.printStackTrace();
    }

    return null;
  }


}
