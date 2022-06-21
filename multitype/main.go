package main

import (
	"fmt"
	"strconv"
)

//定义泛型切片
// v := vector[int]{10, 20, 30}
type vector[T any] []T

//定义泛型函数
// printslice(v)
// printslice([]int{1, 2, 3, 4, 5, 6})
func printslice[T any](s []T) {
	for _, v := range s {
		fmt.Printf("%v \t", v)
	}
	fmt.Println()
}

//定义泛型map
// m1 := dict[string, int]{"aaa": 1}
type dict[K NumStr, V any] map[K]V

//定义泛型通道
// c1 := make(Trail[int])
type Trail[T any] chan T

//定义泛型类型限制
type Num interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Double interface {
	~float32 | ~float64
}

type Complex interface {
	~complex64 | ~complex128
}

//泛型类型限制可以组合
type NumStr interface {
	Num | ~string | Double
}

type Oper interface {
	Num | ~string | Double | Complex
}

// fmt.Println(max("aa", "bb"))
func max[T NumStr](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// fmt.Println(add(123, 456))
func add[T Oper](a, b T) T {
	return a + b
}

//泛型函数约束(约束该类型必须包含某函数)
type ShowPrice interface {
	String() string
	//Num | ~string	类型约束和函数约束可以同时存在
}

type Price int

func (i Price) String() string {
	return strconv.Itoa(int(i))
}

func ShowPriceSlice[T ShowPrice](s []T) {
	for _, x := range s {
		fmt.Printf("%s \t", x.String())
	}
	fmt.Println()
}

//保留约束类型 any compareable
// fmt.Println(findFunc([]int{1, 2, 3, 4, 5}, 3))
func findFunc[T comparable](a []T, v T) int {
	for i, x := range a {
		if x == v {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(findFunc([]int{1, 2, 3, 4, 5}, 3))

	//ShowPriceSlice([]Price{1, 2, 3, 4, 5})

	// fmt.Println(max("aa", "bb"))
	// fmt.Println(max(1, 2))
	// fmt.Println(add(123, 456))

	// c1 := make(Trail[int])
	// go func() {
	// 	for x := range c1 {
	// 		fmt.Println(x)
	// 	}
	// }()
	// for i := 0; i < 10; i++ {
	// 	c1 <- i
	// }
	// time.Sleep(1 * time.Second)

	// m1 := dict[string, int]{"aaa": 1}
	// m1["bbb"] = 2
	// fmt.Println(m1)

	// v := vector[int]{10, 20, 30}
	// printslice(v)
	// v2 := vector[string]{"aa", "aa", "aa", "aa", "aa", "aa"}
	// printslice(v2)

	// printslice([]int{1, 2, 3, 4, 5, 6})
	// printslice([]float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.5})
	// printslice([]string{"aa", "aa", "aa", "aa", "aa", "aa"})
}
