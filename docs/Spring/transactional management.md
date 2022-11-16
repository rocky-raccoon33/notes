# `Spring Transaction management`

## 1 基本原理

> `Spring` 事务的本质其实就是数据库对事务的支持。对于纯JDBC操作数据库，想要用到事务，可以按照以下步骤进行：

1. 获取连接 `Connection con = DriverManager.getConnection()`
2. 开启事务 `con.setAutoCommit(true/false);`
3. 执行 `CRUD`
4. 提交事务/回滚事务 `con.commit() / con.rollback();`
5. 关闭连接 `conn.close();`

> 通过 `Spring` 事务管理，可以自动完成2，4：
___

> `Spring Boot`自动开启了对注解事务的支持

1. 在相关的类和方法上添加 `@Transactional`
2. `Spring` 启动时解析生成相关 `bean` 并生成代理，根据`@Transactional`的相关参数进行配置注入，在代理中开启正常提交事务以及异常回滚事务
3. 数据库的事务提交和回滚是通过`binlog`和`redo log`实现


> `Spring` 的事务机制提供了 `PlatformTransactionManager` 接口，不同的数据访问技术的事务使用不同的借口实现


- `JDBC` - `DataSourceTransactionManager`

- `JPA` - `JapTransactionManager`

- `Hibernate` - `HibernateTransactionManager`

- `JDO` - `JdoTransactionManager`

- `分布式事务` - `JtaTransactionManager`

> **编程式事务**

```java
public void test() {
      TransactionDefinition def = new DefaultTransactionDefinition();
      TransactionStatus status = transactionManager.getTransaction(def);
       try {
         // 事务操作
 
         // 事务提交
         transactionManager.commit(status);
      } catch (DataAccessException e) {
         // 事务回滚
         transactionManager.rollback(status);
         throw e;
      }
}
```

> **声明式事务**

```java
@Transactional
public void saveSomething(Long  id, String name) {
    //数据库操作
}
```

> `TransactionDefinition` 事务属性

- 隔离级别
- 传播行为
```java
public interface TransactionDefinition {
    ...
    // 当前存在事务，则加入该事务，否则创建一个新的事务
    int PROPAGATION_REQUIRED = 0; 
    // 存在事务，则加入事务，否则以非事务的方式继续运行
    int PROPAGATION_SUPPORTS = 1;
    // 当前存在事务，则加入，否则抛出异常
    int PROPAGATION_MANDATORY = 2;
    // 创建一个新的事务，如果当前存在事务，则将当前事务挂起
    int PROPAGATION_REQUIRES_NEW = 3; 
    // 以非事务方式运行，如果存在事务，则将当前事务挂起
    int PROPAGATION_NOT_SUPPORTED = 4;
    // 以非事务方式运行，如果存在事务，则抛出异常
    int PROPAGATION_NEVER = 5;
    // 当前存在事务，就在嵌套事务内执行，否则单独开启一个新的事务
    int PROPAGATION_NESTED = 6;
    ...
}
```
- 回滚规则
- 是否只读
- 事务超时

### 事务注解

> 如果类或者方法的 `public` 方法被标注 `@Transactional` ，Spring 容器在启动的时候为其创建一个代理类，方法调用时实际调用的是 `TransactionInterceptor` 类中的 `invoke()`，在目标方法之前开启事务，执行过程中如果遇到异常时回滚事务，方法调用完成后提交事务。

- `TransactionInterceptor` 的 `invoke()` 方法内部实际调用的是 `TransactionAspectSupport` 的 `invokeWithinTransaction()` 方法

## 2 动态数据源

## 参考

- [1] [Spring事务管理详解](https://juejin.cn/post/6844903608224333838)

- [2] [深入理解Spring事务原理](https://cloud.tencent.com/developer/article/1832182)

