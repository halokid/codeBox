package org.halokid.transactional;

import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.io.IOException;

/**
 * <h1>事务失效的场景</h1>
 * */
@Slf4j
@Service
public class TransactionalLose {

  /** springboot_user dao */
  private final SpringBootUserRepository springBootUserRepository;

  public TransactionalLose(SpringBootUserRepository springBootUserRepository) {
    this.springBootUserRepository = springBootUserRepository;
  }

  //    @Transactional
  @Transactional(rollbackFor = Exception.class)
  public void wrongRollbackFor() throws Exception {

    JpaSpringBootUser springBootUser = new JpaSpringBootUser();
    springBootUser.setUsername("ImoocQinyi-2021-01-01");
    springBootUserRepository.save(springBootUser);

    // ....
    // 由于某种原因抛出了异常
    throw new IOException("throw io exception for check rollback");
  }

  /**
   * <h2>同一个类中的方法调用</h2>
   * */
  public void wrongInnerCall() throws Exception {
    this.wrongRollbackFor();
  }

  @Transactional(rollbackFor = Exception.class)
  public void wrongTryCatch() {

    try {
      JpaSpringBootUser springBootUser = new JpaSpringBootUser();
      springBootUser.setUsername("ImoocQinyi-2021-01-01");
      springBootUserRepository.save(springBootUser);

      // ....
      // 由于某种原因抛出了异常
      throw new IOException("throw io exception for check rollback");
    } catch (Exception ex) {
      log.error("has some error: [{}]", ex.getMessage(), ex);
    }
  }
}



