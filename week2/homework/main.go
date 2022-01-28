package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	pkgerrors "github.com/pkg/errors"
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

func main() {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	database, err := sql.Open("mysql", path)
	//database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
	err = Query("select username from scholar_userinfo where id =?", Db)
	fmt.Printf("original error : %T %v\n", pkgerrors.Cause(err), pkgerrors.Cause(err))
	defer Db.Close() // 注意这行代码要写在上面err判断的下面
}

func Query(sqlStr string, Db *sql.DB) error {
	err := sqlFacade(sqlStr, 1, Db)
	return err
}

type sqlError struct {
	sqlStr string
	err    error
	res    interface{}
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
	if errors.Is(err, sql.ErrNoRows) {
		err = nil
	}
	return &sqlError{sqlStr, err, res}
}
