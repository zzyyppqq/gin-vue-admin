package main

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"
	"time"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// 这部分 @Tag 设置用于排序, 需要排序的接口请按照下面的格式添加
// swag init 对 @Tag 只会从入口文件解析, 默认 main.go
// 也可通过 --generalInfo flag 指定其他文件
// @Tag.Name        Base
// @Tag.Name        SysUser
// @Tag.Description 用户

// @title                       Gin-Vue-Admin Swagger API接口文档
// @version                     v2.8.0
// @description                 使用gin+vue进行极速开发的全栈开发基础平台
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
func main() {
	// 记录开始时间
	start := time.Now()
	global.GVA_VP = core.Viper() // 初始化Viper
	//configJson, _ := json.MarshalIndent(global.GVA_CONFIG, "", "  ")
	//fmt.Println("global.GVA_CONFIG: ", string(configJson))

	initialize.OtherInit()
	global.GVA_LOG = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	initialize.Timer()
	initialize.DBList()
	if global.GVA_DB != nil {
		// 迁移表耗时很长，大概1分钟，若不改变表结构，则无须执行
		//initialize.RegisterTables() // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	// 计算构建时间
	fmt.Printf("初始化时间: %s\n", time.Since(start))
	core.RunWindowsServer()

}

func testMongodb() {
	global.GVA_VP = core.Viper() // 初始化Viper

	// 初始化mongodb
	err := initialize.Mongo.Initialization()
	if err != nil {
		zap.L().Error(fmt.Sprintf("%+v", err))
	}
}
