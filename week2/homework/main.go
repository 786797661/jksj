package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

var Db *sql.DB

//数据库配置
const (
	userName = "yzc"
	password = "123456"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "MySql002"
)

type sqlError struct {
	sqlStr string
	err    error
	res    interface{}
}

func main() {
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	database, err := sql.Open("mysql", path)
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
	err = Query("select username from userinfo where id =?", 1, Db)
	te, _ := err.(*sqlError)
	if te.err != nil {
		//处理除了sql.ErrNoRows的错误
	} else if te.res == nil {
		//处理没有数据的情况
	} else if te.res != nil {
		//处理正常数据
	}
	defer Db.Close()
}

func Query(sqlStr string, args interface{}, Db *sql.DB) error {
	err := sqlFacade(sqlStr, args, Db)
	return err
}

func (e *sqlError) UnWrap() error {
	return e.err
}

func (e *sqlError) Error() string {
	return e.err.Error()
}

func sqlFacade(sqlStr string, args interface{}, db *sql.DB) error {
	var res interface{}
	err := db.QueryRow(sqlStr, args).Scan(&res)
	//如果是sql.ErrNoRows 降级返回nil
	if errors.Is(err, sql.ErrNoRows) {
		err = nil
	}
	return &sqlError{sqlStr, err, res}
}
