package utils

/*******************
* Import
*******************/
import (
	"sync"
	"sync/atomic"

	"github.com/petermattis/goid"
)

/*******************
* Lock
*******************/
type S_Lock struct {
	mutex     sync.RWMutex
	lockGoID  int64
	lockCount int32
}

func (lock *S_Lock) Lock() {
	goID := goid.Get()
	if atomic.LoadInt64(&lock.lockGoID) == goID {
		// Déjà verrouillé par cette goroutine
		atomic.AddInt32(&lock.lockCount, 1)
		return
	}
	lock.mutex.Lock()
	atomic.StoreInt64(&lock.lockGoID, goID)
	atomic.StoreInt32(&lock.lockCount, 1)

}

func (lock *S_Lock) Unlock() {
	goID := goid.Get()
	if atomic.LoadInt64(&lock.lockGoID) == goID {
		if atomic.AddInt32(&lock.lockCount, -1) == 0 {
			atomic.StoreInt64(&lock.lockGoID, 0)
			lock.mutex.Unlock()
		}
	}
}
