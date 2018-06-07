package zfx

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

func AppSettingString(key string) string {
	return beego.AppConfig.String(key)
}
func AppSettingInt(key string, def int) int {
	return beego.AppConfig.DefaultInt(key, def)
}
func AppSettingBool(key string, def bool) bool {
	return beego.AppConfig.DefaultBool(key, def)
}

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		beego.Debug(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func Get_internal() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops:" + err.Error())
		os.Exit(1)
	}
	ip := ""
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				//os.Stdout.WriteString(ipnet.IP.String() + "\n")
				ip = ipnet.IP.String()
				//break
			}
		}
	}
	//os.Exit(0)
	return ip
}

//func GetCurrentRouter(c BaseController) string {
//	scheme := "http://"
//	if c.Ctx.Request.TLS != nil {
//		scheme = "https://"
//	}
//	//fmt.Println("self:", scheme, c.Ctx.Request.Host, c.Ctx.Request.RequestURI)
//	//fmt.Println("self:", scheme, c.Ctx.Request.Host, AppSettingString("ufolder"))
//	resUrl := strings.Join([]string{scheme, c.Ctx.Request.Host, AppSettingString("ufolder"), resName}, "/")
//}

func Md5Sum(file string) string {
	f, err := os.Open(file)
	if err != nil {
		return ""
	}
	defer f.Close()
	r := bufio.NewReader(f)
	h := md5.New()
	_, err = io.Copy(h, r)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

func ToInt(obj interface{}) int {
	val, err := strconv.Atoi(obj.(string))
	if err != nil {
		panic(err)
	}
	return val
}

/*
type Person struct {
	Name   string
	Gender bool
	Age    int
}
p := Person{Name: "zhangxx", Gender: true, Age: 35}
fmt.Println(p)
maps := ToMap(p)
fmt.Println(maps)
*/
func ToMap(obj interface{}) map[string]interface{} {
	types := reflect.TypeOf(obj)
	value := reflect.ValueOf(obj)
	var data = make(map[string]interface{})

	for i := 0; i < types.NumField(); i++ {
		key := types.Field(i).Name
		val := value.Field(i).Interface()
		typ := reflect.TypeOf(val)

		if typ == reflect.TypeOf(BaseModel{}) {
			childs := ToMap(val)
			for key, value := range childs {
				//typ := reflect.TypeOf(value)
				//fmt.Println(typ, "@", key, ":", value)
				data[key] = value
			}
			continue
		}
		//fmt.Println(typ, "@", key, ":", val)
		data[key] = val
	}
	return data
}
