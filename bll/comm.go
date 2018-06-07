package bll

import (
	. "ManualApi/zfx"
	"fmt"
)

// 公共相关
type CommController struct {
	BaseController
}

// @Title		获取行业列表
// @Description	获取行业列表
// @Param		accessKey query       string  true        "参数描述"
// @Success		200       {object}    utils.Resp
// @router		/GetIndustryList [post]
func (c *CommController) GetIndustryList() {
	Try(func() {
		Log.Println("GetIndustryList()")
		accessKey := c.GetString("accessKey")
		ai := Rbac.AccessKeyIsValid(accessKey)
		fmt.Println(ai)
		//if ai.Role != UserRoles_Manager {
		//	Throw("DeniedAccess")
		//}
		// TODO
		//jsonText := AppSettingString("industrys")
		//industryList := []string{}
		//jsonText := `["金属生产贸易商", "煤焦钢生产贸易商", "化工行业", "农业", "养殖业", "建材行业", "互联网/电子商务/IT", "金融业", "制造业", "广告/媒体", "教育/培训", "房地产/建筑", "服务业", "物流/运输", "制药/医疗", "政府/非盈利/其他"]`
		//if err := c.FromJson(jsonText, &industryList); err != nil {
		//	Throw(err.Error())
		//}
		//c.JsonData(industryList)
		c.JsonData(map[int]string{
			1:  "金属生产贸易商",
			2:  "煤焦钢生产贸易商",
			3:  "化工行业",
			4:  "农业",
			5:  "养殖业",
			6:  "建材行业",
			7:  "互联网/电子商务/IT",
			8:  "金融业",
			9:  "制造业",
			10: "广告/媒体",
			11: "教育/培训",
			12: "房地产/建筑",
			13: "服务业",
			14: "物流/运输",
			15: "制药/医疗",
			16: "政府/非盈利/其他",
		})

	}, Catch(func(err error) {
		if err.Error() != "" {
			c.Error(err.Error())
		}
	}))
}

// @Title		获取资源类型
// @Description	获取资源类型
// @Param		accessKey query       string  true        "参数描述"
// @Success		200       {object}    utils.Resp
// @router		/GetResourceTypes [post]
func (c *CommController) GetResourceTypes() {
	Try(func() {
		Log.Println("GetIndustryList()")
		accessKey := c.GetString("accessKey")
		ai := Rbac.AccessKeyIsValid(accessKey)
		fmt.Println(ai)
		//if ai.Role != UserRoles_Manager {
		//	Throw("DeniedAccess")
		//}

		// TODO
		c.JsonData(map[int]string{
			1: "音频", // Audio
			2: "视频", // Video
			3: "活动", // Activities
		})

	}, Catch(func(err error) {
		if err.Error() != "" {
			c.Error(err.Error())
		}
	}))
}

// @Title		获取资源类型
// @Description	获取资源类型
// @Param		accessKey query       string  true        "参数描述"
// @Success		200       {object}    utils.Resp
// @router		/GetChargingRules [post]
func (c *CommController) GetChargingRules() {
	Try(func() {
		Log.Println("GetIndustryList()")
		accessKey := c.GetString("accessKey")
		ai := Rbac.AccessKeyIsValid(accessKey)
		fmt.Println(ai)
		//if ai.Role != UserRoles_Manager {
		//	Throw("DeniedAccess")
		//}

		// 计费规则(1.免费资源,2.会员专享,3.付费资源)
		// TODO
		c.JsonData(map[int]string{
			1: "免费资源",
			2: "会员专享",
			3: "付费资源",
		})

	}, Catch(func(err error) {
		if err.Error() != "" {
			c.Error(err.Error())
		}
	}))
}
