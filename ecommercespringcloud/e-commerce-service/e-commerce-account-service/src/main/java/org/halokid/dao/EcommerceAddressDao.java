package org.halokid.dao;

import org.halokid.entity.EcommerceAddress;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.List;

public interface EcommerceAddressDao extends JpaRepository<EcommerceAddress, Long> {

  /**
   * 根据用户id查询地址信息
   */
  List<EcommerceAddress> findAllByUserId(Long userId);
}


