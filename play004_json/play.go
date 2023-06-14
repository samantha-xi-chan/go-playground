package play004_json

import (
	"encoding/json"
	"fmt"
)

func TestJson() {
	// 定义一个包含整数的切片
	numbers := []int{1, 2, 3, 4, 5}

	// 将切片转换为JSON字符串
	jsonData, err := json.Marshal(numbers)
	if err != nil {
		fmt.Println("转换为JSON时出错:", err)
		return
	}

	// 打印JSON字符串
	fmt.Println(string(jsonData))

	// 将JSON字符串解析为切片
	var parsedNumbers []int
	err = json.Unmarshal(jsonData, &parsedNumbers)
	if err != nil {
		fmt.Println("解析JSON时出错:", err)
		return
	}

	// 打印解析后的切片
	fmt.Println(parsedNumbers)
}
