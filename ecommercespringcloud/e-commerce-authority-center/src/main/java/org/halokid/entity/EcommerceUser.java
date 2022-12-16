package org.halokid.entity;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;
import org.springframework.data.annotation.CreatedDate;
import org.springframework.data.annotation.LastModifiedDate;
import org.springframework.data.jpa.domain.support.AuditingEntityListener;

import javax.persistence.*;
import java.io.Serializable;
import java.util.Date;

@Data
@NoArgsConstructor
@AllArgsConstructor
@Entity
@EntityListeners(AuditingEntityListener.class)
@Table(name = "t_ecommerce_user")
public class EcommerceUser implements Serializable {

  @Id
  @GeneratedValue(strategy = GenerationType.IDENTITY)
  @Column(name = "id",nullable = false)
  private Long id;

  @Column(name= "username",nullable = false)
  private String username;

  /** MD5 密码 */
  @Column(name= "password", nullable = false)
  private String password;

  /** 额外的信息，json */
  @Column(name= "extra_info", nullable = false)
  private String extraInfo;

  @CreatedDate
  @Column(name = "create_time", nullable = false)
  private Date createTime;


  @LastModifiedDate
  @Column(name= "update_time",nullable = false)
  private Date updateTime;
}


