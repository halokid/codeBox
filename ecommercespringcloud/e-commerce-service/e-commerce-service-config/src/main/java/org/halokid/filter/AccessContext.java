package org.halokid.filter;

import org.halokid.vo.LoginUserInfo;

/**
 * <h1>使用 ThreadLocal 去单独存储每一个线程携带的LoginUserInfo 信息</h1>
 * 要及时的清理我们保存到ThreadLocal 中的用户信息：
 * 1.保证没有资源泄漏
 * 2.保证线程 在重用时，不会出现数据混乱
 */
public class AccessContext {
  private static final ThreadLocal<LoginUserInfo> loginUserInfo = new ThreadLocal<>();

  public static LoginUserInfo getLoginUserInfo() {
    return loginUserInfo.get();
  }

  public static void setLoginUserInfo(LoginUserInfo loginUserInfo_) {
    loginUserInfo.set(loginUserInfo_);
  }

  public static void clearLoginUserInfo() {
    loginUserInfo.remove();
  }
}



