完成 CLI 工具最后一公里的体验优化：


支持 Ctrl+C 优雅退出

1. sigChan := make(chan os.Signal, 1)
这一行创建了一个带缓冲区的通道(channel)，用于接收操作系统信号。

make(chan os.Signal, 1) 创建一个元素类型为 os.Signal 的通道
参数 1 表示这个通道有一个缓冲区，意味着它可以存储一个信号而不阻塞
缓冲区的作用是防止在没有人接收信号时丢失信号
2. signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
这一行注册了我们想要捕获的系统信号。


signal.Notify 是 Go 标准库 os/signal 包中的函数
它告诉运行时将指定的信号转发到 sigChan 通道
syscall.SIGINT 是中断信号，通常由用户按下 Ctrl+C 产生
syscall.SIGTERM 是终止信号，通常由系统管理员或进程管理器发送

3. 匿名 goroutine
go func() {
    sig := <-sigChan
    log.Printf("接收到系统信号: %v, 正在优雅退出...", sig)
    cancel()
}()
这部分启动了一个新的 goroutine (轻量级线程)，它在后台运行，专门用于处理系统信号：
go func() { ... }() 创建并立即执行一个匿名函数作为新的 goroutine
sig := <-sigChan 这一行会阻塞 goroutine，直到从 sigChan 通道接收到一个信号
当接收到信号后，它会打印一条日志消息，表明收到了哪种信号
最后调用 cancel() 函数，这会触发上下文的取消操作

整体工作流程
1.程序创建一个信号通道并注册 SIGINT 和 SIGTERM 信号
2.启动一个 goroutine 在后台等待这些信号
3.程序正常运行主逻辑
4.如果用户按下 Ctrl+C 或系统发送终止信号：
    操作系统向程序发送 SIGINT 或 SIGTERM 信号
    Go 运行时将该信号转发到 sigChan
    后台 goroutine 接收到信号并调用 cancel() 函数
    cancel() 函数取消上下文，通知所有使用该上下文的操作
    程序的各个部分检测到上下文被取消，开始清理并退出

这种模式在 Go 应用程序中非常常见，尤其是在开发服务器和长期运行的应用时。它允许程序在收到终止信号时优雅地关闭，而不是突然终止，从而有机会正确清理资源（关闭文件、完成正在进行的事务、关闭网络连接等）。


*   CLI 命令结构整理（模块化、参数校验）
*   输出统一格式（日志/错误/结果美化）