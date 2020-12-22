package base_tool

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"log"
	"math"
	"math/rand"
	"net/url"
	"runtime"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// 获取正在运行的函数名
// @param skip 函数调用层级，1：当前调用函数，2：当前函数上层调用函数，以此类推
func RunFuncName(skip ...int) string {
	if len(skip) == 0 {
		skip = []int{2}
	} else {
		skip[0] += 1
	}

	pc := make([]uintptr, 1)
	runtime.Callers(skip[0], pc)

	return runtime.FuncForPC(pc[0]).Name()
}

//URL字符解码
func UrlDecode(encodeStr string) string {
	unescapeUrl, _ := url.PathUnescape(encodeStr)
	return unescapeUrl
}

//PrintJSON 将struct序列化json打印日志
func PrintJSON(inter interface{}) {
	bt, _ := json.Marshal(inter)
	log.Println("json：", string(bt))
}

//生成md5字串
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//生成32位Guid字串
func GetGuid32() string {
	return strings.ReplaceAll(GetGuid36(), "-", "")
}

//生成36位Guid字串
func GetGuid36() string {
	return uuid.New().String()
}

//元转分
func YuanToFen(f float64) int {
	decimalValue := decimal.NewFromFloat(f)
	decimalValue = decimalValue.Mul(decimal.NewFromInt(100))
	res, _ := decimalValue.Float64()

	return int(res)
}

//获取随机字符串，指定长度
func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]byte, l)
	for i := 0; i < l; i++ {
		result[i] = bytes[r.Intn(len(bytes))]
	}
	return string(result)
}

//全角转换半角
func DBCtoSBC(s string) string {
	retstr := ""
	for _, i := range s {
		inside_code := i
		if inside_code == 12288 {
			inside_code = 32
		} else {
			inside_code -= 65248
		}
		if inside_code < 32 || inside_code > 126 {
			retstr += string(i)
		} else {
			retstr += string(inside_code)
		}
	}
	return retstr
}

// 计算地球上两点间距离
func EarthDistance(lat1, lng1, lat2, lng2 float64) float64 {
	radius := 6371.0 // 6378137
	rad := math.Pi / 180.0

	lat1 = lat1 * rad
	lng1 = lng1 * rad
	lat2 = lat2 * rad
	lng2 = lng2 * rad

	theta := lng2 - lng1
	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))

	return dist * radius
}
