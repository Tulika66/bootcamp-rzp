package Config


import (
	"fmt"
	"gorm.io/gorm"
)

var Db *gorm.DB


type DBConfig struct {
	Host string
	Port int
	User string
	DBName string
	Password string
}

func BuildConfig()*DBConfig{
	dbconfig:=DBConfig{
		Host:     "localhost",
		Port:     3306,
		User:     "root",
		Password: "tulikamanek",
		DBName:   "commerce",
	}
	return &dbconfig
}

func DBUrl(dbConfig *DBConfig) string{
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,

	)
}