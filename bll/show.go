package bll

import (
	. "ManualApi/zfx"
	"fmt"
)

// 业务相关
type ShowController struct {
	BaseController
}

type GetTopicList struct {
	PageIndex int `valid:"Range(0, 1000)"`
	PageSize  int `valid:"Range(1, 100)"`
}

// @Title		获取栏目列表
// @Description	获取栏目列表
// @Param		accessKey query       string  true        "参数描述"
// @Param		args      body        bll.GetTopicList  true        "参数描述"
// @Success		200       {object}    utils.Resp
// @router		/GetTopicList [post]
func (c *ShowController) GetTopicList() {
	Try(func() {
		Log.Println("GetTopicList()")
		accessKey := c.GetString("accessKey")
		ai := Rbac.AccessKeyIsValid(accessKey)
		fmt.Println(ai)
		//if ai.Role != UserRoles_Manager {
		//	Throw("DeniedAccess")
		//}

		var args GetTopicList
		if err := c.ParseBody(&args); err != nil {
			Throw(err.Error())
		}
		if err := c.IsValid(args); err != nil {
			Throw(err.Error())
		}
		// TODO
		sql := fmt.Sprintf("SELECT {uniqueId,title,addDate,addTime} FROM topics WHERE isDeleted = 0 ORDER BY addUtc ASC LIMIT %d,%d", args.PageIndex*args.PageSize, args.PageSize)
		csql := fmt.Sprintf("SELECT COUNT(*) AS {countUsers} FROM topics WHERE isDeleted = 0")
		fmt.Println("args:", args)
		fmt.Println("sql:", sql)
		result := Db.QMaps(sql)
		total := ToInt(Db.QField(csql))
		c.JsonDatas(result, total, args.PageIndex, args.PageSize)
	}, Catch(func(err error) {
		if err.Error() != "" {
			c.Error(err.Error())
		}
	}))
}

type GetSpecialList struct {
	PageIndex int `valid:"Range(0, 1000)"`
	PageSize  int `valid:"Range(1, 100)"`
	TopicName string
}

// @Title		获取专题列表
// @Description	获取专题列表
// @Param		accessKey query       string  true        "参数描述"
// @Param		args      body        bll.GetSpecialList  true        "参数描述"
// @Success		200       {object}    utils.Resp
// @router		/GetSpecialList [post]
func (c *ShowController) GetSpecialList() {
	Try(func() {
		Log.Println("GetSpecialList()")
		accessKey := c.GetString("accessKey")
		ai := Rbac.AccessKeyIsValid(accessKey)
		fmt.Println(ai)
		//if ai.Role != UserRoles_Manager {
		//	Throw("DeniedAccess")
		//}

		var args GetSpecialList
		if err := c.ParseBody(&args); err != nil {
			Throw(err.Error())
		}
		if err := c.IsValid(args); err != nil {
			Throw(err.Error())
		}
		// TODO
		sql := fmt.Sprintf("SELECT {uniqueId,title,topicName,provider,cover,introduce,addDate,addTime} FROM specials WHERE isDeleted = 0 ORDER BY addUtc ASC LIMIT %d,%d", args.PageIndex*args.PageSize, args.PageSize)
		csql := fmt.Sprintf("SELECT COUNT(*) AS {countUsers} FROM  specials WHERE isDeleted = 0")
		if args.TopicName != "" {
			sql = fmt.Sprintf("SELECT {uniqueId,title,topicName,provider,cover,introduce,addDate,addTime} FROM specials WHERE isDeleted = 0 AND topicName = '%s' ORDER BY addUtc ASC LIMIT %d,%d", args.TopicName, args.PageIndex*args.PageSize, args.PageSize)
			csql = fmt.Sprintf("SELECT COUNT(*) AS {countUsers} FROM  specials WHERE isDeleted = 0 AND topicName = '%s'", args.TopicName)
			fmt.Println(sql)
		}
		fmt.Println("args:", args)
		fmt.Println("sql:", sql)
		result := Db.QMaps(sql)
		total := ToInt(Db.QField(csql))
		c.JsonDatas(result, total, args.PageIndex, args.PageSize)
	}, Catch(func(err error) {
		if err.Error() != "" {
			c.Error(err.Error())
		}
	}))
}

type GetContentList struct {
	PageIndex       int    `valid:"Range(0, 1000)"`
	PageSize        int    `valid:"Range(1, 100)"`
	SpecialUniqueId string `valid:"Required; Length(16)"`
}

// @Title		获取内容列表
// @Description	获取内容列表
// @Param		accessKey query       string  true        "参数描述"
// @Param		args      body        bll.GetContentList  true        "参数描述"
// @Success		200       {object}    utils.Resp
// @router		/GetContentList [post]
func (c *ShowController) GetContentList() {
	Try(func() {
		Log.Println("GetContentList()")
		accessKey := c.GetString("accessKey")
		ai := Rbac.AccessKeyIsValid(accessKey)
		fmt.Println(ai)
		//if ai.Role != UserRoles_Manager {
		//	Throw("DeniedAccess")
		//}

		var args GetContentList
		if err := c.ParseBody(&args); err != nil {
			Throw(err.Error())
		}
		if err := c.IsValid(args); err != nil {
			Throw(err.Error())
		}
		// TODO
		sql := fmt.Sprintf("SELECT {specialUniqueId,price,chargingRule,days,resourceType,resourceUrl} FROM contents WHERE specialUniqueId = '%s' AND isDeleted = 0 ORDER BY addUtc ASC LIMIT %d,%d", args.SpecialUniqueId, args.PageIndex*args.PageSize, args.PageSize)
		csql := fmt.Sprintf("SELECT COUNT(*) AS {countUsers} FROM contents WHERE specialUniqueId = '%s' AND isDeleted = 0 ", args.SpecialUniqueId)
		fmt.Println("args:", args)
		fmt.Println("sql:", sql)
		result := Db.QMaps(sql)
		total := ToInt(Db.QField(csql))
		c.JsonDatas(result, total, args.PageIndex, args.PageSize)
	}, Catch(func(err error) {
		if err.Error() != "" {
			c.Error(err.Error())
		}
	}))
}
