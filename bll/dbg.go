package bll

import (
	. "ManualApi/zfx"
	"fmt"
)

// 调试相关
type DbgController struct {
	BaseController
}

//https://beego.me/docs/advantage/docs.md

//type MyDebug struct {
//	codes string
//}

// @Title		函数标题
// @Description	函数描述
// @Param		accessKey query       string  true        "参数描述"
// @Param		args      formData        string  true        "参数描述"
// @Success		200       {object}    []orm.ParamsList
// @router		/MyDebug [post]
func (c *DbgController) MyDebug() {
	return
	Try(func() {
		Log.Println("MyDebug()")
		accessKey := c.GetString("accessKey")
		ai := Rbac.AccessKeyIsValid(accessKey)
		fmt.Println(ai)
		if ai.Role != UserRoles_Manager {
			Throw("DeniedAccess")
		}
		//		var args MyDebug
		//		if err := c.ParseBody(&args); err != nil {
		//			Throw(err.Error())
		//		}
		//		if err := c.IsValid(args); err != nil {
		//			Throw(err.Error())
		//		}
		//		// TODO
		//		result := Db.Query(args.codes)
		//		c.JsonData(result)
		codes := c.GetString("args")
		fmt.Println("codes:", codes)
		result := Db.Query(codes)
		c.JsonData(result)

		c.Return()
	}, Catch(func(err error) {
		if err.Error() != "" {
			c.Error(err.Error())
		}
	}))
}
