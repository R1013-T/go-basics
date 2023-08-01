package main

import (
	"errors"
	"fmt"
	"go-basics/calculator"
	"os"
	"strings"
	"unsafe"

	"golang.org/x/exp/constraints"
)

type Os int

const (
	Mac Os = iota + 1
	Windows
	Linux
)

type Task struct {
	Title    string
	Estimate int
}

type controller interface {
	speedUp() int
	speedDown() int
}
type vehicle struct {
	speed       int
	enginePower int
}
type bicycle struct {
	speed      int
	humanPower int
}

func (v *vehicle) speedUp() int {
	v.speed += 10 * v.enginePower
	return v.speed
}
func (v *vehicle) speedDown() int {
	v.speed -= 10 * v.enginePower
	return v.speed
}

func (b *bicycle) speedUp() int {
	b.speed += 10 * b.humanPower
	return b.speed
}
func (b *bicycle) speedDown() int {
	b.speed -= 10 * b.humanPower
	return b.speed
}

func speedUpAndDown(c controller) {
	fmt.Printf("current speed: %d\n", c.speedUp())
	fmt.Printf("current speed: %d\n", c.speedDown())
}

func (v vehicle) String() string {
	return fmt.Sprintf("speed: %d, enginePower: %d", v.speed, v.enginePower)
}

var ErrCustom = errors.New("custom error")


type customConstrains interface {
	~int | int16 | float32 | float64 | string
}
type NewInt int

