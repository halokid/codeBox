
import java.util.*;

public class Collection {
  public static void main(String[] args) throws Exception {
    List<String> list = Arrays.asList("apple", "pear", "banana");
    for (int i = 0; i < list.size(); i++) {
      String s = list.get(i);
      System.out.println(s);
    }

    // -------------------------------------------------------
    System.out.println("-----------------------------------------");
    for (Iterator<String> it = list.iterator(); it.hasNext();) {
      String s = it.next();
      System.out.println(s);
    }

    System.out.println("-----------------------------------------");
    for (String s : list) {
      System.out.println(s);
    }

    System.out.println("-----------------------------------------");
    List<String> listx = Arrays.asList("A", "B", "C");
    System.out.println(listx.contains(new String("C")));
    System.out.println(listx.indexOf(new String("C")));
    System.out.println(listx.indexOf("C"));

    System.out.println("-----------------------------------------");
    Queue<String> q = new PriorityQueue<>();
    // 添加3个元素到队列:
    q.offer("apple");
    q.offer("pear");
    q.offer("banana");
    System.out.println(q.poll()); // apple
    System.out.println(q.poll()); // banana
    System.out.println(q.poll()); // pear
    System.out.println(q.poll()); // null,因为队列为空

    System.out.println("-----------------------------------------");
    Deque<String> deque = new LinkedList<>();
//    deque.offerLast("A"); // A
//    deque.offerLast("B"); // A <- B
//    deque.offerFirst("C"); // C <- A <- B
//    System.out.println(deque.pollFirst()); // C, 剩下A <- B
//    System.out.println(deque.pollLast()); // B, 剩下A
//    System.out.println(deque.pollFirst()); // A
//    System.out.println(deque.pollFirst()); // null

    // todo: stack datastructure
    deque.push("A"); // A
    deque.push("B"); // A
    deque.push("C"); // A
    System.out.println(deque.pop()); // null
    System.out.println(deque.pop()); // null
    System.out.println(deque.pop()); // null

    System.out.println("-----------------------------------------");
    List<String> listy = new ArrayList<>();
    listy.add("apple");
    listy.add("pear");
    listy.add("orange");
    // 排序前:
    System.out.println(listy);
    Collections.sort(listy);
    // 排序后:
    System.out.println(listy);
  }
}





