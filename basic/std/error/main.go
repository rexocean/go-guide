package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	// 1. 演示errors.New基础用法
	fmt.Println("=== 1. errors.New基础用法 ===")
	basicErrorDemo()
	fmt.Println()

	// 2. 演示错误判断与处理模式
	fmt.Println("=== 2. 错误判断与处理模式 ===")
	errorHandlingDemo()
	fmt.Println()

	// 3. 演示自定义错误类型（实现error接口）
	fmt.Println("=== 3. 自定义错误类型 ===")
	customErrorDemo()
	fmt.Println()

	// 4. 演示自定义错误的类型断言与扩展信息提取
	fmt.Println("=== 4. 自定义错误的类型断言 ===")
	customErrorAssertionDemo()
	fmt.Println()

	// 5. 演示错误在实际业务中的应用
	fmt.Println("=== 5. 业务场景中的错误应用 ===")
	businessErrorDemo()
}

// 1. errors.New基础用法：创建简单字符串错误
func basicErrorDemo() {
	// 使用errors.New创建错误实例
	err := errors.New("这是一个基础错误")
	fmt.Printf("错误内容: %v\n", err)
	fmt.Printf("错误类型: %T\n", err) // *errors.errorString

	// 空错误判断（成功场景）
	successErr := getSuccessStatus()
	if successErr == nil {
		fmt.Println("getSuccessStatus: 操作成功（错误为nil）")
	}
}

// 辅助函数：返回nil表示成功
func getSuccessStatus() error {
	return nil // 无错误时返回nil
}

// 2. 错误判断与处理的标准模式
func errorHandlingDemo() {
	// 测试正常情况
	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("除法运算失败: %v\n", err)
	} else {
		fmt.Printf("10 / 2 = %v\n", result)
	}

	// 测试错误情况（除数为0）
	result, err = divide(10, 0)
	if err != nil {
		fmt.Printf("除法运算失败: %v\n", err) // 会执行此分支
	} else {
		fmt.Printf("10 / 0 = %v\n", result)
	}
}

// 辅助函数：除法运算，除数为0时返回错误
func divide(a, b int) (int, error) {
	if b == 0 {
		// 使用errors.New返回错误信息
		return 0, errors.New("除数不能为0")
	}
	return a / b, nil
}

// 3. 自定义错误类型（实现error接口）
func customErrorDemo() {
	// 调用返回自定义错误的函数
	err := simulateFailure()
	if err != nil {
		fmt.Printf("捕获到错误: %v\n", err)
		// 自定义错误会调用其Error()方法返回字符串
	}
}

// 定义自定义错误类型（包含更多上下文信息）
type OperationError struct {
	Operation string    // 操作名称
	Time      time.Time // 错误发生时间
	Reason    string    // 错误原因
}

// 实现error接口的Error()方法
func (e *OperationError) Error() string {
	return fmt.Sprintf("操作失败 [%s] - 时间: %v, 原因: %s",
		e.Operation, e.Time.Format("2006-01-02 15:04:05"), e.Reason)
}

// 辅助函数：模拟一个操作失败并返回自定义错误
func simulateFailure() error {
	return &OperationError{
		Operation: "文件写入",
		Time:      time.Now(),
		Reason:    "磁盘空间不足",
	}
}

// 4. 自定义错误的类型断言与信息提取
func customErrorAssertionDemo() {
	err := simulateNetworkError()
	if err != nil {
		fmt.Printf("错误信息: %v\n", err)

		// 使用类型断言判断错误类型
		if netErr, ok := err.(*NetworkError); ok {
			// 提取自定义错误中的详细信息
			fmt.Printf("错误详情 - 代码: %d, 地址: %s, 重试建议: %v\n",
				netErr.Code, netErr.Address, netErr.Retryable)
		} else {
			fmt.Println("这不是NetworkError类型的错误")
		}
	}
}

// 定义网络错误类型（包含状态码、地址等信息）
type NetworkError struct {
	Code      int    // 错误代码（如500、404）
	Address   string // 访问的地址
	Retryable bool   // 是否可重试
}

// 实现error接口
func (e *NetworkError) Error() string {
	return fmt.Sprintf("网络错误 %d: 访问 %s 失败", e.Code, e.Address)
}

// 辅助函数：模拟网络错误
func simulateNetworkError() error {
	return &NetworkError{
		Code:      503,
		Address:   "https://example.com/api",
		Retryable: true, // 503错误通常可重试
	}
}

// 5. 业务场景中的错误应用示例
func businessErrorDemo() {
	// 测试用户注册业务
	_, err := registerUser("", "123456")
	if err != nil {
		fmt.Printf("注册失败: %v\n", err)
	}

	_, err = registerUser("testuser", "123")
	if err != nil {
		fmt.Printf("注册失败: %v\n", err)
	}

	username, err := registerUser("validuser", "123456")
	if err == nil {
		fmt.Printf("注册成功，用户名: %s\n", username)
	}
}

// 业务错误类型：用户注册错误
type RegisterError struct {
	Field   string // 错误字段（如用户名、密码）
	Message string // 错误信息
}

func (e *RegisterError) Error() string {
	return fmt.Sprintf("注册错误（%s）: %s", e.Field, e.Message)
}

// 辅助函数：用户注册业务逻辑
func registerUser(username, password string) (string, error) {
	// 验证用户名
	if username == "" {
		return "", &RegisterError{
			Field:   "用户名",
			Message: "不能为空",
		}
	}

	// 验证密码长度
	if len(password) < 6 {
		return "", &RegisterError{
			Field:   "密码",
			Message: "长度不能少于6位",
		}
	}

	// 验证通过
	return username, nil
}
