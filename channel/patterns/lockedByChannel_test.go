package patterns

import "testing"

func TestLockedByChannel(t *testing.T) {

	defer func(){
		if e:=recover();e!=nil{
			t.Log("panic recovered:",e)
		}
	}()
	m := NewMutex()
	m.Lock()
	t.Log("m is locked:",m.IsLocked())
	m.Unlock()
	t.Log("m is Unlocked:",m.IsLocked())
	success:=m.TryLock()
	t.Log("is locked success:",success)
	success2:=m.TryLock()
	t.Log("is locked success again?",success2)
	t.Log("m is locked",m.IsLocked())
	m.Unlock()
	t.Log("m is locked after unlook()",m.IsLocked())
	//m.Unlock()   // this will trigger a un-recoverable fatalthrow()
	//panic("panic here")    //recoverable
	t.Log("suppose never reach here, fatalthrow() by m.Unlock()")
}