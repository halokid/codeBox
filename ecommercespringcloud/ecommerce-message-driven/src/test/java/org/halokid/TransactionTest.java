package org.halokid;

import org.halokid.transactional.JpaSpringBootUser;
import org.halokid.transactional.SpringBootUserRepository;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.annotation.Rollback;
import org.springframework.test.context.junit4.SpringRunner;
import org.springframework.transaction.annotation.Transactional;

@RunWith(SpringRunner.class)
@SpringBootTest(classes = Application.class)
public class TransactionTest {

  @Autowired
  private SpringBootUserRepository springBootUserRepository;

  /**
   * <h2>测试保存数据表记录的事务问题</h2>
   * 1. 只有 @Test, 不会回滚
   * 2. 加上 @Transactional, 会回滚
   * 3. 如果已经有了 @Transactional 注解在类上面, 但是还是想要不回滚某个单测, 加上 @Rollback(value = false)
   * */
  @Test
  @Transactional
  @Rollback(value = false)
  public void testCreateSpringBootUser() {

    JpaSpringBootUser springBootUser = new JpaSpringBootUser();
    springBootUser.setUsername("Halokid");
    springBootUserRepository.save(springBootUser);
  }
}


