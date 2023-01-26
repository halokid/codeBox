package org.halokid.pojo;

import java.time.LocalDate;

public class EmployeePojo {

  public String firstNaame;
  public String lastName;
  private LocalDate startDate;

  public EmployeePojo(String firstNaame, String lastName, LocalDate startDate) {
    this.firstNaame = firstNaame;
    this.lastName = lastName;
    this.startDate = startDate;
  }

  public String name() {
    return this.firstNaame + " " + this.lastName;
  }

  public LocalDate getStart() {
    return this.startDate;
  }
}



