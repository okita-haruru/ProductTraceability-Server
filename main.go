package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"product-trace-server/Controllor"
	"product-trace-server/Service"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("error when reading config")
	}
	PORT := viper.GetString("PORT")
	LogFileLocation :=viper.GetString("LOG_FILE_LOCATION")

	gin.SetMode(gin.ReleaseMode)

	engine := gin.Default()
	engine.Use(CORSMiddleware())
	engine.GET( "/ping" ,  func (ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message" :  "pong" ,
		})
	})

	log := logrus.New()
	log.Out = os.Stdout
	log.Level = 4 // Info

	logfile, err := os.OpenFile(LogFileLocation, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("failed to open file for log")
	} else {
		log.Out = logfile
		log.Formatter = &logrus.JSONFormatter{}
	}
	ts := Service.NewTraceService(
		log,
	)

	/*
		Initialize Controllers
	*/
	traceController := Controllor.NewTraceController(log,ts)


	engine.GET( "/checkID" ,  traceController.HandleCheckUserID)
	engine.GET( "/toponym" ,  traceController.HandleGetFullToponym)
	engine.POST( "/unit" ,  traceController.HandleCreateUnit)
	engine.GET( "/unit" ,  traceController.HandleGetUnit)
	engine.GET("/record" , traceController.HandleTransRecord)
	engine.GET("/time" , traceController.HandleTransTime)

	engine.Run(PORT)

}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, X-Auth-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
