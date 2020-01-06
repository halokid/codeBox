package golang_in_action

import (
  "errors"
  "time"
)

type ReusableObj struct {

}


type ObjPool struct {
  bufChan  chan *ReusableObj
}

func NewObjPool(numOfObj int) *ObjPool {
  objPool := ObjPool{}
  objPool.bufChan = make(chan *ReusableObj, numOfObj)
  for i := 0; i < numOfObj; i++ {
    objPool.bufChan <-&ReusableObj{}
  }
  return &objPool
}

func (p *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error) {
  select {
  case ret := <-p.bufChan:
    return ret, nil
  case <-time.After(timeout):
    return nil, errors.New("获取Obj超时")
  }
}

func (p *ObjPool) Release(obj *ReusableObj) error {
  select {
  case p.bufChan <-obj:
    return nil
  default:
    return errors.New("channel满了，放不进去")
  }
}




