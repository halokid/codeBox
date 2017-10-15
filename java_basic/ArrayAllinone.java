/*
 * @Author: r00x.tactx 
 * @Date: 2017-10-15 11:48:46 
 * @Last Modified by: xx.tactx
 * @Last Modified time: 2017-10-15 15:31:47

   Top 10 Methods for Java Arrays
 */


// package org.r00txx.tese;

import java.util.Arrays;

public class ArrayAllinone {
  
  public static void main (String[] args) {
    defileArray();  

    outputArray();
  }

  //define a array 
  public static void defileArray() {
    String[] aArray = new String[5];

    String[] bArray = {"a", "b", "c", "d", "e"};

    String[] cArray = new String[]{"a", "b", "c", "d", "e"};

    System.out.println("------------- defined array ----------------");
    System.out.println(Arrays.toString(aArray));
    System.out.println(Arrays.toString(bArray));
    System.out.println(Arrays.toString(cArray));

  }



  public static void outputArray() {
    System.out.println("------------- outpuy array ----------------");

    int[] intArray = {1, 2, 3, 4, 5};
    String intArrayString = Arrays.toString(intArray);
    System.out.println(intArrayString);
  }


  public static void  arrayToList() {
    String[] stringArray = {"a", "b", "c", "d", "e"};
    ArrayList<String> arrayList = new ArrayList<String>(Arrays.asList(stringArray));
    System.out.println(arrayList);
  }


  public static void checkExists() {
    String[] stringArray = {"a", "b", "c", "d", "e"};
    boolean b = Arrays.asList(stringArray).contains("a");
    System.out.println(b);
  }



  public static void connectArray() {
    int[] intArray = {1, 2, 3, 4, 5};
    int[] intArray2 = {6, 7, 8, 9, 10};
    int[] combinedintArray = ArrayUtils.addAll(intArray, intArray2);
  }

  public static void defineInlineArray() {
    //method (new String[] {"a", "b", "c"});
  }


  public static void bytesArray() {
    byte[] bytes = ByteBuffer.allocate(4).putlnt(8).array();

    for (byte t : bytes) {
      System.out.println("0x%x", t);
    }
  }

}