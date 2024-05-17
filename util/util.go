package util

var LunarMonthNameToMonth = map[string]int{
	"正":  1,
	"二":  2,
	"三":  3,
	"四":  4,
	"五":  5,
	"六":  6,
	"七":  7,
	"八":  8,
	"九":  9,
	"十":  10,
	"十一": 11,
	"十二": 12,
}

var IntToWeekday = map[int]string{
	0: "天",
	1: "一",
	2: "二",
	3: "三",
	4: "四",
	5: "五",
	6: "六",
}

var DateLayout = "2006-01-02"

func Contains(elems []string, v string) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}
