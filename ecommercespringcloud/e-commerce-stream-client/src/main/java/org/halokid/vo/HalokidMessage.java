package org.halokid.vo;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class HalokidMessage {

  private Integer id;
  private String projectName;
  private String org;
  private String author;
  private String version;

  public static HalokidMessage defaultMessage() {
    return new HalokidMessage(
        1,
        "e-commerce-stream-client",
        "halokid.org",
        "HalokidStudy",
        "1.0"
    );
  }
}



