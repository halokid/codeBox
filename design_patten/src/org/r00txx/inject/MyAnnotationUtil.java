package org.r00txx.inject;



import java.lang.reflect.InvocationTargetException;
import java.lang.reflect.Method;

public class MyAnnotationUtil {

  MyAnnotationUtil(){
    di(this);
  }

  public static void di(Object obj){

    try {

      Method[] methods = obj.getClass().getDeclaredMethods();

      for (Method method : methods) {

        if(method.isAnnotationPresent(MyDiAnnotation.class)){//只处理包含MyDiAnnotation的方法

          MyDiAnnotation diAnnotation = method.getAnnotation(MyDiAnnotation.class);

          String class_key = diAnnotation.value();

          if(null==class_key || class_key.trim().length()==0){//key值默认为setXXX中的XXX且第一个字母要转换为小写
            class_key = method.getName().substring(3);
            class_key = class_key.substring(0, 1).toLowerCase() + class_key.substring(1);
          }

          method.invoke(obj, Class.forName(PropertiesUtil.getValue(class_key)).newInstance());

        }
      }
    } catch (SecurityException e) {
      e.printStackTrace();
    } catch (IllegalArgumentException e) {
      e.printStackTrace();
    } catch (IllegalAccessException e) {
      e.printStackTrace();
    } catch (InvocationTargetException e) {
      e.printStackTrace();
    } catch (InstantiationException e) {
      e.printStackTrace();
    } catch (ClassNotFoundException e) {
      e.printStackTrace();
    }
  }


