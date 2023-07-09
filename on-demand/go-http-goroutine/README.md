# Go 的 HTTP 服务器是否使用单独 Goroutine 处理每个请求<!-- omit in toc -->

## 1. 疑问

Go 语言的并发模型和协程机制使其非常适合用于构建高并发和高性能的应用程序。对此，不由地产生一个疑问：Golang 的 HTTP 服务器是如何处理每个请求的？我根据 Golang 的特点推测，应该是对每个请求都使用单独的 Goroutine 进行处理。

## 2. 思路

使用 Go 编写一个 HTTP 服务，在请求处理函数中打印当前的进程 ID、线程 ID，以及 Goroutine ID，并阻塞该请求，然后通过多次请求接口的方式即可看到打印的 Goroutine ID 情况。

## 3. 核实

1. 运行程序

运行 `./main.go`（[详见代码>>](./main.go)），然后多次访问 `http://127.0.0.1/`，控制台输出：

```sh
进程ID:25396 线程ID:11640 GoroutineId:19
进程ID:25396 线程ID:11640 GoroutineId:6
进程ID:25396 线程ID:22820 GoroutineId:8
进程ID:25396 线程ID:22820 GoroutineId:21
进程ID:25396 线程ID:22820 GoroutineId:23
进程ID:25396 线程ID:22820 GoroutineId:25
进程ID:25396 线程ID:22820 GoroutineId:27
进程ID:25396 线程ID:11640 GoroutineId:29
进程ID:25396 线程ID:11640 GoroutineId:10
进程ID:25396 线程ID:22820 GoroutineId:12
进程ID:25396 线程ID:22820 GoroutineId:31
```

2. 结果分析

   - Go 程序运行在单个进程中，如上述输出所示，进程 ID 为 25396。
   - 在程序的开始，使用了 `runtime.GOMAXPROCS(2)` 设置用于执行 goroutine 的最大线程数，所以上述的输出可以看到，线程 ID 徘徊在 22820 和 11640 之间。
   - 如上述输出所示，Go 的 HTTP 服务器对每个请求确实使用了单独的 Goroutine 进行处理。
