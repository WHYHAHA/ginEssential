package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
	"why/ginessential/common"
)


func main() {
	InitConfig()
	common.InitDB()

	r := gin.Default()
	r = CollectRoute(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	println(r.Run())
	 // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func InitConfig()  {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
