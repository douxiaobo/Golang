package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://douxiaobo:postgres@localhost/testdb?sslmode=disable")
	//postgres://<username>:<password>@<host>:<port>/<dbname>?sslmode=<ssl_mode>
	if err != nil {
		log.Fatal("Error connecting to database: %v", err)
		return
	}
	// _, err := db.Ping()
	// if err != nil {
	// 	fmt.Println("Error pinging database:", err)
	// 	return
	// }
	// fmt.Println("Connected to the database")
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO mytable (name, age) VALUES ($1, $2) returning id")
	if err != nil {
		log.Fatalf("Prepare:Error preparing statement: %v", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec("John Doe", 30)
	if err != nil {
		log.Fatalf("Exec:Error executing statement: %v", err)
		return
	}

	var lateInsertId int
	err = db.QueryRow("INSERT INTO mytable (name, age) VALUES ($1, $2) returning id", "Antiony Davis", 25).Scan(&lateInsertId)
	if err != nil {
		log.Fatalf("QueryRoow:Error executing statement: %v", err)
		return
	}
	fmt.Println("Late insert id:", lateInsertId)

	// _, err = db.Exec("INSERT INTO mytable (name, age) VALUES ($1, $2)", "John Doe", 30)
	// if err != nil {
	// 	fmt.Println("Error executing statement")
	// 	return
	// }

	fmt.Println("Record inserted successfully")
}

// douxiaobo@192 lib_pq % go mod init lib_pq
// go: creating new go.mod: module lib_pq
// douxiaobo@192 lib_pq % go get -u github.com/lib/pq
// go: downloading github.com/lib/pq v1.10.9
// go: added github.com/lib/pq v1.10.9
// douxiaobo@192 lib_pq %

// 安装PostgreSQL以后不能使用PostgreSQL很困难。

// douxiaobo@192 lib_pq % go run insert.go
// Error preparing statement
// douxiaobo@192 lib_pq % sudo systemctl status postgresql
// Password:
// sudo: systemctl: command not found
// douxiaobo@192 lib_pq % psql -h localhost -U douxiaobo -d testdb -c "INSERT INTO mytable (name, age) VALUES ('John Doe', 30);"
// ERROR:  relation "mytable" does not exist
// 第1行INSERT INTO mytable (name, age) VALUES ('John Doe', 30);
//                  ^
// douxiaobo@192 lib_pq % go run insert.go
// Error preparing statement
// douxiaobo@192 lib_pq % psql -h localhost -U douxiaobo -d testdb -c "CREATE TABLE mytable (name VARCHAR(100), age INT);"
// CREATE TABLE
// douxiaobo@192 lib_pq % go run insert.go
// Record inserted successfully
// douxiaobo@192 lib_pq % psql -h localhost -U douxiaobo -d testdb -c "INSERT INTO mytable (name, age) VALUES ('John Doe', 30);"
// INSERT 0 1
// douxiaobo@192 lib_pq %

// douxiaobo@192 lib_pq % /opt/homebrew/Cellar/postgresql@16/16.3/bin/initdb --locale=C -E UTF-8 /opt/homebrew/var/postgr
// 属于此数据库系统的文件宿主为用户 "douxiaobo".
// 此用户也必须为服务器进程的宿主.
// 数据库簇将使用本地化语言 "C"进行初始化.
// 缺省的文本搜索配置将会被设置到"english"

// 禁止为数据页生成校验和.

// 创建目录 /opt/homebrew/var/postgr ... 成功
// 正在创建子目录 ... 成功
// 选择动态共享内存实现 ......posix
// 选择默认最大联接数 (max_connections) ... 100
// 选择默认共享缓冲区大小 (shared_buffers) ... 128MB
// 选择默认时区 ... Asia/Shanghai
// 创建配置文件 ... 成功
// 正在运行自举脚本 ...成功
// 正在执行自举后初始化 ...成功
// 同步数据到磁盘...成功

// initdb: 警告: 为本地连接启用"trust"身份验证
// initdb: hint: You can change this by editing pg_hba.conf or using the option -A, or --auth-local and --auth-host, the next time you run initdb.

// 成功。您现在可以用下面的命令开启数据库服务器：

//     '/opt/homebrew/Cellar/postgresql@16/16.3/bin/pg_ctl' -D /opt/homebrew/var/postgr -l 日志文件 start

// douxiaobo@192 lib_pq % /opt/homebrew/Cellar/postgresql@16/16.3/bin/pg_ctl -D /opt/homebrew/var/postgr -l 日 志文件 restart
// pg_ctl: PID 文件 "/opt/homebrew/var/postgr/postmaster.pid" 不存在
// 服务器进程是否正在运行?
// 尝试启动服务器进程
// pg_ctl: 无法读取文件 "/opt/homebrew/var/postgr/postmaster.opts"
// douxiaobo@192 lib_pq % /opt/homebrew/Cellar/postgresql@16/16.3/bin/pg_ctl -D /opt/homebrew/var/postgr -l /path/to/your/logfile start

// 等待服务器进程启动 ..../bin/sh: /path/to/your/logfile: No such file or directory
//  已停止等待
// pg_ctl: 无法启动服务器进程
// 检查日志输出.
// douxiaobo@192 lib_pq % cd ~
// douxiaobo@192 ~ % touch postgres.log
// douxiaobo@192 ~ % /opt/homebrew/Cellar/postgresql@16/16.3/bin/pg_ctl -D /opt/homebrew/var/postgr -l ~/postgres.log start

// 等待服务器进程启动 .... 完成
// 服务器进程已经启动
// douxiaobo@192 ~ % psql postgres
// psql (16.3 (Homebrew))
// 输入 "help" 来获取帮助信息.

// postgres=# create database testdb;
// CREATE DATABASE
// postgres=# use testdb;
// ERROR:  syntax error at or near "use"
// 第1行use testdb;
//      ^
// postgres=# create user postgres with password 'postgres'
// postgres-# grant all privileges on database testdb to postgres
// postgres-# grant connect on database testdb to postgres
// postgres-# create table mytable(name text,age integer);
// ERROR:  syntax error at or near "grant"
// 第2行grant all privileges on database testdb to postgres
//      ^
// postgres=# grant connect on database testdb to postgres;
// ERROR:  role "postgres" does not exist
// postgres=# create table mytable(name text,age integer);
// CREATE TABLE
// postgres=# select * from mytable;
//  name | age
// ------+-----
// (0 行记录)

// douxiaobo@192 lib_pq % pg_ctl -D /opt/homebrew/var/postgr start
// 等待服务器进程启动 ....2024-06-01 18:06:25.007 CST [5338] LOG:  starting PostgreSQL 16.3 (Homebrew) on aarch64-apple-darwin23.4.0, compiled by Apple clang version 15.0.0 (clang-1500.3.9.4), 64-bit
// 2024-06-01 18:06:25.009 CST [5338] LOG:  listening on IPv6 address "::1", port 5432
// 2024-06-01 18:06:25.009 CST [5338] LOG:  listening on IPv4 address "127.0.0.1", port 5432
// 2024-06-01 18:06:25.009 CST [5338] LOG:  listening on Unix socket "/tmp/.s.PGSQL.5432"
// 2024-06-01 18:06:25.013 CST [5341] LOG:  database system was shut down at 2024-05-31 21:50:15 CST
// 2024-06-01 18:06:25.017 CST [5338] LOG:  database system is ready to accept connections
//  完成
// 服务器进程已经启动

// douxiaobo@192 lib_pq % go run insert.go
// 2024/06/01 18:34:57 Exec:Error executing statement: pq: null value in column "id" of relation "mytable" violates not-null constraint
// exit status 1
// douxiaobo@192 lib_pq % go run insert.go
// 2024/06/01 18:37:27 Exec:Error executing statement: pq: null value in column "id" of relation "mytable" violates not-null constraint
// exit status 1
// douxiaobo@192 lib_pq % go run insert.go
// Late insert id: 2
// Record inserted successfully

// douxiaobo@192 lib_pq % psql postgres
// psql: 错误: 连接到套接字"/tmp/.s.PGSQL.5432"上的服务器失败:No such file or directory
// 	Is the server running locally and accepting connections on that socket?
// douxiaobo@192 lib_pq % psql postgres douxiaobo
// psql: 错误: 连接到套接字"/tmp/.s.PGSQL.5432"上的服务器失败:No such file or directory
// 	Is the server running locally and accepting connections on that socket?
// douxiaobo@192 lib_pq % pg_ctl -D postgr -l start
// pg_ctl: 没有指定操作
// 请用 "pg_ctl --help" 获取更多的信息.
// douxiaobo@192 lib_pq % pg_ctl -D /usr/local/var/postgres start
// pg_ctl: 目录 "/usr/local/var/postgres" 不存在
// douxiaobo@192 lib_pq % pg_ctl
// pg_ctl: 没有指定操作
// 请用 "pg_ctl --help" 获取更多的信息.
// douxiaobo@192 lib_pq % pg_ctl -D /opt/homebrew/var/postgr start
// 等待服务器进程启动 ....2024-06-01 18:06:25.007 CST [5338] LOG:  starting PostgreSQL 16.3 (Homebrew) on aarch64-apple-darwin23.4.0, compiled by Apple clang version 15.0.0 (clang-1500.3.9.4), 64-bit
// 2024-06-01 18:06:25.009 CST [5338] LOG:  listening on IPv6 address "::1", port 5432
// 2024-06-01 18:06:25.009 CST [5338] LOG:  listening on IPv4 address "127.0.0.1", port 5432
// 2024-06-01 18:06:25.009 CST [5338] LOG:  listening on Unix socket "/tmp/.s.PGSQL.5432"
// 2024-06-01 18:06:25.013 CST [5341] LOG:  database system was shut down at 2024-05-31 21:50:15 CST
// 2024-06-01 18:06:25.017 CST [5338] LOG:  database system is ready to accept connections
//  完成
// 服务器进程已经启动
// douxiaobo@192 lib_pq % launchctl load ~/Library/LaunchAgents/com.postgresql.postgres.plist
// Load failed: 5: Input/output error
// Try running `launchctl bootstrap` as root for richer errors.
// douxiaobo@192 lib_pq % psql postgres
// psql (16.3 (Homebrew))
// 输入 "help" 来获取帮助信息.

// postgres=# \l
//                                                         数据库列表
//    名称    |  拥有者   | 字元编码 | Locale Provider | 校对规则 | Ctype | ICU Locale | ICU Rules |        存取权限
// -----------+-----------+----------+-----------------+----------+-------+------------+-----------+-------------------------
//  postgres  | douxiaobo | UTF8     | libc            | C        | C     |            |           |
//  template0 | douxiaobo | UTF8     | libc            | C        | C     |            |           | =c/douxiaobo           +
//            |           |          |                 |          |       |            |           | douxiaobo=CTc/douxiaobo
//  template1 | douxiaobo | UTF8     | libc            | C        | C     |            |           | =c/douxiaobo           +
//            |           |          |                 |          |       |            |           | douxiaobo=CTc/douxiaobo
//  testdb    | douxiaobo | UTF8     | libc            | C        | C     |            |           |
// (4 行记录)

// postgres=# \c testdb
// 您现在已经连接到数据库 "testdb",用户 "douxiaobo".
// testdb=# drop table mytable
// testdb-# ;
// DROP TABLE
// testdb=# create table mytable(
// testdb(# uid serial not null,
// testdb(# name character vary2024-06-01 18:11:25.083 CST [5339] LOG:  checkpoint starting: time
// ing2024-06-01 18:11:25.798 CST [5339] LOG:  checkpoint complete: wrote 10 buffers (0.1%); 0 WAL file(s) added, 0 removed, 0 recycled; write=0.708 s, sync=0.003 s, total=0.715 s; sync files=8, longest=0.001 s, average=0.001 s; distance=40 kB, estimate=40 kB; lsn=0/19E8420, redo lsn=0/19E83E8

// testdb(# create table mytable(
// testdb(# uid serial not null,
// testdb(# name character varying(100) not null,
// testdb(# age integer,
// testdb(# uid int primary key not null,
// testdb(# );
// testdb(#
// testdb(# quit
// 使用\q 退出.
// testdb(# \q
// douxiaobo@192 lib_pq % psql postgres
// psql (16.3 (Homebrew))
// 输入 "help" 来获取帮助信息.

// postgres=# \d testdb;
// 没有找到任何名称为 "testdb" 的关联.
// postgres=# \c testdb;
// 您现在已经连接到数据库 "testdb",用户 "douxiaobo".
// testdb=# select * from mytable;
// 2024-06-01 18:14:26.770 CST [5460] ERROR:  relation "mytable" does not exist at character 15
// 2024-06-01 18:14:26.770 CST [5460] STATEMENT:  select * from mytable;
// ERROR:  relation "mytable" does not exist
// 第1行select * from mytable;
//                    ^
// testdb=# \d
// 没有找到任何关系.
// testdb=# create table mytable(
// testdb(# id int primary key not null,
// testdb(# name text not null,
// testdb(# age int not null);
// CREATE TABLE
// testdb=# 2024-06-01 18:16:25.802 CST [5339] LOG:  checkpoint starting: time
// 2024-06-01 18:16:29.444 CST [5339] LOG:  checkpoint complete: wrote 37 buffers (0.2%); 0 WAL file(s) added, 0 removed, 0 recycled; write=3.634 s, sync=0.005 s, total=3.643 s; sync files=32, longest=0.001 s, average=0.001 s; distance=152 kB, estimate=152 kB; lsn=0/1A0E6C0, redo lsn=0/1A0E688
// \d
//                 关联列表
//  架构模式 |  名称   |  类型  |  拥有者
// ----------+---------+--------+-----------
//  public   | mytable | 数据表 | douxiaobo
// (1 行记录)

// testdb=# \d mytable
//            数据表 "public.mytable"
//  栏位 |  类型   | 校对规则 |  可空的  | 预设
// ------+---------+----------+----------+------
//  id   | integer |          | not null |
//  name | text    |          | not null |
//  age  | integer |          | not null |
// 索引：
//     "mytable_pkey" PRIMARY KEY, btree (id)

// testdb=# 2024-06-01 18:23:01.831 CST [5754] FATAL:  role "postgres" does not exist
// \d mytable
//            数据表 "public.mytable"
//  栏位 |  类型   | 校对规则 |  可空的  | 预设
// ------+---------+----------+----------+------
//  id   | integer |          | not null |
//  name | text    |          | not null |
//  age  | integer |          | not null |
// 索引：
//     "mytable_pkey" PRIMARY KEY, btree (id)

// testdb=# select * from mytable;
//  id | name | age
// ----+------+-----
// (0 行记录)

// testdb=# 2024-06-01 18:25:03.061 CST [5831] FATAL:  role "postgres" does not exist
// 2024-06-01 18:26:06.272 CST [5910] FATAL:  role "postgres" does not exist
// 2024-06-01 18:26:18.862 CST [5974] FATAL:  role "postgre" does not exist
// 2024-06-01 18:26:39.243 CST [6043] ERROR:  null value in column "id" of relation "mytable" violates not-null constraint
// 2024-06-01 18:26:39.243 CST [6043] DETAIL:  Failing row contains (null, John Doe, 30).
// 2024-06-01 18:26:39.243 CST [6043] STATEMENT:  INSERT INTO mytable (name, age) VALUES ($1, $2) RETURNING id
// select * from mytable;
//  id | name | age
// ----+------+-----
// (0 行记录)

// testdb=# 2024-06-01 18:34:57.772 CST [6413] ERROR:  null value in column "id" of relation "mytable" violates not-null constraint
// 2024-06-01 18:34:57.772 CST [6413] DETAIL:  Failing row contains (null, John Doe, 30).
// 2024-06-01 18:34:57.772 CST [6413] STATEMENT:  INSERT INTO mytable (name, age) VALUES ($1, $2) RETURNING id
// 2024-06-01 18:37:27.345 CST [6503] ERROR:  null value in column "id" of relation "mytable" violates not-null constraint
// 2024-06-01 18:37:27.345 CST [6503] DETAIL:  Failing row contains (null, John Doe, 30).
// 2024-06-01 18:37:27.345 CST [6503] STATEMENT:  INSERT INTO mytable (name, age) VALUES ($1, $2)
// ALTER TAB
// testdb=# ALTER TABLE mytable ALTER COLUMN id SET DATA TYPE SERIAL;
// 2024-06-01 18:38:36.597 CST [5460] ERROR:  type "serial" does not exist
// 2024-06-01 18:38:36.597 CST [5460] STATEMENT:  ALTER TABLE mytable ALTER COLUMN id SET DATA TYPE SERIAL;
// ERROR:  type "serial" does not exist
// testdb=# drop mytable
// testdb-# create table mytable(
// testdb(# id serial not null,
// testdb(# name text nut null,
// testdb(# age int null);
// 2024-06-01 18:40:34.099 CST [5460] ERROR:  syntax error at or near "mytable" at character 6
// 2024-06-01 18:40:34.099 CST [5460] STATEMENT:  drop mytable
// 	create table mytable(
// 	id serial not null,
// 	name text nut null,
// 	age int null);
// ERROR:  syntax error at or near "mytable"
// 第1行drop mytable
//           ^
// testdb=# DROP TABLE IF EXISTS mytable;
// CREATE TABLE mytable (
//     id SERIAL NOT NULL,
//     name TEXT NOT NULL,
//     age INT
// );
// DROP TABLE
// CREATE TABLE
// testdb=# 2024-06-01 18:41:25.417 CST [5339] LOG:  checkpoint starting: time
// 2024-06-01 18:41:29.280 CST [5339] LOG:  checkpoint complete: wrote 39 buffers (0.2%); 0 WAL file(s) added, 0 removed, 0 recycled; write=3.853 s, sync=0.005 s, total=3.864 s; sync files=29, longest=0.001 s, average=0.001 s; distance=158 kB, estimate=158 kB; lsn=0/1A360E8, redo lsn=0/1A360B0

// testdb=# select * from mytable;
//  id |     name      | age
// ----+---------------+-----
//   1 | John Doe      |  30
//   2 | Antiony Davis |  25
// (2 行记录)

// testdb=#
