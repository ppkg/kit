package rp_kit

import (
	"fmt"
	"strconv"

	"github.com/gogrpc/glog"
	"github.com/shopspring/decimal"
)

//元转分
func YuanToFen(f interface{}) int {
	var decimalValue decimal.Decimal
	switch f.(type) {
	case float32:
		decimalValue = decimal.NewFromFloat32(f.(float32))
		decimalValue = decimalValue.Mul(decimal.NewFromInt(100))
	case float64:
		decimalValue = decimal.NewFromFloat(f.(float64))
		decimalValue = decimalValue.Mul(decimal.NewFromInt(100))
	case int32:
		decimalValue = decimal.NewFromInt32(f.(int32))
		decimalValue = decimalValue.Mul(decimal.NewFromInt(100))
	case int64:
		decimalValue = decimal.NewFromInt(f.(int64))
		decimalValue = decimalValue.Mul(decimal.NewFromInt(100))
	case int:
		decimalValue = decimal.NewFromInt(int64(f.(int)))
		decimalValue = decimalValue.Mul(decimal.NewFromInt(100))
	default:
		panic("can't suppot type")
	}

	res, _ := decimalValue.Float64()
	return int(res)
}

//分转元
func FenToYuan(i interface{}) float64 {
	var decimalValue decimal.Decimal
	switch i.(type) {
	case int32:
		decimalValue = decimal.NewFromInt32(i.(int32))
		decimalValue = decimalValue.Div(decimal.NewFromInt32(100))
	case int64:
		decimalValue = decimal.NewFromInt(i.(int64))
		decimalValue = decimalValue.Div(decimal.NewFromInt(100))
	case int:
		decimalValue = decimal.NewFromInt(int64(i.(int)))
		decimalValue = decimalValue.Div(decimal.NewFromInt(100))
	default:
		panic("can't suppot type")
	}

	res, _ := decimalValue.Float64()
	return res
}

//float32转float64
func Float32ToFloat64(f float32) float64 {
	str := fmt.Sprintf("%f", f)
	v, err := strconv.ParseFloat(str, 64)
	if err != nil {
		glog.Error("Float32ToFloat64 err:", err, f)
	}
	return v
}
