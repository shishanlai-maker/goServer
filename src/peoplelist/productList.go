package peoplelist

import (
	"encoding/json"
	"fmt"
	"utils"
)

const ProductNeedType = "productNeed"

type ProductData struct {
	Funcation string `json:"funcation"`
	Params    Product
}

type Product struct {
	Name      string `json:"name"`
	Version   string `json:"version"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	NeedType  string `json:"needType"`
	Desc      string `json:"desc"`
}

//返回结果的结构体
type ProductResponseData struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (this *Product) ProductFunc(product ProductData) []byte {
	//完成搜索可初始化的逻辑
	var mysqlConn utils.SqlConn
	conn := mysqlConn.SetupConnect()
	//准备sql
	insertSql := fmt.Sprintf("INSERT INTO `product_list`(`name`, `version`, `startTime`, `endTime`, `needType`,`desc`) VALUES ('%s','%s','%s','%s','%s','%s')", product.Params.Name,
		product.Params.Version, product.Params.StartTime, product.Params.EndTime, product.Params.NeedType, product.Params.Desc)
	//开始插入数据
	// fmt.Println("sql=", insertSql)
	mysqlConn.Insert(conn, insertSql)
	//查询如果有就是成功，没有就是失败
	looksql := fmt.Sprintf("SELECT name FROM product_list WHERE name = '%s';", product.Params.Name)
	rows := mysqlConn.Query(conn, looksql)
	var errInsertData bool = true           //定义errInsertData，如果插入成功,还是true，没成功就改为false
	var responseallData ProductResponseData //定义返回数据的结构体
	for rows.Next() {
		if err := rows.Scan(&product.Params.Name); err != nil {
			errInsertData = false
			fmt.Println("插入数据失败～")
		}
	}
	if !errInsertData {
		//表示插入数据失败
		responseallData.Code = 204
		responseallData.Msg = "插入需求失败～"
	} else {
		//插入数据成功
		responseallData.Code = 200
		responseallData.Msg = "插入需求成功～"
	}
	abc, err := json.Marshal(responseallData)
	errorHandler(err)
	return abc
}
