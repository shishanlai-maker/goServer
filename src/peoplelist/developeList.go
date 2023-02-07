package peoplelist

import (
	"encoding/json"
	"fmt"
	"utils"

	_ "github.com/go-sql-driver/mysql"
)

const DevelopeSearchType = "developeSearch"
const InsertDeveloperType = "insertDeveloper"
const DeleteDeveloperType = "developerDelete"
const UpdataloperType = "updataDeveloper"

type DevelopeResponse struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Age          int    `json:"age"`
	WorktimeLong int    `json:"worktimeLong"`
	BugNum       int    `json:"bugNum"`
	WorkNum      int    `json:"workNum"`
}

type InsertDevelope struct {
	Funcation string `json:"funcation"`
	Params    DevelopeResponse
}

//返回结果的结构体
type DevelopeResponseData struct {
	Code     int                `json:"code"`
	Data     []DevelopeResponse `json:"data"`
	Msg      string             `json:"msg"`
	TotalNum int                `json:"totalNum"`
}

//处理搜索的结构体
type DevelopeRequese struct {
	Name     string `json:"name"`
	BugNum   string `json:"bugNum"`
	WorkNum  string `json:"workNum"`
	PageSize string `json:"pageSize"`
	PageNo   string `json:"pageNo"`
}

func errorHandler(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (this DevelopeRequese) Developfunc(name string, bugNum string, workNum string, pageSize string, pageNo string) []byte {
	//完成搜索可初始化的逻辑
	var mysqlConn utils.SqlConn
	conn := mysqlConn.SetupConnect()
	//sql拼接
	where := "1"
	if name != "" {
		conditionA := fmt.Sprintf(" and name = '%s'", name)
		where += conditionA
	}
	if bugNum != "" {
		conditionB := fmt.Sprintf(" and bugNum = '%s'", bugNum)
		where += conditionB
	}
	if workNum != "" {
		conditionC := fmt.Sprintf(" and workNum = '%s'", workNum)
		where += conditionC
	}
	allNum := mysqlConn.QuerySumCount("develope_list", where) //查询得到返回数据的数量
	//处理分页
	sqlPage := utils.HandlePaging(pageSize, pageNo)
	where += sqlPage
	//sql整合
	var QUERY_DATA = fmt.Sprintf(`SELECT * FROM develope_list WHERE %s;`, where)
	developeResponse := DevelopeResponse{}
	//开始查询
	rows := mysqlConn.Query(conn, QUERY_DATA)
	var dataList []DevelopeResponse
	for rows.Next() {
		if err := rows.Scan(&developeResponse.Id, &developeResponse.Name, &developeResponse.Age, &developeResponse.WorktimeLong, &developeResponse.BugNum, &developeResponse.WorkNum); err != nil {
			fmt.Println(err)
		}
		dataList = append(dataList, developeResponse)
	}
	//需要返回切片
	var responseallData DevelopeResponseData
	responseallData.Code = 200
	responseallData.Msg = "查询成功～"
	responseallData.TotalNum = allNum
	responseallData.Data = dataList
	abc, err := json.Marshal(responseallData)
	errorHandler(err)
	return abc
}

func (this InsertDevelope) Insertdeve(developres InsertDevelope) []byte {
	//完成搜索可初始化的逻辑
	var mysqlConn utils.SqlConn
	conn := mysqlConn.SetupConnect()
	//准备sql
	insertSql := fmt.Sprintf("INSERT INTO develope_list(name,age,worktimeLong,bugNum,workNum) VALUES('%s','%d','%d','%d','%d');", developres.Params.Name, developres.Params.Age, developres.Params.WorktimeLong, developres.Params.BugNum, developres.Params.WorkNum)
	//开始插入
	mysqlConn.Insert(conn, insertSql)
	//查询如果有就是成功，没有就是失败
	looksql := fmt.Sprintf("SELECT name FROM develope_list WHERE name = '%s';", developres.Params.Name)
	rows := mysqlConn.Query(conn, looksql)
	var errInsertData bool = true            //定义errInsertData，如果插入成功,还是true，没成功就改为false
	var responseallData DevelopeResponseData //定义返回数据的结构体
	for rows.Next() {
		if err := rows.Scan(&developres.Params.Name); err != nil {
			errInsertData = false
			fmt.Println("插入数据失败～")
		}
	}
	if !errInsertData {
		//表示插入数据失败
		responseallData.Code = 204
		responseallData.Msg = "插入开发人员失败～"
	} else {
		//插入数据成功
		responseallData.Code = 200
		responseallData.Msg = "插入开发人员成功～"
	}
	abc, err := json.Marshal(responseallData)
	errorHandler(err)
	return abc
}
func (this InsertDevelope) Deletedeve(developres InsertDevelope) []byte {
	//完成搜索可初始化的逻辑
	var mysqlConn utils.SqlConn
	conn := mysqlConn.SetupConnect()
	//准备sql
	insertSql := fmt.Sprintf("DELETE FROM `develope_list` WHERE `develope_list`.`id` = '%d'", developres.Params.Id)
	//开始删除
	mysqlConn.Delete(conn, insertSql)
	//查询如果有就是成功，没有就是失败
	looksql := fmt.Sprintf("SELECT name FROM develope_list WHERE id = '%d';", developres.Params.Id)
	rows := mysqlConn.Query(conn, looksql)
	var errDeleteData bool = false           //定义errInsertData，如果删除成功,还是false，没成功就改为true
	var responseallData DevelopeResponseData //定义返回数据的结构体
	for rows.Next() {
		if err := rows.Scan(&developres.Params.Id); err != nil {
			errDeleteData = true
			fmt.Println("删除数据失败～")
		}
	}
	if errDeleteData {
		//表示删除数据成功
		responseallData.Code = 200
		responseallData.Msg = "删除开发人员成功～"
	} else {
		//删除数据失败
		responseallData.Code = 204
		responseallData.Msg = "删除开发人员失败～"
	}
	abc, err := json.Marshal(responseallData)
	errorHandler(err)
	return abc
}

func (this InsertDevelope) Updatadeve(developres InsertDevelope) []byte {
	//完成搜索可初始化的逻辑
	var mysqlConn utils.SqlConn
	conn := mysqlConn.SetupConnect()
	//准备sql
	updataSql := fmt.Sprintf("UPDATE develope_list SET name='%s',age='%d',worktimeLong='%d',bugNum='%d',workNum='%d' WHERE id = '%d';", developres.Params.Name, developres.Params.Age, developres.Params.WorktimeLong, developres.Params.BugNum, developres.Params.WorkNum, developres.Params.Id)
	//开始更新
	mysqlConn.Update(conn, updataSql)
	var responseallData DevelopeResponseData //定义返回数据的结构体
	responseallData.Code = 200
	responseallData.Msg = "更新开发人员失败～"
	abc, err := json.Marshal(responseallData)
	errorHandler(err)
	return abc
}
