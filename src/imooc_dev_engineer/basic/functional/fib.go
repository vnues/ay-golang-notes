package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

//斐波那契数列
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func fibonacci2() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type intGen func() int

//实现io.Reader接口
func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 100 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)

}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
func main() {
	//f := fibonacci()
	//for i := 0; i < 10; i++ {
	//	fmt.Println(f())
	//}

	f := fibonacci2()
	//f实现了reader接口，所以可以传入下述函数
	printFileContents(f)
}
