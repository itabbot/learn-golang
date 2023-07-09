# 问题：Golang HTTP 服务是否使用单独 Goroutine 处理每个请求？

## 1. 运行程序

运行 `./main.go` 后，多次访问 `http://127.0.0.1/` 输出：

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

## 2. 理解

1. Go 程序运行在单个进程中，如上述输出所示，进程 ID 为 25396。
2. 在程序的开始，使用了 `runtime.GOMAXPROCS(2)` 设置用于执行 goroutine 的最大线程数，所以上述的输出可以看到，线程 ID 徘徊在 22820 和 11640 之间。
3. 如上述输出所示，Go 的 HTTP 服务器对每个请求都使用了单独的 Goroutine 进行处理。
