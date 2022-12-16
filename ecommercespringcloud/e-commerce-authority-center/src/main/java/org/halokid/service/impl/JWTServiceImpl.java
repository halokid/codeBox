package org.halokid.service.impl;

import com.alibaba.fastjson.JSON;
import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.SignatureAlgorithm;
import lombok.extern.slf4j.Slf4j;
import org.halokid.constant.AuthorityConstant;
import org.halokid.constant.CommonConstant;
import org.halokid.dao.EcommerceUserDao;
import org.halokid.entity.EcommerceUser;
import org.halokid.service.IJWTService;
import org.halokid.vo.LoginUserInfo;
import org.halokid.vo.UsernameAndPassword;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;
import sun.misc.BASE64Decoder;

import java.security.KeyFactory;
import java.security.PrivateKey;
import java.security.spec.PKCS8EncodedKeySpec;
import java.time.LocalDate;
import java.time.ZoneId;
import java.time.ZonedDateTime;
import java.time.temporal.ChronoUnit;
import java.util.Date;
import java.util.UUID;

@Slf4j
@Service
@Transactional
public class JWTServiceImpl implements IJWTService {

  private EcommerceUserDao ecommerceUserDao;

  @Override
  public String generateToken(String username, String password) throws Exception {
    return null;
  }

  @Override
  public String generateToken(String username, String password, int expire) throws Exception {
    EcommerceUser ecommerceUser = ecommerceUserDao.findByUsernameAndPassword(username, password);
    if (null == ecommerceUser) {
      log.error("can not find user: {}, {}", username, password);
      return null;
    }

    LoginUserInfo loginUserInfo = new LoginUserInfo(ecommerceUser.getId(), ecommerceUser.getUsername());

    if (expire <= 0) {
      expire = AuthorityConstant.DEFAULD_EXPIRE_DAY;
    }

    ZonedDateTime zdt = LocalDate.now().plus(expire, ChronoUnit.DAYS).atStartOfDay(ZoneId.systemDefault());
    Date expirDate = Date.from(zdt.toInstant());
    return Jwts.builder()
        .claim(CommonConstant.JWT_USER_INFO_KEY, JSON.toJSONString(loginUserInfo))
        .setId(UUID.randomUUID().toString())
        .setExpiration(expirDate)
        .signWith(getPrivateKet(), SignatureAlgorithm.RS256)
        .compact();
  }

  @Override
  public String registerUserAndGenerateToken(UsernameAndPassword usernameAndPassword) {
    return null;
  }

  private PrivateKey getPrivateKet() throws Exception {
    PKCS8EncodedKeySpec periPKCS8 = new PKCS8EncodedKeySpec(
        new BASE64Decoder().decodeBuffer(AuthorityConstant.PRIVATE_KEY)
    );
    KeyFactory keyFactory = KeyFactory.getInstance("RSA");
    return keyFactory.generatePrivate(periPKCS8);
  }
}



