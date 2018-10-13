package raw
import (
	"context"
	"fmt"
	"time"
	)

type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}


type Arith int
func (t *Arith) Mul(ctx context.Context, args Args, reply *Reply) error {
	reply.C = args.A * args.B
	return nil
}
func (t *Arith) Add(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A + args.B
	return nil
}
type Echo string
func (s *Echo) Echo(ctx context.Context, args string, reply *string) error {
	*reply = fmt.Sprintf("Hello %s !", args)
	return nil
}
type TimeS struct{}
func (s *TimeS) Time(ctx context.Context, args time.Time, reply *time.Time) error {
	*reply = time.Now()
	return nil
}
