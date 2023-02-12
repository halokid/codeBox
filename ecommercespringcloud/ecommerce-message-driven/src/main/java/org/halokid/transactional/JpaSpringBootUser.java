package org.halokid.transactional;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import javax.persistence.*;

@Data
@NoArgsConstructor
@AllArgsConstructor
@Entity
@Table(name = "springboot_user")
public class JpaSpringBootUser {

  // 主键 id
  @Id
  @GeneratedValue(strategy = GenerationType.IDENTITY)
  @Column(name = "id", nullable = false)
  private Integer id;

  // 用户名
  @Basic
  @Column(name = "user_name", nullable = false)
  private String username;
}






