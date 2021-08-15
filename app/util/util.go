package util

import (
	"crypto/md5"
	"fmt"
	"time"
)


// 1970年01月01日00时00分00秒起至现在的秒数
func Unixtime() int32 {
	return int32(time.Now().Unix())
}

// 1970年01月01日00时00分00秒起至现在的毫秒数
func Millisecond() int64 {
	return time.Now().UnixNano() / 1e6
}

// 时间戳转年月日时分秒
func Time2Datetime(unixtime int64) string {
	return time.Unix(unixtime, 0).Format("2006-01-02 15:04:05")
}

// md5 散列
func Md5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

// 三目运算模拟
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}
