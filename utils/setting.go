package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	AppMode     string
	HttpPort    string
	Db          string
	DbHost      string
	DbPort      string
	DbUser      string
	DbPassWord  string
	DbName      string
	JwtKey      string
	AccessKey   string
	SecretKey   string
	Bucket      string
	QiniuServer string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径", err)
	}
	LoadServer(file)
	LoadData(file)
	LoadQiniu(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":8777")
	JwtKey = file.Section("server").Key("JwtKey").MustString("Usu84DWl9M")
}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("172.29.238.2")
	DbPort = file.Section("database").Key("DbPort").MustString("23306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("Usu84DWl_q2l_n9M")
	DbName = file.Section("database").Key("DbName").MustString("ginblog_szp")
}

func LoadQiniu(file *ini.File) {
	AccessKey = file.Section("qiniu").Key("AccessKey").String()
	SecretKey = file.Section("qiniu").Key("SecretKey").String()
	Bucket = file.Section("qiniu").Key("Bucket").String()
	QiniuServer = file.Section("qiniu").Key("QiniuServer").String()
}
