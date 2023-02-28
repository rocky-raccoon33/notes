# CHAPTER 4 - 抽象：线程

> 关键问题：操作系统如何提供有多个 `CPU` 的假象？

{== 虚拟化CPU：通常让一个进程只运行一个时间片，然后切换到其他进程 [时分共享（time sharing）CPU技术] ==}，潜在的开销是性能损失。

操作系统为正在运行的程序提供的抽象：进程
进程 API：

- 创建（ create ）
- 销毁（ destroy ）
- 等待（ wait ）
- 其他控制（ miscellaneous control）
- 创建（ status ）

