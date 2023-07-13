import time,os
import pickle
from selenium import webdriver
from selenium.webdriver.common.by import By

# 大麦网主页
damai_url = 'https://www.damai.cn/'
# 登录
login_url = 'https://passport.damai.cn/login?ru=https%3A%2F%2Fwww.damai.cn%2F'
# 抢票目标页
target_url = 'https://detail.damai.cn/item.htm?id=710818567651'


class Concert:
  def __init__(self):
    self.number = '13588888888'   # 登录需要的手机号
    self.num = 1  # 需要抢的票数
    self.name = '小陈'
    self.status = 0  # 状态，表示当前操作执行到了哪个步骤
    self.login_method = 1  # {0:模拟登录， 1:cookie登录}自行选择登录的方式
    self.options = webdriver.ChromeOptions()
    # 隐藏"Chrome正在受到自动软件的控制"
    self.options.add_experimental_option('useAutomationExtension', False)  # 去掉开发者警告
    self.options.add_experimental_option('excludeSwitches', ['enable-automation'])

    self.options.add_argument('--disable-blink-features=AutomationControlled')
    self.driver = webdriver.Chrome(options=self.options)    # 当前浏览器对象

  def enter_concert(self):
    """打开浏览器"""
    print('###打开浏览器，进入大麦网###')
    # 调用登录
    self.login()
    # self.driver.refresh()  # 刷新页面
    self.status = 2  # 登录成功标识
    print("###登录成功###")

  def get_cookie(self):
    """假如说本地有cookies.pkl 那么直接获取"""
    cookies = pickle.load(open('cookies.pkl', 'rb'))
    for cookie in cookies:
      cookie_dict = {
        'domain': '.damai.cn',  # 必须要有的，否则就是假登录
        'name': cookie.get('name'),
        'value': cookie.get('value')
      }
      self.driver.add_cookie(cookie_dict)
    print("###载入cookie###")

  def login(self):
    """登录"""
    if self.login_method == 0:
      self.driver.get(login_url)
      print('###开始登录###')
    elif self.login_method == 1:
      # 创建文件夹，文件是否存在
      if not os.path.exists('cookies.pkl'):
        self.set_cookies()   # 没有文件的情况下，登录一下
      else:
        self.driver.get(target_url)  # 跳转到抢票页
        self.get_cookie()

  def set_cookies(self):
    """cookies: 登录网站时出现的 记录用户信息用的"""
    self.driver.get(damai_url)
    self.driver.find_element(By.XPATH, '/html/body/div[2]/div/div[3]/div[1]/div[1]/span').click()
    iframe = self.driver.find_element(By.XPATH, '//div[@class="mui-zebra-module"]/div[1]/div[1]/div[1]/iframe')
    self.driver.switch_to.frame(iframe)
    self.driver.find_element(By.XPATH, '//*[@id="login-tabs"]/div[2]').click()
    self.driver.find_element(By.XPATH, '//*[@id="fm-sms-login-id"]').send_keys(self.number)
    self.driver.find_element(By.XPATH, '//*[@id="login-form"]/div[2]/div[3]/a').click()
    print("###请输入验证码###")
    while self.driver.title != '大麦网-全球演出赛事官方购票平台-100%正品、先付先抢、在线选座！':
      time.sleep(1)
    print("###登录成功###")
    pickle.dump(self.driver.get_cookies(), open('cookies.pkl', 'wb'))
    print("###cookie保存成功###")
    self.driver.get(target_url)

  def choose_num(self):
    for i in range(1, self.num):
      self.driver.find_element(By.XPATH, '/html/body/div[2]/div/div[1]/div[1]/div/div[2]/div[3]/div[6]/div[2]/div/div/a[2]').click()

  def choose_ticket(self):
    """下单操作"""
    if self.status == 2:
      print('=' * 30)
      self.choose_num()
      print("###准备购买{}张票###".format(self.num))
      print('=' * 30)
      print('###正在进行抢票###')
      while self.driver.title != '订单确认页':
        try:
          self.driver.find_element(By.XPATH, '/html/body/div[2]/div/div[1]/div[1]/div/div[2]/div[4]/div[10]/div/div[3]/div[3]').click()
        except:
          print('###没有跳转到订单结算界面###')

  def finifsh(self):
    """结束，退出浏览器"""
    self.driver.quit()


if __name__ == '__main__':
  concert = Concert()
  try:
    concert.enter_concert()  # 打开浏览器
    concert.choose_ticket()  # 抢票
  except Exception as e:
    print(e)
    concert.finifsh()


