package com.inif;

import java.sql.*;

/**
 * Created by r00xx on 2017/10/9.
 */
public class DbConn {

  public static Connection getUmsConnection() throws  Exception
  {
    //1: load mysql jdbc driver
    Class.forName("com.mysql.jdbc.Driver");


    //2: create connection instance for mysql
    Connection con = DriverManager.getConnection(url, username, password);

    return  con;
  }

  public static Connection getF2cConnection() throws Exception
  {
    Class.forName("com.mysql.jdbc.Driver");

    String url = "jdbc:mysql://xxxxx:3306/xxxxx";

    Connection con = DriverManager.getConnection(url, username, password);
    return  con;
  }




}
