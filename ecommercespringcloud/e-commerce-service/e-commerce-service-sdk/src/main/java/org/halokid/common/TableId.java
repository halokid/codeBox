package org.halokid.common;

import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.List;

/**
 * <h1>主键 ids</h1>
 */
@ApiModel(description = "通用 id 对象")
@Data
@NoArgsConstructor
@AllArgsConstructor
public class TableId {

  @ApiModelProperty(value = "数据表记录主键")
  private List<Id> ids;


  @ApiModel(description = "数据表记录主键对象")
  @Data
  @NoArgsConstructor
  @AllArgsConstructor
  public static class Id{
    @ApiModelProperty(value = "数据表记录主键")
    private Long id;
  }
}

