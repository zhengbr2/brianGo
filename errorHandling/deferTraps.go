package main

func f() (result int) {

	ret := 0
	defer func() {
		ret++ // only take effect when defer update the returned parameters
	}()
	return ret
}

func f1() (result int) {

	ret := 0
	defer func() {
		result++
	}()
	return ret
}

func f2() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5 //masked
	}(r)
	return 99
}

func main() {
	i := f()
	println("f() 0 or 1 ?:", i)

	i = f1()
	println("f1() 0 or 1 ?:", i)

	i = f2()
	println("f2() 0 or 1 ?:", i)

	i = f3()
	println("f3() 0 or 1 ?:", i)
}
