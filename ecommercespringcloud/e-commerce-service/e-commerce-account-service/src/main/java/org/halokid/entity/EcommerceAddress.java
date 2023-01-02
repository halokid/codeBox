package org.halokid.entity;

import org.halokid.account.AddressInfo;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;
import org.springframework.data.annotation.CreatedDate;
import org.springframework.data.annotation.LastModifiedDate;
import org.springframework.data.jpa.domain.support.AuditingEntityListener;

import javax.persistence.*;
import java.util.Date;

/**
 * <h1>用户地址表实体类定义</h1>
 */
@Data
@NoArgsConstructor
@AllArgsConstructor
@Entity
@EntityListeners(AuditingEntityListener.class)
@Table(name="t_ecommerce_address")
public class EcommerceAddress {
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

  @Column(name = "username",nullable = false)
  private String username;

  @Column(name = "phone",nullable = false )
  private String phone;

  @Column(name = "province",nullable = false)
  private String province;

  @Column(name="city",nullable = false)
  private String city;

  @Column(name = "address_detail",nullable = false)
  private String addressDetail;

  @CreatedDate
  @Column(name = "create_time",nullable = false)
  private Date createTime;

  @LastModifiedDate
  @Column(name = "update_time",nullable = false)
  private Date updateTime;

  /**
   * <h2>根据userId +AddressItem 得到EcommerceAddress</h2>
   * @return
   */
  public static EcommerceAddress to(Long userId, AddressInfo.AddressItem addressItem){
    EcommerceAddress ecommerceAddress = new EcommerceAddress();
    ecommerceAddress.setUserId(userId);
    ecommerceAddress.setUsername(addressItem.getUsername());
    ecommerceAddress.setPhone(addressItem.getPhone());
    ecommerceAddress.setProvince(addressItem.getProvince());
    ecommerceAddress.setCity(addressItem.getCity());
    ecommerceAddress.setAddressDetail(addressItem.getAddressDetail());

    return ecommerceAddress;
  }

  /**
   * <h2>将 EcommerceAddress 对象转成AddressInfo</h2>
   * @return
   */
  public AddressInfo.AddressItem toAddressItem(){
    AddressInfo.AddressItem addressItem = new AddressInfo.AddressItem();

    addressItem.setId(this.id);
    addressItem.setUsername(this.username);
    addressItem.setPhone(this.phone);
    addressItem.setProvince(this.province);
    addressItem.setCity(this.city);
    addressItem.setAddressDetail(this.addressDetail);
    addressItem.setCreateTime(this.createTime);
    addressItem.setUpdateTime(this.updateTime);
    return addressItem;
  }
}