package hello;

/**
 * Created by r00xx<82049406@qq.com> on 2016/12/28.
 */
public class Customer {
  private long id;
  private String firstName, lastName;

  public Customer(long id, String firstName, String lastName) {
    this.id = id;
    this.firstName = firstName;
    this.lastName = lastName;
  }

  @Override
  public String toString() {
    return String.format("Customer[id=%d, firstName='%s', lastName='%s'", id, firstName, lastName);
  }


}
