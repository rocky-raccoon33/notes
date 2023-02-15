# QA

## `1` Jvm SafePoint

> JVM 会在以下情况进入 `safePoint`，确保所有线程都暂停，以便进行关键操作，保证系统的正确性

- `GC` 操作：`STW`
- `JIT`（即时编译优化）：JIT编译器需要自方法去进行内存重构
- 线程操作：创建、终止或者暂停线程
- 系统调用：如创建文件或网络连接

## `2` Jvm 调优

1. 选择合适的 `垃圾收集器（GC）`：根据应用程序的特征，选择合适的 `GC` 可以提高性能
2. 调整内存大小：通过设置 JVM 参数，例如 `-Xmx` 和 `-Xms`，可以调整 `JVM` 可用的最大内存和最小内存。如果内存不够，`JVM` 就会进行垃圾回收，从而影响性能
3. 启用编译器优化：`JVM` 中的 `Just-In-Time (JIT)` 编译器可以在运行时将字节码编译为本地代码。启用 `JIT` 优化可以提高代码执行速度。
4. 使用内存池：`JVM` 中的内存池可以避免内存碎片和提高内存利用率。使用合适的内存池可以提高性能。
5. 使用性能监控工具：使用性能监控工具，例如 `JVisualVM` 或 `JProfiler`，了解 `jvm` 的运行情况

## `3` Jvm 进入 `safePoint`

> JVM 会在以下情况进入 `safePoint`，确保所有线程都暂停，以便进行关键操作，保证系统的正确性

- `GC` 操作：`STW`
- `JIT`（即时编译优化）：JIT编译器需要自方法去进行内存重构
- 线程操作：创建、终止或者暂停线程
- 系统调用：如创建文件或网络连接

## `4` Java 线程池的状态

1. `RUNNING（运行状态）`：线程池处于正常的工作状态，可以接收新的任务，也可以处理已有的任务。

2. `SHUTDOWN（关闭状态）`：线程池停止接收新的任务，但是会继续处理已有的任务，等所有任务都完成之后，线程池就会转换到TERMINATED状态。

3. `STOP（停止状态）`：线程池停止接收新的任务，并且会尝试中断正在处理的任务，等所有任务都完成之后，线程池就会转换到TERMINATED状态。

4. `TIDYING（整理状态）`：线程池中的所有任务都已经完成，但是线程池中的线程还没有全部销毁，正在执行一些清理工作。

5. `TERMINATED（终止状态）`：线程池中的所有任务都已经完成，线程池中的所有线程都已经销毁。

> **线程池的状态转换过程如下：**

- 初始状态 `RUNNING`
- `shutdown()` -> `SHUTDOWN`
- `shutdownNow()` -> `STOP`
- 当线程池中的所有任务都完成之后，进入 `TIDYING` 状态，等待清理工作完成
- 清理工作完成之后，线程池会进入 `TERMINATED` 状态

## `5` Volatile

> **answer in** [stackoverflow](https://stackoverflow.com/questions/106591/what-is-the-volatile-keyword-useful-for)

- Shared Multiprocessor Architectur

![](img/cpu.webp)

`volatile` has semantics for memory visibility. Basically, the value of a volatile field becomes visible to all readers (other threads in particular) after a write operation completes on it. Without `volatile`, readers could see some non-updated value.

___

## How does  `compare and set` in `AtomicInteger` works

>  **answer in** [stackoverflow](https://stackoverflow.com/questions/32634280/how-does-compare-and-set-in-atomicinteger-works)

`AtomicInteger` works with two concepts: `CAS` and `volatile` variable.

Using `volatile` variable insures that the current value will be visible to all threads and it will not be cached.

But I am confused over CAS（compare AND set）concept which is explained below:

```java
public final int getAndIncrement() {
    for (;;) {
        int current = get();
        int next = current + 1;
        if (compareAndSet(current, next))
            return current;
    }
 }
```

My question is that what if `compareAndSet(current, next` returns false? Will the value not be updated? In this case what will happen when a Thread is executing the below case:

```java
private AtomicInteger count = new AtomicInteger();
count.incrementAndGet();
```


The atomic objects make use of [Compare and Swap mechanism](https://www.wikiwand.com/en/Compare-and-swap) to make them atomic - i.e. it is possible to guarantee that the value was as specified and is now at the new value.

The code you posted continually tries to set the current value to one more than it was before. Remember that another thread could also have performed a get and is trying to set it too. If two threads race each other to change the value it is possible for one of the increments to fail.

Consider the following scenario:

- Thread 1 calls get and gets the value 1.
- Thread 1 calculates next to be 2.
- Thread 2 calls get and gets the value 1.
- Thread 2 calculates next to be 2.
- Both threads try to write the value.

Now because of atomics - **only one thread will succeed**, the other will recieve false from the compareAndSet and go around again.

If this mechanism was not used it would be quite possible for both threads to increment the value resulting in only one increment actually being done.

The confusing infinite loop `for(;;)` will only really loop if many threads are writing to the variable at the same time. Under very heavy load it may loop around several times but it should complete quite quickly.

