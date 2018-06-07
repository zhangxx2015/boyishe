package main

import (
	_ "ManualApi/routers" // 必须
	. "ManualApi/zfx"
	"fmt"

	"github.com/astaxie/beego" // 必须
)

func test() {
	go func() {
		AbsTime.Sleep(1)
		Try(func() {
			fmt.Println("")
			resp, err := ApiClient.Post("/v1/users/RegValidatePhone", Dict{
				"LoginName": "13593869486",
				"Password":  "123456",
				"UserRole":  1,
			})
			if err != nil {
				Throw(err.Error())
			}
			fmt.Println(resp.Data)
		}, func(err error) {
			println("err:" + err.Error())
		})

	}()
}


func main() {

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.SetStaticPath(AppSettingString("urouter"), AppSettingString("ufolder"))
	beego.Run()
}
