/*******
* @Author:qingmeng
* @Description:
* @File:InitMysql
* @Date2021/12/2
 */

package Mysql

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

var(
	name="root"
	pwd="@XUEHUI."
	host="localhost"
	port="3306"
	dbname="userandmessage"
)

func InitMysql() {
	dsn:= name +":"+ pwd +"@tcp("+ host +":"+ port +")/"+ dbname
	db,err:=sql.Open("mysql",dsn)
	if err!=nil{
		fmt.Printf("mysql connect failed:%v", err)
		return
	}
	DB =db
}
