// AUTOGENERATED FILE: rpcx server stub code
// compilable during generation.

package main

import (
  "flag"
  "time"

  metrics "github.com/rcrowley/go-metrics"
  "github.com/smallnest/rpcx/server"
  "github.com/smallnest/rpcx/serverplugin"
  "brianGo/web/rpcx/raw"
)


var (

	addr     = flag.String("addr", "localhost:8972", "server address")
	basePath = flag.String("base", "/rpcx", "prefix path")
)
	
func main() {
	flag.Parse()

	_ = time.Second
	_ = metrics.UseNilMetrics
	_ = serverplugin.GetFunctionName

	s := server.NewServer()
	addRegistryPlugin(s)

	registerServices(s)

	s.Serve("tcp", *addr)
}
func registerServices(s *server.Server) {
	s.Register(new(raw.Arith), "")
	s.Register(new(raw.Echo), "")
	s.Register(new(raw.TimeS), "")
}
func addRegistryPlugin(s *server.Server) {

		
}
