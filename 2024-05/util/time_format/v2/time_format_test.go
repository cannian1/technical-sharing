package time_format

import (
	"fmt"
	"testing"
)

func TestGenerateRecentDays(t *testing.T) {

	got := GenerateRecentDays(7, WithTimeReverse(false), WithFormat("2006-01-02"))
	fmt.Println(got)

}

func TestGenerateRecentMonth(t *testing.T) {
	got := GenerateRecentMonth(6, WithTimeReverse(false), WithFormat("200"))
	fmt.Println(got)
}
