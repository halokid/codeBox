/**
下面这段代码仅适用于linux，golang在windows下面 mmap函数有不同的API
**/
package main

import (
    "fmt"
    "os"
    "syscall"
    "unsafe"
)

func main() {
    const n = 1e3
    t := int(unsafe.Sizeof(0)) * n

    map_file, err := os.Create("/tmp/test.dat")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    _, err = map_file.Seek(int64(t-1), 0)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    _, err = map_file.Write([]byte(" "))
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    //把文件句柄作为参数传入MMAP的函数里面， 然后改变变量 mmap 的时候，这个变量的变动
    //就会自动更新到 这个文件里面去
    mmap, err := syscall.Mmap(int(map_file.Fd()), 0, int(t), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    //mmap是  & 引用（指针的用法）， 那么即是改变 map_array的话也就是改变mmap变量了
    map_array := (*[n]int)(unsafe.Pointer(&mmap[0]))

    for i := 0; i < n; i++ {
        map_array[i] = i * i
    }

    fmt.Println(*map_array)

    err = syscall.Munmap(mmap)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    err = map_file.Close()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
