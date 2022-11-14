# Spring Transaction management

## 基本原理

> `Spring` 事务的本质其实就是数据库对事务的支持。对于纯JDBC操作数据库，想要用到事务，可以按照以下步骤进行：

1. 获取连接 `Connection con = DriverManager.getConnection()`
2. 开启事务 `con.setAutoCommit(true/false);`
3. 执行 `CRUD`
4. 提交事务/回滚事务 `con.commit() / con.rollback();`
5. 关闭连接 `conn.close();`

> 通过Spring事务管理，可以自动完成2，4：
___

> Spring Boot自动开启了对注解事务的支持

1. 在相关的类和方法上添加 `@Transactional`
2. Spring 启动时解析生成相关 `bean` 并生成代理，根据`@Transactional`的相关参数进行配置注入，在代理中开启正常提交事务以及异常回滚事务
3. 数据库的事务提交和回滚是通过`binlog`和`redo log`实现


> `Spring` 的事务机制提供了 `PlatformTransactionManager` 接口，不同的数据访问技术的事务使用不同的借口实现


`JDBC` - `DataSourceTransactionManager`
`JPA` - `JapTransactionManager`
`Hibernate` - `HibernateTransactionManager`
`JDO` - `JdoTransactionManager`
`分布式事务` - `JtaTransactionManager`

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
