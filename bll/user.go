package bll

import (
	sms "ManualApi/sdk"
	. "ManualApi/zfx"
	"fmt"

	"github.com/astaxie/beego"
)

var vcodeLength int = beego.AppConfig.DefaultInt("vcodeLength", 6)
var freezeSeconds int = beego.AppConfig.DefaultInt("freezeSeconds", 60)
var tokenExpires int = beego.AppConfig.DefaultInt("tokenExpires", 3600)

// 用户相关
type UserController struct {
	BaseController
}

//https://beego.me/docs/advantage/docs.md
/*
type Sample struct {
	LoginName string `valid:"Mobile "`
}

// @Title		函数标题
// @Description	函数描述
// @Param		args      body        bll.Sample  true        "参数描述"
// @Success		200       {object}    utils.Resp
// @router		/Sample [post]
func (c *UserController) Sample() {
	Try(func() {
		Log.Println("Sample()")
		var args Sample
		if err := c.ParseBody(&args); err != nil {
			Throw(err.Error())
		}
		if err := c.IsValid(args); err != nil {
			Throw(err.Error())
		}
		// TODO
		c.Return()
	}, Catch(func(err error) {
		if err.Error() != "" {
			c.Error(err.Error())
		}
	}))
}
*/
type VerifyPhone struct {
	//// 输出基类
	//BaseModel
	//Name      string `valid:"Required;Match(/^Bee.*/)"` // Name 不能为空并且以 Bee 开头
	Phone string `valid:"Mobile "`
	//Password  string `valid:"Required; MinSize(6)"`
	//UserRole  int    `valid:"Range(1, 3)"`
}

// @Title		注册
// @Description	验证手机
//             |参数名
//             |         |参数类型，可以有的值是 formData、query、path、body、header，formData 表示是 post 请求的数据，query 表示带在 url 之后的参数，path 表示请求路径上得参数，例如上面例子里面的 key，body 表示是一个 raw 数据请求，header 表示带在 header 信息中得参数。
//             |         |           |参数类型
//             |         |           |                     |是否必须
//             |         |           |                     |           |注释
// @Param		args      body        bll.VerifyPhone  true        "user info"
//             |status code
//             |         |返回的类型，必须使用 {} 包含
//             |         |           |返回的对象或者字符串信息，如果是 {object} 类型，那么 bee 工具在生成 docs 的时候会扫描对应的对象，这里填写的是想对你项目的目录名和对象，例如 models.ZDTProduct.ProductList 就表示 /models/ZDTProduct 目录下的 ProductList 对象。
// @Success		200       {object}    utils.Resp
// @router		/VerifyPhone [post]
func (c *UserController) VerifyPhone() {
	Try(func() {
		Log.Println("VerifyPhone()")
		var args VerifyPhone
		if err := c.ParseBody(&args); err != nil {
			Throw(err.Error())
		}
		if err := c.IsValid(args); err != nil {
			Throw(err.Error())
		}
		if Db.Exist("SELECT * FROM users WHERE loginName = ?", args.Phone) {
			Throw("['%s'] is exist", args.Phone)
		}
		if !Db.Exist("SELECT vcode FROM vcodes WHERE loginName = ?", args.Phone) {
			vcode := JobNumber.VCode(vcodeLength)
			if !Db.InsertAsBaseModel(Dict{
				"loginName": args.Phone,
				"vcode":     vcode,
				"freeze":    AbsTime.CurrUtcTo(freezeSeconds),
			}, "vcodes") {
				Throw("failed to INSERT")
			}
			if err := sms.SendVCode(args.Phone, vcode); err != nil {
				Throw("failed to send sms")
			}
			c.Return()
		}
		if v := Db.QField("SELECT {vcode} FROM vcodes WHERE loginName = ? AND freeze <= ?", args.Phone, AbsTime.CurrUtc()); v == nil {
			Throw("operation too frequent,try again later")
		}
		vcode := JobNumber.VCode(vcodeLength)
		if !Db.Exec("UPDATE vcodes SET vcode = ?,freeze = ? WHERE loginName = ?", vcode, AbsTime.CurrUtcTo(freezeSeconds), args.Phone) {
			Throw("failed to UPDATE")
		}
		if err := sms.SendVCode(args.Phone, vcode); err != nil {
			Throw("failed to send sms")
		}
		c.Return()
	}, Catch(func(err error) {
		if err.Error() != "" {
			c.Error(err.Error())
		}
	}))
}

type RegisterUser struct {
	Phone    string `valid:"Mobile "`
	Vcode    string `valid:"Length(6)"`
	Password string `valid:"Required; MinSize(6); MaxSize(20)"`
	UserName string `valid:"Required; MinSize(1); MaxSize(20)"`
}

// @Title		注册
// @Description	注册用户
// @Param		args      body        bll.RegisterUser           true        "user info"
// @Success		200       {object}    utils.Resp
// @router		/RegisterUser [post]
func (c *UserController) RegisterUser() {
	Try(func() {
		Log.Println("RegisterUser()")
		var args RegisterUser
		if err := c.ParseBody(&args); err != nil {
			Throw(err.Error())
		}
		if err := c.IsValid(args); err != nil {
			Throw(err.Error())
		}

		if Db.Exist("SELECT * FROM users WHERE loginName = ?", args.Phone) {
			Throw("['%s'] is exist", args.Phone)
		}
		row := Db.TopRow("SELECT {vcode} FROM vcodes WHERE loginName = ?", args.Phone)
		if row == nil {
			Throw("wrong [phone number]")
		}
		if args.Vcode != row["vcode"].(string) {
			Throw("wrong [vcode]")
		}
		if !Db.InsertAsBaseModel(Dict{
			"loginName": args.Phone,
			"password":  args.Password,
			"userRole":  1,
			"UserName":  args.UserName,
		}, "users") {
			c.Error("failed to INSERT")
			return
		}
		c.Return()

	}, Catch(func(err error) {
		if err.Error() != "" {
			c.Error(err.Error())
		}
	}))
}

