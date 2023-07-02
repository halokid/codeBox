package org.halokid;

public class Main {

  public static void main(String[] args) {
    java.util.ArrayList a = new java.util.ArrayList<>();

    ArrayList<Person> persons = new ArrayList<>();
    persons.add(new Person(10, "Jack"));
    persons.add(null);
    persons.add(new Person(12, "James"));
    persons.add(null);
    persons.add(new Person(15, "Rose"));
    System.out.println(persons);

    System.out.println(persons.indexOf(null));

//    persons.clear();

    // TODO: remind JVM to do `garbage collection`
//    System.gc();
  }

  public static void test(String[] args) {

    int array[] = new int[] {11, 22, 33};

    // TODO: in java, all the class inherit java.lang.Object

    // TODO: ArrayList support a `genernic` type, if want to use `int` element, need declare <E> is <integer>
//    ArrayList<Integer> list = new ArrayList();

    ArrayList<Integer> ints = new ArrayList<>();
    ints.add(10);
    ints.add(11);
    ints.add(22);
    ints.add(33);
    System.out.println(ints);

//    TODO: for test
//    list.get(0);

//    /*
//    list.add(99);
//    list.add(88);
//    list.add(77);
//    list.add(66);
//    list.add(55);
//
//    for (int i =0; i < 30; i++) {
//      list.add(i);
//    }
//
//    list.set(3, 80);
//    Assert.test(list.get(3) == 80);

//    System.out.println(list);
//     */
  }
}