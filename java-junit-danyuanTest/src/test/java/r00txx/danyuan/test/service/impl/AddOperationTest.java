package r00txx.danyuan.test.service.impl;

import static org.junit.Assert.*;

/**
 * Created by r00xx on 2017/9/8.
 */
public class AddOperationTest {

  @org.junit.Before
  public void setUp() throws Exception {

  }

  @org.junit.After
  public void tearDown() throws Exception {

  }

  @org.junit.Test
  public void testAdd() throws Exception {
    System.out.println("add");

    int x = 0;
    int y = 0;

    AddOperation instance = new AddOperation();
    int expResult = 0;
    int result = instance.add(x, y);
    assertEquals(expResult, result);
  }
}



