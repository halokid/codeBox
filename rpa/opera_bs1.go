package main
// 打开系统默认浏览器

import (
    "fmt"
    "os/exec"
    "runtime"
    "log"
    "path/filepath"
    "os"
)

var commands = map[string]string{
    "windows": "start",
    "darwin":  "open",
    "linux":   "xdg-open",
}

func Open(uri string) error {
    run, ok := commands[runtime.GOOS]
    if !ok {
        return fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)
    }
    log.Println("run---", run)

    cmd := exec.Command(run, uri)

    run2 := `"C:\Program Files (x86)\Microsoft\Edge\Application\msedge.exe"`
    cmd = exec.Command(run2, uri)

    // cmd := exec.Command("cmd.exe", "/c", "D:\\docker\\kubectl.exe", "version")
    cmd = exec.Command("cmd.exe", "/c", "C:\\Program Files (x86)\\Microsoft\\Edge\\Application\\msedge.exe", "http://bing.com")
    return cmd.Start()
}

func main() {
	err := Open("http://baidu.com")
	log.Println(err)

    ex, err := os.Executable()
    if err != nil {
        panic(err)
    }
    exPath := filepath.Dir(ex)
    fmt.Println(exPath)
}


