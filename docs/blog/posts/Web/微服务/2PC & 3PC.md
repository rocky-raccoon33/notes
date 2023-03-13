---
title: 2PC & 3PC
comments: true
description: 单体服务中，事务可以通过数据库事务机制自行解决，例如：lock、redo log、undo log、ACID特性等。但是在微服务架构中，一个本地逻辑执行单位被拆分到多个独立的微服务中...
date: 2023-03-13
category: 分布式技术
section: blog
tags:
    - XA协议
---

术语：

TM：Transacion Manager（事务管理器），也称为 Coodinator

RM: Resource Manager（Cohort），例如订单系统、支付系统等

## `1` 2PC - Two-Phase Commit

{== 只有 TM 有超时 ==}

> Work Flow

### 第一阶段：pre-commit

TM 发送提交请求给所有相关的 RM，RMs 开启事务但不会提交，事务开启后所有资源加锁

### 第二阶段：do-commit

- 所有 RM 在预提交阶段成功，TM 向 RMs 发送 commit 指令，提交事务

![](./img/3-2pc-success.png)

- 存在一个或者多个 RM 在预提交阶段执行失败（网络异常、TM 超时等），TM 向 RMs 发送 abort 指令，回滚事务

![](./img/4-2pc-failure.png)

Images From [Patterns for distributed transactions within a microservices architecture](https://developers.redhat.com/blog/2018/10/01/patterns-for-distributed-transactions-within-a-microservices-architecture)

#### PROS

相比于服务间的直接调用，2PC 能保证更高的提交成功率，在 TM 发送 commit 指令前，第一阶段 pre-commit 已经保证 RM 能正常提交，各个服务大概率处于较好的状态

#### CONS

1. 单点失败：整个流程由TM管理，如果 TM 发生异常，分布式事务会被破坏
2. 阻塞正常业务：两个阶段资源都加锁，如果出现服务异常或者网络抖动，可能导致其他业务长时间无法访问相关资源
3. 数据不一致：第二阶段中，TM 发送 commit 指令给 RM1 并且 RM1 成功提交，在 RM2 接收到 commit 指令前 TM 和 RM1 服务异常，此时 RM1 提交但是 RM2 未提交，出现数据不一致的问题

## `2` 3PC - Three-Phase Commit
