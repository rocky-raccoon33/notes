
`关键问题：操作系统如何提供有多个 CPU 的假象？`

{== 虚拟化CPU：通常让一个进程只运行一个时间片，然后切换到其他进程 [时分共享（time sharing）CPU技术] ==}，潜在的开销是性能损失。

操作系统为正在运行的程序提供的抽象：进程

进程 API：

- 创建 `create`
- 销毁 `destroy`
- 等待 `wait`
- 其他控制 `miscellaneous control`
- 创建 `status`

#### fork()：创建子进程
#### wait()：等待子进程执行完毕
#### exec()：执行与父进程不同的程序

```c
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

int main(int argc, char *argv[])
{
	printf("hello world (pid:%d)\n", (int)getpid());
	int rc = fork();
	if (rc < 0)
	{
		fprintf(stderr, "fork failed\n");
	}
	else if (rc == 0)
	{
		printf("hello, I am child(pid:%d)\n", (int)getpid());
	}
	else
	{
		printf("hello, I am parent of %d(pid:%d)\n", rc, (int)getpid());
	}

	return 0;
}
```

输出结果：

```dotnetcli
hello world (pid:33414)
hello, I am parent of 33435(pid:33414)
hello world (pid:33414)
hello, I am child(pid:33435)
```