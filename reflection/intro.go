package main

import (
	"reflect"
	"fmt"
	"strconv"
)


type User struct{
	Name string
	Age int
}
func (p User) Greet() string{
	return "hello:" + p.Name + ";" + strconv.Itoa(p.Age)
}

func (u User) Print(prfix string){
	fmt.Printf("%s:Name is %s,Age is %d",prfix,u.Name,u.Age)
}

type Person struct{
	Name string
	Age int
}



func (p *User) Greet2() string{
	return "hello:" + p.Name + ";" + strconv.Itoa(p.Age)
}

func main() {
	u:= User{"张三",20}
	t:=reflect.TypeOf(u)
	fmt.Println(t)

	v:=reflect.ValueOf(u)
	fmt.Println(v)

	{
		v := reflect.ValueOf(&u)
		fmt.Println(v)
	}

	t2:=v.Type()
	fmt.Println(t2)
	fmt.Println(t2==t)    //true
	fmt.Println("kind:",t2.Kind())


	fmt.Printf("%T\n",u)
	fmt.Printf("%v\n",u)

	user,ok:=v.Interface().(User)

	fmt.Println("ok?:",ok)
	fmt.Println("user:",user)

	{
		user, ok := v.Interface().(Person)

		fmt.Println("person ok?:", ok)   //false
		fmt.Println("person:", user)  //person: { 0}
	}

	fmt.Println("println all the field")   //false
	for i:=0;i<t.NumField();i++ {
		fmt.Println(t.Field(i).Name)
	}

	fmt.Println("println all the Method")   //no Greet2()
	for i:=0;i<t.NumMethod() ;i++  {
		fmt.Println(t.Method(i).Name)
	}

	{
		x:=2
		v:=reflect.ValueOf(&x)
		v.Elem().SetInt(100)
		fmt.Println("new x:",x)
	}

	{
		mPrint:=v.MethodByName("Print")
		args:=[]reflect.Value{reflect.ValueOf("UserInfo")}
		rets:=mPrint.Call(args)
		if(len(rets)>0) {
			fmt.Println(rets[0])
		}
	}
	
	{
		mPrint:=v.MethodByName("Greet")
		args:=[]reflect.Value{}
		fmt.Println("\n");
		fmt.Println(mPrint.Call(args)[0])
	}

}