func add[T customConstrains](x, y T) T {
	return x + y
}
func min[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func main() {
	//? 6. module・package
	godotenv.Load()
	fmt.Println(os.Getenv("GO_ENV"))
	fmt.Println(calculator.Offset)
	fmt.Println(calculator.Sum(1, 2))
	fmt.Println(calculator.Multiply(1, 2))

	//? 7. variables
	i := int(2)
	ui := uint(4)

	fmt.Printf("i: %v %T\n", i, i)
	fmt.Printf("i: %[1]v %[1]T ui: %[2]v %[2]T\n", i, ui)

	f := 1.23456
	s := "hello"
	b := true
	fmt.Printf("f: %v %T\n", f, f)
	fmt.Printf("s: %v %T\n", s, s)
	fmt.Printf("b: %v %T\n", b, b)

	pi, title := 3.14, "Go"
	fmt.Printf("pi: %v %T\n", pi, pi)
	fmt.Printf("title: %v %T\n", title, title)

	x := 10
	y := 1.23
	z := float64(x) + y
	fmt.Printf("z: %v %T\n", z, z)

	fmt.Println(Mac, Windows, Linux)

	fmt.Println("i: ", i)

	//? 8. pointer・shadowing
	var ut1 uint16
	fmt.Printf("memory address of ut1: %p\n", &ut1)
	var ut2 uint16
	fmt.Printf("memory address of ut2: %p\n", &ut2)
	var p1 *uint16
	fmt.Printf("value of p1: %v\n", p1)
	p1 = &ut1
	fmt.Printf("value of p1: %v\n", p1)
	fmt.Printf("size of p1: %d[bytes]\n", unsafe.Sizeof(p1))
	fmt.Printf("memory address of p1: %p\n", &p1)
	fmt.Printf("valie of ut1(dereference): %v\n", *p1)
	*p1 = 1
	fmt.Printf("valie of ut1: %v\n", ut1)
	fmt.Printf("valie of ut1(dereference): %v\n", *p1)

	var pp1 **uint16 = &p1
	fmt.Printf("value of pp1: %v\n", pp1)
	fmt.Printf("memory address of pp1: %p\n", &pp1)
	fmt.Printf("size of pp1: %d[bytes]\n", unsafe.Sizeof(pp1))
	fmt.Printf("valie of p1(dereference): %v\n", *pp1)
	fmt.Printf("valie of ut1(dereference): %v\n", **pp1)

	ok, result := true, "A"
	if ok {
		result := "B"
		println(result)
	} else {
		result := "C"
		println(result)
	}
	println(result)

	//? 9. slice・map
	var a1 [2]int
	var a2 = [3]int{1, 2, 3}
	var a3 = [...]int{1, 2, 3}
	var a4 = [...]int{0: 1, 1: 2, 2: 3}
	a5 := [...]int{1, 2, 3}
	fmt.Printf("%v %v %v %v %v\n", a1, a2, a3, a4, a5)

	fmt.Printf("len: %d\n", len(a1))
	fmt.Printf("cap: %d\n", cap(a1))

	fmt.Printf("%T %T\n", a1, a2)

	var s1 []int
	s2 := []int{}
	fmt.Printf("s1: %[1]T %[1]v %v %v\n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %[1]T %[1]v %v %v\n", s2, len(s2), cap(s2))
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	s1 = append(s1, 1, 2, 3)
	fmt.Printf("s1: %[1]T %[1]v %v %v\n", s1, len(s1), cap(s1))

	s3 := []int{4, 5, 6}
	s1 = append(s1, s3...)
	fmt.Printf("s1: %[1]T %[1]v %v %v\n", s1, len(s1), cap(s1))

	s4 := make([]int, 0, 2)
	fmt.Printf("s4: %[1]T %[1]v %v %v\n", s4, len(s4), cap(s4))
	s4 = append(s4, 1, 2, 3)
	fmt.Printf("s4: %[1]T %[1]v %v %v\n", s4, len(s4), cap(s4))

	s5 := make([]int, 4, 6)
	fmt.Printf("s5: %[1]T %[1]v %v %v\n", s5, len(s5), cap(s5))
	s6 := s5[1:3]
	s6[1] = 10
	fmt.Printf("s5: %[1]T %[1]v %v %v\n", s5, len(s5), cap(s5))
	fmt.Printf("s6: %[1]T %[1]v %v %v\n", s6, len(s6), cap(s6))
	s6 = append(s6, 2)
	fmt.Printf("s5: %[1]T %[1]v %v %v\n", s5, len(s5), cap(s5))
	fmt.Printf("s6: %[1]T %[1]v %v %v\n", s6, len(s6), cap(s6))

	sc6 := make([]int, len(s5[1:3])) // capを指定しない場合は元のcapを引き継ぐ
	fmt.Printf("s5: %[1]T %[1]v %v %v\n", s5, len(s5), cap(s5))
	fmt.Printf("sc6: %[1]T %[1]v %v %v\n", sc6, len(sc6), cap(sc6))
	copy(sc6, s5[1:3])
	fmt.Printf("sc6: %[1]T %[1]v %v %v\n", sc6, len(sc6), cap(sc6))
	sc6[1] = 1
	fmt.Printf("s5: %[1]T %[1]v %v %v\n", s5, len(s5), cap(s5))
	fmt.Printf("sc6: %[1]T %[1]v %v %v\n", sc6, len(sc6), cap(sc6))

	var m1 map[string]int
	m2 := map[string]int{}
	fmt.Printf("m1: %[1]T %[1]v %v\n", m1, m1 == nil)
	fmt.Printf("m2: %[1]T %[1]v %v\n", m2, m2 == nil)
	m2["A"] = 10
	m2["B"] = 20
	m2["C"] = 30
	fmt.Printf("m2: %v %v %v\n", m2, len(m2), m2["A"])

	delete(m2, "A")
	fmt.Printf("m2: %v %v %v\n", m2, len(m2), m2["A"])
	v, ok := m2["A"]
	fmt.Printf("m2: %v %v %v %v\n", m2, len(m2), v, ok)
	v, ok = m2["B"]
	fmt.Printf("m2: %v %v %v %v\n", m2, len(m2), v, ok)

	for k, v := range m2 {
		fmt.Printf("%v: %v\n", k, v)
	//}

	//? 10. struct・receiver
	task1 := Task{
		Title:    "Task 1",
		Estimate: 3,
	}
	task1.Title = "Task 1+"
	fmt.Printf("%[1]T %+[1]v\n", task1)

	var task2 Task = task1
	task2.Title = "Task 2"
	fmt.Printf("%[1]T %+[1]v\n", task1)
	fmt.Printf("%[1]T %+[1]v\n", task2)

	task1p := &Task{
		Title:    "Task 1p",
		Estimate: 3,
	}
	fmt.Printf("task1p: %T %+v %v\n", task1p, task1p, unsafe.Sizeof(task1p))
	//(*task1p).Title = "Task 1p+" // 省略可能
	task1p.Title = "Task 1p+"
	fmt.Printf("task1p: %T %+v %v\n", task1p, task1p, unsafe.Sizeof(task1p))

	var task2p *Task = task1p
	task2p.Title = "Task 2p"
	fmt.Printf("task1: %+v\n", *task1p)
	fmt.Printf("task2: %+v\n", *task2p)

	task1.extendEstimate()
	fmt.Printf("task1: %+v\n", task1)
	//(&task1).extendEstimatePointer() // 省略可能
	task1.extendEstimatePointer()
	fmt.Printf("task1: %+v\n", task1)

	//? 11. function・closure
	funcDefer()
	files := []string{"a.txt", "b.txt", "c.txt"}
	fmt.Println(trimExtension(files...))
	name, err := fileChecker("a.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(name)

	i := 1
	func(i int) {
		fmt.Println(i)
	}(i)
	// 1

	f1 := func(i int) int {
		return i + 1
	}
	fmt.Println(f1(i))
	// 2

	f2 := func(file string) string {
		return file + ".txt"
	}
	fmt.Println(f2("a"))

	f3 := multiply()
	fmt.Println(f3(2))

	f4 := countUp()
	for i := 1; i <= 5; i++ {
		v := f4(2)
		fmt.Println(v)
	}

	//? 12. interface
	v := &vehicle{0, 5}
	speedUpAndDown(v)
	b := &bicycle{0, 1}
	speedUpAndDown(b)
	fmt.Println(v)

	var i1 interface{}
	var i2 any
	fmt.Printf("%[1]T %[1]v %v\n", i1, unsafe.Sizeof(i1))
	fmt.Printf("%[1]T %[1]v %v\n", i2, unsafe.Sizeof(i2))
	i1 = 10
	i2 = "A"
	checkType(i1)
	checkType(i2)

	//? 13. if・for・switch
		a := -1
		if a == 0 {
			fmt.Println("a is 0")
		} else if a > 0 {
			fmt.Println("a is positive")
		} else {
			fmt.Println("a is negative")
		}

		for i := 0; i < 5; i++ {
			fmt.Println(i)
		}
		//for {
		//	fmt.Println("infinite loop")
		//	time.Sleep(1 * time.Second)
		//}
		//var i int
		//for {
		//	if i > 3 {
		//		break
		//	}
		//	fmt.Println(i)
		//	i++
		//	time.Sleep(1 * time.Second)
		//}

	loop:
		for i := 0; i < 10; i++ {
			switch i {
			case 2:
				continue
			case 3:
				continue
			case 8:
				break loop
			default:
				fmt.Printf("%v ", i)
			}
		}
		fmt.Printf("\ndone\n")

		items := []item{
			{price: 10.},
			{price: 20.},
			{price: 30.},
		}
		for _, item := range items {
			item.price *= 2
		}
		fmt.Printf("%+v\n", items)
		for item := range items {
			items[item].price *= 2
		}
		fmt.Printf("%+v\n", items)

	//? 14. errors
	err01 := errors.New("error01")
	fmt.Printf("%[1]p %[1]T %[1]v\n", err01)
	fmt.Println(err01.Error())
	fmt.Println(err01)

	err0 := fmt.Errorf("add info: %w", errors.New("original error"))
	fmt.Printf("%[1]p %[1]T %[1]v\n", err0)
	fmt.Println(errors.Unwrap(err0))
	fmt.Printf("%T\n", errors.Unwrap(err0))
	err1 := fmt.Errorf("add info: %v", errors.New("original error"))
	fmt.Println(err1)

	err2 := fmt.Errorf("in repogitory layer: %w", ErrCustom)
	fmt.Println(err2)
	err2 = fmt.Errorf("in service layer: %w", err2)
	fmt.Println(err2)

	if errors.Is(err2, ErrCustom) {
		fmt.Println("err2 is ErrCustom")
	}

	file := "dummy.txt"
	err3 := fileChecker(file)
	if err3 != nil {
		if errors.Is(err3, os.ErrNotExist) {
			fmt.Println("file not found :", file)
		} else {
			fmt.Println("error", err3)
		}
	} else {
		fmt.Println("file found :", file)
	}

	//? 15. generics
	fmt.Printf("%v\n", add(1, 2))
	fmt.Printf("%v\n", add(1.1, 2.2))
	fmt.Printf("%v\n", add("A", "B"))
	var i1, i2 NewInt = 1, 2
	fmt.Printf("%v\n", add(i1, i2))
	fmt.Printf("%v\n", min(1, 2))
}

func (task Task) extendEstimate() {
	task.Estimate += 10
}
func (taskp *Task) extendEstimatePointer() {
	taskp.Estimate += 10
}

func funcDefer() {
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	fmt.Println("done")
}

func trimExtension(files ...string) []string {
	out := make([]string, 0, len(files))
	for _, f := range files {
		out = append(out, strings.TrimSuffix(f, ".txt"))
	}
	return out
}

func fileChecker(name string) (string, error) {
	f, err := os.Open(name)
	if err != nil {
		return "", errors.New("file not found")
	}
	defer f.Close()
	return name, nil
}

func addExt(f func(file string) string, name string) {
	fmt.Println(f(name))
}

func multiply() func(int) int {
	return func(n int) int {
		return n * 2
	}
}

func countUp() func(int) int {
	count := 0
	return func(n int) int {
		count += n
		return count
	}
}

func checkType(i any) {
	switch i.(type) {
	case nil:
		fmt.Println("nil")
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown")
	}
}

type item struct {
	price float32
}

func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("fileChecker err: %w", err)
	}
	defer f.Close()
	return nil
}
