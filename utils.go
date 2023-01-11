package grbtree

import (
	"bytes"
	"strconv"
)

func converToBianry(n int) string {
    result := ""
    for ; n > 0; n /= 2 {
        lsb := n % 2
        result = strconv.Itoa(lsb) + result     
    }
    return result
}

func StrCopy(s string, n int) string {
	var b bytes.Buffer
	for i := 0; i < n; i ++ {
		b.WriteString(s)
	}
	return b.String()
}

// 字符串 左边填充
func StrLeftFilling(s string, width int, filling string) string {
    f_count := width - len(s) 
    if f_count <= 0 {
        return s
    }else{
        return StrCopy(filling, f_count) + s
    }
}

// 字符串 右边填充
func StrRightFilling(s string, width int, filling string) string {
    f_count := width - len(s) 
    if f_count <= 0 {
        return s
    }else{
        return s + StrCopy(filling, f_count)
    }
}

// int64 转 16 进制
func Int64ToHexStr(n int64) string {
    i := int64(n)
    return "0x" + strconv.FormatInt(i, 16)
}