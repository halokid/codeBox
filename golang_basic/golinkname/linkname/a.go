package linkname

import (
  "log"
  _ "unsafe"
)

//go:linkname hello  ../outer.World
func hello()  {
  log.Println("hello, world!")
}



