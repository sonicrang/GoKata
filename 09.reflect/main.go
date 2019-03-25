package main

import (
	"fmt"
	"reflect"
)

type People struct {
	name string
	sex  string
}

func (o People) ShowName() {
	fmt.Println(o.name)
}

type Man struct {
	People
	beard string
}

func (o Man) ShowName() {
	fmt.Println("我叫" + o.name)
}

func (o Man) Shave(shaver string) {
	if shaver == "" {
		fmt.Println("shaver is nil")
	} else {
		fmt.Println(o.name + "使用" + shaver + "刮" + o.beard)
	}
}

func Reflect(o interface{}) {
	t := reflect.TypeOf(o)
	v := reflect.ValueOf(o)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	fmt.Println("TypeName:", t.Name())
	fmt.Println("\nField:")

	numFiled := t.NumField()
	for i := 0; i < numFiled; i++ {
		f := t.Field(i)
		val := v.Field(i)
		fmt.Printf("%s(%s) : %s\n", f.Name, f.Type, val)
	}

	fmt.Println("\nMethod:")
	numMethod := t.NumMethod()
	for i := 0; i < numMethod; i++ {
		method := t.Method(i)
		fmt.Printf("%s:%s\n", method.Name, method.Type)
	}

	fmt.Println("\nCall Method ShowName:")
	method := v.MethodByName("ShowName")
	if method.IsValid() {
		method.Call(nil)
	} else {
		fmt.Println("Can not find method ShowName!")
	}

	fmt.Println("\nCall Method People.ShowName:")

	method = v.FieldByName("People").MethodByName("ShowName")
	if method.IsValid() {
		method.Call(nil)
	} else {
		fmt.Println("Can not find method People.ShowName!")
	}

	fmt.Println("\nCall Method Shave:")
	method = v.MethodByName("Shave")
	if method.IsValid() {
		method.Call([]reflect.Value{reflect.ValueOf("飞利浦")})
	} else {
		fmt.Println("Can not find method ShowName!")
	}

}

func main() {
	man := Man{People{"张三", "男"}, "络腮胡"}
	fmt.Println("--- Value Object ---")
	Reflect(man)
	fmt.Println("--- Point Object ---")
	Reflect(&man)
}
