

------------------------------------ 类型 --------------------------------

// ***** 基础类型 *******


整型:
int8, uint8, int16, uint16, int32, uint32, int64, uint64, int, rune, byte, complex128, complex64

byte == int8



浮点型:
float32, float64



复数类型:
complex64, complex128


字符串:
string


字符类型:
rune (int32的别名)


错误类型:
error



// **** 复合类型 ****
指针        (pointer)
数组        (array)
切片        (slice)
字典        (map)
通道        (chan)
结构体      (struct)
接口        (interface)





---------------------- 关键字---------------------------------

break     default     func    interface     select
case      defer       go      map           struct
chan      else        goto    package       switch
const     fallthrough if      range         type
continue  for         import  return        var




---------------------------------------------------------------

var a int
a = 2


a := 2

var (
  a int
  b bool
)

a, b := 2, true



------------------------------------------------------


len
cap
close
append
copy
delete


/array
a := [3]int{1, 2, 3}

a := [...]int{1, 2, 3}


//slice
s := make([]int, 3)
s := append(s, 1)
s := append(s, 2)


//map
m := make(map[string]int)
m["golang"] = 7



--------------------------------------------------------

//golang不支持函数的重载，但是可以通过下面的方式去实现类似效果

func myfunc(args ...int) {
  //TODO
}


myfunc(2)

myfunc(1, 3, 5)




//函数的多值返回
func (file *file) Read(b []byte) (n int, err error)

//我们可以通过下划线 (_) 来忽略某个返回值
n, _ := f.Read(buf)



//匿名函数
func (a, b int) bool {
  return a < b
}


f := func(a, b int) bool {
  return a < b
}

func (a, b int) bool {
  return a < b
} (3, 4)   //花括号后直接跟参数列表表示函数调用



------------------------- 闭包 ----------------------

package main

import "fmt"

func main() {
  j := 5
  a := func() func() {
    i := 10
    return func() {
      fmt.Println("i, j:  %d, %d\n", i, j)
    }
  } ()
  
  a()
  j *= 2
  a()
}



-------------------- 错误处理 -------------------------

package main

import "fmt"

func main() {
  defer func() {
    fmt.Println("recovered: ", recover())
  } ()
  
  panic("not good")
}




----------------- 面向对象 ---------------------------
type Name struct {
  First   string
  Middle  string
  Last    string
}


type SimpleName string

func (s SimpleName) String() string { return string(s) }
//或者
func (s string) NoWay()




----------------  sample code --------------------------



------ 生成器 Generator -------


#python version
def fib(n):
  a, b = 0, 1
  for i in range(n):
    a, b = b, a+b
    yield a
    

for x in fib(10):
  print x
  
print 'done'




#golang version
package main

import "fmt"

func fib(n int) chan int {
  c := make(chan int)
  go func() {
    a, b := 0, 1
    for i := 0; i < n; i++ {
      a, b = b, a+b
      c <- a
    }
    close(c)
  } ()
  return c
}


func main() {
  for x := range fib(10) {
    fmt.Println(x)
  }
}







------- 装饰器 Decorator -------

#python version
from urlparse import urlparse, parse_qs
from BaseHTTPServer import HTTPServer, BaseHTTPRequestHandler

def auth_required(myfunc):
  def checkuser(self):
    user = parse_qs(urlparse(self.path).query).get('user')
    if user:
      self.user = user[0]
      myfunc(self)
    else:
      self.wfile.write('unknow user')
    
  return checkuser



class myHandler(BaseHTTPRequestHandler):
  @auth_required
  def do_GET(self):
    self.wfile.write('Hello, %s!' % self.user)
    
    
if __name__ == '__main__':
  try:
    server = HTTPServer(('localhost', 8080), myHandler)
    server.serve_forever()
  except KeyboardInterrupt:
    server.socket.close()
    
    
    
#golang version
package main

import (
  "fmt"
  "net/http"
)
  

var hiHandler = authRequired(
  func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hi, %v", r.FormValue("user"))
  },
)



func authRequired(f http.HandlerFunc) http.HandlerFunc {  
    return func(w http.ResponseWriter, r *http.Request) {
        if r.FormValue("user") == "" {
            http.Error(w, "unknown user", http.StatusForbidden)
            return
        }
        f(w, r)
    }
}

func main() {  
    http.HandleFunc("/hi", hiHandler)
    http.ListenAndServe(":8080", nil)
}



----------------- 猴子补丁 -------------------------

#python version
import urllib

def say_hi(usr):
  if auth(usr):
    print 'Hi, %s' % usr
  else
    print 'unknow user %s' % usr
    
    
def auth(usr):
  try:
    auth_url = 'localhost'
    r = urllib.urlopen(auth_url + '/' + usr)
    return r.getcode() == 200
  except:
    return false;


def sayhitest():
  #Test authenticated user
  globals()['auth'] = lambda x: True
  say_hi('John')
  
  #Test unauthenticated user
  globals()['auth'] = lambda x: False
  say_hi('john')
  
if __name__ == '__main__':
  sayhitest()

    
    
    

#golang version
package main

import (
  "fmt"
  "net/http"
)

func sayHi(user string) {
  if !auth(user) {
    fmt.Println("unknow user %v\n", user)
    return
  }
  fmt.Println("Hi, %v\n", user)
}


var auth = func(user string) bool {
  authURL := "localhost"
  res, err := http.Get(authURL + "/" + user)
  return err == nil && res.StarusCode == http.StatusOK
}


func testSayHi() {
  auth = func(string) bool { return true }
  sayHi("John")
  
  auth = func (string) bool { return false }
  sayHi("John")
}



func main() {
  testSayHi()
}






    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    








  


































































































