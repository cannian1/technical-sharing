package service

import "net/http"

func learnGo() {
	// ...

	// 不稳定，可能会导致测试失败
	http.Post("...", "application/json", nil)

	// ...
}
