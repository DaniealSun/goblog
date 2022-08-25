package main

import (
	"goblog/model"
	"goblog/routes"
)

func main() {
	// 引用数据
	model.InitDb()
	// 引用路由
	routes.InitRouter()
}
