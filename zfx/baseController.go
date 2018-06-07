package zfx

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/validation"
)

//////////////////////////////////// 控制器基类
type BaseController struct {
	beego.Controller
}

//////////////////////////////////// 序列化
func (c *BaseController) ToJson(obj interface{}) (string, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (c *BaseController) FromJson(jsonText string, obj interface{}) error {
	return json.Unmarshal([]byte(jsonText), obj)
}

//////////////////////////////////// 反序列化
func (c *BaseController) ParseBody(args interface{}) error {
	return json.Unmarshal(c.Ctx.Input.RequestBody, args)
}

//////////////////////////////////// Resp 辅助函数
func (c *BaseController) Error(err string) {
	fmt.Println("throw err:", err)
	c.Ctx.Output.SetStatus(200)
	//c.Data["json"] = &Resp{Total: 0, PageMax: 0, PageIndex: 0, PageSize: 0, Request: "", Status: "Error", ResType: "Error", Data: nil, Error: err}
	c.Data["json"] = &Resp{Total: 0, PageMax: 0, PageIndex: 0, PageSize: 0, Status: "Error", ResType: "Error", Data: nil, Error: err}
	c.ServeJSON()
}
func (c *BaseController) JsonData(obj interface{}) {
	c.Ctx.Output.SetStatus(200)
	//c.Data["json"] = &Resp{Total: 0, PageMax: 0, PageIndex: 0, PageSize: 0, Request: "", Status: "OK", ResType: "JsonData", Data: obj, Error: ""}
	c.Data["json"] = &Resp{Total: 0, PageMax: 0, PageIndex: 0, PageSize: 0, Status: "OK", ResType: "JsonData", Data: obj, Error: ""}
	c.ServeJSON()
}
func (c *BaseController) JsonDatas(objs []map[string]interface{}, total int, pageIndex int, pageSize int) {
	c.Ctx.Output.SetStatus(200)
	//c.Data["json"] = &Resp{Total: total(), PageMax: max(), PageIndex: pageIndex, PageSize: pageSize, Request: "", Status: "OK", ResType: "JsonData", Data: obj, Error: ""}
	pageMax := total / pageSize
	if pageSize > total {
		pageMax = 1
	}
	c.Data["json"] = &Resp{Total: total, PageMax: pageMax, PageIndex: pageIndex, PageSize: pageSize, Status: "OK", ResType: "JsonData", Data: objs, Error: ""}
	c.ServeJSON()
}

//////////////////////////////////// 模型验证
func (c *BaseController) IsValid(args interface{}) error {
	valid := validation.Validation{}
	b, err := valid.Valid(args)
	if err != nil {
		fmt.Println("IsValid@50:", err)
		return err
	}
	if !b {
		for _, err := range valid.Errors {
			//			err.Message = fmt.Sprintf("[%s]%s", err.Key, err.Message)
			//			return err
			fmt.Println("IsValid:", err.Key, err.Message)
			return errors.New(fmt.Sprintf("[%s]%s", err.Key, err.Message))
		}
	}
	return nil
}
