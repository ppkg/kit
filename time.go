package base_tool

import (
	"log"
	"time"
)

const (
	DATETIME_LAYOUT      = "2006-01-02 15:04:05"
	DATE_LAYOUT          = "2006-01-02"
	DATE_LAYOUT_SHORT_CN = "01月02日"
	DATE_LAYOUT_SHORT_EN = "01-02"
	TIME_LAYOUT          = "15:04:05"
	TIME_LAYOUT_SHORT    = "15:04"
)

//获取当前标准格式时间
func GetTimeNow(time2 ...time.Time) string {
	if len(time2) == 0 {
		return time.Now().Format(DATETIME_LAYOUT)
	}
	log.Println(RunFuncName())
	return time2[0].Format(DATETIME_LAYOUT)
}
