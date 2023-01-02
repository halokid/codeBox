package org.halokid.service.impl;

import com.alibaba.fastjson.JSON;
import lombok.extern.slf4j.Slf4j;
import org.halokid.account.AddressInfo;
import org.halokid.common.TableId;
import org.halokid.dao.EcommerceAddressDao;
import org.halokid.entity.EcommerceAddress;
import org.halokid.filter.AccessContext;
import org.halokid.service.IAddressService;
import org.halokid.vo.LoginUserInfo;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.util.CollectionUtils;

import java.util.Collections;
import java.util.List;
import java.util.stream.Collectors;

@Slf4j
@Service
@Transactional(rollbackFor = Exception.class)   // todo: 事务处理
public class AddressServiceImpl implements IAddressService {

  private final EcommerceAddressDao addressDao;

  public AddressServiceImpl(EcommerceAddressDao addressDao) {
    this.addressDao = addressDao;
  }

  /**
   * <h2>存储多个地址信息</h2>
   * @param addressInfo
   * @return
   */
  @Override
  public TableId createAddressInfo(AddressInfo addressInfo) {
    // 不能直接从参数获取id 可能串改
    LoginUserInfo loginUserInfo = AccessContext.getLoginUserInfo();
    // 将传递的参数转换成实体对象
    List<EcommerceAddress> ecommerceAddressList = addressInfo.getAddressItems().stream()
        .map(a->EcommerceAddress.to(loginUserInfo.getId(),a))
        .collect(Collectors.toList());
    // 保存到数据表并把返回记录的 id 给调用方
    List<EcommerceAddress> savedRecords=addressDao.saveAll(ecommerceAddressList);
    List<Long> ids= savedRecords.stream()
        .map(EcommerceAddress::getId).collect(Collectors.toList());
    log.info("create address info: [{}],[{}]", loginUserInfo.getId(), JSON.toJSONString(ids));

    return new TableId(
        ids.stream().map(TableId.Id::new).collect(Collectors.toList())
    );
  }

  @Override
  public AddressInfo getCurrentAddressInfo() {
    LoginUserInfo loginUserInfo = AccessContext.getLoginUserInfo();
    // 根据userId 查询用户的地址信息, 再实现转换
    List<EcommerceAddress>ecommerceAddresses =addressDao.findAllByUserId(
        loginUserInfo.getId()
    );
    List<AddressInfo.AddressItem> addressItems = ecommerceAddresses.stream()
        .map(EcommerceAddress::toAddressItem)
        .collect(Collectors.toList());
    return new AddressInfo(loginUserInfo.getId(),addressItems);
  }

  @Override
  public AddressInfo getAddressInfoById(Long id) {

    EcommerceAddress ecommerceAddress = addressDao.findById(id).orElse(null);
    if(null==ecommerceAddress){
      throw new RuntimeException("address is not exist");
    }
    return new AddressInfo(
        ecommerceAddress.getUserId(),
        Collections.singletonList(ecommerceAddress.toAddressItem())
    );
  }

  @Override
  public AddressInfo getAddressInfoByTableId(TableId tableId) {

    List<Long> ids = tableId.getIds().stream()
        .map(TableId.Id::getId)
        .collect(Collectors.toList());

    log.info("get address info by table id:[{}]",JSON.toJSONString(ids));
    List<EcommerceAddress> ecommerceAddresses = addressDao.findAllById(ids);

    if(CollectionUtils.isEmpty(ecommerceAddresses)){
      return new AddressInfo(-1L,Collections.emptyList());
    }
    List<AddressInfo.AddressItem> addressItems = ecommerceAddresses.stream()
        .map(EcommerceAddress::toAddressItem)
        .collect(Collectors.toList());
    return new AddressInfo(
        AccessContext.getLoginUserInfo().getId(),
        addressItems);
  }
}



