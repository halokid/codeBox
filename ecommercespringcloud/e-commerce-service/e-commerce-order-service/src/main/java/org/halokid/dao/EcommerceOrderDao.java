package org.halokid.dao;

import org.halokid.entity.EcommerceOrder;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.repository.PagingAndSortingRepository;

public interface EcommerceOrderDao extends PagingAndSortingRepository<EcommerceOrder, Long> {

  /**
   * <h2>Queries paging orders based on userids</h2>
   * select * from t_ecommerce_order where user_id = ? order by ... desc/asc limit x offset y
   */
  Page<EcommerceOrder> findAllByUserId(Long userId, Pageable pageable);
}
