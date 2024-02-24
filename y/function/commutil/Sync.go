package commutil

import "sync"

func fastWorkerDo(wg *sync.WaitGroup, fn func()) {
	fn()
	wg.Done()
}

func FastWorker(fn ...func()) {
	fnNum := len(fn)
	if fnNum < 0 {
		return
	}
	var wg sync.WaitGroup
	wg.Add(fnNum)
	for _, f := range fn {
		go fastWorkerDo(&wg, f)
	}
	wg.Wait()
}
