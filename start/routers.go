/*******
* @Author:qingmeng
* @Description:
* @File:routers
* @Date2021/12/2
 */

package start

import "github.com/gin-gonic/gin"

func Routers() {
	router:=gin.Default()
	router.GET("/login",Login)
	router.POST("/register",Register)
	router.GET("/AddSecurity",Auth,AddHighSchool)
	router.POST("/ChangePassword",ChangePassword)
	router.POST("/SendMessage",Auth,SendMessage)
	router.POST("/Reply",Auth,Reply)
	router.GET("/ViewMessage",ViewMessage)
	router.Run()
}
