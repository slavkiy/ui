package reactive

import "sync"

var (
	batchMu    sync.Mutex
	batchDepth int
	batchQueue []func()
)

// Batch groups notifications until the outermost batch completes.
func Batch(fn func()) {
	if fn == nil {
		return
	}
	batchMu.Lock()
	batchDepth++
	batchMu.Unlock()
	defer func() {
		batchMu.Lock()
		batchDepth--
		if batchDepth != 0 {
			batchMu.Unlock()
			return
		}
		queue := batchQueue
		batchQueue = nil
		batchMu.Unlock()
		for _, update := range queue {
			update()
		}
	}()
	fn()
}

func enqueue(fn func()) {
	batchMu.Lock()
	if batchDepth > 0 {
		batchQueue = append(batchQueue, fn)
		batchMu.Unlock()
		return
	}
	batchMu.Unlock()
	fn()
}
