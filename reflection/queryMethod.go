package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type User struct {
	Name string
	Age  int
}

func (p User) Greet() string {
	return "hello:" + p.Name + ";" + strconv.Itoa(p.Age)
}

func (u User) Print(prfix string) {
	fmt.Printf("%s:Name is %s,Age is %d", prfix, u.Name, u.Age)
}

func (p *User) Greet2() string {
	return "hello:" + p.Name + ";" + strconv.Itoa(p.Age)
}


func main() {
	u := User{"张三", 20}
	t := reflect.TypeOf(u)
	v := reflect.ValueOf(u)
	fmt.Println(t,v)

	{//&{张三 20}
		v := reflect.ValueOf(&u)
		fmt.Println(v)
		fmt.Println("reflect.ValueOf(&u)==reflect.ValueOf(u)?:",reflect.ValueOf(&u)==reflect.ValueOf(u))   //false
	}
	{//{张三 20}
		v := reflect.ValueOf(&u).Elem()  //for interface or pointer
		fmt.Println(v)
		fmt.Println("reflect.ValueOf(&u).Elem()==reflect.ValueOf(u)?:",reflect.ValueOf(&u).Elem()==reflect.ValueOf(u)) //false!
	}

	user, ok := v.Interface().(User) // Value->Interface-> Type Instance
	//user, ok = v.(User)  // compile error
	fmt.Println("ok?:", ok, "user:", user)

	fmt.Println("println all the Method") //no Greet2()
	for i := 0; i < t.NumMethod(); i++ {
		fmt.Println(t.Method(i).Name)    //Greet, Print
	}

	t2 := reflect.TypeOf(&u)    // has Greet2()
	for i := 0; i < t2.NumMethod(); i++ {
		fmt.Println(t2.Method(i).Name)   //Greet, Greet2, Print
	}


	{
		mPrint := v.MethodByName("Print")
		Greet := v.MethodByName("Greet")   //Greet is not nil
		Greet2 := v.MethodByName("Greet2") //Greet2 is nil
		Greet2b := reflect.ValueOf(&u).MethodByName("Greet2") //Greet2b is NOT nil
		_,_,_=Greet,Greet2,Greet2b
		args := []reflect.Value{reflect.ValueOf("UserInfo")}
		rets := mPrint.Call(args)
		if len(rets) > 0 {
			fmt.Println(rets[0])
		}
	}

	{
		mPrint := v.MethodByName("Greet")
		args := []reflect.Value{}
		fmt.Println("\n")
		fmt.Println(mPrint.Call(args)[0])
	}

	{// set value by passing pointer's Value
		x := 2
		v := reflect.ValueOf(&x)
		v.Elem().SetInt(100)
		fmt.Println("new x:", x)
	}
}
