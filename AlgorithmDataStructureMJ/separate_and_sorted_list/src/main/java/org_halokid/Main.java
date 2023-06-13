package org_halokid;

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

  ListNode partition(ListNode head, int x) {
    // put head of less than `x` list
    ListNode dummy1 = new ListNode(-1);

    // put head of greater than `x` list
    ListNode dummy2 = new ListNode(-1);

    ListNode p1 = dummy1, p2 = dummy2;

    ListNode p = head;
    while (p != null) {
      if (p.val >= x) {
        p2.next = p;
        p2 = p2.next;
      } else {
        p1.next = p;
        p1 = p1.next;
      }

      // TODO: if `p.next` is null, do need to continue loop
      ListNode temp = p.next;
      p.next = null;
      p = temp;
    }
    // link two lists
    p1.next = dummy2.next;

    return dummy1.next;
  }
}







