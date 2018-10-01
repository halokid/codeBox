### java泛型

- 下面的程序在编译阶段不会发现错误， 但是在运行时会出现错误

  ```
  java.lang.ClassCastException: java.lang.Integer cannot be cast to java.lang.String
  ```

  ```java
  List array_list = new ArrayList();
  array_list.add("aaaa");
  array_list.add(100);

  for (int i = 0; i < array_list.size(); i++) {
    String item = (String)array_list.get(i);
    Log.d("泛型测试", "item=" + item)
  }
  ```



- 我们要让程序在编译阶段就可以发现这个错误， 程序可以改写成

  ```java
  List<String> array_list = new ArrayList<String>();
  array_list.add("aaaa");
  array_list.add(100);

  for (int i = 0; i < array_list.size(); i++) {
    String item = (String)array_list.get(i);
    Log.d("泛型测试", "item=" + item)
  }
  ```

这样上面的程序就会在编译阶段就发现错误， 避免产生bug。

- 泛型的特性

  ```java
  List<String> string_array_list = new ArrayList<String>();
  List<Integer> integer_array_list = new ArrayList<Integer>();

  Class class_string_array_list = string_array_list.getClass();
  Calss class_integer_array_list = integer_array_list.getClass(); 

  if (class_string_array_list.equals(class_integer_array_list)) {
    Log.d("泛型测试", "输出结果为类型相同")
  }

  ```


**泛型类型在逻辑上看以看成是多个不同的类型，实际上都是相同的基本类型。**


### 泛型有三种使用方式，分别为：泛型类、泛型接口、泛型方法

- 泛型类

```java
# 一个普通的泛型类

//此处T可以随便写为任意标识，常见的如T、E、K、V等形式的参数常用于表示泛型
//在实例化泛型类时，必须指定T的具体类型

public class Generic<T> {
  //key这个成员变量的类型为T,T的类型由外部指定  
  private T key;

   //泛型构造方法形参key的类型也为T，T的类型由外部指定
  public Generic(T key) {
    this.key = key;
  }

  //泛型方法getKey的返回值类型为T，T的类型由外部指定
  public T getKey() {
    return this.key;
  }
}


// 实例化上面的类

// 整型
Generic<Integer> generic_integer = new Generic<Integer>(12345);
Log.d("泛型测试", "key is " + generic_integer.getKey());


// 字符型
Generic<String> generic_string = new Generic<String>("hello");
Log.d("泛型测试", "key is " + generic_string.getKey());


Generic generic = new Generic("111111");
Generic generic1 = new Generic(4444);
Generic generic2 = new Generic(55.55);
Generic generic3 = new Generic(false);

Log.d("泛型测试","key is " + generic.getKey());
Log.d("泛型测试","key is " + generic1.getKey());
Log.d("泛型测试","key is " + generic2.getKey());
Log.d("泛型测试","key is " + generic3.getKey());


D/泛型测试: key is 111111
D/泛型测试: key is 4444
D/泛型测试: key is 55.55
D/泛型测试: key is false

```
 

- 泛型接口

```java

# 定义一个泛型接口
public interface Generator<T> {
  public T next();
}


/**
 * 未传入泛型实参时，与泛型类的定义相同，在声明类的时候，需将泛型的声明也一起加到类中
 * 即：class FruitGenerator<T> implements Generator<T>{
 * 如果不声明泛型，如：class FruitGenerator implements Generator<T>，编译器会报错："Unknown class"
 */
class FruitGenerator<T> implements Generator<T>{
    @Override
    public T next() {
        return null;
    }
}


public class FruitGenerator implements Generator<String> {
  private String[] fruits = new String[]{"apple", "banana", "pear"};

  @Override
  public String next() {
    Random random = new Random();
    return fruits[random.nextInt(3)];
  }
}

```


- 泛型通配符












































