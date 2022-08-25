# goblog

├─ .gitignore
│  go.mod // 项目依赖
│  go.sum
│  main.go //主程序
│  README.md
│          
├─api         
├─config // 项目配置入口   
├─log  // 项目日志
├─middleware  // 中间件
├─model // 数据模型层
├─routes
│   └─router.go // 路由入口    
├─static // 打包静态文件
│  ├─admin  // 后台管理页面 (已废弃，打包静态文件在web/admin/dist下)         
│  └─front  // 前端展示页面 (待开发) 
├─upload   // 接到七牛云
├─utils // 项目公用工具库
│  │  setting.go 
│  ├─errmsg   
│  └─validator         
└─web // 前端开发源码（VUECLI项目源文件)
    ├─admin             
    └─front  // 前端（待开发）
