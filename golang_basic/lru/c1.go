package lru

import "container/list"

type Cache struct {
  MaxExtries    int
  ll            *list.List
  cache         map[string]*list.Element
}

type entry struct {
  key     string
  value   interface{}
}

func New(max int) *Cache {
  return &Cache{
    MaxExtries:     max,
    ll:             list.New(),
    cache:          make(map[string]*list.Element),
  }
}

func (c *Cache) Set(key string, value interface{}) {
  // if cache is nil
  if c.cache == nil {
    c.cache = make(map[string]*list.Element)
    c.ll = list.New()
  }

  // if key is exists
  if ee, ok := c.cache[key]; ok {
    c.ll.MoveToFront(ee)
    ee.Value.(*entry).value = value
    return
  }

  // add new key to double linked list & push to front
  ele := c.ll.PushFront(&entry{key: key, value: value})
  c.cache[key] = ele
}



