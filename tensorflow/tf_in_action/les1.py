#coding=utf-8
import tensorflow as tf


#计算公式:   y = a*x + b

#创建指定维度的零张量
zero_tsr = tf.zeros([row_dim, col_dim])

#创建指定维度的单位张量
ones_tsr = tf.ones([row_dim, col_dim])

#创建指定维度的常数填充的张量
filled_tsr = tf.fill([row_dim, col_dim], 42)

#用已知常数张量创建一个张量
constant_tsr = tf.constant([1, 2, 3])



