package patterns

import "time"

type MutexTM struct {
	ch chan struct{}
}

func NewMutex2() *MutexTM {
	mu := &MutexTM{make(chan struct{}, 1)}
	mu.ch <- struct{}{}
	return mu
}
func (m *MutexTM) Lock() {
	<-m.ch
}
func (m *MutexTM) Unlock() {
	select {
	case m.ch <- struct{}{}:
	default:
		panic("unlock of unlocked mutex")
	}
}
func (m *MutexTM) TryLock(timeout time.Duration) bool {
	timer := time.NewTimer(timeout)
	select {
	case <-m.ch:
		timer.Stop()
		return true
	case <-timer.C:
	}
	return false
}
func (m *MutexTM) IsLocked() bool {
	return len(m.ch) == 0
}
