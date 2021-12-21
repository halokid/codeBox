package org.r00txx.models;

import javax.persistence.*;
import javax.swing.text.StringContent;
import java.util.StringJoiner;

/**
 * Created by r00xx<82049406@qq.com> on 2016/12/29.
 */

@Entity
@Table(name = "movie")
public class Movie {

  @Id
  @GeneratedValue(strategy = GenerationType.AUTO)
  private long id;
  private String title;
  private int year;

 public Movie() {

 }

  public Movie(String title, int year) {
    this.title = title;
    this.year = year;

  }

  public long getId() {
    return id;
  }

  public void setId(long id) {
    this.id = id;
  }

  public String getTitle() {
    return title;
  }

  public void setTitle(String title) {
    this.title = title;
  }


  public int getYear() {
    return year;
  }

  public void setYear(int year) {
    this.year = year;
  }


  @Override
  public String toString() {
    return String.format("Movie[id=%d, title='%s', year='%d']", id, title, year);
  }

}







