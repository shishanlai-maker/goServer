package utils

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type SqlConn struct {
	SqlDB *sql.DB
}

func errorHandler(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

// 建立数据库连接
func (this SqlConn) SetupConnect() *sql.DB {
	db, err := sql.Open("mysql", "xwjmysql:!xi5jie@com#@(47.99.200.102:3306)/test")
	errorHandler(err)
	return db
}

// 查询数据
func (this SqlConn) Query(db *sql.DB, sql string) *sql.Rows {
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println(err)
	}
	return rows
}

// 插入数据
func (this SqlConn) Insert(db *sql.DB, sql string) {
	db.Exec(sql)
}

// 删除记录
func (this SqlConn) Delete(db *sql.DB, sql string) {
	db.Exec(sql)
}

// 修改数据
func (this SqlConn) Update(db *sql.DB, sql string) {
	db.Exec(sql)
	fmt.Println("修改成功")
}

// 删除表
func (this SqlConn) DeleteTable(db *sql.DB, sql string) {
	db.Exec(sql)
}

func (this SqlConn) QuerySumCount(dataName string, wherelateSql string) int {
	sql := fmt.Sprintf(`SELECT COUNT(1) FROM %s WHERE %s;`, dataName, wherelateSql)
	var mysqlConn SqlConn
	conn := mysqlConn.SetupConnect()
	rowsNum := mysqlConn.Query(conn, sql)
	var allNum int
	for rowsNum.Next() {
		if err := rowsNum.Scan(&allNum); err != nil {
			fmt.Println(err)
		}
	}
	return allNum
}

//处理分页逻辑，返回的字符串，拼接sql使用
func HandlePaging(pageSize string, pageNo string) string {
	pagesize, _ := strconv.Atoi(pageSize)
	pageno, _ := strconv.Atoi(pageNo)
	sqlPage := fmt.Sprintf(" LIMIT %d,%d", (pageno-1)*pagesize, pagesize)
	return sqlPage
}
