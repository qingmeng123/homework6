/*******
* @Author:qingmeng
* @Description:
* @File:test
* @Date2021/12/1
 */

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var(
	name="root"
	pwd="@XUEHUI."
	host="localhost"
	port="3306"
	dbname="student"
)

type info struct {
	id int
	name string
	age int
	password string
}

var db *sql.DB

func main() {
	initDb()
	QueryRowDemo()
	queryMultiRowDemo()
	updateRowDemo()
	queryMultiRowDemo()
	deleteRowDemo()
	prepareQueryDemo()
}

func initDb(){
	dsn:=name+":"+pwd+"@tcp("+host+":"+port+")/"+dbname
	d,err:=sql.Open("mysql",dsn)
	//d,err:=sql.Open("mysql","root:@XUEHUI.@(localhost:3306)/student")
	if err!=nil{
		log.Fatal(err)
		return
	}
	db=d
}

// QueryRowDemo 查询单条数据示例
func QueryRowDemo() {
	sqlStr:="select id,name,age from studentinfo where id = ?"
	var u info
	err:=db.QueryRow(sqlStr,1).Scan(&u.id,&u.name,&u.age)
	if err!=nil{
		fmt.Printf("scan failed,err:%v\n",err)
		return
	}
	fmt.Println(u)
}

// 查询多条数据示例
func queryMultiRowDemo(){
	sqlStr:="select id,name,age from studentinfo where id >?"
	rows,err:=db.Query(sqlStr,1)
	if err!=nil{
		fmt.Printf("query failed,err:%v\n",err)
		return
	}
	defer rows.Close()
	for rows.Next(){
		var u info
		err:=rows.Scan(&u.id,&u.name,&u.age)
		if err!=nil {
			fmt.Printf("scan failed,err:%v\n",err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n",u.id,u.name,u.age)
	}
}

//插入、更新、删除都会使用Exec()方法

// 更新数据
func updateRowDemo(){
	sqlStr:="update studentinfo set age=? where id =?"
	ret,err:=db.Exec(sqlStr,39,3)
	if err!=nil{
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n,err:=ret.RowsAffected()
	if err!=nil{
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

// 删除数据
func deleteRowDemo(){
	sqlStr:="delete from studentinfo where id =?"
	ret,err:=db.Exec(sqlStr,111)
	if err!=nil{
		fmt.Printf("delete failed,err:%v\n",err)
		return
	}
	n,err:=ret.RowsAffected()
	if err!=nil{
		fmt.Printf("get RowsAffected failed,err:%v\n",err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}

// 预处理查询示例
func prepareQueryDemo() {
	sqlStr := "select id, name, age from studentinfo where id > ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	// 循环读取结果集中的数据
	for rows.Next() {
		var u info
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
}