package bll

import (
	. "ManualApi/zfx"
	"fmt"
	"path/filepath"
	"strings"
)

// 管理相关
type MgrController struct {
	BaseController
}

//https://beego.me/docs/advantage/docs.md

// @Title		上传资源
// @Description	上传资源
// @Param		accessKey query       string  true        "参数描述"
// @Param		args      formData    file  true        "参数描述"
// @Success		200       {object}    utils.Resp
// @router		/UploadResource [post]
func (c *MgrController) UploadResource() {
	Try(func() {
		Log.Println("UploadResource()")
		// TODO
		accessKey := c.GetString("accessKey")
		ai := Rbac.AccessKeyIsValid(accessKey)
		fmt.Println(ai)
		if ai.Role != UserRoles_Manager {
			Throw("DeniedAccess")
		}

		f, h, err := c.GetFile("args")
		defer f.Close()
		if err != nil {
			Throw(err.Error())
		}
		resName := JobNumber.New() + filepath.Ext(h.Filename)
		c.SaveToFile("args", AppSettingString("ufolder")+"/"+resName)
		//hash := Md5Sum("ufiles/" + resName)
		//fmt.Println("hash:", hash)

		scheme := "http:/"
		if c.Ctx.Request.TLS != nil {
			scheme = "https:/"
		}
		//fmt.Println("self:", scheme, c.Ctx.Request.Host, c.Ctx.Request.RequestURI)
		//fmt.Println("self:", scheme, c.Ctx.Request.Host, AppSettingString("ufolder"))
		resUrl := strings.Join([]string{scheme, c.Ctx.Request.Host, AppSettingString("ufolder"), resName}, "/")

		c.JsonData(resUrl)
	}, Catch(func(err error) {
		if err.Error() != "" {
			c.Error(err.Error())
		}
	}))
}

type CreateCarousel struct {
	Title            string
	ImageUrl         string
	ShowcaseUniqueId string
	SortIndex        int
}

// @Title		创建轮播图
// @Description	创建轮播图
// @Param		accessKey query       string  true        "参数描述"
// @Param		args      body        bll.CreateCarousel  true        "参数描述"
// @Success		200       {object}    utils.Resp
// @router		/CreateCarousel [post]
func (c *MgrController) CreateCarousel() {
	Try(func() {
		Log.Println("CreateCarousel()")
		accessKey := c.GetString("accessKey")
		ai := Rbac.AccessKeyIsValid(accessKey)
		fmt.Println(ai)
		if ai.Role != UserRoles_Manager {
			Throw("DeniedAccess")
		}
		var args CreateCarousel
		if err := c.ParseBody(&args); err != nil {
			Throw(err.Error())
		}
		if err := c.IsValid(args); err != nil {
			Throw(err.Error())
		}
		// TODO
		if !Db.InsertAsBaseModel(Dict{
			"Title":            args.Title,
			"ImageUrl":         args.ImageUrl,
			"ShowcaseUniqueId": args.ShowcaseUniqueId,
			"SortIndex":        0,
		}, "vcodes") {
			Throw("failed to INSERT")
		}
		c.Return()
	}, Catch(func(err error) {
		if err.Error() != "" {
			c.Error(err.Error())
		}
	}))
}

type GetUserList struct {
	Phone     string
	PageIndex int `valid:"Range(0, 1000)"`
	PageSize  int `valid:"Range(1, 100)"`
}

// @Title		获取用户列表
// @Description	获取用户列表
// @Param		accessKey query       string  true        "参数描述"
// @Param		args      body        bll.GetUserList  true        "参数描述"
// @Success		200       {object}    utils.Resp
// @router		/GetUserList [post]
func (c *MgrController) GetUserList() {
	Try(func() {
		Log.Println("GetUserList()")

		accessKey := c.GetString("accessKey")
		ai := Rbac.AccessKeyIsValid(accessKey)
		fmt.Println("ai:", ai)
		if ai.Role != UserRoles_Manager {
			Throw("DeniedAccess")
		}

		var args GetUserList
		if err := c.ParseBody(&args); err != nil {
			Throw(err.Error())
		}
		if err := c.IsValid(args); err != nil {
			Throw(err.Error())
		}

		// TODO
		sql := fmt.Sprintf("SELECT {uniqueId,loginName,userName,userRole,addDate,addTime,addUtc} FROM users ORDER BY addUtc ASC LIMIT %d,%d", args.PageIndex*args.PageSize, args.PageSize)
		csql := fmt.Sprintf("SELECT COUNT(*) AS {countUsers} FROM users")
		if args.Phone != "" {
			fmt.Println("phone is empty!!!")
			if !Db.Exist("SELECT * FROM users WHERE loginName = ?", args.Phone) {
				//Throw("['%s'] is not exist", args.Phone)
				c.Return()
			}
			sql = fmt.Sprintf("SELECT {uniqueId,loginName,userName,userRole,addDate,addTime,addUtc} FROM users WHERE loginName = %s ORDER BY addUtc ASC LIMIT %d,%d", args.Phone, args.PageIndex*args.PageSize, args.PageSize)
			csql = fmt.Sprintf("SELECT COUNT(*) AS {countUsers} FROM users WHERE loginName = %s", args.Phone)
		}
		fmt.Println("args:", args)
		fmt.Println("sql:", sql)
		result := Db.QMaps(sql)
		total := ToInt(Db.QField(csql))

		c.JsonDatas(result, total, args.PageIndex, args.PageSize)

	}, Catch(func(err error) {
		fmt.Print("GetUserList@104")
		if err.Error() != "" {
			c.Error(err.Error())
		}
	}))
}

