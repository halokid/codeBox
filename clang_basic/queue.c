#include <stdio.h>
#include <stdlib.h>

/**
TODO: tips
the important point of this program is, we set two point for find the front &
rear of the Queue
**/

typedef struct Queue {
  int capacity;
  int size;
  int front;
  int rear;
  int *elements;
} Queue;


Queue * createQueue (int maxElements) {
  //create a Queue
  Queue *Q;
  Q = (Queue *)malloc(sizeof(Queue));
  //init its properties
  Q->elements = (int *)malloc(sizeof(int)*maxElements);
  Q->size = 0;
  Q->capacity = maxElements;
  Q->front = 0;
  Q->rear = -1;
  //return the point of the Queue
  return Q;
}

void Dequeue (Queue *Q) {
  //if Queue size is zero then it is empty, so we cannot pop
  if (Q->size == 0) {
    printf("Queue is empty\n");
    return;
  }
  //removing an element is equivalent to incrementing index of front by one
  else {
    Q->size--;
    Q->front++;
    //as we fill elements in circular fashion
    if (Q->front == Q->capacity) {
      Q->front = 0;
    }
  }
  return;
}


int front (Queue *Q) {
  if (Q->size == 0) {
    printf("Queue is empty\n");
    exit(0);
  }
  //return the element which is at the front
  return Q->elements[Q->front];
}

void Enqueue(Queue *Q, int element) {
  //if the Queue is full, we cannot push an element into it as there is no space for it
  if (Q->size == Q->capacity) {
    printf("Queue is full\n");
  } else {
    Q->size++;
    Q->rear = Q->rear + 1;
    //as we fill the Queue in circular fashion
    if (Q->rear == Q->capacity) {
      Q->rear = 0;
    }
    //insert the element in its rear side
    Q->elements[Q->rear] = element;
  }
  return;
}

int main() {
  Queue *Q = createQueue(5);
  Enqueue(Q, 1);
  Enqueue(Q, 2);
  Enqueue(Q, 3);
  Enqueue(Q, 4);
  printf("Front element is %d\n", front(Q));
  
  Enqueue(Q, 5);
  Dequeue(Q);
  Enqueue(Q, 6);
  printf("Front element is %d\n", front(Q) );
}



