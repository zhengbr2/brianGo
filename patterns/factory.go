package main

import "brianGo/patterns/factory"

func main(){
	s := data.NewStore(data.MemoryStorage)
	f, _ := s.Open("file")
	_, _ = f.Write([]byte("data"))
	defer f.Close()

}
