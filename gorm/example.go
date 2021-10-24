package gorm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main()  {
	//hostname := viper.GetString("mysql.hostname")
	//user := viper.GetString("mysql.user")
	//port := viper.GetString("mysql.port")
	//password := viper.GetString("mysql.password")
	//db := viper.GetString("mysql.db")
	hostname := "localhost"
	user := "root"
	password := "123"
	port := "3306"
	db := "test"
	//https://gorm.io/zh_CN/docs/connecting_to_the_database.html
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		hostname,
		port,
		db,
	)

	mysqlDb,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err!= nil {
		fmt.Println(err)
	}
	mysqlDb.Select("")
}
