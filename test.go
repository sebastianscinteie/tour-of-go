package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func nameOf(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func print(f func(...any) any, args ...any) {
	fmt.Print(nameOf(f), ": ", f(args...))
}

func pass[T any](x T) T {
	return x
}

func test[T any](f func(T) T, x T) any {
	return f(x)
}

func testspr[T any](f func(T) T, x T) any {
	return f(x)
}

func sum(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func swap_named(x, y string) (a, b string) {
	a = y
	b = x
	return
}

func main() {
	// fmt.Println(sum(10, 12))
	// fmt.Println(swap("a", "b"))
	// fmt.Println(swap_named("x", "y"))
	// a := func(x int) int {
	// 	return x
	// }
	// print(sum, 10, 12)
	fmt.Println(test(pass[int], 5))
	fmt.Println(testspr(pass[int], 5))
}
