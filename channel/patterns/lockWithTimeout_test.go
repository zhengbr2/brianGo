package patterns

import (
	"time"
	"testing"
)



func TestLockWithTimeout(t *testing.T) {

	defer func(){
		if e:=recover();e!=nil{
			t.Log("panic recovered:",e)
		}
	}()
	m := NewMutex2()
	m.Lock()
	t.Log("m is locked:",m.IsLocked())
	m.Unlock()
	t.Log("m is Unlocked:",m.IsLocked())
	success:=m.TryLock(time.Duration(time.Millisecond))
	t.Log("is locked success:",success)
	success2:=m.TryLock(time.Duration(time.Millisecond))
	t.Log("is locked success again?",success2)
	t.Log("m is locked",m.IsLocked())
	m.Unlock()
	t.Log("m is locked after unlook()",m.IsLocked())
	m.Unlock()   // this will trigger a recoverable panic

	t.Log("suppose never reach here, fatalthrow() by m.Unlock()")
}