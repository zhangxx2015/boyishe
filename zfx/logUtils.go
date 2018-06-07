package zfx

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type _logger struct {
}

func init() {

	// 设置适配器(文件)
	logs.SetLogger(logs.AdapterFile, `{"filename":"`+beego.AppConfig.String("logFile")+`","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
	// 设置适配器(控制台)第二个, 同时输出
	logs.SetLogger(logs.AdapterConsole)
	// 输出调用的文件名和文件行号
	logs.EnableFuncCallDepth(true)
	// 异步(存在丢失概率)缓存大小
	//logs.Async(1024)
	//logs.Async()
}

//////////////////////////////////// 日志函数
func (inst *_logger) Println(text string) {
	l := logs.GetLogger() // 改系统时间,查看日志自动分段
	l.Println(fmt.Sprintf("%s\r\n", text))
}
