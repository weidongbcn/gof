package util

import (
	"strconv"
	"strings"
)

var (
	BoolExt *boolExtend   = &boolExtend{}
	StrExt  *stringExtend = &stringExtend{}
)

type (
	boolExtend struct {
	}
	stringExtend struct {
	}
)

func threeCondition(b bool, i1, i2 interface{}) interface{} {
	if b {
		return i1
	}
	return i2
}

func (e *boolExtend) TInt(b bool, v1, v2 int) int {
	return threeCondition(b, v1, v2).(int)
}
func (e *boolExtend) TInt32(b bool, v1, v2 int32) int32 {
	return threeCondition(b, v1, v2).(int32)
}
func (e *boolExtend) TString(b bool, v1, v2 string) string {
	return threeCondition(b, v1, v2).(string)
}

func (e *stringExtend) String(v interface{}) string {
	return Str(v)
}

// 字符串转为int32切片
func (e *stringExtend) I32Slice(s string, delimer string) []int32 {
	arr := []int32{}
	sArr := strings.Split(s, delimer)
	for _, v := range sArr {
		i, err := strconv.ParseInt(v, 10, 32)
		if err == nil {
			arr = append(arr, int32(i))
		}
	}
	return arr
}

// 字符串转为int切片
func (e *stringExtend) IntSlice(s string, delimer string) []int {
	arr := []int{}
	sArr := strings.Split(s, delimer)
	for _, v := range sArr {
		i, err := strconv.ParseInt(v, 10, 64)
		if err == nil {
			arr = append(arr, int(i))
		}
	}
	return arr
}
