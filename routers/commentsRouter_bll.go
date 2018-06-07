package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["ManualApi/bll:CommController"] = append(beego.GlobalControllerRouter["ManualApi/bll:CommController"],
		beego.ControllerComments{
			Method: "GetIndustryList",
			Router: `/GetIndustryList`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ManualApi/bll:CommController"] = append(beego.GlobalControllerRouter["ManualApi/bll:CommController"],
		beego.ControllerComments{
			Method: "GetResourceTypes",
			Router: `/GetResourceTypes`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ManualApi/bll:CommController"] = append(beego.GlobalControllerRouter["ManualApi/bll:CommController"],
		beego.ControllerComments{
			Method: "GetChargingRules",
			Router: `/GetChargingRules`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ManualApi/bll:DbgController"] = append(beego.GlobalControllerRouter["ManualApi/bll:DbgController"],
		beego.ControllerComments{
			Method: "MyDebug",
			Router: `/MyDebug`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ManualApi/bll:MgrController"] = append(beego.GlobalControllerRouter["ManualApi/bll:MgrController"],
		beego.ControllerComments{
			Method: "UploadResource",
			Router: `/UploadResource`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ManualApi/bll:MgrController"] = append(beego.GlobalControllerRouter["ManualApi/bll:MgrController"],
		beego.ControllerComments{
			Method: "CreateCarousel",
			Router: `/CreateCarousel`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ManualApi/bll:MgrController"] = append(beego.GlobalControllerRouter["ManualApi/bll:MgrController"],
		beego.ControllerComments{
			Method: "GetUserList",
			Router: `/GetUserList`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ManualApi/bll:MgrController"] = append(beego.GlobalControllerRouter["ManualApi/bll:MgrController"],
		beego.ControllerComments{
			Method: "CreateTopic",
			Router: `/CreateTopic`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ManualApi/bll:MgrController"] = append(beego.GlobalControllerRouter["ManualApi/bll:MgrController"],
		beego.ControllerComments{
			Method: "CreateSpecial",
			Router: `/CreateSpecial`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ManualApi/bll:MgrController"] = append(beego.GlobalControllerRouter["ManualApi/bll:MgrController"],
		beego.ControllerComments{
			Method: "CreateContent",
			Router: `/CreateContent`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ManualApi/bll:ShowController"] = append(beego.GlobalControllerRouter["ManualApi/bll:ShowController"],
		beego.ControllerComments{
			Method: "GetTopicList",
			Router: `/GetTopicList`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ManualApi/bll:ShowController"] = append(beego.GlobalControllerRouter["ManualApi/bll:ShowController"],
		beego.ControllerComments{
			Method: "GetSpecialList",
			Router: `/GetSpecialList`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ManualApi/bll:ShowController"] = append(beego.GlobalControllerRouter["ManualApi/bll:ShowController"],
		beego.ControllerComments{
			Method: "GetContentList",
			Router: `/GetContentList`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ManualApi/bll:UserController"] = append(beego.GlobalControllerRouter["ManualApi/bll:UserController"],
		beego.ControllerComments{
			Method: "VerifyPhone",
			Router: `/VerifyPhone`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ManualApi/bll:UserController"] = append(beego.GlobalControllerRouter["ManualApi/bll:UserController"],
		beego.ControllerComments{
			Method: "RegisterUser",
			Router: `/RegisterUser`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ManualApi/bll:UserController"] = append(beego.GlobalControllerRouter["ManualApi/bll:UserController"],
		beego.ControllerComments{
			Method: "ForgotPassword",
			Router: `/ForgotPassword`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ManualApi/bll:UserController"] = append(beego.GlobalControllerRouter["ManualApi/bll:UserController"],
		beego.ControllerComments{
			Method: "ResetPassword",
			Router: `/ResetPassword`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ManualApi/bll:UserController"] = append(beego.GlobalControllerRouter["ManualApi/bll:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/Login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ManualApi/bll:UserController"] = append(beego.GlobalControllerRouter["ManualApi/bll:UserController"],
		beego.ControllerComments{
			Method: "GetUserInfo",
			Router: `/GetUserInfo`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
