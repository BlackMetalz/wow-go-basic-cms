package config

import (
	"fmt"
	"github.com/spf13/viper"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func GetMysqlConfig() (db *sql.DB) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper_err := viper.ReadInConfig()

	if viper_err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", viper_err))
	}

	server := viper.GetString("database.server")
	port := viper.GetString("database.port")
	dbName := viper.GetString("database.dbName")
	user := viper.GetString("database.user")
	password := viper.GetString("database.password")

	connectStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",user,password,server,port,dbName)
	fmt.Println("Connection String: ", connectStr)
	db, err := sql.Open("mysql", connectStr)
	if err != nil {
		fmt.Println(err)
	}
	return

}
