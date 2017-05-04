

------------------------------------ ���� --------------------------------

// ***** �������� *******


����:
int8, uint8, int16, uint16, int32, uint32, int64, uint64, int, rune, byte, complex128, complex64

byte == int8



������:
float32, float64



��������:
complex64, complex128


�ַ���:
string


�ַ�����:
rune (int32�ı���)


��������:
error



// **** �������� ****
ָ��        (pointer)
����        (array)
��Ƭ        (slice)
�ֵ�        (map)
ͨ��        (chan)
�ṹ��      (struct)
�ӿ�        (interface)





---------------------- �ؼ���---------------------------------

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

//golang��֧�ֺ��������أ����ǿ���ͨ������ķ�ʽȥʵ������Ч��

func myfunc(args ...int) {
  //TODO
}


myfunc(2)

myfunc(1, 3, 5)




//�����Ķ�ֵ����
func (file *file) Read(b []byte) (n int, err error)

//���ǿ���ͨ���»��� (_) ������ĳ������ֵ
n, _ := f.Read(buf)



//��������
func (a, b int) bool {
  return a < b
}


f := func(a, b int) bool {
  return a < b
}

func (a, b int) bool {
  return a < b
} (3, 4)   //�����ź�ֱ�Ӹ������б��ʾ��������



------------------------- �հ� ----------------------

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



-------------------- ������ -------------------------

package main

import "fmt"

func main() {
  defer func() {
    fmt.Println("recovered: ", recover())
  } ()
  
  panic("not good")
}




----------------- ������� ---------------------------
type Name struct {
  First   string
  Middle  string
  Last    string
}


type SimpleName string

func (s SimpleName) String() string { return string(s) }
//����
func (s string) NoWay()




----------------  sample code --------------------------



------ ������ Generator -------


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







------- װ���� Decorator -------

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











  


































































































