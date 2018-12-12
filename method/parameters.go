package main

func main (){

	println( sum(1,2,3,4,5))
}

func sum( args ...int) int {

	//abc :=[3] int{1,2,3}
	//abc=append(abc, 3)

	s:=0
	//args=append(args, 6)  //args is slice

	for _,a:=range args{
		s=s+a
	}
	return s
}
