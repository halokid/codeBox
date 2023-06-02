package org.halokid;

public class Main {

  public class ListNode {
    int val;
    ListNode next;

    ListNode() {
    }

    ListNode(int val) {
      this.val = val;
    }

    ListNode(int val, ListNode next) {
      this.val = val;
      this.next = next;
    }
  }

  ListNode mergeTwoLists(ListNode l1, ListNode l2) {
    // 虚拟头结点
    ListNode dummy = new ListNode(-1), p = dummy;
    ListNode p1 = l1, p2 = l2;

    while (p1 != null && p2 != null) {
      // compare `p1` and `p2` points
      // use the smaller val put in the `p` point
      if (p1.val > p2.val) {
        p.next = p2;
        p2 = p2.next;
      } else {
        p.next = p1;
        p1 = p1.next;
      }
      // `p` point keep moving forward
      p = p.next;
    }

    if (p1 != null) {
      p.next = p1;
    }

    if (p2 != null) {
      p.next = p2;
    }

    return dummy.next;
  }
}









