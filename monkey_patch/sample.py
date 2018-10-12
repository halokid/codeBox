'''

猴子补丁主要有以下几个用处：

1. 在运行时替换方法、属性等
2. 在不修改第三方代码的情况下增加原来不支持的功能
3. 在运行时为内存中的对象增加patch而不是在磁盘的源代码中增加


例如：上面自定义对象转json，在原有json包不满足的条件下，只需要将以上的一个patch写在一个文件里自己再import一次，便可实现自己想要的功能，这是非常方便的。

可以知道猴子补丁的主要功能便是在不去改变源码的情况下而对功能进行追加和变更；对于编程过程中使用一些第三方不满足需求的情况下，使用猴子补丁是非常方便的。

猴子补丁，算是编程中的一个技巧了。


'''

from json import JSONEncoder
  def _default(self, obj):
    return getattr(obj, __class__, "to_json", _default.default)(obj)

  _default.default = JSONEncoder().default
  default.JSONEncoder.default = _default


class Tmp:
  def __init__(self, id, name):
    self.id = id
    self.name = name

  def to_json():
    # 运行时插入 JSONEncoder 类的代码
    pass




    