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

```java

public void showKeyValue1(Generic<Number> obj){
    Log.d("泛型测试","key value is " + obj.getKey());
}

Generic<Integer> gInteger = new Generic<Integer>(123);
Generic<Number> gNumber = new Generic<Number>(456);

showKeyValue(gNumber);

// showKeyValue这个方法编译器会为我们报错：Generic<java.lang.Integer> 
// cannot be applied to Generic<java.lang.Number>
// showKeyValue(gInteger);

通过提示信息我们可以看到Generic<Integer>不能被看作为`Generic<Number>的子类。由此可以看出:同一种泛型可以对应多个版本（因为参数类型是不确定的），不同版本的泛型类实例是不兼容的。


// 正确的通配符写法
public void showKeyValue1(Generic<?> obj) {
  Log.d("范型测试", "key value is " + obj.getKey());
}

// 类型通配符一般是使用？代替具体的类型实参，注意了，此处’？’是类型实参，而不是类型形参 。重要说三遍！此处’？’是类型实参，而不是类型形参 ！
// 此处’？’是类型实参，而不是类型形参 ！再直白点的意思就是，此处的？和Number、String、Integer一样都是一种实际的类型，可以把？看成所有类型的父类。是一种真实的类型。

```  


### 范型方法

```java

public class GenericTest {
  // 这个类是范型类, 跟上面一样
  public class Generic<T> {
    private T key;

    public Generic (T key) {
      this.key = key;
    }

  //我想说的其实是这个，虽然在方法中使用了泛型，但是这并不是一个泛型方法。
  //这只是类中一个普通的成员方法，只不过他的返回值是在声明泛型类已经声明过的泛型。
  //所以在这个方法中才可以继续使用 T 这个泛型。 
    public T getkey() {
      return this.key
    } 

  }

  /** 
     * 这才是一个真正的泛型方法。
     * 首先在public与返回值之间的<T>必不可少，这表明这是一个泛型方法，并且声明了一个泛型T
     * 这个T可以出现在这个泛型方法的任意位置.
     * 泛型的数量也可以为任意多个 
     *    如：public <T,K> K showKeyName(Generic<T> container){
     *        ...
     *        }
     */

  // public 后面的  <T>  是表明该函数返回的类型是 T 类型， 从逻辑上面来看 就是不确定的类型， 实际上 <T> 也是一种类型
  public <T> T showKeyName (Generic<T> container) {
    System.out.println("container key: " + container.getKey());
    T test = container.getKey();
    return test;
  }

  //这也不是一个泛型方法，这就是一个普通的方法，只是使用了Generic<Number>这个泛型类做形参而已。
    public void showKeyValue1(Generic<Number> obj){
        Log.d("泛型测试","key value is " + obj.getKey());
    }

  //这也不是一个泛型方法，这也是一个普通的方法，只不过使用了泛型通配符?
    //同时这也印证了泛型通配符章节所描述的，?是一种类型实参，可以看做为Number等所有类的父类
    public void showKeyValue2(Generic<?> obj){
        Log.d("泛型测试","key value is " + obj.getKey());
    }

     /**
     * 这个方法是有问题的，编译器会为我们提示错误信息："UnKnown class 'E' "
     * 虽然我们声明了<T>,也表明了这是一个可以处理泛型的类型的泛型方法。
     * 但是只声明了泛型类型T，并未声明泛型类型E，因此编译器并不知道该如何处理E这个类型。
     */
    public <T> T showKeyName(Generic<E> container){

    }  

    /**
     * 这个方法也是有问题的，编译器会为我们提示错误信息："UnKnown class 'T' "
     * 对于编译器来说T这个类型并未项目中声明过，因此编译也不知道该如何编译这个类。
     * 所以这也不是一个正确的泛型方法声明。
     */
    public void showkey(T genericObj){

    }

    public static void main(String[] args) {

    }

}

```




### golang的范型
```go 
package main 

import (
  "fmt"
)

func bubbleSort(array []int) {
  for i := 0; i < len(array); i++ {
    for j := 0; j < len(array) - i - 1; j++ {
      if array[j] > array[j+1] {
        // swap
        array[j], array[j+1] = array[j+1], array[j]
      } 
    }
  }
}

