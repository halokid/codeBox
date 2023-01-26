package org.halokid.pojo;

import org.apache.commons.beanutils.PropertyUtils;

import java.beans.PropertyDescriptor;

/**
 * @Description reflect POJO and JavaBean
 */
public class ReflectionExample {
  public static void main(String[]args) {
    System.out.println("Fields for ExmployeePojo are:");
    for (PropertyDescriptor pd: PropertyUtils.getPropertyDescriptors(EmployeePojo.class)) {
      System.out.println(pd.getDisplayName());
    }

    System.out.println("Fields for ExmployeeBean are:");
    for (PropertyDescriptor pd: PropertyUtils.getPropertyDescriptors(EmployeeBean.class)) {
      System.out.println(pd.getDisplayName());
    }
  }
}
