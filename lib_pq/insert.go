package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://douxiaobo:postgres@localhost/testdb?sslmode=disable")
	//postgres://<username>:<password>@<host>:<port>/<dbname>?sslmode=<ssl_mode>
	if err != nil {
		fmt.Println("Error connecting to database")
		return
	}
	// _, err := db.Ping()
	// if err != nil {
	// 	fmt.Println("Error pinging database:", err)
	// 	return
	// }
	// fmt.Println("Connected to the database")
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO mytable (name, age) VALUES ($1, $2) ")
	if err != nil {
		fmt.Println("Error preparing statement")
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec("John Doe", 30)
	if err != nil {
		fmt.Println("Error executing statement")
		return
	}

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
