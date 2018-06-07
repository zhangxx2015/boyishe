package zfx

import (
	"encoding/json"
	"errors"

	"github.com/astaxie/beego"
)

type _apiClient struct {
}

type Dict map[string]interface{}

func (inst *_apiClient) ThisHost() string {
	return "http://localhost:" + beego.AppConfig.String("httpport")
}

func (inst *_apiClient) Post(path string, args Dict) (Resp, error) {
	t := Resp{}
	resp, err := HttpClient.Post(inst.ThisHost()+path, args)
	if err != nil {
		return t, err
	}
	err = json.Unmarshal([]byte(resp), &t)
	if err != nil {
		return t, err
	}
	if t.Status != "OK" {
		return t, errors.New(t.Error)
	}
	return t, nil
}
