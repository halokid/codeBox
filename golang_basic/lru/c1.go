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

// set cache
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

  // if cache len larger than max-len of LRU, then remove double list last ele
  if c.MaxExtries != 0 && c.ll.Len() > c.MaxExtries {
    c.RemoveOldest()
  }
}

// get cache
func (c *Cache) Get(key string) (value interface{}, ok bool) {
  if c.cache == nil {
    return
  }

  if ele, hit := c.cache[key]; hit {
    c.ll.MoveToFront(ele)
    return ele.Value.(*entry).value, true
  }
  return
}


func (c *Cache) Remove(key string) {
  if c.cache == nil {
    return
  }

  if ele, hit := c.cache[key]; hit {
    c.removeElement(ele)
  }
}

// remove double list oldest(最远) element
func (c *Cache) RemoveOldest() {
  if c.cache == nil {
    return
  }

  ele := c.ll.Back()      // 获取最远的元素
  if ele != nil {
    c.removeElement(ele)
  }
}

func (c *Cache) removeElement(e *list.Element) {
  c.ll.Remove(e)
  kv := e.Value.(*entry)
  delete(c.cache, kv.key)
}

func (c *Cache) Len() int {
  if c.cache == nil {
    return 0
  }
  return c.ll.Len()
}

func (c *Cache) Clear() {
  c.ll = nil
  c.cache = nil
}




