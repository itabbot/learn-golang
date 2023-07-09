package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"syscall"
	"time"
)

// 获取当前的 Goroutine ID
func getGoroutineId() int64 {
	// 获取当前 goroutine 的调用栈信息
	var buf [64]byte
	n := runtime.Stack(buf[:], false)

	// 截取 goroutine id
	id := int64(-1)
	_, _ = fmt.Sscanf(string(buf[:n]), "goroutine %d", &id)

	return id
}

// 获取当前线程 ID
// 这里是 Windows 系统的获取方式
// Linux 系统直接 return syscall.Gettid()
func getThreadId() int64 {
	threadId := int64(-1)

	// 加载动态链接库 kernel32.dll 文件
	kernel32, err := syscall.LoadLibrary("kernel32.dll")
	if err != nil {
		return threadId
	}
	defer func(handle syscall.Handle) {
		_ = syscall.FreeLibrary(handle)
	}(kernel32)

	// 获取动态链接库（DLL）文件中 GetCurrentThreadId 函数的地址
	proc, err := syscall.GetProcAddress(kernel32, "GetCurrentThreadId")
	if err != nil {
		return threadId
	}

	// 调用动态链接库（DLL）文件中的函数
	rst, _, _ := syscall.SyscallN(proc, 0, 0, 0, 0)

	return int64(rst)
}

// 程序入口
func main() {
	// 设置用于执行 goroutine 的最大线程数
	runtime.GOMAXPROCS(2)

	// 请求的处理逻辑
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 打印当前的系统信息
		fmt.Printf(
			"进程ID:%d 线程ID:%d GoroutineId:%d \n",
			os.Getpid(), getThreadId(), getGoroutineId(),
		)
		// 阻塞请求
		time.Sleep(time.Second * 100)
	})

	// 启动 HTTP 服务
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println("HTTP 服务启动失败", err.Error())
	}
}