func main() {
  a1 := []int{3, 2, 6, 10, 7, 4, 6, 5}
  bubbleSort(a1)
  fmt.Println(a1)
}


/**

那么，我们如果希望这个bubbleSort能够同时支持float类型数据排序，或者是按照字符串的长度来排序应该怎么做呢？在其他的例如java语言中，我们可以将bubbleSort定义为支持泛型的排序，但是Go里面就不行了。为了达到这个目的，我们可以使用interface来实现相同的功能。

针对上面的排序问题，我们可以分析一下排序的步骤：

查看切片长度，以用来遍历元素(Len)；
比较切片中的两个元素(Less)；
根据比较的结果决定是否交换元素位置(Swap)。
到这里，或许你已经明白了，我们可以把上面的函数分解为一个支持任意类型的接口，任何其他类型的数据只要实现了这个接口，就可以用这个接口中的函数来排序了。

**/

// 实际的作用就是到时候， 定义一个 Sortable 的实例的时候， 比如定义了两个 Sortable 的实例a, b, 假如 a 是 int[] 类型， 那么可以定义属于 int[] 类型的属性和方法
// 假如 b  是属于  string[] 类型， 那么可以针对  string[] 类型来定义属于  string[] 类型的 属性和方法
type Sortable interface {
  Len() int
  Less(int, int) bool
  Swap(int, int)
}

func bubbleSort(array Sortable) {
  for i := 0; i < array.Len(); i++ {
    for j := 0; j < array.Len()-i-1; j++ {
      if array.Less(j+1, j) {
        array.Swap(j, j+1)
      }
    }
  }
}



// 实现接口的整型切片
// 这里的意思就是, 假如 InArr 是  []int 的意思
type InArr []int

func (array InArr) Len() int {
  return len(array) 
}

func (array InArr) Less(i int, j int) bool {
  return array[i] < array[j]
}

func (array InArr) Swap(i int, j int) {
  array[i], array[j] = array[j], array[i]
}



//实现接口的字符串，按照长度排序
// 这里的意思就是， 假如 InArr 是 []string 的意思
type StringArr []string

func (array StringArr) Len() int {
  return len(array)
}

func (array StringArr) Less(i int, j int) bool {
  return len(array[i]) < len(array[j])
}

func (array StringArr) Swap(i int, j int) {
  array[i], array[j] = array[j], array[i]
}

//测试
func main() {
  intArray1 := IntArr{3, 4, 2, 6, 10, 1}
  bubbleSort(intArray1)
  fmt.Println(intArray1)

  stringArray1 := StringArr{"hello", "i", "am", "go", "lang"}
  bubbleSort(stringArray1)
  fmt.Println(stringArray1)
}


```



### java装饰器

```java 


public interface Person {
  void doCoding(); 
}


public class Employee implements Person {
  @Override
  public void doCoding() {
    System.out.println("程序员加班写代码啊， 写代码， 终于写完了......")
  }
}


// ---------------------------------------------------------------




// abstract 关键字证明该类一定会被再继承
public abstract class Manager implements Person {
  // 装饰器增加功能
  public abstract void doCoding();
}

public class ManagerA extends Manager {
  private Person person;    // 给雇员升职

  public ManagerA (Person person) {
    super();
    this.person = person;
  }

  @Override
  public void doCoding() {
    doEarlyWork();
    person.doCoding();
  }


  public void doEarlyWork() {
    System.out.println("项目经理A做需求分析")
  }
}


// --------------------------------------------------

public class ManagerB extends Manager {
  private Person person;      // 给雇员升职

  public ManagerB(Person person) {
    super();
    this.person = person;
  }

  @Override
  public void doCoding() {
    person.doCoding();
    doEndWork();
  }

  public void doEndWork() {
    System.out.println("项目经理B 再做收尾工作")
  }
}



public class Client {
  public static void main(String[] args) {
    Person employee = new Employee();

    employee = new ManagerA(employee);    //赋予程序猿项目经理A职责
    employee = new ManagerB(employee);    //赋予程序猿项目经理B职责 

    employee.doCoding();
  }
}

//输出
//项目经理A做需求分析
// 程序员加班写程序啊，写程序，终于写完了。。。
// 项目经理B 在做收尾工作






```


















