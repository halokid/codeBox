microkernel架构模式的总结
==========================


- micro kernel本身实现的方法，  pulgin也是继承
```

// -------------- kernel ---------------------
type Kernel struct {
  pgs    *[]Plugin
}


func (k *Kernel) Start() {
  for _, pg := range k.pgs {
    pg.Start()
  }
}

func (k *Kernel) Stop() {

}


// -------------- plugin -------------------
type Plugin struct {
  name      string
}

func (p *Plugin) Start() {

}


// ---------- 模式调用 -----------
pg1 := &Plugin{"plugin1"}
pg2 := &Plugin{"plugin2"}

k := &Kernel{ pgs:   *[]Plugin{pg1, pg2} }
k.Start()


```

