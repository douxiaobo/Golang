package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	my_sql_password := os.Getenv("my_sql_password")
	if my_sql_password == "" {
		fmt.Println("my_sql_password enviroment variable is not, please set my_sql_password environment variable")
		return
	}
	db, err := sql.Open("mysql", "root:"+my_sql_password+"@tcp(localhost:3306)/testdb?charset=utf8")
	checkErr(err)
	defer db.Close()
	// stmt, err := db.Prepare("INSERT INTO employeetest(first_name,last_name,age,sex,income) VALUES(?,?,?,?,?)")	//OK
	stmt, err := db.Prepare("INSERT employeetest SET first_name=?,last_name=?,age=?,sex=?,income=?") //OK
	checkErr(err)
	res, err := stmt.Exec("John", "Doe", 30, "M", 50000) //panic: Error 1406 (22001): Data too long for column 'SEX' at row 1
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println("Last inserted ID is:", id)
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// douxiaobo@192 go-sql-driver % go mod init go-sql-driver
// go: creating new go.mod: module go-sql-driver
// go: to add module requirements and sums:
//         go mod tidy
// douxiaobo@192 go-sql-driver % go get github.com/go-sql-driver/mysql
// go: downloading github.com/go-sql-driver/mysql v1.8.1
// go: downloading filippo.io/edwards25519 v1.1.0
// go: added filippo.io/edwards25519 v1.1.0
// go: added github.com/go-sql-driver/mysql v1.8.1
// douxiaobo@192 go-sql-driver %

// douxiaobo@192 go-sql-driver % mysql.server status
//  ERROR! MySQL is not running, but PID file exists
// douxiaobo@192 go-sql-driver % mysql.server start
// Starting MySQL
//  SUCCESS!
// douxiaobo@192 go-sql-driver % mysql.server status
//  SUCCESS! MySQL running (15715)

// douxiaobo@192 go-sql-driver % mysql -u root -p
// Enter password:
// Welcome to the MySQL monitor.  Commands end with ; or \g.
// Your MySQL connection id is 9
// Server version: 8.3.0 Homebrew

// Copyright (c) 2000, 2024, Oracle and/or its affiliates.

// Oracle is a registered trademark of Oracle Corporation and/or its
// affiliates. Other names may be trademarks of their respective
// owners.

// Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

// mysql> show databases
//     -> ;
// +--------------------+
// | Database           |
// +--------------------+
// | information_schema |
// | mysql              |
// | performance_schema |
// | sys                |
// | testdb             |
// +--------------------+
// 5 rows in set (0.01 sec)

// mysql> use testdb;
// Reading table information for completion of table and column names
// You can turn off this feature to get a quicker startup with -A

// Database changed
// mysql> show tables
//     -> ;
// +------------------+
// | Tables_in_testdb |
// +------------------+
// | EMPLOYEETEST     |
// +------------------+
// 1 row in set (0.00 sec)

// mysql> exit
// Bye

// douxiaobo@192 go-sql-driver % go run insert.go
// Last inserted ID is: 0

// douxiaobo@192 go-sql-driver % mysql -u root -p
// Enter password:
// Welcome to the MySQL monitor.  Commands end with ; or \g.
// Your MySQL connection id is 13
// Server version: 8.3.0 Homebrew

// Copyright (c) 2000, 2024, Oracle and/or its affiliates.

// Oracle is a registered trademark of Oracle Corporation and/or its
// affiliates. Other names may be trademarks of their respective
// owners.

// Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

// mysql> use testdb
// Reading table information for completion of table and column names
// You can turn off this feature to get a quicker startup with -A

// Database changed
// mysql> select * from employeetest;
// +------------+-----------+------+------+--------+
// | FIRST_NAME | LAST_NAME | AGE  | SEX  | INCOME |
// +------------+-----------+------+------+--------+
// | Mac        | Mohan     |   20 | M    |   2000 |
// | John       | Doe       |   30 | M    |  50000 |
// +------------+-----------+------+------+--------+
// 2 rows in set (0.00 sec)

// OK

// douxiaobo@192 go-sql-driver % export my_sql_password='root'
