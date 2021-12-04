/*******
* @Author:qingmeng
* @Description:
* @File:handler
* @Date2021/12/2
 */

package start

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"homework6/lv2"
	"homework6/lv3"
	"net/http"
	"strconv"
)

func Login(ctx *gin.Context) {
	username:=ctx.PostForm("username")
	password:=ctx.PostForm("password")
	res:=lv2.Login(username,password)
	if res==nil{
		ctx.SetCookie("username",ctx.PostForm("username"),3600,"/","", false, true)
		ctx.JSON(http.StatusOK,gin.H{
			"message":"欢迎您~"+ctx.PostForm("username"),
		})
	} else {
		ctx.JSON(http.StatusForbidden,gin.H{
			"message":res.Error(),
		})
	}
}

func Register(ctx *gin.Context) {
	username:=ctx.PostForm("username")
	password:=ctx.PostForm("password")
	res:=lv2.Register(username,password)
	if res==nil{
		ctx.JSON(http.StatusOK,gin.H{
			"message":"注册成功",
		})
	}else{
		ctx.JSON(http.StatusForbidden,gin.H{
			"message":res.Error(),
		})
	}
}

//鉴权中间键
func Auth(ctx *gin.Context)  {
	value,err:=ctx.Cookie("username")
	if err!=nil{
		ctx.JSON(http.StatusForbidden,gin.H{
			"message":"认证失败,没有cookie,请先登陆",
		})
		ctx.Abort()
	}else{
		ctx.Set("cookie",value)
	}
}

//增加密保
func AddHighSchool(ctx *gin.Context)  {
	cookie,_:=ctx.Get("cookie")
	username:=cookie.(string)
	highSchool:=ctx.PostForm("highSchool")
	res:=lv2.AddHighSchool(username,highSchool)
	if res==nil{
		ctx.JSON(http.StatusOK,gin.H{
			"message":"成功将"+highSchool+"设为密保问题高中学校的答案",
		})
	}else {
		ctx.JSON(http.StatusForbidden,gin.H{
			"message":res.Error(),
		})
	}
}

//通过密保更改密码
func ChangePassword(ctx *gin.Context)  {
	username:=ctx.PostForm("username")
	highSchool:=ctx.PostForm("highSchool")
	password:=ctx.PostForm("password")
	res:=lv2.ChangePassword(username,highSchool,password)
	if res==nil{
		ctx.JSON(http.StatusOK,gin.H{
			"message":"更改密码成功",
		})
	}else{
		ctx.JSON(http.StatusForbidden,gin.H{
			"message":res.Error(),
		})
	}
}

//发送信息处理
func SendMessage(ctx *gin.Context) {
	cookie,_:=ctx.Get("cookie")
	username:=cookie.(string)
	receiverName:=ctx.PostForm("receiverName")
	content:=ctx.PostForm("content")
	res:=lv3.ReceiveMessage(username,receiverName,content)
	if res==nil{
		ctx.JSON(http.StatusOK,gin.H{
			"message":"发送邮件成功",
		})
	}else{
		ctx.JSON(http.StatusOK,gin.H{
			"message":res.Error(),
		})
	}
}

//回复信息
func Reply(ctx *gin.Context)  {
	cookie,_:=ctx.Get("cookie")
	username:=cookie.(string)
	messageId:=ctx.PostForm("messageId")
	id,err:=strconv.Atoi(messageId)
	if err!=nil{
		fmt.Printf("Atoi failed:%v",err)
		return
	}
	content:=ctx.PostForm("content")
	res:=lv3.Reply(username,content,id)
	if res==nil{
		ctx.JSON(http.StatusOK,gin.H{
			"message":"回复留言成功",
		})
	}else{
		ctx.JSON(http.StatusOK,gin.H{
			"message":res.Error(),
		})
	}
}

//通过messageId查看信息
func ViewMessage(ctx *gin.Context) {
	messageId:=ctx.PostForm("messageId")
	id,err:=strconv.Atoi(messageId)
	if err!=nil{
		fmt.Printf("Atoi failed:%v",err)
		return
	}
	message:=lv3.ViewMessage(id)
	if message!=""{
		ctx.JSON(http.StatusOK,gin.H{
			"message":message,
		})
	}else{
		ctx.JSON(http.StatusOK,gin.H{
			"message":"没有信息",
		})
	}

}