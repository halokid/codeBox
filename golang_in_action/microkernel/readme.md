microkernel架构模式的总结
==========================


- micro kernel本身实现的方法，  pulgin也是继承
```

// -------------- kernel ---------------------
type Kernel struct {

}

func (k *Kernel) Start() {

}

func (k *Kernel) Stop() {

}


// -------------- plugin -------------------
type Plugin struct {

}

func (p *Plugin) Start() {

}


```

