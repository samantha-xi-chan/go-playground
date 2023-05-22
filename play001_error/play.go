package play001_error

import (
	"errors"
	"fmt"
)

type CustomError struct {
	Message string
}

func (e *CustomError) Error() string {
	return e.Message
}

func Test() {
	// 示例使用 errors.Is() 函数
	err := fetchResource()
	if errors.Is(err, ErrNotFound) {
		fmt.Println("Resource not found")
	} else if err != nil {
		fmt.Println("Failed to fetch resource:", err)
	}

	// 示例使用 errors.As() 函数
	err = processError()
	var customErr *CustomError
	if errors.As(err, &customErr) {
		fmt.Println("Custom error occurred:", customErr.Message)
	} else if err != nil {
		fmt.Println("Unknown error occurred:", err)
	}
}

var ErrNotFound = errors.New("not found")

func fetchResource() error {
	// 模拟资源未找到的情况
	return ErrNotFound
}

func processError() error {
	// 模拟发生自定义错误的情况
	customErr := &CustomError{
		Message: "Custom error occurred",
	}
	return customErr
}
