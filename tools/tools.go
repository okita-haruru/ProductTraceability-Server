package tools

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"product-trace-server/Model"
)

var _db *gorm.DB

//包初始化函数，golang特性，每个包初始化的时候会自动执行init函数，这里用来初始化gorm。
func init() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	_err := viper.ReadInConfig()
	if _err != nil {
		fmt.Println("Error when reading config")
		return
	}
	_user := viper.GetString("DB_USER")
	_dbName := viper.GetString("DB_NAME")
	_pass := viper.GetString("DB_PASSWORD")
	_DbUrl := viper.GetString("DB_URL")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", _user, _pass, _DbUrl, _dbName)



	var err error

	_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库失败, error=" + err.Error())
	}

	sqlDB, _ := _db.DB()
	_db.AutoMigrate(Model.ProductUnit{})
	sqlDB.SetMaxOpenConns(100)   //连接池最大连接数
	sqlDB.SetMaxIdleConns(20)   //最大允许的空闲连接数
}

func GetDB() *gorm.DB {
	return _db
}
