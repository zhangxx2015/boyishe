package sdk

import (
	//"errors"
	"fmt"
)

func SendVCode(phone string, vcode string) error {
	fmt.Println("send sms vcode:", vcode)
	//panic(errors.New("--------failed to send sms"))
	return nil
}
