/*******
* @Author:qingmeng
* @Description:
* @File:main
* @Date2021/12/2
 */

package main

import (
	"homework6/Mysql"
	"homework6/start"
)

func main() {
	Mysql.InitMysql()
	start.Routers()
}
