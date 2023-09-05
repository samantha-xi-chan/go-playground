package play260_array

import "fmt"

func Play() {
	arrayA := []string{"item1", "item2", "item3", "item4", "item5"}
	arrayB := make([]string, len(arrayA)+2)

	arrayB[0] = "fir"
	arrayB[1] = "sec"

	// 将 arrayA 的剩余元素拷贝到 arrayB
	copy(arrayB[2:], arrayA[0:])

	fmt.Println(arrayB)

}
