package zfx

import (
	"encoding/base64"
	"strconv"

	"math/rand"
	"time"

	"github.com/satori/go.uuid"
)

type _jobNumber struct {
}

func (inst *_jobNumber) New() string {
	// 67000000:11GB zhangxx @ 2017-07-08
	return base64.URLEncoding.EncodeToString([]byte(uuid.NewV4().String()))[0:8] + uuid.NewV4().String()[0:8]
}

func (inst *_jobNumber) VCode(length int) string {
	rand.Seed(time.Now().UnixNano())
	vcode := ""
	for i := 0; i < length; i++ {
		vcode += strconv.Itoa(rand.Intn(9))
	}
	return vcode
}
