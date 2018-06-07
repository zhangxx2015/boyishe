package zfx

//////////////////////////////////// API 契约包装
type Resp struct {
	// 记录总数
	Total int `json:"total"`
	// 分页上限
	PageMax int `json:"PageMax"`
	// 分页索引
	PageIndex int `json:"pageIndex"`
	// 分页大小
	PageSize int `json:"pageSize"`
	//// 请求内容
	//Request string `json:"Request"`
	// 返回状态
	Status string `json:"Status"`
	// 返回内容
	ResType string `json:"ResType"`
	// 数据
	Data interface{} `json:"Data"`
	// 错误
	Error string `json:"Error"`
}
