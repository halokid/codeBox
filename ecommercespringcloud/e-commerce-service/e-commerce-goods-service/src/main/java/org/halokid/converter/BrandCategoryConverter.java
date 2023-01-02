package org.halokid.converter;

import org.halokid.constant.BrandCategory;

import javax.persistence.AttributeConverter;

/**
 */
public class BrandCategoryConverter implements AttributeConverter<BrandCategory,String> {
  @Override
  public String convertToDatabaseColumn(BrandCategory brandCategory) {
    return brandCategory.getCode();
  }

  @Override
  public BrandCategory convertToEntityAttribute(String s) {
    return BrandCategory.of(s);
  }
}