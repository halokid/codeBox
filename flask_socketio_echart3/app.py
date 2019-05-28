# encoding:utf-8
##import psutil
import time
from threading import Lock

# import paramiko
from flask import Flask, render_template, request
from flask_socketio import SocketIO
import json
async_mode = None
app = Flask(__name__)
app.config['SECRET_KEY'] = 'secret!'
socketio = SocketIO(app, async_mode=async_mode)
thread = None
thread_lock = Lock()
class ReturnJ(object):
  def __init__(self):
    #由于存在setattr,此处必须采用这种方式赋值
    self.__dict__['res'] = {
      'code': 200,
      'msg': '请求成功!'
    }

  def toJson(self):
    return json.dumps(self.res, ensure_ascii=False)

  def __setattr__(self, key, val):
    self.res[key] = val

@app.route("/", methods=['POST','GET'])
def lowerString():
  data = request.values.to_dict()
  ret = ReturnJ()
  ret.data = data
  socketio.emit('response_data',
                {'data': data, 'count': 1},
                namespace='/showData')
  return ret.toJson()


if __name__ == '__main__':
  socketio.run(app, host='0.0.0.0',port=5000,debug=True)





