package manage

import (
	"encoding/json"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

//创建两个常量，保存方法的类型
const RegisterType = "register"
const LoginType = "login"

//接受登陆注册数据的结构体
type ResquestData2 struct {
	UserName     string `json:"userName"`
	UserAccount  int    `json:"userAccount"`
	UserPassword int    `json:"userPassword"`
}

//返回响应数据的结构体
type ResponseAllData struct {
	Code int               `json:"code"`
	Data map[string]string `json:"data"`
	Msg  string            `json:"msg"`
}

//定义一个全局的pool
var pool *redis.Pool

//redis链接池的操作，这样更省时间
func init() {
	pool = &redis.Pool{
		MaxIdle:     8,   //最大空闲连接数
		MaxActive:   0,   //表示和数据库的最大链接数，0表示没有限制
		IdleTimeout: 100, //表示最大空闲时间，100s后不使用久放到最大空闲连接数
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func (this *ResquestData2) Login(userName string, userPassword int) []byte {
	// fmt.Printf("userName=%s,userPassword=%d\n", userName, userPassword)
	//处理登陆的逻辑
	var responseall ResponseAllData
	var resdata []byte
	conn := pool.Get() //先从pool 取出一个链接
	defer conn.Close() //关闭这个链接，用完了就不要占着位置
	data, err := redis.String(conn.Do("Hget", "registerUserData", userName))
	// fmt.Println("data=", data)
	if err != nil {
		fmt.Println("redis.String err=", err)
	}
	if len(data) == 0 {
		responseall.Code = 301
		responseall.Msg = "不存在当前账号，请先注册吧～"
		resdata, err := json.Marshal(responseall)
		if err != nil {
			fmt.Println("json.Marshal(responseall) err=", err)
		}
		fmt.Println("不存在当前账号，请先注册吧～")
		return resdata
	} else {
		//校验密码是否正确
		var requesrData ResquestData2 //声明一个结构体，用来接受反序列化的数据
		//反序列化为
		err := json.Unmarshal([]byte(data), &requesrData)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println("反序列化的data=", requesrData)
		if requesrData.UserPassword != userPassword {
			responseall.Code = 302
			responseall.Msg = "密码不正确，请重新输入～"
			resdata, err := json.Marshal(responseall)
			if err != nil {
				fmt.Println("json.Marshal(responseall) err=", err)
			}
			fmt.Println("密码不正确，请重新输入～")
			return resdata
		} else if requesrData.UserPassword == userPassword {
			responseall.Code = 200
			responseall.Msg = "登陆成功"
			responseall.Data = map[string]string{"userName": userName}
			resdata, err := json.Marshal(responseall)
			if err != nil {
				fmt.Println("json.Marshal(responseall) err=", err)
			}
			fmt.Println("登陆成功")
			return resdata
		}
	}
	return resdata
}

func (this *ResquestData2) Register(requestdata *ResquestData2) []byte {
	//处理注册的逻辑
	var responseall ResponseAllData
	// var resdata []byte
	conn := pool.Get() //先从pool 取出一个链接
	defer conn.Close() //关闭这个链接，用完了就不要占着位置
	data, err := redis.String(conn.Do("Hget", "registerUserData", this.UserName))
	fmt.Println("data=", data)
	if len(data) == 0 { //证明没有这个用户可以完成注册
		_, err = conn.Do("HSet", "registerUserData", this.UserName, fmt.Sprintf("{\"UserAccount\":%d, \"userPassword\":%d}", this.UserAccount, this.UserPassword))
		if err != nil {
			fmt.Println("hset err=", err)
		}
		responseall.Code = 200
		responseall.Msg = "恭喜你注册完成"
		responseall.Data = map[string]string{"userName": requestdata.UserName}
		resdata, err := json.Marshal(responseall)
		if err != nil {
			fmt.Println("json.Marshal(responseall) err=", err)
		}
		fmt.Println("恭喜你注册完成")
		return resdata
	} else {
		responseall.Code = 201
		responseall.Msg = "当前用户已存在，请重新填写"
		responseall.Data = map[string]string{}
		resdata, err := json.Marshal(responseall)
		if err != nil {
			fmt.Println("json.Marshal(responseall) err=", err)
		}
		fmt.Println("当前用户已存在，请重新填写")
		return resdata
	}
}
