package notify_center

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"sync"
	"testing"
)

type student struct {
	Name string
	Age  int
}

func Test_demo0(t *testing.T) {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}

func Test_demo1(t *testing.T) {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhao", Age: 24},
		{Name: "qian", Age: 23},
		{Name: "sun", Age: 22},
	}
	for _, stu := range stus {
		//fmt.Printf("%s %v \n", stu.Name, stu)
		m[stu.Name] = &stu
		//fmt.Printf("%#v \n\n", m[stu.Name])
	}

	fmt.Printf("%#v", m)
}

func Test_demo2(t *testing.T) {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("Ai: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("Bi: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func Test_demo3(t *testing.T) {
	//ints := make([]int, 5)
	//ints = append(ints, 1, 2, 3)
	//fmt.Println(ints)

	ints := make([]int, 5)
	// range 格式可以对 slice、map、数组、字符串等进行迭代循环
	for k, _ := range ints {
		ints[k] = 1
	}
	fmt.Println(ints)
}

func Test_demo4(t *testing.T) {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func Test_demo5(t *testing.T) {
	a := 1
	b := 2
	defer calc("1", b, calc("10", a, b))
	a = 0
	defer func() {
		calc("2", b, calc("20", a, b))
	}()
	b = 1
}

func Test_demo6(t *testing.T) {
	s := make([]int, 3)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

type threadSafeSet struct {
	sync.RWMutex
	s []interface{}
}

func (set *threadSafeSet) Add() {
	set.s = append(set.s, "3")
}
func (set *threadSafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		set.RLock()
		for _, value := range set.s {
			ch <- value
			//fmt.Printf("Iter: %v \n", value)
		}
		close(ch)
		set.RUnlock()
	}()
	return ch
}

func Test_demo7(t *testing.T) {
	th := threadSafeSet{
		s: []interface{}{"1", "2"},
	}
	for b := range th.Iter() {
		fmt.Println("b ===", b)
	}
}

func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}

func Test_demo8(t *testing.T) {
	println(DeferFunc1(1))
	println(DeferFunc2(1))
	println(DeferFunc3(1))
}

func Test_demo9(t *testing.T) {
	list := make([]int, 1)
	list = append(list, 1)
	fmt.Println(list)
}

func Test_demo10(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2...)
	fmt.Println(s1)
}

func Test_demo11(t *testing.T) {
	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	sm1 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}
	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	if reflect.DeepEqual(sm1, sm2) {
		fmt.Println("sm1 == sm2")
	} else {
		fmt.Println("sm1 != sm2")
	}
}

func Foo(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}
func Test_demo12(t *testing.T) {
	var x *int = nil
	Foo(x)
}

const (
	x = iota
	y
	z = "zz"
	k
	p = iota
)

func Test_demo13(t *testing.T) {
	fmt.Println(x, y, z, k, p)
}

func Test_demo14(t *testing.T) {
	type MyInt1 int
	type MyInt2 = int
	var i int = 9
	var i1 MyInt1 = MyInt1(i)
	var i2 MyInt2 = i
	fmt.Println(i1, i2)
}

type User struct {
}
type MyUser1 User
type MyUser2 = User

func (i MyUser1) m1() {
	fmt.Println("MyUser1.m1")
}
func (i User) m2() {
	fmt.Println("User.m2")
}

func Test_demo15(t *testing.T) {
	var i1 MyUser1
	var i2 MyUser2
	i1.m1()
	i2.m2()
}

type T1 struct {
}

func (t T1) m1() {
	fmt.Println("T1.m1")
}

type T2 = T1
type MyStruct struct {
	T1
	T2
}

func Test_demo16(t *testing.T) {
	my := MyStruct{}
	//my.m1()
	my.T1.m1()
	my.T2.m1()
}

var ErrDidNotWork = errors.New("did not work")

func DoTheThing(reallyDoIt bool) (err error) {
	var result string
	if reallyDoIt {
		result, err = tryTheThing()
		if err != nil || result != "it worked" {
			err = ErrDidNotWork
		}
	}
	return err
}

func tryTheThing() (string, error) {
	return "", ErrDidNotWork
}

func Test_demo17(t *testing.T) {
	fmt.Println(DoTheThing(true))
	fmt.Println(DoTheThing(false))
}

func test() []func() {
	var funs []func()
	for i := 0; i < 2; i++ {
		x := i
		funs = append(funs, func() {
			println(&x, x)
		})
	}
	return funs
}

func Test_demo18(t *testing.T) {
	funs := test()
	for _, f := range funs {
		f()
	}
}

func test1(x int) (func(), func()) {
	return func() {
			println(x)
			x += 10
		}, func() {
			println(x)
		}
}

func Test_demo19(t *testing.T) {
	a, b := test1(100)
	a()
	a()
	b()
}

func Test_demo20(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("fatal")
		}
	}()

	defer func() {
		panic("defer panic")
	}()
	panic("panic")
}

func Test_demo21(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("++++")
			f := err.(func() string)
			fmt.Println(err, f(), reflect.TypeOf(err).Kind().String())
		} else {
			fmt.Println("fatal")
		}
	}()

	defer func() {
		panic(func() string {
			return "defer panic"
		})
	}()
	panic("panic")
}

type Param map[string]interface{}

type Show struct {
	*Param
}

func Test_demo22(t *testing.T) {
	var param Param = make(Param)
	param["1"] = 1
	s := new(Show)
	s.Param = &param
	fmt.Println(*s.Param)
}

type People struct {
	Name string `json:"name"`
}

func Test_demo23(t *testing.T) {
	js := `{
        "name":"11"
    }`
	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("people: ", p)
}

func Test_demo24(t *testing.T) {
	five := []string{"Annie", "Betty", "Charley", "Doug", "Edward"}

	for _, v := range five {
		five = five[:2]
		fmt.Printf("v[%s]\n", v)
	}

	fmt.Println(five)
}

func MultipleParam(p ...interface{}) {
	fmt.Println("MultipleParam=", p)
}

func Test_demo25(t *testing.T) {
	iis := []int{1, 2, 3, 4}
	newParam := make([]interface{}, 0)
	newParam = append(newParam, "ssss")
	for _, v := range iis {
		newParam = append(newParam, v)
	}
	MultipleParam(newParam...)
}

func Test_demo26(t *testing.T) {
	iis := []int{1, 2, 3, 4}
	f := MultipleParam
	value := reflect.ValueOf(f)
	pps := make([]reflect.Value, 0, len(iis)+1)
	pps = append(pps, reflect.ValueOf("ssss"))
	for _, ii := range iis {
		pps = append(pps, reflect.ValueOf(ii))
	}
	value.Call(pps)
}
