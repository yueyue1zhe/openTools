package core

import (
	"e.coding.net/zhechat/magic/taihao/library/logutil"
)

func GoSafe(fn func()) {
	go goSafeRun(fn)
}
func goSafeRun(fn func()) {
	defer SafeRunRecover()

	fn()
}
func SafeRunRecover(cleanups ...func()) {
	for _, cleanup := range cleanups {
		cleanup()
	}

	if p := recover(); p != nil {
		logutil.Error(p)
		logutil.Error(logutil.GetStack())
	}
}
