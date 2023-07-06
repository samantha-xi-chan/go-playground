package play000_grammar

import "fmt"

type Aliasint int // 定义int 的别名 为Aliasint

const (
	AA Aliasint = iota //初始化 0
	BB                 // 1
	CC                 // 2
)

func test(m Aliasint) { fmt.Println(m) }

func Play() {
	m := AA
	test(m)
	x := 1
	//test(x)
	// cannot use x (type int) as type Aliasint in argument to test
}
