package org.halokid.dao;

import org.halokid.entity.EcommerceBalance;
import org.springframework.data.jpa.repository.JpaRepository;

/**
 * <h1>EcommerceBalance Dao 接口定义</h1>
 */
public interface EcommerceBalanceDao extends JpaRepository<EcommerceBalance,Long> {
  /** 根据userId 查询EcommerceBalance 对象*/
  EcommerceBalance findByUserId(Long userId);
}
