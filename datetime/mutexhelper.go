package datetime

import "sync"


// mutexHelper - This type encapsulates the sync.RWMutex
// type in order to avoid 'pass by value' warnings for those
// core types requiring Mutex Locking services.
//
type mutexHelper struct {
	lock *sync.Mutex
}

// Initialize - Initializes private member
// 'lock'.
func (mtxHelper *mutexHelper) Initialize() {

	if mtxHelper.lock == nil {
		mtxHelper.lock = new(sync.Mutex)
	}

}

func (mtxHelper *mutexHelper) Lock() {

	if mtxHelper.lock == nil {
		mtxHelper.lock = new(sync.Mutex)
	}

	mtxHelper.lock.Lock()
}

func (mtxHelper *mutexHelper) New() *mutexHelper {

	newMtxHelper := mutexHelper{}

	newMtxHelper.lock = new(sync.Mutex)

	return &newMtxHelper
}

func (mtxHelper *mutexHelper) Unlock() {

	if mtxHelper.lock == nil {
		mtxHelper.lock = new(sync.Mutex)
		return
	}

	mtxHelper.lock.Unlock()
}

