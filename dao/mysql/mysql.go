package mysql

import (
	"fmt"

	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func Init() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", viper.GetString("mysql.user"), viper.GetString("mysql.password"), viper.GetString("mysql.host"), viper.GetInt("mysql.port"), viper.GetString("mysql.dbname"))
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	//sqlx.MustConnect("mysql",dsn) //会直接返回panic。。不需要自己做err判断
	db.SetMaxOpenConns(viper.GetInt("max_open_conn"))
	db.SetMaxIdleConns(viper.GetInt("max_idle_conn"))
	return
}

func Close() {
	_ = db.Close()
}
