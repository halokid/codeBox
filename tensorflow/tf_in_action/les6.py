#coding=utf-8
import os
os.environ['TF_CPP_MIN_LOG_LEVEL']='3'
import tensorflow as tf

v1 = tf.Variable(tf.constant(1.0, shape=[1], name="v1"))
v2 = tf.Variable(tf.constant(2.0, shape=[1], name="v2"))

'''
装载一个已经保存的模型， 注意这里不用 initialize_all_variables
'''

result = v1 + v2

saver = tf.train.Saver()

with tf.Session() as sess:
  saver.restore(sess, "./model.ckpt")
  print(sess.run(result))


