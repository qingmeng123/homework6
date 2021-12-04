/*******
* @Author:qingmeng
* @Description:
* @File:user
* @Date2021/12/2
 */

package lv2

import (
	"errors"
	"fmt"
	"homework6/Mysql"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
	HighSchool string `json:"highSchool"`
}

//错误信息
var (
	errPassword = errors.New("password error")
	errRepeatedAccount=errors.New("repeated account")
	errPasswordLen=errors.New("密码长度应大于6位")
	errAnswer=errors.New("密保答案错误")
	errMissUsername=errors.New("用户名错误，没有该用户")
	errMissAnswer=errors.New("未设密保")
)

//登陆处理
func Login(username string,password string) error {
	stmt,err:=Mysql.DB.Query("select username,password from userandmessage.user where username=?",username)
	if err!=nil{
		fmt.Printf("query failed:%v", err)
		return err
	}
	defer stmt.Close()
	var u user
	for stmt.Next(){
		err=stmt.Scan(&u.Username,&u.Password)
		if err!=nil{
			fmt.Printf("scan failed:%v",err)
			return err
		}
	}
	if u.Password!=password{
		return errPassword
	}
	return err
}

//注册处理
func Register(username string,password string)error {
	err:=CheckRegister(username,password)
	if err!=nil{
		return err
	}
	stmt,err:=Mysql.DB.Prepare("insert into userandmessage.user(username, password)values (?,?)")
	if err!=nil{
		fmt.Printf("mysql prepare failed:%v",err)
		return err
	}
	defer stmt.Close()
	_,err= stmt.Exec(username,password)
	if err != nil {
		fmt.Printf("insert failed:%v",err)
		return err
	}
	return err
}

// CheckRegister 验证注册
func CheckRegister(username string, password string)error {
	if len(password)<=6{
		return 	errPasswordLen
	}
		stmt,err:=Mysql.DB.Query("select username from userandmessage.user")
	if err!=nil{
		fmt.Printf("query failed:%v", err)
		return err
	}
	defer stmt.Close()
	var u user
	for stmt.Next(){
		err=stmt.Scan(&u.Username)
		if err!=nil{
			fmt.Printf("query failed:%v", err)
			return err
		}
		if u.Username==username{
			return errRepeatedAccount
		}
	}
	return nil
}



//增加密保信息处理
func AddHighSchool(username string,highSchool string)error{
	stmt,err:=Mysql.DB.Prepare("update userandmessage.user set highSchool=? where username=?")
	if err!=nil{
		fmt.Printf("mysql prepare failed:%v",err)
		return err
	}
	defer stmt.Close()
	_,err=stmt.Exec(highSchool,username)
	if err!=nil{
		fmt.Printf("update failed:%v",err)
		return err
	}
	return err
}

//通过用户名和密保修改密码处理
func ChangePassword(username string,highSchool string,password string)error  {
	res:=isExist(username,highSchool)
	if res!=nil{
		return res
	}
	if len(password)<=6{
		return 	errPasswordLen
	}

	stmt,err:=Mysql.DB.Prepare("update userandmessage.user set password=? where username=? and highSchool=?")
	if err!=nil{
		fmt.Printf("mysql prepare failed:%v",err)
		return err
	}
	defer stmt.Close()
	_,err=stmt.Exec(password,username,highSchool)
	if err!=nil{
		fmt.Printf("update failed:%v",err)
		return err
	}
	return err
}

//验证用户名和密保
func isExist(username string,highSchool string) error {
	stmt,err:=Mysql.DB.Query("select username,highSchool from userandmessage.user where username=?",username)
	if err!=nil{
		fmt.Printf("query failed:%v",err)
		return err
	}
	defer stmt.Close()
	var u user
	for stmt.Next(){
		err := stmt.Scan(&u.Username,&u.HighSchool)
		if err != nil {
			fmt.Printf("scan failed:%v",err)
			return errMissAnswer
		}
	}
	if u.Username!=username{
		return errMissUsername
	}
	if u.HighSchool!=highSchool{
		return errAnswer
	}
	return nil
}
