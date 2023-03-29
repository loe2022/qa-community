package api

import "github.com/gin-gonic/gin"

func InitRouter() {
	r := gin.Default()
	//设置路由引擎
	u := r.Group("/user")
	{
		//主页面，当浏览器发出/user/home的请求时，执行root函数
		u.GET("/home", root)
		//注册,当浏览器发出/user/register的请求时，执行Register函数
		u.POST("/register", Register)
		//登录,当浏览器发出/user/login的请求时，执行Login函数
		u.POST("/login", Login)
		//找回密码，当浏览器发出/user/forget请求时，执行Forget函数
		u.GET("/forget", Forget)
	}

	m := r.Group("/message")
	{
		//获得一条留言,查看评论
		m.GET("/message", ShowMessage)
		//发送一条留言,发出一个问题或者回答一个问题
		m.POST("/leave_message", LeaveMessage)
		//更新一条留言,修改自己之前发出的内容，答案或者问题
		m.PUT("/alter_message", AlterMessage)
		//删除一条留言，删除自己的回答或者回复
		m.DELETE("/del_message", DeleteMessage)
	}

	_ = r.Run()
}
