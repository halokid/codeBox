package org.r00txx.dao;


import java.sql.SQLException;
import java.util.List;
import java.util.Map;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Component;


/**
 * Created by r00xx<82049406@qq.com> on 2016/12/30.
 */

@Component
public class userDao {

  @Autowired
  private JdbcTemplate jdbcTemplate;

//  public void queryUsers() {
  public String queryUsers() {
    String tmp = null;
    System.out.println("query users");
    String sql = "select id, username, email from users where id=1";

    List<Map<String, Object>> list = jdbcTemplate.queryForList(sql);
    for (Map<String, Object> row : list) {
      tmp = (String) row.get("username");
      System.out.println(row.get("username"));
    }

    return tmp;
  }

}
