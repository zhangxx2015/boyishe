package zfx

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type _dataAccessLayer struct {
}

func init() {
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("conn"))
}

func (inst *_dataAccessLayer) InsertAsBaseModel(maps map[string]interface{}, table string) bool {
	baseMaps := make(map[string]interface{}, 0)
	for k, v := range maps {
		//fmt.Println(k, v)
		baseMaps[k] = v
	}
	baseMaps["UniqueId"] = JobNumber.New()
	baseMaps["AddDate"] = AbsTime.CurrDate()
	baseMaps["AddTime"] = AbsTime.CurrTime()
	baseMaps["AddUtc"] = AbsTime.CurrUtc()
	baseMaps["IsDeleted"] = 0
	baseMaps["Remark"] = ""
	baseMaps["LastUpdateUtc"] = AbsTime.CurrUtc()
	//fmt.Println(baseMaps)
	return inst.Insert(baseMaps, table)
}

//////////////////////////////////// sql 添加记录

func (inst *_dataAccessLayer) Insert(maps map[string]interface{}, table string) bool {
	keyParts := make([]string, 0)
	varParts := make([]string, 0)
	valParts := make([]interface{}, 0)
	for k, v := range maps {
		fmt.Println("[", k, "]", ":", v)
		keyParts = append(keyParts, k)
		varParts = append(varParts, "?")
		valParts = append(valParts, v)
	}
	fields := strings.Join(keyParts, ",")
	placeholder := strings.Join(varParts, ",")
	sql := fmt.Sprintf("INSERT INTO %s (%s) VALUES(%s)", table, fields, placeholder)
	fmt.Println("SQL:", sql)
	fmt.Println(valParts)
	return Db.Exec(sql, valParts)
}

func (inst *_dataAccessLayer) Exist(sql string, args ...interface{}) bool {
	lists := Db.Query(sql, args)
	return len(lists) > 0
}

func (inst *_dataAccessLayer) QField(sql string, args ...interface{}) interface{} {
	row := Db.TopRow(sql, args)
	if len(row) < 1 {
		return nil
	}
	for _, v := range row {
		return v
		break
	}
	return nil
}

func (inst *_dataAccessLayer) TopRow(sql string, args ...interface{}) map[string]interface{} {
	rows := Db.QMaps(sql, args)
	if len(rows) > 0 {
		return rows[0]
	}
	return nil
}

//////////////////////////////////// sql 获取字典列表
/*
////rows := DbQMaps("SELECT {uniqueId,loginName,password} FROM users", []string{"uniqueId", "loginName", "password"})
rows := DbQMaps("SELECT {uniqueId,loginName,password} FROM users")
fmt.Println("loginName", rows[1]["loginName"])
for i := 0; i < len(rows); i++ {
	for key, value := range rows[i] {
		fmt.Println("[", i, "]", key, ":", value)
	}
}
*/
func (inst *_dataAccessLayer) QMaps(sql string, args ...interface{}) []map[string]interface{} {
	ri := strings.LastIndex(sql, "}")
	li := strings.IndexAny(sql, "{")
	if -1 == li || -1 == ri {
		Throw("cannot find 'fields' tags in sql")
	}
	if ri <= (1 + li) {
		Throw("illegal 'fields' tags in sql")
	}
	exp := sql[1+li : ri]
	if -1 != strings.LastIndex(exp, "}") || -1 != strings.LastIndex(exp, "}") {
		Throw("illegal 'fields' tags in sql")
	}
	fields := strings.Split(exp, ",")
	countParts := len(fields)
	if countParts < 1 {
		Throw("not enough arguments in sql 'fields'")
	}
	sql = sql[0:li] + exp + sql[ri+1:]

	lists := Db.Query(sql, args)
	rows := make([]map[string]interface{}, 0)
	for _, dataRow := range lists {
		row := make(map[string]interface{})
		if len(dataRow) != len(fields) {
			Throw("the number of parameters does not match")
		}
		for i := 0; i < len(fields); i++ {
			row[fields[i]] = dataRow[i]
		}
		rows = append(rows, row)
	}
	return rows
}

//////////////////////////////////// sql 执行语句
/*
"UPDATE user SET name = ?", "your"
*/
func (inst *_dataAccessLayer) Exec(sql string, args ...interface{}) bool {
	o := orm.NewOrm()
	res, err := o.Raw(sql, args).Exec()
	if err != nil {
		panic(err)
	}
	affected, _ := res.RowsAffected()
	return affected > 0
}

//////////////////////////////////// sql 获取数据列表
/*
lists := query("SELECT loginName,password FROM users WHERE loginName LIKE ?", "%zhang%")
for index, dataRow := range lists {
	fmt.Printf("index[%d]:", index)
	for _, value := range dataRow {
		fmt.Printf("%s,", value)
	}
	fmt.Println()
}
*/
func (inst *_dataAccessLayer) Query(sql string, args ...interface{}) []orm.ParamsList {
	o := orm.NewOrm()
	var lists []orm.ParamsList
	_, err := o.Raw(sql, args).ValuesList(&lists)
	if err != nil {
		panic(err)
	}
	return lists
}
