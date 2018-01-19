#coding=utf-8

# %matplotib inline
# %config InlineBack.figure_format = 'retina'

import numpy as np
import pandas as pd
import matplotlib.pyplot as plt
import unittest


################## 数据的治理 start ##############################################

#装载数据
data_path = 'Bike-Sharing-Dataset/hour.csv'
rides = pd.read_csv(data_path)
rides.head()

#选取数据
rides[:24 * 10].plot(x = 'dteday', y = 'cnt')
# print rides

#虚拟变量, 哑变量
dummp_fields = ['season', 'weathersit', 'mnth', 'hr', 'weekday']
# dummp_fields = ['season']
for each in dummp_fields:
  dummies = pd.get_dummies(rides[each], prefix=each, drop_first=False)
  rides = pd.concat([rides, dummies], axis=1)

fields_to_drop = ['instant', 'dteday', 'season', 'weathersit',
                  'weekday', 'atemp', 'mnth', 'workingday', 'hr']
data = rides.drop(fields_to_drop, axis=1)
data.head()
# print data

#调整目标变量， 使它们的均值为0， 标准差为1
quant_features = ['casual', 'registered', 'cnt', 'temp', 'hum', 'windspeed']
scaled_features = {}
for each in quant_features:
  #fixme: 搞清楚这个 mean 和 std 是什么？ mean是平均值， std是标准值
  mean, std = data[each].mean(), data[each].std()
  scaled_features[each] = [mean, std]
  data.loc[:, each] = (data[each] - mean) / std


#将数据拆分为训练， 测试和验证数据集
#最后的 21  天 作为训练数据
test_data = data[-21 * 24:]

#除去 test_data 的其他数据，作为验证数据
data = data[:-21 * 24]


target_fields = ['cnt', 'casual', 'registered']

features, targets = data.drop(target_fields, axis=1), data[target_fields]   #验证数据
test_features, test_targets = test_data.drop(target_fields, axis=1), test_data[target_fields]     #训练数据

#训练数据是 除最后的60行之外的所有数据
train_featurea, train_targets = features[:-60 * 24], targets[:-60 * 24]
#验证数据是最后的60行
val_features, val_targets = features[-60 * 24:], targets[-60 * 24:]


################## 数据的治理 end ##############################################



###################### 测试用例 start ###########################

#输入的数据
inputs = np.array([[0.5, -0.2, 0.1]])
#我们要的目标数据
targets = np.array([[0.4]])

#输入隐藏层的权重, 这是两层
test_w_i_h = np.array([[0.1, -0.2],
                      [0.4, 0.5],
                      [-0.3, 0.2]])

#输出隐藏层的权重, 这是一层
test_w_h_o = np.array([[0.3],
                       [-0.1]])

class TestMethods(unittest.TestCase):
  def test_data_path(self):
    self.assertTrue(data_path.lower() == 'xxxxxxx')

  def test_data_loaded(self):
    self.assertTrue(isinstance(rides, pd.DataFrame))


  ######
  #unit tests for neiwork functionality
  ######

  def test_activation(self):
    network = NeuralNetwork(3, 2, 1, 0.5)
    self.assertTrue(np.all(network.activation_function(0.5) == 1/(1 + np.exp(-0.5))))

  def test_train(self):
    #test that weights are updateed correctly on training
    network = NeuralNetwork(3, 2, 1, 0.5)
    network.weights_input_to_hidden = test_w_i_h.copy()
    network.weights_hidden_to_output = test_w_h_o.copy()

    self.assertTrue(np.allclose(network.run(inputs), 0.9998924))

  suite = unittest.TestLoader().loadTestsFromModule(TestMethods())
  unittest.TextTestRunner().run(suite)




###################### 测试用例 end  ###########################


def MSE(y, Y):
  return np.mean((y - Y)**2)


###################### 神经网络类 start  ###########################
class NeuralNetwork(object):
  def __init__(self, input_nodes, hidden_nodes, output_nodes, learning_rate):
    self.input_nodes = input_nodes
    self.hidden_nodes = hidden_nodes
    self.output_nodes = output_nodes

    #initialize weights
    self.weights_input_to_hidden = np.random.normal(0.0, self.input_nodes**-0.5, self.input_nodes, self.hidden_nodes)
    self.weights_hidden_to_output = np.random.normal(0.0, self.hidden_nodes**-0.5, self.hidden_nodes, self.output_nodes)
    self.lr = learning_rate

    self.activation_function = lambda x : 0

  def train(self, features, targets):
    n_records = features.shape[0]
    delta_weghts_i_h = np.zeros(self.weights_input_to_hidden.shape)
    delta_weights_h_o = np.zeros(self.weights_hidden_to_output.shape)
    for X, y in zip(features, targets):
      #### implement the forward pass here ####
      ### forword pass ###

      # todo: hidden lqyer
      hidden_inputs = None
      hidden_outputs = None

      # todo: output layer
      final_inputs = None
      final_output = None

      #### implement the backward pass here ####
      ### backword pass ###


      # todo: outpout error
      error = None

      # todo: calcaulate the hidden layer's contribution to the error
      hidden_error = None

      # todo: backpropagated error term
      output_error_term = None
      hidden_error_term = None

      #weights step (input to hidden)
      delta_weights_i_h += None
      #weights step (hidden to output)
      delta_weghts_h_o += None


    #todo: daupte weights, replace these values with your calcualtions
    self.weights_hidden_to_output += None
    self.weights_input_to_hidden += None

  def run(self, features):
    hidden_inputs = None
    hidden_outputs = None

    final_inputs = None
    final_output = None

    return final_output





###################### 神经网络类 end  ###########################
















