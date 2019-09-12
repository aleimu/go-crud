package util

import (
	"reflect"
	"regexp"
)

/************************* 自定义类型 ************************/
//数字+字母  不限制大小写 6~30位
func IsID(str ...string) bool {
	var b bool
	for _, s := range str {
		b, _ = regexp.MatchString("^[0-9a-zA-Z]{6,30}$", s)
		if false == b {
			return b
		}
	}
	return b
}

//数字+字母+符号 6~30位
func IsPwd(str ...string) bool {
	var b bool
	for _, s := range str {
		b, _ = regexp.MatchString("^[0-9a-zA-Z@.]{6,30}$", s)
		if false == b {
			return b
		}
	}
	return b
}

/************************* 数字类型 ************************/
//纯整数
func IsInteger(str ...string) bool {
	var b bool
	for _, s := range str {
		b, _ = regexp.MatchString("^[0-9]+$", s)
		if false == b {
			return b
		}
	}
	return b
}

//纯小数
func IsDecimals(str ...string) bool {
	var b bool
	for _, s := range str {
		b, _ = regexp.MatchString("^\\d+\\.[0-9]+$", s)
		if false == b {
			return b
		}
	}
	return b
}

//手提电话（不带前缀）最高11位
func IsMobile(str ...string) bool {
	var b bool
	for _, s := range str {
		b, _ = regexp.MatchString("^1[0-9]{10}$", s)
		if false == b {
			return b
		}
	}
	return b
}

//家用电话（不带前缀） 最高8位
func IsTelephone(str ...string) bool {
	var b bool
	for _, s := range str {
		b, _ = regexp.MatchString("^[0-9]{8}$", s)
		if false == b {
			return b
		}
	}
	return b
}

/************************* 英文类型 *************************/
//仅小写
func IsEngishLowCase(str ...string) bool {
	var b bool
	for _, s := range str {
		b, _ = regexp.MatchString("^[a-z]+$", s)
		if false == b {
			return b
		}
	}
	return b
}

//仅大写
func IsEnglishCap(str ...string) bool {
	var b bool
	for _, s := range str {
		b, _ = regexp.MatchString("^[A-Z]+$", s)
		if false == b {
			return b
		}
	}
	return b
}

//大小写混合
func IsEnglish(str ...string) bool {
	var b bool
	for _, s := range str {
		b, _ = regexp.MatchString("^[A-Za-z]+$", s)
		if false == b {
			return b
		}
	}
	return b
}

func Match(p string, s string) bool {
	b, _ := regexp.MatchString(p, s)
	return b
}

//邮箱 最高30位
func IsEmail(str ...string) bool {
	var b bool
	for _, s := range str {
		b, _ = regexp.MatchString("^([a-z0-9_\\.-]+)@([\\da-z\\.-]+)\\.([a-z\\.]{2,6})$", s)
		if false == b {
			return b
		}
	}
	return b
}

// 是否是email
func IsEmail2(email string) bool {
	if email == "" {
		return false
	}
	ok, _ := regexp.MatchString(`^([a-zA-Z0-9]+[_|\_|\.]?)*[a-zA-Z0-9]+@([a-zA-Z0-9]+[_|\_|\.]?)*[a-zA-Z0-9]+\.[0-9a-zA-Z]{2,3}$`, email)
	return ok
}

//是否Map类型
func IsMap(v interface{}) bool {
	return reflect.ValueOf(&v).Elem().Elem().Kind() == reflect.Map
}
