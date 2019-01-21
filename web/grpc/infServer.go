package main

import (
	"log"
	"net"
	"runtime"
	"strconv"

	"brianGo/web/grpc/inf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = "41005"
)

type Data struct{}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	//起服务
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	rpcsvr := grpc.NewServer()
	inf.RegisterDataServer(rpcsvr, &Data{})
	reflection.Register(rpcsvr)
	if err = rpcsvr.Serve(lis); err != nil {
		log.Fatalln("faile serve at: " + ":" + port)
	}

	log.Println("grpc server in: %rpcsvr", port)

}

// 定义方法
func (t *Data) GetUser(ctx context.Context, request *inf.UserRq) (response *inf.UserRp, err error) {
	response = &inf.UserRp{
		Name: strconv.Itoa(int(request.Id)) + ":test",
	}
	return response, err
}
