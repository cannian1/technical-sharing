package date

import (
	"time"
)

const (
	defaultMonthStyle = "2006-01"
	defaultDayStyle   = "2006-01-02"
)

type TimeOptionFunc func(*TimeOptions)

type TimeOptions struct {
	Format  string
	Reverse bool
}

func WithTimeReverse(reverse bool) TimeOptionFunc {
	return func(opts *TimeOptions) {
		opts.Reverse = reverse
	}
}

func WithFormat(format string) TimeOptionFunc {
	return func(opts *TimeOptions) {
		if _, err := time.Parse(format, format); err != nil {
			opts.Format = defaultMonthStyle
		} else {
			opts.Format = format
		}
	}
}

// generateTimeSlice 生成一个时间字符串的切片
//   - num: 要生成的时间字符串的数量(几天、几个月)
//   - generator 接受一个整数参数（表示偏移量）并返回一个 time.Time 类型的值。这个函数负责根据给定的偏移量生成时间
//   - format: 时间格式
//   - reverse: 是否逆序
func generateTimeSlice(num int, generator func(int) time.Time, format string, reverse bool) []string {
	if num <= 0 {
		return []string{}
	}

	times := make([]string, num)
	for i := 0; i < num; i++ {
		offset := i

		// 外部函数控制的闭包
		t := generator(offset).Format(format)
		if !reverse {
			times[num-1-i] = t
		} else {
			times[i] = t
		}
	}

	return times
}

// GenerateRecentMonth 生成最近几个月的月份（含本月）
func GenerateRecentMonth(num int, opts ...TimeOptionFunc) []string {

	options := &TimeOptions{}
	for _, opt := range opts {
		opt(options)
	}

	if options.Format == "" {
		options.Format = defaultMonthStyle
	}

	return generateTimeSlice(num, func(i int) time.Time {
		return time.Now().AddDate(0, -i, 0)
	}, options.Format, options.Reverse)
}

// GenerateRecentDays 生成最近几天的日期(含今日)
func GenerateRecentDays(num int, opts ...TimeOptionFunc) []string {

	options := &TimeOptions{}
	for _, opt := range opts {
		opt(options)
	}

	if options.Format == "" {
		options.Format = defaultDayStyle
	}

	return generateTimeSlice(num, func(i int) time.Time {
		return time.Now().AddDate(0, 0, -i)
	}, options.Format, options.Reverse)
}
