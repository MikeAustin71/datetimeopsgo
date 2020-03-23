package datetime

import "sync"

// rWMutexHelper - This type encapsulates the sync.RWMutex
// type in order to avoid 'pass by value' warnings for those
// core types requiring Read/Write Mutex services.
//
type rWMutexHelper struct {
	lock *sync.RWMutex
}

func (rWMtxHelper *rWMutexHelper) New() *rWMutexHelper {

	newMtxHelper := rWMutexHelper{}

	newMtxHelper.lock = new(sync.RWMutex)

	return &newMtxHelper
}

func (rWMtxHelper *rWMutexHelper) Lock() {

	if rWMtxHelper.lock == nil {
		rWMtxHelper.lock = new(sync.RWMutex)
	}

	rWMtxHelper.lock.Lock()
}

func (rWMtxHelper *rWMutexHelper) RLock() {

	if rWMtxHelper.lock == nil {
		rWMtxHelper.lock = new(sync.RWMutex)
	}

	rWMtxHelper.lock.RLock()
}

func (rWMtxHelper *rWMutexHelper) RUnlock() {

	if rWMtxHelper.lock == nil {
		rWMtxHelper.lock = new(sync.RWMutex)
		return
	}

	rWMtxHelper.lock.RUnlock()
}

func (rWMtxHelper *rWMutexHelper) Unlock() {

	if rWMtxHelper.lock == nil {
		rWMtxHelper.lock = new(sync.RWMutex)
		return
	}

	rWMtxHelper.lock.Unlock()
}