type CreateTopic struct {
	Title string `valid:"Required; MinSize(2); MaxSize(20)"`
}

// @Title		创建栏目
// @Description	创建栏目
// @Param		accessKey query       string  true        "参数描述"
// @Param		args      body        bll.CreateTopic  true        "参数描述"
// @Success		200       {object}    utils.Resp
// @router		/CreateTopic [post]
func (c *MgrController) CreateTopic() {
	Try(func() {
		Log.Println("CreateTopic()")
		accessKey := c.GetString("accessKey")
		ai := Rbac.AccessKeyIsValid(accessKey)
		fmt.Println(ai)
		if ai.Role != UserRoles_Manager {
			Throw("DeniedAccess")
		}
		var args CreateTopic
		if err := c.ParseBody(&args); err != nil {
			Throw(err.Error())
		}
		if err := c.IsValid(args); err != nil {
			Throw(err.Error())
		}
		// TODO
		if !Db.InsertAsBaseModel(Dict{
			"Title": args.Title,
		}, "topics") {
			Throw("failed to INSERT")
		}
		c.Return()
	}, Catch(func(err error) {
		if err.Error() != "" {
			c.Error(err.Error())
		}
	}))
}

type CreateSpecial struct {
	Title     string `valid:"Required; MinSize(2); MaxSize(255)"`
	TopicName string `valid:"Required; MinSize(2); MaxSize(20)"`
	Provider  string `valid:"Required; MinSize(2); MaxSize(20)"`
	Cover     string `valid:"Required; MinSize(2); MaxSize(255)"`
	Introduce string `valid:"Required; MinSize(2); MaxSize(255)"`
}

// @Title		创建专题
// @Description	创建专题
// @Param		accessKey query       string  true        "参数描述"
// @Param		args      body        bll.CreateSpecial  true        "参数描述"
// @Success		200       {object}    utils.Resp
// @router		/CreateSpecial [post]
func (c *MgrController) CreateSpecial() {
	Try(func() {
		Log.Println("CreateSpecial()")
		accessKey := c.GetString("accessKey")
		ai := Rbac.AccessKeyIsValid(accessKey)
		fmt.Println(ai)
		if ai.Role != UserRoles_Manager {
			Throw("DeniedAccess")
		}
		var args CreateSpecial
		if err := c.ParseBody(&args); err != nil {
			Throw(err.Error())
		}
		if err := c.IsValid(args); err != nil {
			Throw(err.Error())
		}
		// TODO
		if !Db.InsertAsBaseModel(Dict{
			"title":     args.Title,
			"topicName": args.TopicName,
			"provider":  args.Provider,
			"cover":     args.Cover,
			"introduce": args.Introduce,
		}, "specials") {
			Throw("failed to INSERT")
		}
		c.Return()
	}, Catch(func(err error) {
		if err.Error() != "" {
			c.Error(err.Error())
		}
	}))
}

type CreateContent struct {
	SpecialUniqueId string
	//Title           string `valid:"Required; MinSize(2); MaxSize(255)"`
	//Description     string `valid:"Required; MinSize(2); MaxSize(20)"`
	Price        int    `valid:"Range(1, 10000)"`
	ChargingRule int    `valid:"Range(1, 3)"`
	Days         int    `valid:"Range(1, 365)"`
	ResourceType int    `valid:"Range(1, 3)"`
	ResourceUrl  string `valid:"Required; MinSize(2); MaxSize(255)"`
}

// @Title		创建内容
// @Description	创建内容
// @Param		accessKey query       string  true        "参数描述"
// @Param		args      body        bll.CreateContent  true        "参数描述"
// @Success		200       {object}    utils.Resp
// @router		/CreateContent [post]
func (c *MgrController) CreateContent() {
	Try(func() {
		Log.Println("CreateContent()")
		accessKey := c.GetString("accessKey")
		ai := Rbac.AccessKeyIsValid(accessKey)
		fmt.Println(ai)
		if ai.Role != UserRoles_Manager {
			Throw("DeniedAccess")
		}
		var args CreateContent
		if err := c.ParseBody(&args); err != nil {
			Throw(err.Error())
		}
		if err := c.IsValid(args); err != nil {
			Throw(err.Error())
		}
		// TODO
		if !Db.InsertAsBaseModel(Dict{
			"specialUniqueId": args.SpecialUniqueId,
			"price":           args.Price,
			"chargingRule":    args.ChargingRule,
			"days":            args.Days,
			"resourceType":    args.ResourceType,
			"resourceUrl":     args.ResourceUrl,
		}, "contents") {
			Throw("failed to INSERT")
		}
		c.Return()
	}, Catch(func(err error) {
		if err.Error() != "" {
			c.Error(err.Error())
		}
	}))
}
