package org.halokid.util;

import com.alibaba.fastjson.JSON;
import org.halokid.constant.CommonConstant;
import org.halokid.vo.LoginUserInfo;
import io.jsonwebtoken.Claims;
import io.jsonwebtoken.Jws;
import io.jsonwebtoken.Jwts;
import sun.misc.BASE64Decoder;

import java.security.KeyFactory;
import java.security.PublicKey;
import java.security.spec.X509EncodedKeySpec;
import java.util.Calendar;

/**
 * <h1>JWT Token 解析工具类</h1>
 */
public class TokenParseUtil {

  /**
   * <h2> 从JWT Token中解析LoginUserInfo对象</h2>
   * @param token
   * @return
   * @throws Exception
   */
  public static LoginUserInfo parseUserInfoFromToken(String token)
      throws Exception{
    if(null == token){
      return null;
    }

    Jws<Claims> claimsJws = parseToken(token, getPublickey());
    Claims body = claimsJws.getBody();
    // 如果Token 已经过期了，返回null
    if(body.getExpiration().before(Calendar.getInstance().getTime())){
      return null;
    }

    // 返回Token 中保存的用户信息
    return JSON.parseObject(
        body.get(CommonConstant.JWT_USER_INFO_KEY).toString(),
        LoginUserInfo.class
    );
  }


  /**
   * <h2>通过公钥去解析 JWT Token</h2>
   * @param token
   * @param publicKey
   * @return
   */
  public static Jws<Claims> parseToken(String token,PublicKey publicKey){
    return Jwts.parser().setSigningKey(publicKey).parseClaimsJws(token);
  }

  /**
   * <h2>根据本地存储的公钥获取到PublicKey 对象</h2>
   * @return
   * @throws Exception
   */
  private static PublicKey getPublickey()throws
      Exception{
    X509EncodedKeySpec keySpec = new X509EncodedKeySpec(
        new BASE64Decoder().decodeBuffer(CommonConstant.PUBLIC_KEY)
    );
    return KeyFactory.getInstance("RSA").generatePublic(keySpec);
  }
}



