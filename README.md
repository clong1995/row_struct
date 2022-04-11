# 将数据库查询结果转化为结构体

## 使用方法

### 安装
` go get github.com/clong1995/row_struct`
### 说明

**定义要查询结果的结构体**

```go
type field struct {
    Id       int64     `field:"id"`
    Name     string    `field:"name"`
    Birthday time.Time `field:"birthday"`
    Alive    bool      `field:"alive"`
    Height   float64   `field:"height"`
    Emoji    string    `field:"emoji"`
}
```  

`field`**为sql查询结果的字段名**

**数据库和go类型对应关系**  

| golang类型  | mysql类型                                  | 
|-----------|------------------------------------------|
| int64     | TINYINT,SMALLINT,INT,BIGINT / (UNSIGNED) | 
| string    | VARCHAR,TEXT                             | 
| float64   | DOUBLE                                   | 
| bool      | 数字,字符串,NULL                              | 
| time.Time | DATE,DATETIME                            | 

### 例子
**表结构**
```mysql
CREATE TABLE `test` (
  `id` bigint unsigned NOT NULL DEFAULT '0',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `birthday` datetime(3) NOT NULL DEFAULT '1000-01-01 00:00:00.000',
  `alive` tinyint unsigned NOT NULL DEFAULT '1',
  `height` double(5,2) unsigned NOT NULL DEFAULT '0.00',
  `emoji` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
```
**表数据**
```mysql
INSERT INTO `test`.`test` (`id`, `name`, `birthday`, `alive`, `height`, `emoji`) VALUES (1506556359976423424, '十一', '1995-09-19 13:46:28.123', 1, 175.55, '😄');
```
**代码**
```go
package main

import (
	"database/sql"
	"github.com/clong1995/row_struct"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func main() {
	//打开数据库链接
	dataSource := "root:密码@tcp(localhost)/test"
	db, err := sql.Open("mysql", dataSource+"?charset=utf8mb4&loc=Asia%2FShanghai&parseTime=true&multiStatements=true")
	if err != nil {
		log.Println(err)
		return
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}(db)

	//struct字段和查询结果的映射关系
	type field struct {
		Id       int64     `field:"id"`
		Name     bool      `field:"name"`
		Birthday time.Time `field:"birthday"`
		Alive    bool      `field:"alive"`
		Height   float64   `field:"height"`
		Emoji    string    `field:"emoji"`
	}

	//查询
	rows, err := db.Query(
		"SELECT id,name,birthday,alive,height,emoji FROM test WHERE id = ? LIMIT 1",
		1506556359976423424,
	)
	if err != nil {
		log.Println(err)
		return
	}

	// 读取
	person := new(field)
	for rows.Next() {
		err = row_struct.Scan(rows, person)
		if err != nil {
			log.Println(err)
			return
		}
		break
	}

	log.Printf("%+v", person)
}
```

### 结果

```
{
    Id:1506556359976423424
    Name:十一 
    Birthday:1995-09-19 13:46:28.123 +0800 CST 
    Alive:true 
    Height:175.55 
    Emoji:😄
}
```