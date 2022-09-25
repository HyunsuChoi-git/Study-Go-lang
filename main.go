package main

import (
	"fmt"
	"strings"
)

func main() {
	println("hi")
	fmt.Println("Hello World!")
	// 라이브러리 내 기능을 export하는 방법. 메소드를 대문자로 시작하기.
	// 대문자 메소드 : public method
	// 소문자 메소드 : private method

	const name string = "hera"
	fmt.Println(name)

	var name2 string = "hera2"
	name3 := "hera3"

	fmt.Println(name2, name3)

	fmt.Println(multiply(3, 4))
	fmt.Println(lenAndUpper(name))
	repeatMe("a", "b", "c", "d", "e")
}

func multiply(a int, b int) int {
	return a * b
}

func lenAndUpper(name string) (lenName int, upperName string) {
	defer fmt.Println("defer: I am done") // func리턴 후 실행되는 finally개념
	// return len(name), strings.ToUpper((name))
	lenName = len(name)
	upperName = strings.ToUpper(name)
	return // naked return 방법!
}

func repeatMe(words ...string) {
	fmt.Println(words)
}
