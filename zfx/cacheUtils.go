package zfx

import (
	"time"

	"github.com/astaxie/beego/cache"
)

type _memoryCache struct {
}

var bm cache.Cache

func init() {
	bm, _ = cache.NewCache("memory", `{"interval":60}`) // 实例化
}

//////////////////////////////////// 缓存
func (inst *_memoryCache) Write(key string, val interface{}, timeoutMillisecond int64) error {
	return bm.Put(key, val, time.Duration(timeoutMillisecond)*time.Millisecond) // 创建(键,值,生命周期)
}

func (inst *_memoryCache) Read(key string) interface{} {
	return bm.Get(key) // 取值
}

func (inst *_memoryCache) Exist(key string) bool {
	return bm.IsExist(key) // 判断键是否存在
}

func (inst *_memoryCache) Delete(key string) error {
	return bm.Delete("astaxie11") // 删除键
}
