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

  public JWTServiceImpl(EcommerceUserDao ecommerceUserDao){
    this.ecommerceUserDao=ecommerceUserDao;
  }

  @Override
  public String generateToken(String username, String password) throws Exception {
    return generateToken(username, password,0);
  }

  @Override
  public String generateToken(String username, String password, int expire) throws Exception {

    // 首先需要验证用户是否能够通过授权校验，即输入的用户名密码能否匹配数据表记录
    EcommerceUser ecommerceUser =ecommerceUserDao.findByUsernameAndPassword(username,password);
    if(null==ecommerceUser){
      log.error("can not find user: [{}],[{}]",  username, password);
      return null;
    }

    // Token 中塞入对象，即 JWT 中存储的信息，后端拿到这些信息就可以知道是哪个用户在操作
    LoginUserInfo loginUserInfo = new LoginUserInfo(
        ecommerceUser.getId(),ecommerceUser.getUsername()
    );

    if(expire <= 0){
      expire= AuthorityConstant.DEFAULD_EXPIRE_DAY;
    }

    // 计算超时时间
    ZonedDateTime zdt = LocalDate.now().plus(expire, ChronoUnit.DAYS)
        .atStartOfDay(ZoneId.systemDefault());
    Date expireDate =Date.from(zdt.toInstant());
    return Jwts.builder()
        // jwt payload -->kv
        .claim(CommonConstant.JWT_USER_INFO_KEY, JSON.toJSONString(loginUserInfo))
        //jwt id
        .setId(UUID.randomUUID().toString())
        //jwt 过期时间
        .setExpiration(expireDate)
        // jwt 签名--> 加密
        .signWith(getPrivateKey(), SignatureAlgorithm.RS256)
        .compact();
  }

  @Override
  public String registerUserAndGenerateToken(UsernameAndPassword usernameAndPassword) throws Exception {

    // 先去校验用户名是否存在，如果存在，不能重复注册
    EcommerceUser oldUser = ecommerceUserDao.findByUsername(usernameAndPassword.getUsername());
    if(null!=oldUser){
      log.error("username is registered: [{}]",usernameAndPassword.getUsername());
      return null;
    }

    EcommerceUser user = new EcommerceUser();
    user.setUsername(usernameAndPassword.getUsername());
    user.setPassword(usernameAndPassword.getPassword());
    user.setExtraInfo("{}");

    // 注册一个新用户
    user = ecommerceUserDao.save(user);
    log.info("register user success: [{}],[{}]",user.getUsername(),user.getId());
    // 生成token 并返回
    return generateToken(user.getUsername(),user.getPassword());
  }

  /**
   * <h2> 根据本地存储的私钥获取到PrivateKey 对象</h2>
   * @return
   * @throws Exception
   */
  private PrivateKey getPrivateKey()throws Exception{
    PKCS8EncodedKeySpec periPKCS8 = new PKCS8EncodedKeySpec(
        new BASE64Decoder().decodeBuffer(AuthorityConstant.PRIVATE_KEY)
    );
    KeyFactory keyFactory = KeyFactory.getInstance("RSA");
    return keyFactory.generatePrivate(periPKCS8);
  }
}




