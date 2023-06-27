package mysql

import (
	"fmt"
	"web_app/settings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func Init(cfg *settings.MysqlConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DbName)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	//sqlx.MustConnect("mysql",dsn) //会直接返回panic。。不需要自己做err判断
	db.SetMaxOpenConns(cfg.MaxOpenConn)
	db.SetMaxIdleConns(cfg.MaxIdleConn)
	return
}

func Close() {
	_ = db.Close()
}
