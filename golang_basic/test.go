package main

func foox(sl []int) {
  sl[0] = 9
}

func fooy() string {
  defer println("fooy defer")
  panic("fooy panic")
  return "fooy func"
}

func main() {
  /**
  sl := make([]int, 10)
  fmt.Println(reflect.TypeOf(sl))

  sl = append(sl, 1)
  fmt.Println(sl)

  foox(sl)
  fmt.Println(sl)
  */

  fooy()
}

