import com.sun.tools.javac.code.Attribute;

import javax.annotation.PostConstruct;
import javax.annotation.Resource;
import java.lang.annotation.ElementType;
import java.lang.annotation.Retention;
import java.lang.annotation.RetentionPolicy;
import java.lang.annotation.Target;
import java.lang.reflect.Field;

public class NoteComment {
  public static void main(String[] args) throws Exception {

  }

  void check(Person person) throws Exception {
    for (Field field: person.getClass().getFields()) {
      Range range = field.getAnnotation(Range.class);
      if (range != null) {
        Object value = field.get(person);
        if (value instanceof String) {
          String s = (String) value;
          if ( s.length() < range.min() || s.length() > range.max() ) {
            throw new IllegalArgumentException("Invalid field: " + field.getName());
          }
        }
      }
    }
  }
}

@Retention(RetentionPolicy.RUNTIME)
@Target(ElementType.FIELD)
@interface Range {
  int min() default 0;
  int max() default 255;
}

class UseAnnon {
  @Range(min = 1, max = 20)
  public String name;

  @Range(max = 10)
  public String city;
}




