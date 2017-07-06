import tensorflow as tf 
import os

os.environ['TF_CPP_MIN_LOG_LEVEL']='3'
matrix1 = tf.constant([[3., 3.]])

matrix2 = tf.constant([[2.], [2.]])

product = tf.matmul(matrix1, matrix2)

#print product

sess = tf.Session()

result = sess.run(product)
print result
sess.close()

print "---------------------------------\n"



state = tf.Variable(0, name="counter")

one = tf.constant(1)
new_value = tf.add(state, one)
update = tf.assign(state, new_value)

init_op = tf.global_variables_initializer()

with tf.Session() as sess:
    sess.run(init_op)
    print (sess.run(state))

    for _ in range(3):
        sess.run(update)
        print (sess.run(state))




