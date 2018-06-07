package zfx

import (
	"fmt"
	"strconv"
)

type _rbac struct {
}

////type UserRoles int

//const (
//	//User UserRoles = 1 + iota
//	User = iota
//	Operation
//	Manager
//)
const (
	UserRoles_User = 1 + iota
	UserRoles_Operation
	UserRoles_Manager
)

type AccessInfo struct {
	LoginName string
	Role      int
	UserLevel int
}

func (inst *_rbac) AccessKeyIsValid(accessKey string) AccessInfo {
	//fmt.Println("AccessKeyIsValid()")
	accessInfo := AccessInfo{}
	Try(func() {
		row := Db.TopRow("SELECT {loginName,userRole,userLevel} FROM users WHERE accessKey = ? AND expires >= ?", accessKey, AbsTime.CurrUtc())
		if row == nil {
			Throw("DeniedAccess")
		}
		if v, err := strconv.Atoi(row["userRole"].(string)); err != nil {
			Throw("DeniedAccess")
		} else {
			accessInfo.Role = v
		}
		if v, err := strconv.Atoi(row["userLevel"].(string)); err != nil {
			Throw("DeniedAccess")
		} else {
			accessInfo.UserLevel = v
		}
		accessInfo.LoginName = row["loginName"].(string)
	}, Catch(func(err error) {
		panic(err)
	}))
	fmt.Println("accessInfo:", accessInfo)
	return accessInfo
}
