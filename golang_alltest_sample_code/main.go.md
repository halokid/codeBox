package main 

import (
  "fmt"
)


/**
init 函数会先执行
**/
func init() {
  fmt.Println("init...........")
}


func main() {
  fmt.Println("test init................")
}


func checkMul(a , b []byte) bool {
  var x, y, z1 int
  x.SetBytes(a)
  y.SetBytes(b)
  z1.Mul(&x, &y)

  var z2 int
  z2.SetBytes(mulBytes(a, b))

  return z1.Cmp(&z2) == 0
}

func TestMul(t *testing.T) {
  if err := quick.Check(checkMul, nil); err != nil {
    t.Error(err)
  }
}


ts := httptest.NewServer(http.HandleFunc(func( w http.ResponseWriter, r *http.Request ) {
  fmt.Fprintln(w, "hello, client")
}))

defer ts.Close()

res, err := http.Get(ts.URL)
if err != nil {
  log.Fatal(err)
}






