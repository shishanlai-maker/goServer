package main

import (
	"encoding/json"
	"io/ioutil"
	"manage"
	"net/http"
	"peoplelist"
	"transfer"
)

//创建一个方法结构体，来识别传入的方法类型
type FuncatonType struct {
	Funcation string `json:"funcation"`
}

//解决跨域问题
func setupCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func HandlerHttp(w http.ResponseWriter, r *http.Request) {
	setupCORS(&w)
	if r.Method == "OPTIONS" {
		return
	}
	b, _ := ioutil.ReadAll(r.Body)
	//声明一个结构体
	var requesrFunc FuncatonType
	//反序列化为
	requestSlice := make(map[string]string) //声明一个切片，类型为string，当为get请求时使用
	err := json.Unmarshal([]byte(b), &requesrFunc)
	if err != nil {
		//证明是get方法
		requesrFunc.Funcation = r.URL.Query().Get("funcation")
		for k := range r.URL.Query() {
			if k != "funcation" {
				requestSlice[k] = r.URL.Query().Get(k)
			}
		}
		// fmt.Println("requestSlice=", requestSlice)
	}
	//遍历走对应的逻辑
	switch requesrFunc.Funcation {
	case manage.RegisterType:
		//处理注册的逻辑
		transfer.Registe(w, b)
	case manage.LoginType:
		//处理登录的逻辑
		transfer.Login(w, b)
	case peoplelist.DevelopeSearchType:
		//处理开发列表
		transfer.Develop(w, requestSlice)
	case peoplelist.InsertDeveloperType:
		//处理新增开发人员逻辑
		transfer.InsertDeveloper(w, b)
	case peoplelist.DeleteDeveloperType:
		//处理删除开发人员逻辑
		transfer.DeleteDeveloper(w, b)
	case peoplelist.UpdataloperType:
		//处理删除开发人员逻辑
		transfer.UpdataDeveloper(w, b)
	case peoplelist.ProductNeedType:
		//处理产品创建需求
		transfer.ProductHandle(w, b)
	}

}

func main() {
	http.HandleFunc("/", HandlerHttp)
	http.ListenAndServe("127.0.0.1:9802", nil)
}
