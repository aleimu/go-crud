package util

import "time"

const (
	OK  = 1000
	ERR = 1500
)

var DayValueIntList = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
var DayValueStrList = []string{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0"}
var DayShowValue = "0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0"
var DayClickValue = "0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0"
var DayShowKey = "day:show:key:"
var DayClickKey = "day:click:key:"
var HourShowKey = "hour:show:key:"
var HourClickKey = "hour:click:key:"

var OneHour = 3600 * time.Second
var OneDay = time.Hour * 24
