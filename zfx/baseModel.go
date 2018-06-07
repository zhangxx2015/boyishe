package zfx

//////////////////////////////////// 数据模型基类
type BaseModel struct {
	//Id            int    `orm:"column(id);auto"`
	UniqueId string `orm:"column(uniqueId);size(20)"`
	AddDate  string `orm:"column(addDate);size(15)"`
	AddTime  string `orm:"column(addTime);size(15)"`
	AddUtc   int    `orm:"column(addUtc)"`
	//AppType       string `orm:"column(appType);size(20)"`
	IsDeleted     uint64 `orm:"column(isDeleted);size(1)"`
	Remark        string `orm:"column(remark);size(255);null"`
	Opatator      string `orm:"column(opatator);size(255);null"`
	LastUpdateUtc int    `orm:"column(lastUpdateUtc)"`
}

//func (m *BaseModel) ToMap() map[string]interface{} {
//	return ToMap(*m)
//}

/*
var args models.AddUser
maps := utils.ApplyBaseFields(args)
if !c.Insert(maps, "users") {
	c.Error("failed to INSERT")
	return
}
*/
//func ApplyBaseFields(model interface{}) map[string]interface{} {
//	baseMaps := make(map[string]interface{}, 0)
//	maps := ToMap(model)
//	for k, v := range maps {
//		baseMaps[k] = v
//	}
//	baseMaps["UniqueId"] = JobNumber()
//	baseMaps["AddDate"] = CurrDate()
//	baseMaps["AddTime"] = CurrTime()
//	baseMaps["AddUtc"] = CurrUtc()
//	baseMaps["IsDeleted"] = 0
//	baseMaps["Remark"] = ""
//	baseMaps["LastUpdateUtc"] = CurrUtc()
//	return baseMaps
//}

//func MergeToBaseModel(maps map[string]interface{}) map[string]interface{} {
//	baseMaps := make(map[string]interface{}, 0)
//	for k, v := range maps {
//		baseMaps[k] = v
//	}
//	baseMaps["UniqueId"] = JobNumber.New()
//	baseMaps["AddDate"] = AbsTime.CurrDate()
//	baseMaps["AddTime"] = AbsTime.CurrTime()
//	baseMaps["AddUtc"] = AbsTime.CurrUtc()
//	baseMaps["IsDeleted"] = 0
//	baseMaps["Remark"] = ""
//	baseMaps["LastUpdateUtc"] = AbsTime.CurrUtc()
//	return baseMaps
//}

//func ToBaseModel(model interface{}) map[string]interface{} {
//	baseMaps := make(map[string]interface{}, 0)
//	maps := ToMap(model)
//	for k, v := range maps {
//		baseMaps[k] = v
//	}
//	baseMaps["UniqueId"] = JobNumber.New()
//	baseMaps["AddDate"] = AbsTime.CurrDate()
//	baseMaps["AddTime"] = AbsTime.CurrTime()
//	baseMaps["AddUtc"] = AbsTime.CurrUtc()
//	baseMaps["IsDeleted"] = 0
//	baseMaps["Remark"] = ""
//	baseMaps["LastUpdateUtc"] = AbsTime.CurrUtc()
//	return baseMaps
//}

//func GenBaseModel() map[string]interface{} {
//	baseMaps := make(map[string]interface{}, 0)
//	baseMaps["UniqueId"] = JobNumber.New()
//	baseMaps["AddDate"] = AbsTime.CurrDate()
//	baseMaps["AddTime"] = AbsTime.CurrTime()
//	baseMaps["AddUtc"] = AbsTime.CurrUtc()
//	baseMaps["IsDeleted"] = 0
//	baseMaps["Remark"] = ""
//	baseMaps["LastUpdateUtc"] = AbsTime.CurrUtc()
//	return baseMaps
//}

//func Copy() {
//	//	maps := ToMap(model)
//	//	for k, v := range maps {
//	//		baseMaps[k] = v
//	//	}
//}

//func (m *BaseModel) ApplyBaseFields() map[string]interface{} {
//	return ApplyBaseFields(*m)
//}
