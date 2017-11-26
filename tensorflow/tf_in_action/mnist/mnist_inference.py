#coding=utf-8

import tensorflow as tf

#定义神经网络结构相关的参数
INPUT_NODE = 784
OUTPUT_NODE = 10
LAYER1_NODE = 500

#通过  tf.get_variable 函数来获取变量。 在训练神经网络时会创建这些变量；在测试时
#会通过保存的模型加载这些变量的值。而且更加方便的是， 因为可以在变量加载时将滑动平均
#重命名， 所以可以直接通过同样的名字在训练时使用变量自身， 而在测试时使用变量的滑动平均
#值。 在这个函数中也会将变量的正则化损失加入损失集合。
def get_weight_variable(shape, regularizer):
  weights = tf.get_variable(
    "weights", shape,
    initializer = tf.truncated_normal_initializer(stddev=0.1))



# 当给出了正则化生成函数时，将当前变量的正则化损失加入名字为 losses 的集合。在这里
# 使用了  add_to_collection 函数将一个张量加入一个集合， 而这个集合的名称为 losses.
# 这里自定义的集合， 不在  tensorflow 自动管理的集合列表中
def inference(input_tensor, regularizer):
  #声明第一层神经网络的变量并完成前向传播过程
  with tf.variable_scope('layer1'):
    weights = get_weight_variable( [INPUT_NODE, LAYER1_NODE], regularizer )
    biases = tf.get_variable( "biases", [LAYER1_NODE],
                              initializer = tf.constant_initializer(0.0))
    layer1 = tf.nn.relu( tf.matmul (input_tensor, weights) + biases )


  with tf.variable_scope('layer2'):
    weights = get_weight_variable(
      [LAYER1_NODE, OUTPUT_NODE], regularizer)
    biases = tf.get_variable(
      "biases", [OUTPUT_NODE], initializer = tf.constant_initializer(0.0))
    layer2 = tf.matmul(layer1, weights) + biases

  return layer2







