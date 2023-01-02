package org.halokid.entity;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;
import org.springframework.data.annotation.CreatedDate;
import org.springframework.data.annotation.LastModifiedDate;
import org.springframework.data.jpa.domain.support.AuditingEntityListener;

import javax.persistence.*;
import java.util.Date;

/**
 */
@Data
@NoArgsConstructor
@AllArgsConstructor
@Entity
@EntityListeners(AuditingEntityListener.class)
@Table(name="t_ecommerce_balance")
public class EcommerceBalance {

  /**
   * 自增主键
   */
  @Id
  @GeneratedValue(strategy = GenerationType.IDENTITY)
  @Column(name = "id",nullable = false)
  private Long id;

  /**
   * 用户id
   */
  @Column(name="user_id",nullable = false)
  private Long userId;

  @Column(name="balance",nullable = false)
  private Long balance;

  @CreatedDate
  @Column(name = "create_time",nullable = false)
  private Date createTime;

  @LastModifiedDate
  @Column(name = "update_time",nullable = false)
  private Date updateTime;
}


