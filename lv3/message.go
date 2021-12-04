/*******
* @Author:qingmeng
* @Description:
* @File:message
* @Date2021/12/3
 */

package lv3

import (
	"errors"
	"fmt"
	"homework6/Mysql"
)

var (
	errMissReceiver =errors.New("没有找到邮件接收者" )
	errMissContent 	=errors.New("未找到邮件内容")
	errMissMessageUser=errors.New("没有找到该邮件id的发送者")
)

type message struct {
	messageId int `json:"messageId"`
	username string `json:"username"`
	content string `json:"content"`
	receiverName string `json:"receiverName"`
}


//接收信息处理
func ReceiveMessage(username string,receiver string,content string)error  {
	err:=CheckReceiver(receiver)
	if err!=nil{
		return err
	}
	if content==""{
		return errMissContent
	}
	stmt,err:=Mysql.DB.Prepare("insert into userandmessage.message(username, content,receiverName) VALUES (?,?,?)")
	if err!=nil{
		fmt.Printf("Prepare failed:%v",err)
		return err
	}
	defer stmt.Close()
	_,err=stmt.Exec(username,content,receiver)
	if err!=nil{
		fmt.Printf("insert failed:%v",err)
		return err
	}
	return err
}

//检验是否存在该接收者
func CheckReceiver(receiver string) error {
	stmt,err:=Mysql.DB.Query("select username from userandmessage.user where username=?",receiver)
	if err!=nil{
		fmt.Printf("Query failed:%v",err)
		return err
	}
	defer stmt.Close()
	var username string
	for stmt.Next(){
		err=stmt.Scan(&username)
		if err!=nil{
			fmt.Printf("Scan failed:%v",err)
			return err
		}
	}

	if username!=receiver{
		return errMissReceiver
	}
	return err
}

//回复信息处理
func Reply(username string,content string,messageId int) error {
	receiver:=FindUser(messageId)
	if receiver==""{
		fmt.Printf(errMissMessageUser.Error())
		return errMissMessageUser
	}
	stmt,err:=Mysql.DB.Prepare("insert into userandmessage.message (username, content, receiverName) values (?,?,?);")
	if err!=nil{
		fmt.Printf("Prepare failed:%v",err)
		return err
	}
	defer stmt.Close()
	_, err= stmt.Exec(username,content,receiver)
	if err != nil {
		fmt.Printf("Insert failed:%v",err)
		return err
	}
	return err
}

//通过messageId获取发送者
func FindUser(messageId int) string {
	stmt,err:=Mysql.DB.Query("select username from userandmessage.message where messageId=?",messageId)
	if err!=nil{
		fmt.Printf("Query failed:%v",err)
		return ""
	}
	defer stmt.Close()
	var m message
	for stmt.Next(){
		err=stmt.Scan(&m.username)
		if err!=nil{
			fmt.Printf("Scan failed:%v",err)
			return ""
		}
	}
	return m.username
}

//通过messageId查看信息
func ViewMessage(messageId int)string{
	stmt,err:=Mysql.DB.Query("select content from userandmessage.message where messageId=?",messageId)
	if err!=nil{
		fmt.Printf("Query failed:%v",err)
		return ""
	}
	defer stmt.Close()
	var m message
	for stmt.Next() {
		err=stmt.Scan(&m.content)
		if err!=nil{
			fmt.Printf("Scan failed:%v",err)
			return ""
		}
	}
	return m.content
}
