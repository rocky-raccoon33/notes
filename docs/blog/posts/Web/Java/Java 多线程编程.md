---
title: Java 多线程编程
comments: true
description: Java 为多线程编程提供了良好、考究并且一致的编程模型，使开发人员能够更加专注于问题的解决，即为所遇到的问题建立合适的模型，而不是绞尽脑汁地考虑如何将其多线程化
date: 2022-11-21
category: Programming
section: blog
tags:
    - Java
---

## 1 Concepts

### 1.1 线程状态

`deadlock`：死锁线程，一般指多个线程调用期间进入了相互资源占用，导致一直等待无法释放的情况。

`runnable`：一般指该线程正在执行状态中。

`blocked`：线程正处于阻塞状态。

`waiting on condition`：线程正处于等待资源或等待某个条件的发生，具体的原因需要结合堆栈信息进行分析。

- 可能是在等待锁资源
- 可能是在 sleep
- 可能是网络资源不足，这个要结合系统的使用信息来看。
- waiting for monitor entry 或 in Object.wait()

:    java中 synchronized 的重量级锁借助 Moniter 来实现。Moniter 有两个队列。一个是 entry set，另一个是 wait set。当 synchronized 升级到重量级锁后，竞争锁失败的线程将记录到 entry set 中，这时线程状态是 waiting for monitor entry；而竞争到锁的线程若调用锁对象的 wait 方法时，则记录 到 wait set 中，这时线程状态是 in Object.wait()。

### 1.2 线程池的状态

`RUNNING`：线程池处于正常的工作状态，可以接收新的任务，也可以处理已有的任务

`SHUTDOWN`：线程池停止接收新的任务，但是会继续处理已有的任务，等所有任务都完成之后，线程池就会转换到TERMINATED状态

`STOP`：线程池停止接收新的任务，并且会尝试中断正在处理的任务，等所有任务都完成之后，线程池就会转换到 TERMINATED 状态

`TIDYING`：线程池中的所有任务都已经完成，但是线程池中的线程还没有全部销毁，正在执行一些清理工作

`TERMINATED`：线程池中的所有任务都已经完成，线程池中的所有线程都已经销毁

线程池的状态转换过程如下：

- 初始状态 RUNNING
- shutdown() -> SHUTDOWN
- shutdownNow() -> STOP
- 当线程池中的所有任务都完成之后，进入 TIDYING 状态，等待清理工作完成
- 清理工作完成之后，线程池会进入 TERMINATED 状态

## 2 线程同步

线程同步可以有效避免多个线程访问共享资源时产生的冲突和不一致问题，从而确保程序的正确性。过多的同步也可能降低程序的性能，使用线程同步需要权衡性能和正确性。

### 2.1 synchronized

`保证多线程访问共享资源时的安全性和一致性`

当一个线程进入 synchronized 块时，它会尝试获取锁，如果该锁被其他线程持有，则该线程就会被阻塞，直到该锁被释放。
当一个线程退出 synchronized 块时，它会释放持有的锁，从而允许其他线程进入该块

每个对象都有一个监视器（monitor）锁
当一个线程进入 synchronized 块时，它获取该对象的监视器锁
只有该线程持有该锁时，其他线程才无法进入该块
Java 的内存模型保证，一个线程释放锁之前，对共享变量所做的修改都将立即对其他线程可见

### 2.2 AQS

`AbstractQueuedSynchronizer` 是 Java 中实现锁的基础框架，提供了一种实现同步状态的机制。
AQS `的实现主要依赖于两个类：Node` 和 `Sync`

####  Node

AQS 中的一个内部类，用于表示一个等待获取锁的线程。在 Node 中，包含了线程的状态（CANCELLED、SIGNAL、CONDITION、PROPAGATE、0），以及等待队列中的前驱节点、后继节点等信息。

#### Sync

AQS 的抽象类，它定义了 AQS 的主要接口和状态，是 AQS 实现的核心。在 Sync 中，定义了一系列方法，用于实现锁的获取和释放，包括`tryAcquire` `tryRelease` `acquire` `release`

#### 实现锁的方式

AQS通过维护一个 FIFO 的等待队列，来实现锁的获取和释放。当一个线程请求获取锁时，如果锁已经被占用，那么线程就会被加入等待队列中，然后自旋等待锁的释放。当锁被释放时，AQS 会从等待队列中取出一个线程，并唤醒它，使其重新尝试获取锁。

1. 尝试获取锁：tryAcquire 尝试获取锁，并返回是否获取成功。
2. 获取成功：直接返回。
3. 自旋等待：如果当前线程的前驱节点是头节点，并且当前线程成功获取到锁，就将当前线程设置为头节点，然后唤醒后继节点，使其重新尝试获取锁。
4. 释放锁：tryRelease 方法会释放锁，并返回是否释放成功。
5. 释放成功：尝试唤醒等待队列中的下一个节点。

AQS提供了一种通用的同步状态机制，可以用于实现多种不同的锁和同步工具
例如 `ReentrantLock` `Semaphore` `CountDownLatch`
