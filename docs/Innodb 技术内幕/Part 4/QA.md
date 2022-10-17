> `1` **主键的生成方式**

```mysql
-- create
mysql > create table t (a int primary key);
-- alter
mysql > alter table t add primary key(a);
```

	bigint `8byte`

> `2` **last_update_time and 索引**

	- 根据时间范围查询
	- 方便分库分表

> `3` **mysql charset=utf8mb4**

	`utf8` 字符长度 = 3 字节，`utfmb4` 字符长度 = 4 byte，能够表示更多的字符

> `4` **collate=`utf8mb4_bin`**

	排序规则 utf8mb4_bin 比较字符的二进制格式`binary`，区分大小写
	`8..`: utf8mb4_0900_ai_ci

```sql
mysql> show variables like 'collation%';
+----------------------+--------------------+
| Variable_name        | Value              |
+----------------------+--------------------+
| collation_connection | utf8mb4_0900_ai_ci |
| collation_database   | utf8mb4_0900_ai_ci |
| collation_server     | utf8mb4_0900_ai_ci |
+----------------------+--------------------+
3 rows in set (0.01 sec)
```

> `5` **存储引擎**

	Innodb 支持 `行锁设计`，`支持外键`，支持类似Oracle的`非锁定读`MVCC，默认`读操作不会产生锁`
	MyISAM、Memory、NDB...

> `6` **禁止使用外键**

	（1）性能影响：修改表数据需要到外键对应的表查询数据
	（2）并发问题：在使用外键的情况下，每次修改数据都需要去另外一个表检查数据,需要获取额外的锁。若
				是在高并发大流量事务场景，使用外键更容易造成死锁。
	（3）扩展性问题：在水平拆分和分库的情况下，外键是无法生效的
	（4）技术问题：将应用程序应该执行的判断逻辑转移到了数据库上

> `7` **Row_format=dynamic**

	feat:把记录中数据量过大的字段值 全部存储到溢出页中
	其他表格式 compact、redundant、compressed

> `8` **禁止主表使用TEXT/BLOB 以及大字段优化**

	（1）溢出的数据不再存储在B + tree中
	（2）溢出的数据使用的是解压缩BLOB页面，并且存储独享
	（3）如果有多个大字段，尽量序列化后，存储在同一列中，避免多次off-page
	（4）将文等大字段从主表中拆分出来，a）存储到键值中b）存储在单独的一张子表中，压缩并且必须保证一
	行记录小于8K

> `9` **`尽量避免NULL` (from `高性能mysql`)**

		很多表都包含可为NULL（空值）的列，即使应用程序并不需要保存NULL也是如此，这是因为可为NULL
	是列的默认属性。通常情况下最好指定列为NOT NULL，除非真的需要存储NULL值。

		如果查询中包含可为NULL的列，对MySql来说更难优化，因为可为NULL的列使得索引、索引统计和值
	比较都更复杂。可为NULL的列会使用更多的存储空间，在MySql里也需要特殊处理。当可为NULL的列被索
	引时，每个索引记录需要一个额外的字节，在MyISAM里甚至还可能导致固定大小的索引（例如只有一个整
	数列的索引）变成可变大小的索引。

		通常把可为NULL的列改为NOT NULL带来的性能提升比较小，所以（调优时）没有必要首先在现有
	schema中查找并修改掉这种情况，除非确定这会导致问题。但是，如果计划在列上建索引，就应该尽量避免
	设计成可为NULL的列。


> `10` **decimal 浮点数**

	float，double为非标准类型，在DB中保存的是近似值，而Decimal则以字符串的形式保存数值。


