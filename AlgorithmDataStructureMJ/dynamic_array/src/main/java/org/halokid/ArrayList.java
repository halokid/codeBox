package org.halokid;

import com.sun.xml.internal.ws.server.sei.SEIInvokerTube;

@SuppressWarnings("unchecked")      // clear the sytnx alarm
public class ArrayList<E> {

  private int size;

  private E[] elements;

  private static final int DEFAULT_CAPACITY = 10;
  private static final int ELEMENT_NOT_FOUND = -1;

  public ArrayList(int capacity) {
    capacity = (capacity < DEFAULT_CAPACITY) ? DEFAULT_CAPACITY : capacity;
    elements = (E[]) new Object[capacity];
  }

  public ArrayList() {
//    elements = new int[DEFAULT_CAPACITY];
//    TODO: above code can become below, make it more graceful
//    TODO: below code it means call the construct function use argument
    this(DEFAULT_CAPACITY);
  }

  public void clear() {
    for (int i = 0; i < size; i++) {
      elements[i] = null;
    }
    size = 0;
  }

  public int size() {
    return size;
  }

  public boolean isEmpty() {
    return size == 0;
  }

  public boolean contains(E element) {
    return indexOf(element) != ELEMENT_NOT_FOUND;
  }

  public void add(E element) {
//    elements[size] = element;
//    size++;
    add(size, element);
  }

  // add element in specify `index` position, `index` is count from `0`
  public void add(int index, E element) {
//    rangeCheckForAdd(index);

    ensureCapacity(size + 1);

    for (int i = size - 1; i >= index; i--) {
      elements[i + 1] = elements[i];
    }
    elements[index] = element;
    size++;
  }

  // make sure enough capacity for add
  private void ensureCapacity(int capacity) {
    int oldCapacity = elements.length;
    if (oldCapacity >= capacity) return;

    // new capacity is old capacity times
    // right move variable `1` places in binary means some `int` divide by `2`, so `1 + 0.5` equals `1.5`
    int newCapacity = oldCapacity + (oldCapacity >> 1);
    E[] newElements = (E[]) new Object[newCapacity];
    for (int i = 0; i < size; i++) {
      newElements[i] = elements[i];
    }
    elements = newElements;

    System.out.println(oldCapacity + " capacity extends to " + newCapacity);
  }

  public E get(int index) {
    // TODO: throw exception is better way in coding
    if (index < 0 || index >= size) {
      throw new IndexOutOfBoundsException("Index: " + index + ", Size: " + size);
    }
    return elements[index];
  }

  public E set(int index, E element) {
    if (index < 0 || index > size) {
      throw new IndexOutOfBoundsException("Index: " + index + ", Size: " + size);
    }

    E old = elements[index];
    elements[index] = element;
    return old;
  }

  public void rangeCheck(int index) {
    if (index > size) {
      throw new IndexOutOfBoundsException("index is bigger than size: " + size);
    }
  }

  public E remove(int index) {
    rangeCheck(index);

    E old = elements[index];
    for (int i = index + 1; i <= size - 1; i++) {
      elements[i - 1] = elements[i];
    }
    size--;

    elements[size] = null;

    return old;
  }

  public int indexOf(E element) {
    if (element == null) {
      for (int i = 0; i < size; i++) {
        if (elements[i] == null) return i;
      }

    } else {
      for (int i = 0; i < size; i++) {
//      if (elements[i] == element) return i;
        if (elements[i].equals(element)) return i;
      }
    }
    return ELEMENT_NOT_FOUND;
  }

  public String toString() {
//    return "xxx";
    StringBuilder output = new StringBuilder();
    output.append("Size = " + this.size);
    output.append(", [");

    for (int i = 0; i < size; i++) {
      if (i == size - 1) {
        output.append(this.get(i));
      } else {
        output.append(this.get(i)).append(", ");
      }
    }

    output.append("]");
    return output.toString();
  }
}










