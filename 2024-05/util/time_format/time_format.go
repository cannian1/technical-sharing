package time_format

import "time"

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
		opts.Format = format
	}
}

// GenerateRecentMonth 生成最近几个月的月份（含本月）
func GenerateRecentMonth(num int, format string, opts ...TimeOptionFunc) []string {

	options := &TimeOptions{Format: format}
	for _, opt := range opts {
		opt(options)
	}

	months := make([]string, 0, num)
	now := time.Now()

	if len(options.Format) < 6 {
		options.Format = "2006-01"
	}

	if options.Reverse {
		for i := 0; i < num; i++ {
			month := now.AddDate(0, -i, 0).Format(options.Format)
			months = append(months, month)
		}
	} else {
		for i := num - 1; i >= 0; i-- {
			month := now.AddDate(0, -i, 0).Format(options.Format)
			months = append(months, month)
		}
	}

	return months
}

// GenerateRecentDays 生成最近几天的日期(含今日)
func GenerateRecentDays(num int, opts ...TimeOptionFunc) []string {

	options := &TimeOptions{}
	for _, opt := range opts {
		opt(options)
	}

	days := make([]string, 0, num)
	now := time.Now()

	if options.Format == "" {
		options.Format = "2006-01-02"
	}

	if options.Reverse {
		for i := 0; i < num; i++ {
			day := now.AddDate(0, 0, -i).Format(options.Format)
			days = append(days, day)
		}
	} else {
		for i := num - 1; i >= 0; i-- {
			day := now.AddDate(0, 0, -i).Format(options.Format)
			days = append(days, day)
		}
	}

	return days
}
