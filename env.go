package kit

import (
	"os"
	"strings"
)

//获取环境变量值（不区分大小写）
func EnvExistsValue(name, value string) bool {
	v := os.Getenv(name)
	if v == "" {
		return false
	}
	v = strings.ToLower(v)
	value = strings.ToLower(value)
	vs := strings.Split(v, ",")
	for _, item := range vs {
		if item == value {
			return true
		}
	}
	return false
}

//是否允许进行任务调度（环境变量 ENVIRONMENT='cron'）
func EnvCanCron() bool {
	return EnvExistsValue("ENVIRONMENT", "cron")
}

//测试环境
func EnvIsTest() bool {
	return EnvExistsValue("ASPNETCORE_ENVIRONMENT", "Staging")
}

//UAT环境
func EnvIsUAT() bool {
	return EnvExistsValue("ASPNETCORE_ENVIRONMENT", "uat")
}
