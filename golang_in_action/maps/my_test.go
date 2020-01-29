package maps

import (
  "testing"
)

func TestComm(t *testing.T) {
  m :=CreateRWLockMap()
  t.Log(m)
}
