from django.shortcuts import render, redirect
from django.http import HttpResponse
from django.template import loader

# Create your views here.

def index(request):
  # process index
  temp = loader.get_template('Book/index.html')
  context = {'name': 'halokid', 'datas': list(range(1, 20))}
  html = temp.render(context)
  return HttpResponse(html)

def login(request):
  return render(request, "Book/login.html")

def login_submit(request):
    # 使用request.POST.get来获取相关的参数
    username = request.POST.get("username")
    password = request.POST.get("password")
    # 模拟判断账号密码是否正确
    if "halokid" == username and "xxxxx" == password:
        # 如果账号密码正确，重定向至首页
        return redirect("/index")
    else:
        # 如果错误重定向到登录页面
        return redirect("/login")