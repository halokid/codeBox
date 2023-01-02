package org.halokid.constant;

import lombok.AllArgsConstructor;
import lombok.Getter;

import java.util.Objects;
import java.util.stream.Stream;

/**
 * <h1>商品类别</h1>
 * 电器-> 手机、电脑
 */

@Getter
@AllArgsConstructor
public enum GoodsCategory {

  DIAN_QI("10001", "电器"),
  JIA_JU("10002", "家具"),
  FU_SHI("10003", "服饰"),
  MU_YIN("10004", "母婴"),
  SHI_PIN("10005", "食品"),
  TU_SHU("10006", "图书"),
  ;

  /**
   * 商品分类编码
   */
  private final String code;
  /**
   * 商品分类描述信息
   */
  private final String description;

  /**
   * <h2>根据code 获取到GoodsCategory</h2>
   *
   * @param code
   * @return
   */
  public static GoodsCategory of(String code) {
    Objects.requireNonNull(code);
    return Stream.of(values())
        .filter(bean -> bean.code.equals(code))
        .findAny()
        .orElseThrow(
            () -> new IllegalArgumentException(code + "not exit")
        );
  }
}



