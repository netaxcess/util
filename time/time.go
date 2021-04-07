package time
import(
	"time"
)

/*日期字符串转时间戳
date日期字符串，默认是年月日2020-02-03
返回时间戳int64
例子:StrtoTime("2020-02-02")
*/
func StrtoTime(date string) int64 {
    //如果传的是年月日，没有时分秒，那么需要拼接
    if len(date) == 10 {
        date = date+" 00:00:00"
    }
	timeLayout := "2006-01-02 15:04:05"                             //转化所需模板
	loc, _ := time.LoadLocation("Local")                            //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, date, loc) //使用模板在对应时区转化为time.time类型
	return theTime.Unix()                                            //转化为时间戳 类型是int64
	
}

/*
获取当前系统时间戳
*/
func CurrentTime() int64 {
	return time.Now().Unix()
}

/*
时间戳转年月日
time_unix   : 当前系统时间戳
date_formate:要转换的格式
例子：Date(1605582958,"2006/01/02 15:04:05")
例子：Date(1605582958,"2006年01月02日 15:04:05")
*/

func Date(time_unix int64, date_format string) string {
	if date_format == "" {
		date_format = "2006-01-02 15:04:05"
	}
	return time.Unix(time_unix, 0).Format(date_format)
}

/*
获取指定某天的零点时间戳
date：时间日期
例子：DateZeroTime("2020-12-09")
*/

func DateZeroTime(date string) int64 {
    //如果是年月日时分秒，就截取年月日
    if len(date) !=10 {
        date = date[0:10]
    }

    the_time, err := time.ParseInLocation("2006-01-02", date, time.Local)
    if err == nil {
       return the_time.Unix()
    }
    return 0
}

/*
获取指定日期是本年度第几周
date：时间日期
例子：GetWeek("2020-12-09")或者GetWeek("2020-12-09 15:04:05")
返回：year年度，week星期
*/
func GetWeek(date string) (year, week int) {
    timeLayout := "2006-01-02 15:04:05"
    if len(date) == 10 {
        timeLayout = "2006-01-02"
    }
    loc, _ := time.LoadLocation("Local")
    tmp, _ := time.ParseInLocation(timeLayout, date, loc)
    return tmp.ISOWeek()
}