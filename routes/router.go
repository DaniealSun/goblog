package routes

import (
	v1 "goblog/api/v1"
	"goblog/middleware"
	"goblog/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	// 引入跨域cors
	r.Use(middleware.Cors())
	// 引入自定义日志
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())

	r.LoadHTMLGlob("static/admin/index.html")
	r.Static("admin/static", "static/admin/static")
	r.StaticFile("admin/favicon", "static/admin/favicon.ico")
	r.GET("admin", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	// 需要认证的路由
	auth := r.Group("api/v1")
	// JWT认证携带的token
	auth.Use(middleware.JwtToken())
	{
		// 用户模块路由接口
		auth.PUT("user/:id", v1.EditUser)           // 编辑用户
		auth.PUT("admin/redopass/:id", v1.RedoPass) // 重置密码
		auth.DELETE("user/:id", v1.DeleteUser)      // 删除用户
		// 分类模块路由接口
		auth.POST("category/add", v1.AddCate)      // 添加分类
		auth.PUT("category/:id", v1.EditCate)      // 编辑分类
		auth.DELETE("category/:id", v1.DeleteCate) // 删除分类
		// 文章模块路由接口
		auth.POST("article/add", v1.AddArt)      // 添加文章
		auth.PUT("article/:id", v1.EditArt)      // 编辑文章
		auth.DELETE("article/:id", v1.DeleteArt) // 删除文章
		// 上传文件
		auth.POST("upload", v1.Upload) //上传缩略图
	}
	// 无需认证的路由
	router := r.Group("api/v1")
	{
		router.POST("user/add", v1.AddUser)           // 添加用户
		router.GET("users", v1.GetUsers)              // 获取用户列表
		router.GET("user/:id", v1.GetUserInfo)        // 查询单个用户信息
		router.GET("category", v1.GetCate)            // 查询分类列表
		router.GET("category/:id", v1.GetCateInfo)    // 查询单个分类
		router.GET("article", v1.GetArt)              // 查询文章列表
		router.GET("article/list/:id", v1.GetCateArt) // 获取分类下所有文章
		router.GET("article/info/:id", v1.GetArtInfo) // 获取单个文章信息
		router.POST("login", v1.Login)                // 登录验证
	}
	r.Run(utils.HttpPort)
}
