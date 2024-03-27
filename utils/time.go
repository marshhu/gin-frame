package utils

import (
	"fmt"
	"math"
	"time"
)

const (
	//时间模型
	AUTO_TIME_ZONE  = true
	DATE_LAYOUT     = "2006-01-02"
	TIME_LAYOUT     = "15:04:05"
	DATETIME_LAYOUT = "2006-01-02 15:04:05"
	TimeLayoutDate  = "2006/01/02"
)

// StringToTime 字符串转时间
func StringToTime(val string) (time.Time, error) {
	t, err := DateTimeFrom(val)
	if err != nil {
		t, err = DateFrom(val)
	}
	return t, err
}

// FormatDate 将time.Time时间格式化成"2006-01-02"
func FormatDate(t time.Time) string {
	if t.IsZero() {
		return ""
	} else {
		return t.Format(DATE_LAYOUT)
	}
}

// FormatTimeDefault 将time.Time时间格式化成"2006-01-02 15:04:05"
func FormatTimeDefault(t time.Time) string {
	if t.IsZero() {
		return ""
	} else {
		return t.Format(DATETIME_LAYOUT)
	}
}

// FormatTime 将time.Time时间格式化成"20060102150405"、"2006-01-02 15:04:05"
func FormatTime(t time.Time, format string) string {
	if t.IsZero() {
		return ""
	} else {
		return t.Format(format)
	}
}

// DateFrom date from
func DateFrom(val string) (time.Time, error) {
	return timeParse(val, DATE_LAYOUT)
}

// TimeFrom time from
func TimeFrom(val string) (time.Time, error) {
	return timeParse(val, TIME_LAYOUT)
}

// DateTimeFrom datetime from
func DateTimeFrom(val string) (time.Time, error) {
	return timeParse(val, DATETIME_LAYOUT)
}

// DateTo data to
func DateTo(val time.Time) string {
	return timeFormat(val, DATE_LAYOUT)
}

// TimeTo time to
func TimeTo(val time.Time) string {
	return timeFormat(val, TIME_LAYOUT)
}

// DateTimeTo datetime to
func DateTimeTo(val time.Time) string {
	return timeFormat(val, DATETIME_LAYOUT)
}

// timeParse
func timeParse(val string, layout string) (time.Time, error) {
	if s, err := time.Parse(layout, val); err == nil {
		if AUTO_TIME_ZONE {
			var duration time.Duration
			_, offset := time.Now().Zone()
			if offset > 0 {
				tz := math.Ceil(float64(offset / 3600))
				if tz > 0 {
					duration = time.Duration(-tz)
				} else {
					duration = time.Duration(tz)
				}
			}

			return s.Add(time.Hour * duration).Local(), nil
		}

		return s, nil
	} else {
		return time.Now(), err
	}
}

// timeFormat
func timeFormat(val time.Time, layout string) string {
	if val.IsZero() {
		return ""
	}
	return val.Format(layout)
}

// ToDataTimeString 时间转换成string
func ToDataTimeString(time time.Time, splitter string) string {
	if splitter == "" {
		splitter = "-"
	}
	var localTime = time.Local()
	var layout = fmt.Sprintf("2006%s01%s02 15:04:05", splitter, splitter)
	return localTime.Format(layout)
}

// ToDateString 时间转换成string
func ToDateString(time time.Time, splitter string) string {
	if splitter == "" {
		splitter = "-"
	}
	var localTime = time.Local()
	var layout = fmt.Sprintf("2006%s01%s02", splitter, splitter)
	return localTime.Format(layout)
}

// ToDocPublishTimeVersion 格式化时间
func ToDocPublishTimeVersion(time time.Time) string {
	var localTime = time.Local()
	var layout = fmt.Sprintf("20060102150405")
	return fmt.Sprintf("v_%s", localTime.Format(layout))
}

// TimeParse string转换time
func TimeParse(val string, layout string) time.Time {
	if s, err := time.Parse(layout, val); err == nil {
		var duration time.Duration
		_, offset := time.Now().Zone()
		if offset > 0 {
			tz := math.Ceil(float64(offset / 3600))
			if tz > 0 {
				duration = time.Duration(-tz)
			} else {
				duration = time.Duration(tz)
			}
		}
		return s.Add(time.Hour * duration).Local()
	} else {
		return time.Now().UTC()
	}
}

func ToLocationDateTime(utcTime time.Time, location string) string {
	// 加载指定时区
	loc, err := time.LoadLocation(location)
	if err != nil {
		return utcTime.Format("2006-01-02 15:04:05")
	}
	// 转为指定时区时间
	localTime := utcTime.In(loc)
	return localTime.Format("2006-01-02 15:04:05")
}

func ToChinaDateTime(utcTime time.Time) string {
	// 加载指定时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return utcTime.Format("2006-01-02 15:04:05")
	}
	// 转为指定时区时间
	localTime := utcTime.In(loc)
	formatTimeStr := localTime.Format("2006-01-02 15:04:05")
	return formatTimeStr
}
