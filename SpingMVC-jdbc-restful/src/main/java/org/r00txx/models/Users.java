package org.r00txx.models;

/**
 * Created by r00xx<82049406@qq.com> on 2016/12/30.
 */
public class Users {

  private int id;
  private String username;
  private String email;

  public Users(int id, String username, String email) {
    this.id = id;
    this.username = username;
    this.email = email;
  }

  public int getId() {
    return id;
  }

  public String getUsername() {
    return username;
  }

  public String getEmail() {
    return email;
  }

}
