package org.r00txx.inject;


import java.lang.annotation.*;


@Retention(RetentionPolicy.RUNTIME)
public @interface MyDiAnnotation {

  String value() default "";  //
}