type ForgotPassword struct {
	Phone string `valid:"Mobile "`
}

// @Title		找回密码
// @Description	找回密码
// @Param		args      body        bll.ForgotPassword  true        "user info"
// @Success		200       {object}    utils.Resp
// @router		/ForgotPassword [post]
func (c *UserController) ForgotPassword() {
	Try(func() {
		Log.Println("ForgotPassword()")
		var args ForgotPassword
		if err := c.ParseBody(&args); err != nil {
			Throw(err.Error())
		}
		if err := c.IsValid(args); err != nil {
			Throw(err.Error())
		}
		if !Db.Exist("SELECT * FROM users WHERE loginName = ?", args.Phone) {
			Throw("['%s'] is not exist", args.Phone)
		}
		vcode := JobNumber.VCode(vcodeLength)
		if !Db.Exec("UPDATE vcodes SET vcode = ? WHERE loginName = ?", vcode, args.Phone) {
			c.Error("failed to UPDATE")
		}

		if err := sms.SendVCode(args.Phone, vcode); err != nil {
			Throw("failed to send sms")
		}
		c.Return()
	}, Catch(func(err error) {
		if err.Error() != "" {
			c.Error(err.Error())
		}
	}))
}

type ResetPassword struct {
	Phone    string `valid:"Mobile "`
	Vcode    string `valid:"Length(6)"`
	Password string `valid:"Required; MinSize(6); MaxSize(20)"`
}

// @Title		重置密码
// @Description	重置密码
// @Param		args      body        bll.ResetPassword  true        "参数描述"
// @Success		200       {object}    utils.Resp
// @router		/ResetPassword [post]
func (c *UserController) ResetPassword() {
	Try(func() {
		Log.Println("ResetPassword()")
		var args ResetPassword
		if err := c.ParseBody(&args); err != nil {
			Throw(err.Error())
		}
		if err := c.IsValid(args); err != nil {
			Throw(err.Error())
		}
		if !Db.Exist("SELECT * FROM users WHERE loginName = ?", args.Phone) {
			Throw("['%s'] is not exist", args.Phone)
		}
		if !Db.Exec("UPDATE users,vcodes SET users.`password` = ?,users.`addDate` = ?,users.`addTime` = ?,users.`addUtc` = ? WHERE users.loginName = ? AND vcodes.loginName = ? AND vcodes.vcode = ?",
			args.Password, AbsTime.CurrDate(), AbsTime.CurrTime(), AbsTime.CurrUtc(), args.Phone, args.Phone, args.Vcode) {
			c.Error("帐号或验证码错误")
		}
		if !Db.Exec("UPDATE vcodes SET vcode = ? WHERE loginName = ?", JobNumber.VCode(vcodeLength), args.Phone) {
			c.Error("failed to UPDATE")
		}
		c.Return()
	}, Catch(func(err error) {
		if err.Error() != "" {
			c.Error(err.Error())
		}
	}))
}

type Login struct {
	LoginName string `valid:"Mobile "`
	Password  string `valid:"Required; MinSize(6); MaxSize(20)"`
}

// @Title		登录
// @Description	登录
// @Param		args      body        bll.Login  true        "参数描述"
// @Success		200       {object}    utils.Resp
// @router		/Login [post]
func (c *UserController) Login() {
	Try(func() {
		Log.Println("Login()")
		var args Login
		if err := c.ParseBody(&args); err != nil {
			Throw(err.Error())
		}
		if err := c.IsValid(args); err != nil {
			Throw(err.Error())
		}

		if !Db.Exist("SELECT * FROM users WHERE loginName = ? AND password = ?", args.LoginName, args.Password) {
			Throw("登录失败")
		}
		accessKey := JobNumber.New()
		if !Db.Exec("UPDATE users SET users.`accessKey` = ?,users.`expires` = ? WHERE loginName = ? AND password = ?",
			accessKey, AbsTime.CurrUtcTo(tokenExpires), args.LoginName, args.Password) {
			c.Error("Failed to UPDATE")
		}
		c.JsonData(accessKey)
		return
	}, Catch(func(err error) {
		if err.Error() != "" {
			c.Error(err.Error())
		}
	}))
}

// @Title		获取用户信息
// @Description	获取用户信息
// @Param		accessKey query       string  true        "参数描述"
// @Success		200       {object}    utils.Resp
// @router		/GetUserInfo [post]
func (c *UserController) GetUserInfo() {
	Try(func() {
		Log.Println("GetUserInfo()")
		accessKey := c.GetString("accessKey")
		ai := Rbac.AccessKeyIsValid(accessKey)
		fmt.Println(ai)
		if ai.Role != UserRoles_Manager {
			Throw("DeniedAccess")
		}

		// TODO

		// 头像,姓名,手机号码,--微信号,行业=行业列表(),公司名称,--省份,城市
		sql := fmt.Sprintf("SELECT {faceIcon,userName,loginName,focusIndustry,company,city} FROM users WHERE loginName = %s", ai.LoginName)

		fmt.Println("sql:", sql)
		result := Db.TopRow(sql)
		c.JsonData(result)

	}, Catch(func(err error) {
		if err.Error() != "" {
			c.Error(err.Error())
		}
	}))
}
