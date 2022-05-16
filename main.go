package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"os"
	"simpleTikTok/DB"
	"simpleTikTok/router"
)

func main() {
	InitConfig()
	db := DB.InitDB()
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			panic("failed to close db, err +" + err.Error())
		}
	}(db)

	r := gin.Default()
	// r.MaxMultipartMemory = 64 << 8 这里可以设置 context 最大的文件大小，默认为32 << 8

	router.InitRouter(r)

	// 读取配置文件中的ip和port
	ip := viper.GetString("server.ip")
	port := viper.GetString("server.port")
	if ip != "" && port != "" {
		panic(r.Run(ip + ":" + port))
	}
	panic(r.Run())
}

// InitConfig 获得application.yml的位置
func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
