package org.halokid.service;

import cn.hutool.core.codec.Base64;
import lombok.extern.slf4j.Slf4j;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit4.SpringRunner;

import java.security.KeyPair;
import java.security.KeyPairGenerator;
import java.security.interfaces.RSAPrivateKey;
import java.security.interfaces.RSAPublicKey;

@Slf4j
@SpringBootTest
@RunWith(SpringRunner.class)
public class RSATest {

  @Test
  public void generateKeyBytes() throws Exception {
    KeyPairGenerator keyPairGenerator = KeyPairGenerator.getInstance("RSA");
    keyPairGenerator.initialize(2048);

    //生成公钥和私钥对
    KeyPair keyPair = keyPairGenerator.genKeyPair();
    RSAPublicKey publicKey = (RSAPublicKey) keyPair.getPublic();
    RSAPrivateKey privateKey = (RSAPrivateKey) keyPair.getPrivate();

    log.info("private key: [{}]", Base64.encode(privateKey.getEncoded()));
    log.info("public key: [{}]", Base64.encode(publicKey.getEncoded()));

    // public key: [MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAveXftbH/uu/Fpu2qpjezPiFHzg/JnN9MLzXDSD8PRunfvocJqnSL9utsY5HT3DFQebCoVFDTlkILFuzSFNCrPBpoIywRaH/ruZjMMOMIX3XsHRh2ltofxMkpjULY5jJb077ZunyjRZ8k/U1jdmgADBtnM9Q7TQf1Het3WWrpQlqMk5RMlNnTuAxMW2/bTaEef+eAJLiIbES8+xzgNz1SXluPhQWxoZIcosDus8FqviNI3t5HTlv9nXyy7+AWYYqKhl6PZIw7TITW4fbfVIlOZjt6Ezcj1H/aXO4iPqmqj/66jbGTzGdnqpKiJ1bBigJXaa3/4j/8Bzfg9OoeUKvMhwIDAQAB]

    // private key: [MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC95d+1sf+678Wm7aqmN7M+IUfOD8mc30wvNcNIPw9G6d++hwmqdIv262xjkdPcMVB5sKhUUNOWQgsW7NIU0Ks8GmgjLBFof+u5mMww4whfdewdGHaW2h/EySmNQtjmMlvTvtm6fKNFnyT9TWN2aAAMG2cz1DtNB/Ud63dZaulCWoyTlEyU2dO4DExbb9tNoR5/54AkuIhsRLz7HOA3PVJeW4+FBbGhkhyiwO6zwWq+I0je3kdOW/2dfLLv4BZhioqGXo9kjDtMhNbh9t9UiU5mO3oTNyPUf9pc7iI+qaqP/rqNsZPMZ2eqkqInVsGKAldprf/iP/wHN+D06h5Qq8yHAgMBAAECggEBAIQE3iQ5XWqe6zCEJpFsQ+a7yGnPvACdlR91JxCBFqPcUDrrMqMYZ34AadhHN5zrg+E2GbCUiKT7wS3s8piaAZHFYgShzHB0DnN3cO3DyLxlBQOtUDTL7C4pFAQrJxPSitiI5GGr6O6sF1Eqji1xXCcOng8AS0HgeF6WpuC5XPVNPOhn+D01okEv+zGlGpX4oWmkbuT++4uf+wgsVOc+cTy6UVYVXwjpg3b5cgCXCzncduTncIm4m7xbVsOaQPSYnmc1HWpi9xEs+KBowyPK35u1sHkcBJC5//clVjw764MmPXejM4xIACoj4gr6bOfvslQPlQbkwoVk5yhxod0KSBECgYEA/JycW56qOeHYkI9lVhYVO06/5Bf9QZmfg9iaOL1R505zAEsrKc4JgU5GNcJbWxQ1Drln/+cCmGq+D/jITCXivCQbq2ZvxGeQdB8IXR4IDQ3k/IEj9ZpnoxBc7pElzHzaqKU2pH5UiCaWLmzVsk6tNCGpRk7BQtCoUfomM/rvHOUCgYEAwHHscRWSo7LWKEYq+QkBYNjz9ZTij0dFKXJMux8ltD63AIUBcIStbgG1544+so56sNnVDB6l1dv8UuSoZsElui4ODHYf9lwa3SbenxcKPjH/6pG5S04jlZvC5bQiEctvrATcuj6lsuy1IIWP3B8zrrd+izgUZe5/k2ZH11QhGPsCgYAefi0lph54RDnnIVgjLyE6+oORXg/1Jj0qJjBfnwLPABpDs88vCa7+C9vPy3lJokPRVjImotUeXuw4c1iCUKRw/47TftcJB1NZoRFaZDCmmsHEdTZJBG10Mpp1NdIxP/UJHIBfFNzRXo8MZZjaIEMqItITKZiMnzFiS+bmqWjJvQKBgGCn121kvGXjqJRJycOStlKCJOWvhasYzK/obIl/N1y0Mw3bv8FXohOiYE6QrwTXct/oSIzxx7EF8FFh7wLM97nu77dqjNsbn5J7t0yIMejTDAowbrRX5p95wAXG12XAm7H2LzaLyIV5Re+zpXGVakvSdYhO/k/D3eeWc3a/FnCJAoGBAPb6NGSbefn13XJn4hmbIgpQ37SyGutPGTGioN29GKxOqNdsvKZo6Zw8tPCm58X2/xJ58RfHsig4lRKZY74aCXgKKlb76sYZlfDV8JfKZEysjXMgqOlZK0htv3/BQs2pyo3z2zmnIOFKn5H2LEbZPNWZJUEGfew+/4Gl5nmAbuqO]

  }
}



