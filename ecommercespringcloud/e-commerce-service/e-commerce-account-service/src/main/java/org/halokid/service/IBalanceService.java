package org.halokid.service;

import org.halokid.account.BalanceInfo;

/**
 * <h1>用于余额相关的服务接口定义</h1>
 */
public interface IBalanceService  {

  /**
   * <h2>获取当前用户余额信息</h2>
   * @return
   */
  BalanceInfo getCurrentUserBalanceInfo();

  /**
   * <h2>扣减用户余额</h2>
   * @param balanceInfo  代表想要扣减的余额
   * @return
   */
  BalanceInfo deductBalance(BalanceInfo balanceInfo);
}
