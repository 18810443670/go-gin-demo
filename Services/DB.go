package Services

import (
	"database/sql"
	"fmt"
	"github.com/go-ini/ini"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// 定义数据库链接
var DB *sql.DB

func InitMysql(ConfigIni *ini.File) {
	fmt.Println("InitMysql....")
	if DB == nil {
		DB_HOST := ConfigIni.Section("mysql").Key("DB_HOST").String()
		DB_PORT := ConfigIni.Section("mysql").Key("DB_PORT").String()
		DB_DATABASE := ConfigIni.Section("mysql").Key("DB_DATABASE").String()
		DB_USERNAME := ConfigIni.Section("mysql").Key("DB_USERNAME").String()
		DB_PASSWORD := ConfigIni.Section("mysql").Key("DB_PASSWORD").String()
		fmt.Println(DB_USERNAME + ":" + DB_PASSWORD + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_DATABASE)
		DB, _ = sql.Open("mysql", DB_USERNAME+":"+DB_PASSWORD+"@tcp("+DB_HOST+":"+DB_PORT+")/"+DB_DATABASE)
		log.Println(DB)
		if DB == nil {
			panic("Mysql init fail")
		}

		err2 := DB.Ping()
		if err2 != nil {
			panic(err2)
		}
		fmt.Println("Successfully connected!")
	}

}

// 操作数据库
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := DB.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}

func InsertDB(sql string, args ...any) (sql.Result, error) {
	//预编译
	stmt, err := DB.Prepare(sql)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	//执行数据
	res, err2 := stmt.Exec(args...)
	if err2 != nil {
		log.Println(err2)
		return nil, err2
	}
	return res, nil
}

// 查询
func QueryRowDB(sql string, args ...any) *sql.Row {
	return DB.QueryRow(sql, args...)
}
