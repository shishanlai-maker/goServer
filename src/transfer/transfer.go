package transfer

import (
	"encoding/json"
	"fmt"
	"manage"
	"net/http"
	"peoplelist"
)

func Handlerr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
func Registe(w http.ResponseWriter, b []byte) {
	fmt.Println("处理注册的逻辑")
	var requesrData manage.ResquestData2
	// //反序列化为
	err := json.Unmarshal([]byte(b), &requesrData)
	Handlerr(err)
	request := &manage.ResquestData2{
		UserName:     requesrData.UserName,
		UserAccount:  requesrData.UserAccount,
		UserPassword: requesrData.UserPassword,
	}
	fmt.Println("request=", request.UserName, request.UserPassword)
	data := request.Register(request)
	w.Write(data)
}

func Login(w http.ResponseWriter, b []byte) {
	fmt.Println("处理登录的逻辑")
	var requesrData manage.ResquestData2
	// //反序列化为
	err := json.Unmarshal([]byte(b), &requesrData)
	Handlerr(err)
	request := &manage.ResquestData2{
		UserName:     requesrData.UserName,
		UserAccount:  requesrData.UserAccount,
		UserPassword: requesrData.UserPassword,
	}
	data := request.Login(request.UserName, request.UserPassword)
	w.Write(data)
}

func Develop(w http.ResponseWriter, b map[string]string) {
	fmt.Println("处理查询开发列表逻辑")
	//处理返回结果,先将[]string转成json，在反序列化为[]byte
	requestMar, err := json.Marshal(b)
	var requesrData peoplelist.DevelopeRequese
	err = json.Unmarshal([]byte(requestMar), &requesrData)
	Handlerr(err)
	request := &peoplelist.DevelopeRequese{}
	// fmt.Println("request", request.BugNum, request.Name, request.WorkNum)
	data := request.Developfunc(requesrData.Name, requesrData.BugNum, requesrData.WorkNum, requesrData.PageSize, requesrData.PageNo)
	w.Write(data)
}

func InsertDeveloper(w http.ResponseWriter, b []byte) {
	fmt.Println("处理插入开发人员逻辑")
	var requesrData peoplelist.InsertDevelope
	// //反序列化为
	err := json.Unmarshal([]byte(b), &requesrData)
	Handlerr(err)
	request := &peoplelist.InsertDevelope{} //为了调用方法的
	data := request.Insertdeve(requesrData)
	w.Write(data)
}

func DeleteDeveloper(w http.ResponseWriter, b []byte) {
	fmt.Println("处理删除开发人员逻辑")
	var requesrData peoplelist.InsertDevelope
	// //反序列化为
	err := json.Unmarshal([]byte(b), &requesrData)
	Handlerr(err)
	request := &peoplelist.InsertDevelope{} //为了调用方法的
	data := request.Deletedeve(requesrData)
	w.Write(data)
}

func UpdataDeveloper(w http.ResponseWriter, b []byte) {
	fmt.Println("处理更新开发人员逻辑")
	var requesrData peoplelist.InsertDevelope
	// //反序列化为
	err := json.Unmarshal([]byte(b), &requesrData)
	Handlerr(err)
	request := &peoplelist.InsertDevelope{} //为了调用方法的
	data := request.Updatadeve(requesrData)
	w.Write(data)
}

func ProductHandle(w http.ResponseWriter, b []byte) {
	fmt.Println("处理产品创建需求逻辑")
	var requesrData peoplelist.ProductData
	// //反序列化为
	err := json.Unmarshal([]byte(b), &requesrData)
	Handlerr(err)
	request := &peoplelist.Product{} //为了调用方法的
	data := request.ProductFunc(requesrData)
	w.Write(data)
}
