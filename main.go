package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"product-trace-server/Controllor"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("error when reading config")
	}
	_port := viper.GetString("PORT")


	engine := gin.Default()
	engine.GET( "/ping" ,  func (ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message" :  "pong" ,
		})
	})
	engine.GET( "/checkID" ,  Controllor.HandleCheckUserID)
	engine.GET( "/fullToponym" ,  Controllor.HandleGetFullToponym)
	engine.GET( "/createUnit" ,  Controllor.HandleCreateUnit)
	engine.GET( "/getUnit" ,  Controllor.HandleGetUnit)

	engine.Run(_port)

}
