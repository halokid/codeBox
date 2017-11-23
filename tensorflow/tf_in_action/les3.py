#coding=utf-8

import tensorflow as tf

#构造层运算的变量的运算逻辑， 下面还要初始化的
w1 = tf.Variable(tf.random_normal([2, 3], stddev=1, seed=1))
w2 = tf.Variable(tf.random_normal([3, 1], stddev=1, seed=1))

#定义一个输入的向量
x = tf.constant([0.7, 0.9])


#下面这个逻辑是经过两层运算
a = tf.matmul(x, w1)
y = tf.matmul(a, w2)


#创建一个会话
sess = tf.Session()

#初始化两个运算逻辑
sess.run(w1.initializer)
sess.run(w2.initializer)

print(sess.run(y))
sess.close()






