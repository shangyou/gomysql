# Golang-MySQL-Opreration

A Golang-MySQL-Opreration for Go's [database/sql](http://golang.org/pkg/database/sql) package

![Go-MySQL-Driver logo](https://raw.github.com/wiki/go-sql-driver/mysql/gomysql_m.png "Golang Gopher holding the MySQL Dolphin")

**Current tagged Release:** Version 1.1 (November 02, 2013)

[![Build Status](https://travis-ci.org/go-sql-driver/mysql.png?branch=master)](https://travis-ci.org/go-sql-driver/mysql) *(master branch)*

---------------------------------------

## Requirements
  * Go 1.1 or higher (use [v1.0](https://github.com/go-sql-driver/mysql/tags) for Go 1.0.x)
  * MySQL (Version 4.1 or higher), MariaDB or Percona Server

---------------------------------------

## Installation
Simple install the package to your [$GOPATH](http://code.google.com/p/go-wiki/wiki/GOPATH "GOPATH") with the [go tool](http://golang.org/cmd/go/ "go command") from shell:
```bash
$ go get github.com/shangyou/gomysql
```
Make sure [Git is installed](http://git-scm.com/downloads) on your machine and in your system's `PATH`.

*`go get` installs the latest tagged release*

## Usage
_Go MySQL Driver_ is an implementation of Go's `database/sql/driver` interface. You only need to import the driver and can use the full [`database/sql`](http://golang.org/pkg/database/sql) API then.

Use `mysql` as `driverName` and a valid [DSN](#dsn-data-source-name)  as `dataSourceName`:
```go
package main

import (
    "github.com/shangyou/gomysql"
    "fmt"
    //"time"
)

func main() {
    mysql := wsy.Mysql{nil}
    err := mysql.Init("root:@tcp(127.0.0.1:3306)/online?charset=utf8")
    fmt.Println(err)

    //created_at := time.Now().Unix()
    //id, err := mysql.Insert("insert into user set username = ?, password = ?, created_at = ?", "shangyou", "e10adc3949ba59abbe56e057f20f883e", created_at)
    //fmt.Println(id, err)

    //count, _ := mysql.Update("update user set enable='Y'")
    //fmt.Println(count)

    //count, err := mysql.Delete("delete from user where id = 2")
    //fmt.Println(count, err)

    //rs := mysql.Fetchrows("select * from user where id = 1")

    rs := mysql.Fetchrows("select * from user")
    fmt.Println(rs)
}
```

#### DNS Examples
```
user@unix(/path/to/socket)/dbname
```

```
root:pw@unix(/tmp/mysql.sock)/myDatabase?loc=Local
```

```
user:password@tcp(localhost:5555)/dbname?tls=skip-verify&autocommit=true
```

TCP via IPv6:
```
user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname?timeout=90s
```

TCP on a remote host, e.g. Amazon RDS:
```
id:password@tcp(your-amazonaws-uri.com:3306)/dbname
```

Google Cloud SQL on App Engine:
```
user@cloudsql(project-id:instance-name)/dbname
```

TCP using default port (3306) on localhost:
```
user:password@tcp/dbname&charset=utf8mb4,utf8&sys_var=esc%40ped
```

Use the default protocol (tcp) and host (localhost:3306):
```
user:password@/dbname
```

No Database preselected:
```
user:password@/
```
[more](https://github.com/go-sql-driver/mysql)

---------------------------------------

## License
Go-MySQL-Driver is licensed under the [Mozilla Public License Version 2.0](https://raw.github.com/go-sql-driver/mysql/master/LICENSE)

Mozilla summarizes the license scope as follows:
> MPL: The copyleft applies to any files containing MPLed code.


That means:
  * You can **use** the **unchanged** source code both in private as also commercial
  * You **needn't publish** the source code of your library as long the files licensed under the MPL 2.0 are **unchanged**
  * You **must publish** the source code of any **changed files** licensed under the MPL 2.0 under a) the MPL 2.0 itself or b) a compatible license (e.g. GPL 3.0 or Apache License 2.0)

Please read the [MPL 2.0 FAQ](http://www.mozilla.org/MPL/2.0/FAQ.html) if you have further questions regarding the license.

You can read the full terms here: [LICENSE](https://raw.github.com/go-sql-driver/mysql/master/LICENSE)

![Go Gopher and MySQL Dolphin](https://raw.github.com/wiki/go-sql-driver/mysql/go-mysql-driver_m.jpg "Golang Gopher transporting the MySQL Dolphin in a wheelbarrow")

