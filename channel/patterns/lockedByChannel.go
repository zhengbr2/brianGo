package patterns

type MutexC struct {
	ch chan struct{}
}

func NewMutex() *MutexC {
	mu := &MutexC{make(chan struct{}, 1)}
	mu.ch <- struct{}{}
	return mu
}
func (m *MutexC) Lock() {
	<-m.ch
}
func (m *MutexC) Unlock() {
	select {
	case m.ch <- struct{}{}:
	default:
		panic("unlock of unlocked mutex")
	}
}
func (m *MutexC) TryLock() bool {
	select {
	case <-m.ch:
		return true
	default:
	}
	return false
}
func (m *MutexC) IsLocked() bool {
	return len(m.ch) == 0
}
